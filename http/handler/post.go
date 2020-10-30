package handler

import (
	"context"
	app2 "github.com/cqasen/gin-demo/pkg/app"
	"github.com/cqasen/gin-demo/pkg/service/dao"
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	pagination2 "github.com/ebar-go/ego/http/pagination"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"strconv"
)

//获取单条
func GetPost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := dao.Post(app.DB()).Get(int32(id))
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	response.WrapContext(ctx).Success(post)
}

//获取列表
func GetPostList(ctx *gin.Context) {
	post, err := dao.Post(app.DB()).GetList()
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	response.WrapContext(ctx).Success(post)
}

//单个推送
func PushPost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := dao.Post(app.DB()).Get(int32(id))
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	postJson, _ := egu.JsonEncode(post)
	client := app2.Elasticsearch()
	indexResponse, err := client.Index().Index("post_v1").Id(strconv.Itoa(int(post.LogID))).BodyJson(postJson).Do(context.Background())
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	log.Println(indexResponse.Status)
	response.WrapContext(ctx).Success(post)
}

//批量推送
func PushPostList(ctx *gin.Context) {
	postList, err := dao.Post(app.DB()).GetList()
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	client := app2.Elasticsearch()
	bulkRequest := client.Bulk()
	for _, post := range postList {
		postJson, _ := egu.JsonEncode(post)
		req := elastic.NewBulkIndexRequest().Index("post_v1").Id(strconv.Itoa(int(post.LogID))).Doc(postJson)
		bulkRequest.Add(req)
	}
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	log.Println("耗时：", bulkResponse.Took, "索引了：", len(bulkResponse.Items))
	response.WrapContext(ctx).Success(postList)
}

//搜索
func SearchPost(ctx *gin.Context) {
	word := ctx.Query("word")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	from := (page - 1) * page_size

	log.Println(word, page, page_size)

	client := app2.Elasticsearch()
	query := elastic.NewBoolQuery()
	query.Filter(elastic.NewTermQuery("log_status", 0))
	if word != "" {
		query.Must(elastic.NewMatchQuery("log_title", word))
	}

	fetchSourceContext := elastic.NewFetchSourceContext(true).Exclude("log_content")

	searchRes, err := client.Search().
		Index("post_v1").
		Query(query).
		From(from).
		Size(page_size).
		Timeout("10ms").
		FetchSourceContext(fetchSourceContext).
		Do(context.Background())
	if err != nil {
		ctx.Abort()
		response.WrapContext(ctx).Error(1, err.Error())
	}
	var post entity.ZbpPost
	var postList []entity.ZbpPost
	for _, item := range searchRes.Each(reflect.TypeOf(post)) { //从搜索结果中取数据的方法
		t := item.(entity.ZbpPost)
		postList = append(postList, t)
	}

	pagination := pagination2.Paginate(int(searchRes.Hits.TotalHits.Value), page, page_size)
	response.WrapContext(ctx).Paginate(postList, &pagination)
}
