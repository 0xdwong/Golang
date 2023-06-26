package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	logger.Info("hello world")
}
