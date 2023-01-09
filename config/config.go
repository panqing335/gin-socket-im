package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sync"
)

var configOnce sync.Once

func Setup() {
	configOnce.Do(func() {
		dir, _ := os.Getwd()
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath(path.Join(dir, "config"))

		fmt.Println("dir:", path.Join(dir, "config"))

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
}
