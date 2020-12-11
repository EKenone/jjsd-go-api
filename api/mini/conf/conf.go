package conf

import (
	"flag"
	"github.com/spf13/viper"
	"sync"
)

var (
	confPath string
	Conf     *Config
	viperIns *viper.Viper
	newOnce  = sync.Once{}
)

func init() {
	flag.StringVar(&confPath, "conf", "mini-example.yaml", "default config path.")
}

func Init() (err error) {
	Conf = new(Config)
	err = GetViper().Unmarshal(&Conf)
	return
}

func GetViper() *viper.Viper {
	newOnce.Do(func() {
		viperIns = newViper()
	})

	return viperIns
}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(confPath)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("无法找到配置文件")
		} else {
			panic(err)
		}
	}
	return v
}

type Config struct {
	Wechat *WeChatConf
	Mysql  *MysqlConf
	Redis  *RedisConf
}

type WeChatConf struct {
	Appid  string `yaml:"appid"`
	Secret string `yaml:"secret"`
}

type MysqlConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type RedisConf struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Network string `yaml:"network"`
}
