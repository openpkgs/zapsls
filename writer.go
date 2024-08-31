package zapsls

import (
	"encoding/json"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/gogo/protobuf/proto"
	"io"
	"strconv"
	"sync"
	"time"
)

const (
	defaultMaxSize = 100
	megabyte       = 1024 * 1024
)

type Writer struct {
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	project         string
	logstore        string
	m               sync.Mutex
	slsProducer     *producer.Producer

	maxSize int64
	topic   string
	source  string
}

func newDefaultWriter(config *Config) io.Writer {
	return &Writer{
		endpoint:        config.Endpoint,
		accessKeyId:     config.AccessKeyId,
		accessKeySecret: config.AccessKeySecret,
		logstore:        config.Logstore,
		project:         config.Project,
		maxSize:         config.MaxSize,
		topic:           config.Topic,
		source:          config.Source,
	}
}

func (w *Writer) Write(p []byte) (int, error) {
	w.m.Lock()
	defer w.m.Unlock()

	length := int64(len(p))
	maxSize := func() int64 {
		if w.maxSize == 0 {
			return defaultMaxSize * megabyte
		}
		return w.maxSize * megabyte
	}()
	if length > maxSize {
		return 0, fmt.Errorf("write length %d exceeds max size %d", length, maxSize)
	}

	if w.slsProducer == nil {
		producerConfig := producer.GetDefaultProducerConfig()
		producerConfig.Endpoint = w.endpoint
		producerConfig.AccessKeyID = w.accessKeyId
		producerConfig.AccessKeySecret = w.accessKeySecret
		slsProducer := producer.InitProducer(producerConfig)
		slsProducer.Start()
		w.slsProducer = slsProducer
	}

	msg := map[string]interface{}{}
	err := json.Unmarshal(p, &msg)
	if err != nil {
		return 0, err
	}
	logDetails := mapInterfaceToString(msg)
	content := []*sls.LogContent{}
	for k, v := range logDetails {
		content = append(content, &sls.LogContent{
			Key:   proto.String(k),
			Value: proto.String(v),
		})
	}

	topic := w.topic
	if topic == "" {
		topic = "topic"
	}
	source := w.source
	if source == "" {
		source = "127.0.0.1"
	}
	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	err = w.slsProducer.SendLog(w.project, w.logstore, topic, source, log)
	if err != nil {
		return 0, err
	}
	return int(length), nil
}

func (w *Writer) Close() error {
	w.m.Lock()
	defer w.m.Unlock()

	if w.slsProducer == nil {
		return nil
	}
	w.slsProducer.SafeClose()
	w.slsProducer = nil
	return nil
}

func interfaceToString(i interface{}) string {
	str := ""
	switch v := i.(type) {
	case string:
		str = v
	case int:
		str = strconv.Itoa(v)
	case int64:
		str = strconv.FormatInt(v, 10)
	case float64:
		str = strconv.FormatFloat(v, 'f', 2, 64)
	case uint64:
		str = strconv.FormatUint(v, 10)
	case error:
		str = v.Error()
	// Add cases for other types that you want to handle...
	default:
		msgBytes, err := json.Marshal(i)
		if err == nil {
			str = string(msgBytes)
		}
	}
	return str
}

func mapInterfaceToString(input map[string]interface{}) map[string]string {
	ret := make(map[string]string, 0)
	for k, v := range input {
		ret[k] = interfaceToString(v)
	}
	return ret
}
