package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

/*
 * @Author       : zhixi.fang (Pop)
 * @Date         : 2021-12-03 11:09:02
 * @LastEditors  : zhixi.fang (Pop)
 * @LastEditTime : 2021-12-03 15:16:54
 */

var (
	Config Configuration
)

type Configuration struct {
	Env     string              `yaml:"env"`
	AppName string              `yaml:"app_name"`
	RConfig RemoteConfiguration `yaml:"remote_conf"`
}

type RemoteConfiguration struct {
	LogPath                   string `yaml:"log_path"`
	LogLevel                  int16  `yaml:"log_level"`
	DbLogEnable               string `yaml:"db_log_enable"`
	Network                   string `yaml:"network"`
	ListenAddress             string `yaml:"listen_address"`
	ListenPort                string `yaml:"listen_port"`
	MaxLongConnectionPoolSize int    `yaml:"max_long_connection_pool_size"`
	HeartbeatSecond           int32  `yaml:"heartbeat_second"`
	SentryDsn                 string `yaml:"sentry_dsn"`
}

func init() {
	configFilePath := "../config/config.yaml"
	fmt.Println("[INFO] reading config path:", configFilePath)
	ParseYaml(configFilePath, &Config)
	fmt.Println(Config)
}

func ParseYaml(file string, configRaw interface{}) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}

	err = yaml.Unmarshal(content, configRaw)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
}
