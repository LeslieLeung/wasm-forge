package compiler

var _ Compiler = (*GoCompiler)(nil)

type GoCompiler struct {
}

func (g GoCompiler) Compile(path string, args ...string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (g GoCompiler) Version() string {
	// TODO implement me
	panic("implement me")
}
