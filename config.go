package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string, file_type string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(file_type); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig(file_type string) error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	if (file_type == "") {
		file_type = "yaml"
	}
	viper.SetConfigType(file_type)   
	viper.AutomaticEnv()     
	viper.SetEnvPrefix("APISERVER") 
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s", e.Name)
	})
}

func Get(key string) string {
	return viper.GetString(key)
}

func GetJson(key string) interface{} {
	return viper.GetStringMap(key)
}

func GetRaw(key string) interface{} {
	return viper.Get(key)
}

func GetFloat(key string) float64 {
	return viper.GetFloat64(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
