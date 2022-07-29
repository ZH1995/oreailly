package main

import (
    "fmt"
    "strconv"
    "time"
    "regexp"
    "github.com/gocolly/colly/v2"
)
func crawler(url string, index int) {
    c := colly.NewCollector(
        colly.AllowedDomains("www.oreilly.com.cn"),
    )
    c.OnHTML("a[title]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        absLink := e.Request.AbsoluteURL(link)
        matched, _ := regexp.MatchString("^《.*?》$", e.Text)
        if matched { 
            fmt.Printf("%d\t%s\t%s\n", index, e.Text, absLink)
        }
    })
    c.Visit(url)
}
func main() {
    for i := 2022; i >= 1999; i -- {
        url := "http://www.oreilly.com.cn/index.php?func=completelist&pubyear=" + strconv.Itoa(i)
        crawler(url, i)
        time.Sleep(time.Duration(2) * time.Second)
    }
}
