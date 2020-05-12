package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Fourier struct {
	period       float64
	t1, t2       float64
	numHarmonics int
	f            func(float64) float64
	an           []float64
	bn           []float64
}

const steps = 1000

// NewFourier initializes and creates a new Fourier for a given function and parameters.
func NewFourier(f func(float64) float64, t1, t2 float64, numHarmonics int) *Fourier {
	fourier := new(Fourier)
	fourier.period = t2 - t1
	fourier.t1 = t1
	fourier.t2 = t2
	fourier.numHarmonics = numHarmonics
	fourier.f = f

	fourier.run()

	return fourier
}

func (f *Fourier) run() {
	f.an = make([]float64, f.numHarmonics)
	f.bn = make([]float64, f.numHarmonics)

	for harmonic := 0; harmonic < f.numHarmonics; harmonic++ {
		f.an[harmonic] = 2.0 / f.period * IntTrapz(
			func(x float64) float64 { return f.f(x) * math.Cos(2.0*math.Pi*x*float64(harmonic)/f.period) },
			f.t1, f.t2, steps)
		f.bn[harmonic] = 2.0 / f.period * IntTrapz(
			func(x float64) float64 { return f.f(x) * math.Sin(2.0*math.Pi*x*float64(harmonic)/f.period) },
			f.t1, f.t2, steps)
	}
}

func (f *Fourier) EvalAt(t float64) float64 {
	var c, cPrime float64
	c = 2.0 * math.Pi * t / f.period
	output := f.an[0] / 2
	for harmonic := 1; harmonic < f.numHarmonics; harmonic++ {
		cPrime = c * float64(harmonic)
		output += f.an[harmonic] * math.Cos(cPrime)
		output += f.bn[harmonic] * math.Sin(cPrime)
	}
	return output
}

func (f *Fourier) CreatePlot(outputName string) {
	// Write the data file.
	file, err := os.Create(outputName + ".dat")
	if err != nil {
		log.Panic(err)
	}

	bufWriter := bufio.NewWriter(file)

	delta := f.period / float64(steps)
	for y := f.t1; y < f.t2; y += delta {
		bufWriter.WriteString(fmt.Sprintf(
			"%f %f %f\n", y, f.EvalAt(y), f.f(y),
		))
	}
	bufWriter.Flush()
	file.Close()

	// Now write the Gnuplot file.
	file, err = os.Create(outputName + ".plt")
	if err != nil {
		log.Panic(err)
	}

	bufWriter = bufio.NewWriter(file)
	bufWriter.WriteString(
		fmt.Sprintf("set title \"Fourier Approximation\"\nplot \"%s\" using 1:2 with lines t \"Fourier Approximation\", '' using 1:3 with lines t \"Original Function\"\npause -1", outputName+".dat"),
	)
	bufWriter.Flush()
	file.Close()

	fmt.Printf("Run \"gnuplot %s\" to see the plot.\n", outputName+".plt")
}

func (f *Fourier) PrintHarmonics() {
	fmt.Printf("n\tan\tbn\n")
	for i := 0; i < f.numHarmonics; i++ {
		fmt.Printf("%d\t% .3f\t% .3f\n", i, f.an[i], f.bn[i])
	}
}
