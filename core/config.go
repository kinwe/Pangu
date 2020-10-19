package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"test/global"
)

var config = "config.yml"

func init()  {

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_Serve); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_Serve); err != nil {
		fmt.Println(err)
	}

	fmt.Println(global.GVA_Serve.SqlCommons)
	global.GVA_VP = v
}
