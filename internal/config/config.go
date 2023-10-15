package config

import "github.com/spf13/viper"

var vp *viper.Viper

func init() {
	vp = viper.New()
	vp.SetDefault(RunMode, RunModeBin)
	vp.SetDefault(Port, 9000)
	vp.SetEnvPrefix("wasmforge")
	vp.BindEnv(RunMode)
	vp.BindEnv(Port)

	vp.BindEnv("TINYGOROOT")
}

func GetConfig() *viper.Viper {
	return vp
}
