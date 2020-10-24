package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
    "fmt"
)

var (
	mouseWin geom.Vec2
)

func setup(w *gfx.Win) error {
	w.GetGlfwWindow().SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
    w.GetGlfwWindow().SetCursorPos(0, 0)
    spawnBoids()
    fmt.Println("running")
	return nil
}

func draw(w *gfx.WinDraw) {
    updateBoids()
    playerUpdate()

	size := w.GetFrameSize()
	ar := size.X / size.Y
	near := float32(0.1)
	perspective := geom.Mat4Perspective(-ar*near, ar*near, -near, near, near, 1000).Product(geom.Mat4Scalar(-1, 1, -1))
	translation := geom.Mat4Translation(player.position.ScaledBy(-1))
	rx := geom.Mat4RotationX(-player.pitch)
	ry := geom.Mat4RotationY(-player.bearing)
	view := perspective.Product(rx).Product(ry).Product(translation)

    drawBoids(w, view)
}

func mouse(w *gfx.Win, ev gfx.MouseEvent) {
	switch e := ev.(type) {
	case gfx.MouseScroll:
	case gfx.MouseMove:
		{
            playerLook(e.Position.X, -e.Position.Y)
			w.GetGlfwWindow().SetCursorPos(0, 0)
		}
	case gfx.MouseButton:
	}
}

func keyboard(w *gfx.Win, ev gfx.KeyEvent) {
	switch ev.Key {
	case glfw.KeyW:
		{
			if ev.Action == glfw.Press {
				keys.w = true
			} else if ev.Action == glfw.Release {
				keys.w = false
			}
		}

	case glfw.KeyA:
		{
			if ev.Action == glfw.Press {
				keys.a = true
			} else if ev.Action == glfw.Release {
				keys.a = false
			}
		}

	case glfw.KeyS:
		{
			if ev.Action == glfw.Press {
				keys.s = true
			} else if ev.Action == glfw.Release {
				keys.s = false
			}
		}

	case glfw.KeyD:
		{
			if ev.Action == glfw.Press {
				keys.d = true
			} else if ev.Action == glfw.Release {
				keys.d = false
			}
		}
	}
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		DrawFunc:  draw,
		MouseFunc: mouse,
		SetupFunc: setup,
		KeyFunc:   keyboard,
		Title:     "Boids",
        Width:     1024,
        Height:    768,
	})
}
