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
	title  = "LearnOpenGL: hello_triangle_exercise2"
)

var logger *zap.Logger

func init() {
	// This is needed to arrange that main() runs on the main thread
	// See documentation for functions that are only allowed to be called from the main thread
	runtime.LockOSThread()
}

func main() {
	// Initialize zap logger
	logger = zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		logger.Fatal("Failed to initialize GLFW", zap.Error(err))
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
		logger.Fatal("Failed to create window", zap.Error(err))
	}
	window.MakeContextCurrent()

	// Initialize GLAD
	if err := gl.Init(); err != nil {
		logger.Fatal("Failed to initialize GLAD", zap.Error(err))
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
	firstTriangle := []float32{
		-0.9, -0.5, 0.0, // left
		-0.0, -0.5, 0.0, // right
		-0.45, 0.5, 0.0, // top
	}
	secondTriangle := []float32{
		0.0, -0.5, 0.0, // left
		0.9, -0.5, 0.0, // right
		0.45, 0.5, 0.0, // top
	}

	var VBOs, VAOs [2]uint32
	gl.GenVertexArrays(2, &VAOs[0])
	gl.GenBuffers(2, &VBOs[0])

	// First triangle setup
	gl.BindVertexArray(VAOs[0])
	gl.BindBuffer(gl.ARRAY_BUFFER, VBOs[0])
	gl.BufferData(gl.ARRAY_BUFFER, len(firstTriangle)*4, gl.Ptr(firstTriangle), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	// Second triangle setup
	gl.BindVertexArray(VAOs[1])
	gl.BindBuffer(gl.ARRAY_BUFFER, VBOs[1])
	gl.BufferData(gl.ARRAY_BUFFER, len(secondTriangle)*4, gl.Ptr(secondTriangle), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(0)

	// Render loop
	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shaderProgram.Use()
		// Draw first triangle using the data from the first VAO
		gl.BindVertexArray(VAOs[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		// Draw the second triangle using the data from the second VAO
		gl.BindVertexArray(VAOs[1])
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	// De-allocate all resources once they've outlived their purpose
	gl.DeleteVertexArrays(2, &VAOs[0])
	gl.DeleteBuffers(2, &VBOs[0])
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func framebufferSizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
