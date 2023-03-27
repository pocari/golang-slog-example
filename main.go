package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	fmt.Println("hello, world")

	logger := slog.New(slog.HandlerOptions{
		Level: slog.LevelDebug,
	}.NewJSONHandler(os.Stdout))
	logger.Enabled(context.Background(), slog.LevelDebug)
	logger.Info("message1", "piyora", "hoge", "piyora2", []int64{1, 2, 3})

	logger.Error("error message")
	logger.Debug("debug message")
	logger.Warn("warn message")

	logger.Info("info message", slog.Group("group key",
		slog.String("hoge", "string value"),
		slog.Group("group subkey", slog.Int("int key", 10)),
		slog.Duration("duration1", 2*time.Hour)))
}
