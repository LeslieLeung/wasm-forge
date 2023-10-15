package compiler

import "testing"

func TestTinyGoCompiler_Version(t1 *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test",
			want: "tinygo version 0.30.0 darwin/amd64 (using go version go1.21.0 and LLVM version 16.0.1)\n",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TinyGoCompiler{}
			if got := t.Version(); got != tt.want {
				t1.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}
