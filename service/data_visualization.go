package service

import (
	"IoTGateWay/consts"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palatte = []color.Color{color.White, color.Black,color.RGBA{0,255,0,255},color.RGBA{255,0,0,255},color.RGBA{0,0,255,255},color.RGBA{255,255,0,255}}
type Putpixel func(x, y int)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palatte)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), consts.BlackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func drawResult(out io.Writer,detectionID int) {
	anim := gif.GIF{LoopCount: consts.Nframes}
	phase := 0.0
	data ,err := DetResultSer.GetResultByDetectionID(detectionID)
	if err != nil {
		Logger.Error("drawResult call DetResultSer.GetResultByDetectionID error:%v",err)
		return
	}
	lenthX := float64(len(data))
	for i := 0; i < consts.Nframes; i++ {
		rect := image.Rect(0, 0, 4*consts.Size+1, 2*consts.Size+1)
		img := image.NewPaletted(rect, palatte)
		xtemp := 0.0
		ytemp := float64(data[0].ResultMark)
		for t := 0.0; int(t) < (i+1)*len(data)/consts.Nframes; t ++ {
			k := data[int(t)]
			colorIndex := consts.BlackIndex
			if k.ResultMark == 1 {
				colorIndex = consts.GreenIndex
			}else if k.ResultMark == 2{
				colorIndex = consts.RedIndex
				Logger.Info("i:%v,t:%v,xtemp:%v,ytemp:%v,x:%v,y:%v",i,t,int(xtemp/lenthX*4*consts.Size), 2*consts.Size-int(ytemp/2*consts.Size), int(float64(t)/lenthX*4*consts.Size), 2*consts.Size-int(float64(k.ResultMark)/2*consts.Size))
				Logger.Info("i:%v,t:%v,xtemp:%v,ytemp:%v,x:%v,y:%v",i,t,xtemp/lenthX*4*consts.Size, ytemp/2*consts.Size, float64(t)/lenthX*4*consts.Size, float64(k.ResultMark)/2*consts.Size)
			}else {
				colorIndex = consts.YellowIndex
			}
			drawline(int(xtemp/lenthX*4*consts.Size), 2*consts.Size-int(ytemp/2*consts.Size), int(float64(t)/lenthX*4*consts.Size), 2*consts.Size-int(float64(k.ResultMark)/2*consts.Size), func(x, y int) {
				img.SetColorIndex(x, y, uint8(colorIndex))
			})
			xtemp = float64(t)
			ytemp = float64(k.ResultMark)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, consts.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}