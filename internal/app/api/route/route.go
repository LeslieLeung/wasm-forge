package route

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	compiler2 "wasm-forge/internal/compiler"
)

func BuildRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/go/:version/build", func(c *gin.Context) {
		code := c.PostForm("code")
		// version := c.Param("version")
		// save to file
		os.WriteFile("runtime/main.go", []byte(code), 0644)
		// compile
		compiler := compiler2.TinyGoCompiler{}
		output, err := compiler.Compile("/runtime/main.go")
		if err != nil {
			c.String(500, "compile error: %v, output: %v", err, string(output))
			return
		}
		// return wasm
		f, err := os.Open("runtime/main.wasm")
		if err != nil {
			c.String(500, "open wasm error: %v", err)
			return
		}
		defer f.Close()
		// write file to response
		fBytes, _ := io.ReadAll(f)
		c.Data(200, "application/wasm", fBytes)
	})

	return r
}
