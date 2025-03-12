package hackernews

import (
	"fmt"
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"time"
)

func Crawle() {

	urlPath := "https://www.daemonology.net/hn-daily/%s.html"
	day := time.Now().Format("2006-01-02")
	urlPath = fmt.Sprintf(urlPath, day)
	pages, err := ParseDailyPage(urlPath)
	if err != nil {
		fmt.Printf("解析页面失败: %v\n", err)
		return
	}
	for _, v := range pages {
		models.CreatePostLink(database.DB, v)
	}
}
