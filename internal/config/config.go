package config

import (
	"log/slog"
	"time"
)

type Configuration struct {
	Addr         string
	LogLevel     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (c *Configuration) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("addr", c.Addr),
		slog.String("logLevel", c.LogLevel),
		slog.Any("rtimeout", c.ReadTimeout),
		slog.Any("wtimeout", c.WriteTimeout),
	)
}
