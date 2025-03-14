package main

import (
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"hot-hacker-new/pkg/hackernews"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/robfig/cron"
)

func main() {
	initLog()
	// 初始化数据库
	dbPath := filepath.Join("data", "hackernews.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		slog.Error("初始化数据库失败:", slog.Any("error", err))
		return
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			slog.Error("关闭数据库连接失败:", slog.Any("error", err))
		}
	}()
	_ = db.AutoMigrate(&models.PostLink{})
	c := cron.New()
	spec := "0 0 17 * * ?" // 每天17点运行一次
	err = c.AddFunc(spec, hackernews.Crawle)
	if err != nil {
		slog.Error("添加定时任务失败", slog.Any("error", err))
		return
	}
	c.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("退出...")
	c.Stop()
	os.Exit(0)
}

func initLog() {
	f, err := os.OpenFile("hackernews.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	handler := slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(handler))
}
