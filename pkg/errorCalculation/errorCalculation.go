package errorCalculation
import(
	"errors"
	"math"
)

func RelativeError(realValue, approxValue float64) (float64, error) {
    if realValue == 0 {
        return 0, errors.New("division by zero: realValue is zero")
    }
    return AbsoluteError(realValue, approxValue) / math.Abs(realValue), nil
}

func AbsoluteError(realValue, approxValue float64) float64 {
    return math.Abs(approxValue - realValue)
}
