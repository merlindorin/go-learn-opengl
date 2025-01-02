package utils

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"os"
)

type shaderType uint32

var (
	ShaderTypeString = map[shaderType]string{
		FRAGMENT_SHADER: "FragmentShaderType",
		VERTEX_SHADER:   "VertexShaderType",
	}
)

func (s shaderType) String() string {
	return ShaderTypeString[s]
}

const (
	FRAGMENT_SHADER shaderType = gl.FRAGMENT_SHADER
	VERTEX_SHADER   shaderType = gl.VERTEX_SHADER
)

type Shader struct {
	handle     uint32
	shaderType shaderType
}

func (s *Shader) Handle() uint32 {
	return s.handle
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.handle)
}

func NewShader(src string, t shaderType) (*Shader, error) {
	handle := gl.CreateShader(uint32(t))
	glSrc, freeFn := gl.Strs(src + "\x00")
	defer freeFn()

	gl.ShaderSource(handle, 1, glSrc, nil)
	gl.CompileShader(handle)

	if err := CheckShaderCompileErrors(handle, t.String()); err != nil {
		return nil, fmt.Errorf("compile shader %s failure: %v", t.String(), err)
	}

	return &Shader{handle: handle, shaderType: t}, nil
}

func NewShaderFromFile(file string, t shaderType) (*Shader, error) {
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s: %v", file, err)
	}

	return NewShader(string(src), t)
}
