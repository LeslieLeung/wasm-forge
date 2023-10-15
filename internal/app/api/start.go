package api

import (
	"wasm-forge/internal/app/api/route"
	"wasm-forge/internal/config"
)

func Start() {
	r := route.BuildRouter()

	r.Run(":" + config.GetConfig().GetString(config.Port))
}
