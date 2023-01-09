package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
	"temp/app/constants/errorCode"
	"temp/app/utils"
)

var producer sarama.AsyncProducer
var topic = "go-msg"

func CloseProducer() {
	if producer != nil {
		producer.Close().Error()
	}
}

// InitProducer 初始化生产者
func InitProducer(topicInput, host string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	client := util.NewResult(sarama.NewClient(strings.Split(host, ","), config)).UnwrapOr(errorCode.SERVER_ERROR, "init kafka client error")
	producer = util.NewResult(sarama.NewAsyncProducerFromClient(client)).UnwrapOr(errorCode.SERVER_ERROR, "init kafka async client error")
}

// Send 发送消息
func Send(data []byte) {
	encoder := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: encoder}
}
