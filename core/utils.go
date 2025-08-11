package coherra

import "math"

// Normalize helper for QRE and mesh score
func Normalize(x float64, min float64, max float64) float64 {
	return math.Max(0.1, math.Min(1.0, (x-min)/(max-min)))
}
