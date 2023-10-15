package compiler

var _ Compiler = (*CCompiler)(nil)

type CCompiler struct {
}

func (C CCompiler) Compile(path string, args ...string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (C CCompiler) Version() string {
	// TODO implement me
	panic("implement me")
}
