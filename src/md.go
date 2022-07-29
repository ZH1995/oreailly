package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    fmt.Println("Hello")
    f, err := os.OpenFile("crawler_ret", os.O_RDONLY, 0)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    fmt.Println("|年份|书名|链接")
    fmt.Println("|--|--|--|")
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        data := strings.Split(sc.Text(), "\t")
        fmt.Printf("|%s|%s|%s|\n", data[0], data[1], data[2])
    }
}
