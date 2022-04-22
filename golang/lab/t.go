package main

import (
	"encoding/json"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type P struct {
	Name    string
	Address string
}

func main() {
	l, _ := zap.NewProduction(zap.AddStacktrace(zapcore.LevelEnabler(zapcore.DebugLevel)))
	defer l.Sync()

	sugar := l.Sugar()

	p := P{
		Name:    "name",
		Address: "address",
	}
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"p", p,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")

	b, _ := json.Marshal(p)
	sugar.Info("json", string(b))
}
