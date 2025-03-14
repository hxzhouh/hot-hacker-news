package hackernews

import (
	"fmt"
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"log/slog"
	"time"
)

func Crawle() {
	slog.Info("开始爬取...")
	runCount := 0
CRAWLE:
	urlPath := "https://www.daemonology.net/hn-daily/%s.html"
	day := time.Now().Format("2006-01-02")
	urlPath = fmt.Sprintf(urlPath, day)
	pages, err := ParseDailyPage(urlPath)
	if err != nil {
		fmt.Printf("解析页面失败: %v\n", err)
		time.Sleep(100 * time.Second)
		if runCount < 100 {
			runCount++
			goto CRAWLE
		} else {
			return
		}
	}
	for _, v := range pages {
		_ = models.CreatePostLink(database.DB, v)
	}
}
