package main

import (
	"time"

	"go.uber.org/zap"
)

func demo1() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := "http://xxx.com"
	sugar.Infow("failed to fetch URL",
		"url", "url",
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("Failed to fetch URL: %s", url)
}

func main() {
	demo1()
}
