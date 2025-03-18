package str

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/shopspring/decimal"

	// "fmt"
	// "math"
	"fmt"
	"strconv"
)

func AddTwoFloat(f1, f2 float64) float64 {
	df1 := decimal.NewFromFloat(f1)
	df2 := decimal.NewFromFloat(f2)
	res := df1.Add(df2)
	r1, _ := res.Float64()
	return Float2Float(r1)
}

func SubTwoFloat(f1, f2 float64) float64 {
	df1 := decimal.NewFromFloat(f1)
	df2 := decimal.NewFromFloat(f2)
	res := df1.Sub(df2)
	r1, _ := res.Float64()
	return Float2Float(r1)
}

func DivTwoFloat(f1, f2 float64) float64 {
	bigF1 := new(big.Float).SetFloat64(f1)
	bigF2 := new(big.Float).SetFloat64(f2)
	mul := new(big.Float).Quo(bigF1, bigF2)
	r1, _ := mul.Float64()
	return Float2Float(r1)
}

func MulTwoFloat(f1, f2 float64) float64 {
	bigF1 := new(big.Float).SetFloat64(f1)
	bigF2 := new(big.Float).SetFloat64(f2)
	mul := new(big.Float).Mul(bigF1, bigF2)
	r1, _ := mul.Float64()
	return Float2Float(r1)
}

func Float2Float(num float64) float64 {
	floatNum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return floatNum
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func MakePriceTotal(coins, perCoins int, price float64) (total float64) {
	total = DivTwoFloat(MulTwoFloat(float64(coins), price), float64(perCoins))
	return total
}

func MakePriceTotalFor64(coins, perCoins int64, price float64) (total float64) {
	total = DivTwoFloat(MulTwoFloat(float64(coins), price), float64(perCoins))
	return total
}

func MulTwoFloatWithFloor(f1, f2 float64) float64 {
	bigF1 := new(big.Float).SetFloat64(f1)
	bigF2 := new(big.Float).SetFloat64(f2)
	mul := new(big.Float).Mul(bigF1, bigF2)
	r1, _ := mul.Float64()
	return RoundFloor(r1)
}

// 不四舍五入保留2位小数
func RoundFloor(f float64) float64 {
	priceRoundFloor, _ := decimal.NewFromFloat(f).RoundFloor(4).Float64()
	return priceRoundFloor
}

func GetRandomNumbers(m, n int) ([]int, error) {
	if n > m {
		return nil, fmt.Errorf("n 不能大于 m")
	}
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, m)
	for i := range numbers {
		numbers[i] = i
	}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return numbers[:n], nil
}
