package utils

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func CheckShaderCompileErrors(shader uint32, shaderType string) error {
	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)

	if success == gl.TRUE {
		return nil
	}

	var logLength int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

	infoLog := make([]byte, logLength)
	gl.GetShaderInfoLog(shader, logLength, nil, &infoLog[0])

	return fmt.Errorf("shader of type `%v` compilation failed: %v", shaderType, string(infoLog))
}

func CheckProgramLinkErrors(program uint32) error {
	var success int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)

	if success == gl.TRUE {
		return nil
	}

	var logLength int32
	gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

	infoLog := make([]byte, logLength)
	gl.GetProgramInfoLog(program, logLength, nil, &infoLog[0])

	return fmt.Errorf("program linking failed: %v", string(infoLog))
}
