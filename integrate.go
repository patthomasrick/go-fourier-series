package main

// IntTrapz numerically integrates the function function f from t1 to t2 using the trapezoid rule.
// Delta is the width of a trapezoid.
func IntTrapz(f func(float64) float64, t1, t2 float64, steps int) float64 {
	var area, delta, y1, y2 float64
	delta = (t2 - t1) / float64(steps)
	for t := t1; t < t2-delta; t += delta {
		y1 = f(t)
		y2 = f(t + delta)
		area += (y1 + 0.5*(y2-y1)) * delta
	}
	return area
}
