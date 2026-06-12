package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"os"
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette

	nframes = 120
	size    = 100
	delay   = 8
)

var palette = []color.Color{color.White, color.Black}

func getFrame(n int) *image.Paletted {
	rect := image.Rect(0, 0, size, size)
	img := image.NewPaletted(rect, palette)

	amplitude := float64(size) / 4
	centerY := float64(size) / 2

	for i := range size {
		phase := float64(n) * 0.2
		y := centerY + amplitude*math.Sin(
			2*math.Pi*float64(i)/float64(size)+phase,
		)
		img.SetColorIndex(i, int(y), blackIndex)
	}

	return img
}

func render() bytes.Buffer {
	var anim = gif.GIF{LoopCount: nframes}
	var buf bytes.Buffer

	for i := range nframes {
		img := getFrame(i)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(&buf, &anim)
	return buf
}

func save(buff bytes.Buffer) {
	err := os.WriteFile("test.gif", buff.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	buff := render()
	save(buff)
}
