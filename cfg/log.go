package cfg

import (
	"github.com/miajio/zlog"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	log := zlog.Logger{
		Path:       "./log",
		MaxSize:    256,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}

	lm := zlog.LogMap{
		"debug": zlog.DebufLevel,
		"info":  zlog.InfoLevel,
		"error": zlog.ErrorLevel,
	}

	log.Generate(lm)
	Log = log.Log
}
