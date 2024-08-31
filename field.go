package zapsls

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/http2curl"
	"net/http"
	"net/http/httputil"
)

func ErrMsg(errMsg string) zap.Field {
	return zap.Field{Key: "errMsg", Type: zapcore.StringType, String: errMsg}
}

func CURL(req *http.Request) zap.Field {
	command, _ := http2curl.GetCurlCommand(req)
	curl := command.String()
	return zap.Field{Key: "curl", Type: zapcore.StringType, String: curl}
}

func HTTPRequest(req *http.Request) zap.Field {
	dump, _ := httputil.DumpRequest(req, true)
	return zap.Field{Key: "HTTPRequest", Type: zapcore.StringType, String: string(dump)}
}

func HTTPResponse(resp *http.Response) zap.Field {
	dump, _ := httputil.DumpResponse(resp, true)
	return zap.Field{Key: "HTTPResponse", Type: zapcore.StringType, String: string(dump)}
}
