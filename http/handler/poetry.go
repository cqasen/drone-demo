package handler

import (
	"context"
	"github.com/cqasen/gin-demo/pkg/app"
	"github.com/ebar-go/ego/http/pagination"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)

var indexName = "poetry"

type PoetryItem struct {
	Author   string `json:"author"`
	Contents string `json:"contents"`
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

type PoetryList []  struct {
	PoetryItem
}

//推送
func PushPoetry(ctx *gin.Context) {
	dir, _ := os.Getwd()
	filename := dir + "./resources/300.json"

	// 读取JSON文件内容 返回字节切片
	bytes, _ := ioutil.ReadFile(filename)

	var poetryList PoetryList
	// 将字节切片映射到指定结构上
	egu.JsonDecode(bytes, &poetryList)

	client := app.Elasticsearch()
	bulkRequest := client.Bulk()
	for _, poetryItem := range poetryList {
		postJson, _ := egu.JsonEncode(poetryItem)
		req := elastic.NewBulkIndexRequest().Index(indexName).Id(strconv.Itoa(int(poetryItem.ID))).Doc(postJson)
		bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	log.Println("耗时：", bulkResponse.Took, "索引了：", len(bulkResponse.Items))
	response.WrapContext(ctx).Success("ok")

}

//搜索
func SearchPoetry(ctx *gin.Context) {
	word := ctx.Query("word")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	from := (page - 1) * page_size

	log.Println(word, page, page_size)

	client := app.Elasticsearch()
	query := elastic.NewBoolQuery()

	if word != "" {
		query.Must(elastic.NewMultiMatchQuery(word, "title^0.5", "author^0.8", "contents^0.1", "type^0.1").
			Type("best_fields").
			Operator("and").
			TieBreaker(0.3))
	} else {
		query.Must(elastic.NewMatchAllQuery())
	}

	searchRes, err := client.Search().
		Index(indexName).
		Query(query).
		From(from).
		Size(page_size).
		Timeout("10ms").
		Do(context.Background())
	if err != nil {
		ctx.Abort()
		response.WrapContext(ctx).Error(1, err.Error())
	}
	var poetryItem PoetryItem
	var poetryList []PoetryItem
	for _, item := range searchRes.Each(reflect.TypeOf(poetryItem)) { //从搜索结果中取数据的方法
		t := item.(PoetryItem)
		poetryList = append(poetryList, t)
	}

	pagination := pagination.Paginate(int(searchRes.Hits.TotalHits.Value), page, page_size)
	response.WrapContext(ctx).Paginate(poetryList, &pagination)
}
