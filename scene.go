package main

import (
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

const (
	vAccelation = 2.25
)

type scene struct {
	ufo *ufo
	bg  *sdl.Texture
}

type ufo struct {
	time   int
	tX, tY int32
	x, y   int32
	W, H   int32
	t      *sdl.Texture
	vSpeed int
}

func (u *ufo) lift() {
	u.y -= 100
}

func (u *ufo) posX(hAcc float32) int32 {
	u.x = int32(float32(u.x) + hAcc)
	if u.x < 0 {
		u.x = 0
	}
	return u.x
}

func (u *ufo) posY(vAcc float32) int32 {
	u.y = int32(float32(u.y) + vAcc)
	if u.y <= 0 {
		u.y = 0
	}
	return u.y
}

func newUfo(r *sdl.Renderer, sprite string, w, h int32) (*ufo, error) {
	var err error
	u := new(ufo)
	u.t, err = img.LoadTexture(r, sprite)
	if err != nil {
		return nil, err
	}
	u.W, u.H = w, h
	return u, nil

}

func (u *ufo) paint(r *sdl.Renderer) error {
	u.time++
	rectInSprite := &sdl.Rect{
		X: u.tX * u.W,
		Y: u.tY * u.H,
		W: u.W,
		H: u.H,
	}

	err := r.Copy(u.t,
		rectInSprite,
		&sdl.Rect{X: u.posX(0.0), Y: u.posY(vAccelation), W: u.W, H: u.H},
	)
	if err != nil {
		return err
	}

	if u.time%10 == 0 {
		u.time = 0
		u.tX = (u.tX + 1) % 9
		u.tY = (u.tY + 1) % 6

	}

	return nil
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/bg.png")
	if err != nil {
		return nil, err
	}
	ufo, err := newUfo(r, "res/imgs/ufo_sprite.png", 64, 64)
	if err != nil {
		return nil, err
	}

	s := &scene{bg: bg, ufo: ufo}
	return s, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	err := r.Copy(s.bg, nil, nil)
	if err != nil {
		return err
	}
	err = s.ufo.paint(r)
	if err != nil {
		return err
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.ufo.t.Destroy()
}
