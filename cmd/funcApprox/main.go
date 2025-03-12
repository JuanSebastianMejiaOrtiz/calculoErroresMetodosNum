package main
import (
	"math"
	"fmt"
	"github.com/JuanSebastianMejiaOrtiz/errorCalculosMetodosNum/internal/functionApproximation"
)

func main() {
	approxSin := func(x float64) float64 {
		return x - math.Pow(x, 3)/6
	}
	x := []float64{0.1, 0.3, 0.5, 1}

    errorMagnitude := functionApproximation.ApproximationErrorFunction(x, math.Sin, approxSin)
    fmt.Println(errorMagnitude)
}
