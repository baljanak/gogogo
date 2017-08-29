package sqrt

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot find real square root for: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return float64(math.NaN()), ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func RealSqrt(x float64) float64 {
	v, err := Sqrt(x)
	if err != nil {
		panic(err)
	}
	return v
}
