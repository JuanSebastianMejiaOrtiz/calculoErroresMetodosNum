package numberApproximated
import (
    "fmt"
    "errors"
	"github.com/JuanSebastianMejiaOrtiz/errorCalculosMetodosNum/pkg/errorCalculation"
)

func ErrorInNumber(realNum float64, approxNums []float64) ([]float64, error) {
    if realNum == 0 {
        return []float64{}, errors.New("Invalid value, realNum has to be different from 0")
    }
    errorMagnitude := make([]float64, len(approxNums))
    for i, approximation := range approxNums {
        relErr, err := errorCalculation.RelativeError(realNum, approximation)
        
        if err != nil {
            fmt.Printf("Error calculating relative error for x = %v : %v\n", approxNums[i], err)
            errorMagnitude[i] = -1
            continue
        }
        
        errorMagnitude[i] = relErr
    }
    return errorMagnitude, nil
}
