package utils

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	handle  uint32
	shaders []*Shader
}

func (prog *Program) Delete() {
	for _, shader := range prog.shaders {
		shader.Delete()
	}
	gl.DeleteProgram(prog.handle)
}

func (prog *Program) Attach(shaders ...*Shader) {
	for _, shader := range shaders {
		gl.AttachShader(prog.handle, shader.handle)
		prog.shaders = append(prog.shaders, shader)
	}
}

func (prog *Program) Use() {
	gl.UseProgram(prog.handle)
}

func (prog *Program) Link() error {
	gl.LinkProgram(prog.handle)
	err := CheckProgramLinkErrors(prog.handle)
	if err != nil {
		return fmt.Errorf("program link error: %v", err)
	}

	return nil
}

func NewProgram(shaders ...*Shader) (*Program, error) {
	prog := &Program{handle: gl.CreateProgram()}
	prog.Attach(shaders...)

	if err := prog.Link(); err != nil {
		return nil, err
	}

	return prog, nil
}
