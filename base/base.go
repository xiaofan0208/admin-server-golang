package base

import (
	"admin-server-golang/base/config"
	"github.com/Unknwon/goconfig"
)

var (
	AppConfig *BaseAppConfig
	ConfigFile *goconfig.ConfigFile
)


func init(){
	var err error
	ConfigFile, err = goconfig.LoadConfigFile("conf/app.conf")
	if err != nil{
		panic("配置读取错误")
	}
}


type BaseAppConfig struct {
	innerConfig config.Configer
}

func(cfg *BaseAppConfig)  String(section , key string) string {
	value, err := ConfigFile.GetValue(section,key)
	if err != nil{
		panic("配置读取错误")
	}
	return value
}
func(cfg *BaseAppConfig)  Int( section , key string ) int {
	value, err := ConfigFile.Int(section,key)
	if err != nil{
		panic("配置读取错误")
	}
	return value
}

func(cfg *BaseAppConfig)  Int64( section , key string ) int64 {
	value, err := ConfigFile.Int64(section,key)
	if err != nil{
		panic("配置读取错误")
	}
	return value
}




