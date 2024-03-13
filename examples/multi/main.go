package main

import (
	"fmt"
	imgui "github.com/damntourists/cimgui-go-lite"
	"github.com/damntourists/cimgui-go-themes"
)

var (
	backend imgui.Backend[imgui.GLFWWindowFlags]

	theme1 *imthemes.Theme
	theme2 *imthemes.Theme
	theme3 *imthemes.Theme
)

func loop() {
	theme1Fin := theme1.Apply()
	imgui.Begin("Future Dark")
	imgui.Text("hello world!")
	imgui.End()
	theme1Fin()

	theme2Fin := theme2.Apply()
	imgui.Begin("Light theme")
	imgui.Text("hello world!")
	imgui.End()
	theme2Fin()

	theme3Fin := theme3.Apply()
	imgui.ShowDemoWindow()
	theme3Fin()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	t1, err := imthemes.GetThemeByName("Future Dark")
	panicOnError(err)
	t2, err := imthemes.GetThemeByName("Light")
	panicOnError(err)
	t3, err := imthemes.GetThemeByName("Soft Cherry")
	panicOnError(err)

	theme1 = t1
	theme2 = t2
	theme3 = t3

	backend, _ = imgui.CreateBackend(imgui.NewGLFWBackend())

	backend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	backend.CreateWindow("Hello from cimgui-go", 1200, 900)

	backend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	backend.SetCloseCallback(func(b imgui.Backend[imgui.GLFWWindowFlags]) {
		fmt.Println("window is closing")
	})

	backend.Run(loop)
}
