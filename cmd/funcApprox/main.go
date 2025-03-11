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
	x := []float64{1, 2, 3, 4, 5}

    errorMagnitude, err := functionApproximation.ApproximationErrorFunction(x, math.Sin, approxSin)
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    fmt.Println(errorMagnitude)
}
