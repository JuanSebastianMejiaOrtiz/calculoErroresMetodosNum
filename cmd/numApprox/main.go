package main
import (
	"math"
	"fmt"
	"github.com/JuanSebastianMejiaOrtiz/errorCalculosMetodosNum/internal/numberApproximated"
)

func main() {
    realNum := math.Pi
    approxNums := []float64{3.0, 3.14, 22/7}
    errorMagnitude, err := numberApproximated.ErrorInNumber(realNum, approxNums)
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    fmt.Println(errorMagnitude)
}
