package mq

import (
	"github.com/spf13/viper"
	"temp/app/common/kafka"
	"temp/app/constants/contentType"
)

func Setup() {
	if viper.GetString("msgChannel.type") == contentType.KAFKA {
		kafka.Setup()
	}
}
