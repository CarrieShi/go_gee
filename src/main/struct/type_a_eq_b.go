package main

import "fmt"

func main() {
	var news = otherNews{
		HeaderLine: "boom",
	}

	news.Report()
}

type News struct {
	HeaderLine string
}

func (news News) Report() {
	fmt.Printf("News headerline is " + news.HeaderLine + "\n")
}

type otherNews = News
