package functionApproximation
import (
    "fmt"
	"github.com/JuanSebastianMejiaOrtiz/errorCalculosMetodosNum/pkg/errorCalculation"
)

func ApproximationErrorFunction(x []float64, realFunc func(float64) float64, approxFunc func(float64) float64) []float64 {
    approxVals := make([]float64, len(x))
    realVals := make([]float64, len(x))
    
    for i, value := range x {
        approxVals[i] = approxFunc(value)
        realVals[i] = realFunc(value)
    }
    
    errorMagnitude := make([]float64, len(x))
    for i := range errorMagnitude {
        relError, err := errorCalculation.RelativeError(realVals[i], approxVals[i])
        
        if err != nil {
            fmt.Printf("Error calculating relative error for x = %v : %v\n", x[i], err)
            errorMagnitude[i] = -1
            continue
        }
        
        errorMagnitude[i] = relError
    }
    return errorMagnitude
}

