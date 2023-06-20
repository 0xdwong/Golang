# zap

## 安装
go get -u go.uber.org/zap

## 快速开始
```
logger, _ := zap.NewProduction()
defer logger.Sync() // flushes buffer, if any

sugar := logger.Sugar()

url := "http://xxx.com"
sugar.Infow("failed to fetch URL",
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
# 或者
sugar.Infof("Failed to fetch URL: %s", url)
```