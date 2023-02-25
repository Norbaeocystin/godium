package godium

import (
	"errors"
	"math"
)

// Series is a container for a series of data
type Series []Coordinate

// Coordinate holds the data in a series
type Coordinate struct {
	X, Y float64
}

// calculate gradient, intercept and angle
func LinearRegression(s Series) (float64, float64, float64, error) {

	if len(s) == 0 {
		return 0, 0, 0, errors.New("empty series")
	}

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)
	angle := math.Atan(gradient) * (180.0 / math.Pi)

	return gradient, intercept, angle, nil
}
