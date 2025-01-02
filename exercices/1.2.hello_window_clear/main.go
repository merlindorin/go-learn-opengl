package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go.uber.org/zap"
)

const (
	width  = 800
	height = 600
	title  = "LearnOpenGL: hello_window_clear"
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

	// Render loop
	for !window.ShouldClose() {
		// Process input
		processInput(window)

		// Render
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Swap buffers and poll IO events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func framebufferSizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
