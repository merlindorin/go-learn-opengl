package main

import (
	"learn-opengl/utils"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go.uber.org/zap"
)

const (
	width  = 800
	height = 600
	title  = "LearnOpenGL: hello_triangle_exercise1"
)

var logger *zap.Logger

func init() {
	// This is needed to arrange that main() runs on the main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	// Initialize zap logger
	logger = zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		logger.Fatal("failed to initialize glfw", zap.Error(err))
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// Darwin-specific hint
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	// Create window
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		logger.Fatal("failed to create window", zap.Error(err))
	}
	window.MakeContextCurrent()

	// Initialize GLAD
	if err := gl.Init(); err != nil {
		logger.Fatal("failed to initialize glad", zap.Error(err))
	}

	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	vertexShader, err := utils.NewShaderFromFile("vertex_shader.glsl", utils.VERTEX_SHADER)
	if err != nil {
		logger.Fatal("failed to create shader", zap.Error(err))
	}

	fragmentShader, err := utils.NewShaderFromFile("fragment_shader.glsl", utils.FRAGMENT_SHADER)
	if err != nil {
		logger.Fatal("failed to create shader", zap.Error(err))
	}

	shaderProgram, err := utils.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		logger.Fatal("failed to create program", zap.Error(err))
	}
	defer shaderProgram.Delete()

	// Set up vertex data (and buffer(s)) and configure vertex attributes
	vertices := []float32{
		// First triangle
		-0.9, -0.5, 0.0,
		-0.0, -0.5, 0.0,
		-0.45, 0.5, 0.0,
		// Second triangle
		0.0, -0.5, 0.0,
		0.9, -0.5, 0.0,
		0.45, 0.5, 0.0,
	}

	var VBO, VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)

	gl.BindVertexArray(VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0) // Unbind the VBO

	gl.BindVertexArray(0) // Unbind the VAO

	// Render loop
	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shaderProgram.Use()
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	// De-allocate all resources once they've outlived their purpose
	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func framebufferSizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
