package compiler

import (
	"os"
	"wasm-forge/internal/command"
)

var _ Compiler = (*TinyGoCompiler)(nil)

type TinyGoCompiler struct {
}

func (t TinyGoCompiler) Version() string {
	version, err := command.RunCmd("tinygo", "version")
	if err != nil {
		return ""
	}
	return string(version)
}

func (t TinyGoCompiler) Compile(path string, args ...string) ([]byte, error) {
	currentDir, _ := os.Getwd()
	return command.RunCmd("tinygo", "build", "-o", currentDir+"/runtime/main.wasm", "-target", "wasm", currentDir+path)
}
