package compiler

type Compiler interface {
	Version() string
	Compile(path string, args ...string) ([]byte, error)
}
