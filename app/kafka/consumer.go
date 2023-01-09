package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"temp/app/constants/errorCode"
	"temp/app/utils"
)

var consumer sarama.Consumer

type ConsumerCallback func(data []byte)

// InitConsumer 初始化消费者
func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client := util.NewResult(sarama.NewClient(strings.Split(hosts, ","), config)).UnwrapOr(errorCode.SERVER_ERROR, "init kafka consumer client error")
	consumer = util.NewResult(sarama.NewConsumerFromClient(client)).UnwrapOr(errorCode.SERVER_ERROR, "init kafka consumer error")
}

// ConsumerMsg 消费消息
func ConsumerMsg(callBack ConsumerCallback) {
	// 处理来自给定主题和分区的Kafka消息
	partitionConsumer := util.NewResult(consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)).UnwrapOr(errorCode.SERVER_ERROR, "iConsumePartition error")
	//defer func() {
	//	err := partitionConsumer.Close()
	//	if err != nil {
	//		fmt.Printf("partitionConsumer.Close err: %v\n", err.Error())
	//		return
	//	}
	//}()
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		fmt.Printf("partitionConsumer: %v\n", string(msg.Value))
		if callBack != nil {
			callBack(msg.Value)
		}
	}
}

func CloseConsumer() {
	if consumer != nil {
		consumer.Close().Error()
	}
}
