package main

import (
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

type scene struct {
	ufo struct {
		lastX int32
		lastY int32
		t     *sdl.Texture
	}

	bg *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/bg.png")
	if err != nil {
		return nil, err
	}

	ufot, err := img.LoadTexture(r, "res/imgs/ufo_sprite.png")
	if err != nil {
		return nil, err
	}

	s := &scene{bg: bg}
	s.ufo.t = ufot
	return s, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	err := r.Copy(s.bg, nil, nil)
	if err != nil {
		return err
	}

	srcRect := &sdl.Rect{X: s.ufo.lastX * 64, Y: s.ufo.lastY * 64, W: 64, H: 64}
	s.ufo.lastX = (s.ufo.lastX + 1) % 9
	s.ufo.lastY = (s.ufo.lastY + 1) % 6

	err = r.Copy(s.ufo.t,
		srcRect,
		&sdl.Rect{X: 10, Y: 200 - 64/2, W: 64, H: 64},
	)
	if err != nil {
		return err
	}

	r.Present()

	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
