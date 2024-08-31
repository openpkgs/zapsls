package zapsls

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Logstore        string
	Project         string

	MaxSize  int64
	LogLevel zapcore.Level
	Topic    string
	Source   string
}

func (c *Config) Validate() error {
	if c.Endpoint == "" {
		return fmt.Errorf("endpoint must not be blank")
	}
	if c.AccessKeyId == "" {
		return fmt.Errorf("accessKeyId must not be blank")
	}
	if c.AccessKeySecret == "" {
		return fmt.Errorf("accessKeySecret must not be blank")
	}
	if c.Logstore == "" {
		return fmt.Errorf("logstore must not be blank")
	}
	if c.Project == "" {
		return fmt.Errorf("project must not be blank")
	}
	return nil
}
