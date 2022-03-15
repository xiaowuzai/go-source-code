package logs

import (
	"io"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Discard io.Writer = discard{}

type discard struct{}

func (discard) Write(p []byte) (int, error) {
	return len(p), nil
}

func (discard) WriteString(s string) (int, error) {
	return len(s), nil
}

func BenchmarkLogrus(b *testing.B) {
	b.ReportAllocs() // 开启内存分配统计, 相当于执行了 -test.benchmem。但是仅影响当前的 Benchmark
	b.StopTimer()    // 关闭 timer 对性能的影响
	logger := logrus.New()
	//logger.SetOutput(io.Discard) // go1.16 以上版本
	logger.SetOutput(io.Discard)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"url":     "http://foo.com",
			"attempt": 3,
			"backoff": time.Second,
		}).Info("failed to fetch URL")
	}
}

func BenchmarkZap(b *testing.B) {
	b.ReportAllocs() // 开启内存分配统计, 相当于执行了 -test.benchmem。但是仅影响当前的 Benchmark
	b.StopTimer()    // 关闭 timer 对性能的影响

	cfg := zap.NewProductionConfig()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)

	logger := zap.New(core)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Failed to fetch URL",
			zap.String("url", "http://foo.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
