package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
    "regexp"
)

func main() {
    fmt.Println("Hello")
    f, err := os.OpenFile("crawler_ret", os.O_RDONLY, 0)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    re := regexp.MustCompile("[A-Za-z]{1,}")
    sc := bufio.NewScanner(f)
    hash := make(map[string]int)
    for sc.Scan() {
        data := strings.Split(sc.Text(), "\t")
        // parse book name
        slice := re.FindStringSubmatch(data[1])
        if len(slice) != 0 {
            hash[slice[0]] ++
        } else {
            hash["other"] ++
        }
    }
    for k, v := range hash {
        fmt.Println(k, v)
    }
}
