package zapsls_test

import (
	"bytes"
	"go.uber.org/zap"
	"net/http"
	"testing"
	"zapsls"
)

func TestInitLogger(t *testing.T) {
	conf := &zapsls.Config{
		Endpoint:        "",
		AccessKeyId:     "",
		AccessKeySecret: "",
		Logstore:        "",
		Project:         "",
	}
	zapsls.InitLogger(conf)

	logger := zapsls.NewLogger().
		With(zap.Int("uid", 398)).
		With(zap.Int("parent_uid", 1398))

	req, _ := http.NewRequest("PUT", "http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu",
		bytes.NewBufferString(`{"hello":"world","answer":42}`))
	client := http.Client{}
	resp, _ := client.Do(req)
	logger.Info("测试错误",
		zap.Int64("id", 123),
		zap.String("name", "Luca"),
		zapsls.CURL(req))
	logger.Info("测试错2d误",
		zap.String("name", "allen"),
		zapsls.CURL(req),
		zapsls.HTTPResponse(resp))
	logger.Info("测试43错误",
		zap.Int64("id", 45),
		zapsls.CURL(req),
		zapsls.HTTPResponse(resp))
	select {}
}
