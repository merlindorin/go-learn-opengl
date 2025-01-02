package main

import (
	"go.uber.org/zap"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	title  = "LearnOpenGL: 1.1.hello_window"
	width  = 800
	height = 600
)

var logger *zap.Logger

func init() {
	// This is needed to arrange that main() runs on main thread.
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
		logger.Fatal("Failed to create window", zap.Error(err))
	}

	// It must be called before gl.Init
	window.MakeContextCurrent()

	// Initialize GLAD
	if err := gl.Init(); err != nil {
		logger.Fatal("Failed to init glad", zap.Error(err))
	}

	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	// Render loop
	for !window.ShouldClose() {
		// Process input
		processInput(window)

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
