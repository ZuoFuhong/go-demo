// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

// 基本概念
//  Node：单个Elastic实例称为一个节点
//  Cluster：一组节点构成一个集群
//  Index：Elastic 会索引所有字段，经过处理后写入一个反向索引。Elastic数据管理的顶层单位就叫做 Index。
//  Document：Index 里面单条的记录称为 Document（文档）
//  Type：Document 可以分组，不同的 Type 应该有相似的结构（schema）
//
// 文档：https://pkg.go.dev/gopkg.in/olivere/elastic.v5?tab=doc
type Cli struct {
	client *elastic.Client
}

func NewCli(addr string) *Cli {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		panic(err)
	}
	cli := new(Cli)
	cli.client = client
	return cli
}

func (cli *Cli) Ping() {
	result, code, err := cli.client.Ping("http://47.98.199.80:9200").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("code = %d, ret = %v", code, result)
}

func (cli *Cli) PutDocument() {
	book := `{"name": "微信小程序", "price": "10"}`
	resp, err := cli.client.Index().Index("book").Type("science").Id("1").BodyString(book).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("id = %s, index = %s, type = %s\n", resp.Id, resp.Index, resp.Type)
}

func (cli *Cli) GetDocument(id string) {
	resp, err := cli.client.Get().Index("book").Type("science").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(*resp.Source))
}

func (cli *Cli) SearchDocument() {
	query := elastic.NewMatchQuery("name", "微信")
	resp, err := cli.client.Search().
		Index("book").Type("science").Query(query).From(0).Size(1).Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("hits.total = %d, _source = %s", resp.Hits.TotalHits, string(*resp.Hits.Hits[0].Source))
}

func main() {
	cli := NewCli("http://47.98.199.80:9200")
	cli.Ping()
	cli.PutDocument()
	cli.GetDocument("1")
	cli.SearchDocument()
}
