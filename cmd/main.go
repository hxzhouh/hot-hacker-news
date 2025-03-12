package main

import (
	"fmt"
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"hot-hacker-new/pkg/hackernews"
	"log"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/robfig/cron"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

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
	db.AutoMigrate(&models.PostLink{})
	c := newWithSeconds()
	spec := "0 0 17 * * ?" // 每天17点运行一次
	c.AddFunc(spec, hackernews.Crawle())
	c.Start()
	timeAfter := time.Now().AddDate(0, 0, -500)
	i := 0
	for {
		urlPath := "https://www.daemonology.net/hn-daily/%s.html"
		day := timeAfter.Format("2006-01-02")
		urlPath = fmt.Sprintf(urlPath, day)
		timeAfter.AddDate(0, 0, 1)
		i += 1
		pages, err := hackernews.ParseDailyPage(urlPath)
		if err != nil {
			slog.Error("解析页面失败: %v", err)
		}
		for _, v := range pages {
			models.CreatePostLink(database.DB, v)
		}
	}
}
