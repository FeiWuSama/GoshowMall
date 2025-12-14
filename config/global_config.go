package config

import (
	"flag"
	"github.com/spf13/viper"
	"strconv"
)

var (
	configPath   string
	GlobalConfig Config
)

type Config struct {
	Server      Server              `yaml:"server"`
	MySql       MySql               `yaml:"mysql"`
	Redis       Redis               `yaml:"redis"`
	LarkGroupID string              `yaml:"group_id" mapstructure:"group_id"`
	AppConfig   map[int32]AppConfig `yaml:"app_config"`
}

type Server struct {
	Port        int    `yaml:"port"`
	EnablePprof bool   `yaml:"enable_pprof"`
	LogLevel    string `yaml:"log_level"`
}

type MySql struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Charset     string `yaml:"charset"`
	ShowSql     bool   `yaml:"show_sql"`
	MaxOpenConn int    `yaml:"max_open_conn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

type AppConfig struct {
	AppType   string `yaml:"app_type"`
	AppName   string `yaml:"app_name"`
	AppId     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

func init() {
	flag.StringVar(&configPath, "c", "application.yml", "default config path")
}

func InitConfig() Config {
	var (
		err        error
		tempConfig = &Config{}
		vipConfig  = viper.New()
	)
	flag.Parse()
	tempConfig, err = getConfig(vipConfig)
	if err != nil {
		panic(err)
	}
	return *tempConfig
}

func getConfig(vipConfig *viper.Viper) (*Config, error) {
	vipConfig.SetConfigFile(configPath)
	if err := vipConfig.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vipConfig.Unmarshal(&GlobalConfig); err != nil {
		return nil, err
	}
	// 获取所有设置的原始数据
	allSettings := vipConfig.AllSettings()

	// 处理 app_config
	GlobalConfig.AppConfig = make(map[int32]AppConfig)

	if appConfigRaw, ok := allSettings["app_config"].(map[string]interface{}); ok {
		for appIdStr, appDataRaw := range appConfigRaw {
			// 转换 appId 为 int32
			appIdInt, err := strconv.ParseInt(appIdStr, 10, 32)
			if err != nil {
				// 如果不是数字，跳过
				continue
			}
			appId := int32(appIdInt)

			// 解析单个 AppConfig
			if appData, ok := appDataRaw.(map[string]interface{}); ok {
				appConfig := AppConfig{}

				// 使用类型断言获取值
				if val, ok := appData["app_type"].(string); ok {
					appConfig.AppType = val
				}
				if val, ok := appData["app_name"].(string); ok {
					appConfig.AppName = val
				}
				if val, ok := appData["app_id"].(string); ok {
					appConfig.AppId = val
				}
				if val, ok := appData["app_secret"].(string); ok {
					appConfig.AppSecret = val
				}

				GlobalConfig.AppConfig[appId] = appConfig
			}
		}
	}

	return &GlobalConfig, nil
}
