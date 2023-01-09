package kafka

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"temp/app/common/socket"
	"temp/app/kafka"
)

var kafkaOnce sync.Once

func Setup() {
	kafka.InitProducer(viper.GetString("kafka.topic"), viper.GetString("kafka.host"))
	kafka.InitConsumer(viper.GetString("kafka.host"))
	kafkaOnce.Do(func() {
		fmt.Println("kafka start")
		go kafka.ConsumerMsg(socket.ConsumerKafkaMsg)
	})
}
