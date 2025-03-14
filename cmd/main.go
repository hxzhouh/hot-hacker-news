package main

import (
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"hot-hacker-new/pkg/hackernews"
	"log"
	"log/slog"
	"path/filepath"

	"github.com/robfig/cron"
)

func main() {

	// 初始化数据库
	dbPath := filepath.Join("data", "hackernews.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		slog.Error("初始化数据库失败: %v", err)
		return
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		}
	}()
	_ = db.AutoMigrate(&models.PostLink{})
	c := cron.New()
	spec := "0 0 17 * * ?" // 每天17点运行一次
	err = c.AddFunc(spec, hackernews.Crawle)
	if err != nil {
		slog.Error("添加定时任务失败: %v", err)
		return
	}
	c.Start()

	// timeAfter := time.Now().AddDate(0, 0, -500)
	// i := 0
	// for i < 500 {
	// 	urlPath := "https://www.daemonology.net/hn-daily/%s.html"
	// 	day := timeAfter.Format("2006-01-02")
	// 	urlPath = fmt.Sprintf(urlPath, day)
	// 	i += 1
	// 	pages, err := hackernews.ParseDailyPage(urlPath)
	// 	if err != nil {
	// 		slog.Error("解析页面失败: %v", err)
	// 		return
	// 	}
	// 	for _, v := range pages {
	// 		_ = models.CreatePostLink(database.DB, v)
	// 	}
	// 	timeAfter = timeAfter.AddDate(0, 0, 1)
	// 	fmt.Println(timeAfter.Format("2006-01-02"))
	// }
}
