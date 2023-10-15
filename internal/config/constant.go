package config

const (
	RunMode = "run_mode"
	Port    = "port"
)

type RunModeType string

const (
	RunModeAPI RunModeType = "api"
	RunModeBin RunModeType = "bin"
)
