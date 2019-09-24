package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mmcdole/gofeed"
	"os"
	"os/user"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("command not found")
	}
	if os.Args[1] == "createCsv" {
		if len(os.Args) < 3 {
			panic("number of command is not enouth")
		}
		ranking_csv(os.Args[2])
	} else {
		panic("command not found")
	}
}
func ranking_csv(filepath string) {
	usr, _ := user.Current()
	f := strings.Replace(filepath, "~", usr.HomeDir, 1)
	wfile, err := os.Create(f + time.Now().Format("2006-01-02_15:04:05.000") + ".csv")
	if err != nil {
		fmt.Println(err)
		panic("file error.")
	}
	fp := gofeed.NewParser()
	url := "https://www.nicovideo.jp/ranking/genre/entertainment?tag=%E3%81%AB%E3%81%98%E3%81%95%E3%82%93%E3%81%98&rss=2.0&lang=ja-jp"
	feed, _ := fp.ParseURL(url)
	writer := csv.NewWriter(wfile)

	writer.Write([]string{"title", "link"})
	for _, item := range feed.Items {
		writer.Write([]string{item.Title, item.Link})
	}
	writer.Flush()
}
