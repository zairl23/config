package config 

import (
	"github.com/spf13/viper"
)

// config manager
type Configer interface {
	Add(name string, cfg string) error
	Get(name string) *viper.Viper
	Init(cfg string, v *viper.Viper) error
}