package log

import (
	"fmt"
	"os"
	"sync"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	LOG_LEVEL = ""
	once      sync.Once
)

func init() {
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
}

type logger struct {
	logger kitlog.Logger
}

// インスタンスを保持する変数も小文字にすることでエクスポートを行わない
var instance *logger

func getInstanceOnce() {
	// if instance == nil {
	instance = &logger{}
	instance.logger = kitlog.NewLogfmtLogger(os.Stderr)
	switch LOG_LEVEL {
	case "DEBUG":
		instance.logger = level.NewFilter(instance.logger, level.AllowDebug())
	case "INFO":
		instance.logger = level.NewFilter(instance.logger, level.AllowInfo())
	case "WARN":
		instance.logger = level.NewFilter(instance.logger, level.AllowWarn())
	case "ERROR":
		instance.logger = level.NewFilter(instance.logger, level.AllowError())
	}
	instance.logger = kitlog.With(instance.logger, "ts", kitlog.DefaultTimestamp)
	// }
}

// インスタンス取得用の関数のみエクスポートしておき、ここでインスタンスが
// 一意であることを保証する
func getInstance() *logger {
	once.Do(getInstanceOnce)
	return instance
}

const (
	InfoColor    = "[1;34m"
	NoticeColor  = "[1;36m"
	WarningColor = "[1;33m"
	ErrorColor   = "[1;31m"
	DebugColor   = "[0;36m"
	Escape       = "\033"
	End          = "[0m"
)

func Log(level2, msg2 string) {
	fmt.Print(Escape)
	switch level2 {
	case "DEBUG":
		fmt.Print(DebugColor)
		level.Debug(getInstance().logger).Log("msg", msg2)
	case "INFO":
		fmt.Print(InfoColor)
		level.Info(getInstance().logger).Log("msg", msg2)
	case "WARN":
		fmt.Print(WarningColor)
		level.Warn(getInstance().logger).Log("msg", msg2)
	case "ERROR":
		fmt.Print(ErrorColor)
		level.Error(getInstance().logger).Log("msg", msg2)
	}
	fmt.Print(Escape)
	fmt.Print(End)
}
