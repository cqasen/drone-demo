# gin-demo

- 采用了[ebar-go/ego](https://github.com/ebar-go/ego)框架 进行编码（基于Gin的web微服务框架）
- 部署采用了github的workflows/docker-image,将docker容器打包上传到docker hub仓库



## ES相关


```
# 创建_ingest
PUT _ingest/pipeline/indexed_at
{
  "description" : "向文档添加时间戳",
    "processors" : [
      {
        "set" : {
          "field" : "_source.@timestamp",
          "value" : "{{_ingest.timestamp}}"
        }
      },
      {
        "script" : {
          "source" : "ctx.cont_length = ctx.contents.length();"
        }
      }
    ]
  
}



# 查看_ingest
GET _ingest/pipeline/indexed_at

# 查询索引模板
GET _template/my_poetry_template

# 删除索引模板
DELETE _template/my_poetry_template

# 创建索引模板
PUT _template/my_poetry_template
{
  "order": 0,
  "index_patterns": [
    "poetry*"
  ],
  "settings": {
    "index": {
      "number_of_replicas": "1",
      "default_pipeline": "indexed_at",
      "refresh_interval": "1s"
    }
  },
  "mappings": {
    "properties": {
      "cont_length": {
        "type": "long"
      },
      "contents": {
        "fielddata": true,
        "analyzer": "ik_max_word",
        "type": "text",
        "fields": {
          "field": {
            "type": "keyword"
          }
        }
      },
      "author": {
        "analyzer": "ik_max_word",
        "type": "text",
        "fields": {
          "field": {
            "type": "keyword"
          }
        }
      },
      "title": {
        "analyzer": "ik_max_word",
        "type": "text",
        "fields": {
          "field": {
            "type": "keyword"
          }
        }
      },
      "type": {
        "analyzer": "ik_max_word",
        "type": "text",
        "fields": {
          "field": {
            "type": "keyword"
          }
        }
      },
      "timestamp": {
        "type": "date"
      }~~~~
    }
  },
  "aliases": {
    "poetry": {}
  }
}

# 创建索引
PUT poetry_v1

# 查询
GET poetry/_search
{
  "query": {
    "multi_match": {
      "query": "张九龄",
      "fields": [
        "contents^0.3",
        "author^0.9",
        "type^0.1",
        "title^0.5"
      ],
      "type": "cross_fields", 
      "operator": "and",
      "tie_breaker": 0.3
    }
  },
  "sort": [
    {
      "_score": {
        "order": "desc"
      }
    },
    {
      "timestamp": {
        "order": "desc"
      }
    }
  ]
}
```