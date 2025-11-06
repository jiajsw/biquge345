package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

func main() {
	fmt.Println("开始")
	// https://www.biquge345.com/book/953622/
	url := "https://www.biquge345.com/book/229622/"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}
	// 获取标题
	title := htmlquery.InnerText(htmlquery.FindOne(doc, "//h1"))

	status := "连载中"
	fmt.Println(status)
	// 创建文件
	f, _ := os.Create(title + "-" + status + ".txt")
	defer f.Close()
	fmt.Println(title)

	list := htmlquery.Find(doc, "//div[@class=\"border\"]//li//a/@href")
	// for _, node := range list {
	// 	fmt.Println(htmlquery.InnerText(node))
	// }
	// 跳过前 60
	time.Sleep(1 * time.Second)
	list = list[70:]
	prefix := "https://www.biquge345.com"
	for index, node := range list {
		url := prefix + htmlquery.SelectAttr(node, "href")
		fmt.Println(url)
		doc, err = htmlquery.LoadURL(url)
		if err != nil {
			panic(err)
		}
		// 获取标题
		t := htmlquery.InnerText(htmlquery.FindOne(doc, "//h1"))
		f.WriteString(strings.TrimSpace(t) + "\n")
		fmt.Println(t)
		contentNode := htmlquery.Find(doc, "//div[@id=\"txt\"]//text()")
		text := []string{}

		for _, c := range contentNode {
			i := htmlquery.InnerText(c)
			// 去掉空白字符,如果不为空字符串 添加
			i = strings.TrimSpace(i)
			if i != "" {
				text = append(text, i)
			}
		}
		// 写入文件
		for _, i := range text {
			f.WriteString(i + "\n")
		}
		// 已完成 index / len(list)
		fmt.Println("完成", index, "/", len(list))

		time.Sleep(1 * time.Second)
		// 获取内容
	}

	fmt.Println("结束")

}
