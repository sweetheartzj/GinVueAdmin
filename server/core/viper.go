package core

import (
	"GinVueAdmin/server/core/internal"
	"GinVueAdmin/server/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func parseCommandArgs() string {
	var config string
	flag.StringVar(&config, "c", "", "choose config file")
	flag.Parse()
	return config
}

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) != 0 {
		config = path[0]
	} else if config = parseCommandArgs(); config == "" {
		if configEnv := os.Getenv(internal.ConfigEnv); configEnv != "" {
			switch gin.Mode() {
			case gin.DebugMode:
				config = internal.ConfigDebugFile
			case gin.TestMode:
				config = internal.ConfigTestFile
			case gin.ReleaseMode:
				config = internal.ConfigReleaseFile
			default:
				config = internal.ConfigDefaultFile
			}
		} else {
			config = configEnv
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	_ = v.Unmarshal(&global.GVA_CONFIG)

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.Unmarshal(&global.GVA_CONFIG)
	})

	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
