package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	log.Fatal(run(800, 400))
}

func run(w, h int) error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}

	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		return err
	}
	defer ttf.Quit()

	window, r, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	defer window.Destroy()

	s, err := newScene(r)
	if err != nil {
		return err
	}

	err = drawTitle("Adventure Time", 20, r)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 100; i++ {
		err = s.paint(r)

		if err != nil {
			return err
		}

		time.Sleep(300 * time.Millisecond)
	}

	s.destroy()

	return nil
}

func drawTitle(txt string, size int, r *sdl.Renderer) error {
	f, err := ttf.OpenFont("res/fonts/rpg__.ttf", size)
	if err != nil {
		return err
	}

	s, err := f.RenderUTF8_Solid(txt, sdl.Color{R: 255, B: 0, G: 100, A: 255})
	if err != nil {
		return err
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return err
	}
	defer t.Destroy()

	err = r.Copy(t, nil, nil)
	if err != nil {
		return err
	}

	r.Present()

	return nil
}

func exitOnErr(err error, msg string) {
	if err != nil {
		log.Printf("[E] %v\n", msg)
		log.Fatal(err)
	}
}
