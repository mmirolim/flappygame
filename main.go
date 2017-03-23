package main

import (
	"log"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	runtime.LockOSThread() // required by

	w, h := 800, 600
	err := sdl.Init(sdl.INIT_EVERYTHING)
	exitOnErr(err, "sdl init")

	defer sdl.Quit()

	err = ttf.Init()
	exitOnErr(err, "ttf init")

	defer ttf.Quit()

	window, r, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	exitOnErr(err, "create window and renderer")

	defer window.Destroy()

	s, err := newScene(r)
	exitOnErr(err, "new scene err")

	err = drawTitle("Adventure Time", 20, r)
	exitOnErr(err, "drawTitle")

	nextWorldTime := time.Tick(10 * time.Millisecond)
	events := make(chan sdl.Event)
	go func() {
		for {
			events <- sdl.WaitEvent()
		}

	}()

GAME_LOOP:
	for {

		select {
		case e := <-events:
			switch v := e.(type) {
			case *sdl.KeyUpEvent:
				s.ufo.lift()
			case *sdl.QuitEvent:
				break GAME_LOOP
			default:
				// fmt.Printf("EVENT %T\n", v) // output for debug

			}
		case <-nextWorldTime:
			if err := s.paint(r); err != nil {
				log.Fatal("paint error", err)
			}
		}

	}

	s.destroy()
	log.Println("bye bye exiting")
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
