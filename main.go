package main

import (
	"log"
	"math"
	"time"
)

func main() {
	// Fourier series of a cosine wave, should be unexciting.
	fSin := NewFourier(math.Cos, 0.0, 2.0*math.Pi, 5)
	fSin.PrintHarmonics()
	fSin.CreatePlot("plot_sin")

	// Find the Fourier series of a square wave.
	tStart := time.Now()
	f := NewFourier(squareWave, 0.0, 1.0, 100)
	tElapsed := time.Now().Sub(tStart)
	log.Printf("Took %v to approximate the Fourier series.\n", tElapsed)
	f.CreatePlot("plot")
}

func squareWave(t float64) float64 {
	if math.Mod(t, 1.0) < 0.5 {
		return 0
	}
	return 1
}

func triangleWave(t float64) float64 {
	if math.Mod(t, 1.0) < 0.5 {
		return 2 * t
	}
	return 2 - 2*t
}
