package config

import (
	"fmt"
	"shadowCloud/internal/tool"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 用于存放所有配置对应的结构体
type Config struct {
	App      App
	Logger   Logger
	Database Database
	Redis    Redis
}

func GetConfig() (c *Config) {
	// 获取项目根目录
	rootDir := tool.GetRootDir()
	println("GetConfig-rootDir: ", rootDir)
	//实例化viper，并根据地址读取配置文件
	v := viper.New()
	v.SetConfigFile(rootDir + "/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ConfigReadError: %s", err))
	}
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("reloading config...", in.Name)
		// 重载配置
		if err := v.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}
	})
	// 将读取到的配置文件绑定到返回参数c
	err := v.Unmarshal(&c)
	if err != nil {
		fmt.Println("ConfigUnmarshalError: ", err)
		return
	}
	return
}
