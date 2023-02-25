package raydium

import (
	"encoding/binary"
	"errors"
	bin "github.com/gagliardetto/binary"
	"log"
	"math"
	"math/big"
)

var Q64BI = new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
var Q64 = new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil))

func CalculateLiquidity(tokenAMax, tokenBMax uint64, tick, upperTick, lowerTick int64) *big.Int {
	low := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(lowerTick))))
	current := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(tick))))
	upper := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(upperTick))))
	var diff0 *big.Int
	if current.Cmp(upper) == 1 {
		diff0 = new(big.Int).Sub(current, upper)
	} else {
		diff0 = new(big.Int).Sub(upper, current)
	}
	var diff1 *big.Int
	if current.Cmp(low) == 1 {
		diff1 = new(big.Int).Sub(current, low)
	} else {
		diff1 = new(big.Int).Sub(low, current)
	}
	liq0 := new(big.Int).Div(new(big.Int).Div(new(big.Int).Mul(big.NewInt(int64(tokenAMax)), new(big.Int).Mul(current, upper)), Q64BI), diff0)
	liq1 := new(big.Int).Div(new(big.Int).Mul(big.NewInt(int64(tokenBMax)), Q64BI), diff1)
	if liq0.Cmp(liq1) == -1 {
		return liq0
	}
	return liq1
}

/*
def calc_amount0(liq, pa, pb):
    if pa > pb:
        pa, pb = pb, pa
    return int(liq * q96 * (pb - pa) / pa / pb)


def calc_amount1(liq, pa, pb):
    if pa > pb:
        pa, pb = pb, pa
    return int(liq * (pb - pa) / q96)

amount0 = calc_amount0(liq, sqrtp_upp, sqrtp_cur)
amount1 = calc_amount1(liq, sqrtp_low, sqrtp_cur)
(amount0, amount1)
> (998976618347425408, 5000000000000000000000)

*/

// inverse function to CalculateLiquidity
func CalculateAmounts(liquidity *big.Int, lowerTick, upperTick, tick int64) (*big.Int, *big.Int) {
	low := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(lowerTick))))
	current := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(tick))))
	upper := CalculateSqrtPriceQ64(big.NewFloat(math.Pow(1.0001, float64(upperTick))))
	var diff0 *big.Int
	if current.Cmp(upper) == 1 {
		diff0 = new(big.Int).Sub(current, upper)
	} else {
		diff0 = new(big.Int).Sub(upper, current)
	}
	var diff1 *big.Int
	if current.Cmp(low) == 1 {
		diff1 = new(big.Int).Sub(current, low)
	} else {
		diff1 = new(big.Int).Sub(low, current)
	}
	l0 := new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(liquidity, diff0), Q64BI), new(big.Int).Mul(current, upper))
	l1 := new(big.Int).Div(new(big.Int).Mul(liquidity, diff1), Q64BI)
	return l0, l1
}

// calculate squared root price with Q96
// example of input 0.86 matic/usdc what is 0.86e-12
func CalculateSqrtPriceQ64(price *big.Float) *big.Int {
	priceSqrt := new(big.Float).Sqrt(price)
	sqrtQ64Price := new(big.Float).Mul(priceSqrt, Q64)
	result := new(big.Int)
	sqrtQ64Price.Int(result)
	return result
}

func CalculatePriceFromSQRPriceQ64(priceSQRTQ64 *big.Int) float64 {
	floatPriceSQRTQ64 := new(big.Float).SetInt(priceSQRTQ64)
	floatPriceSQRT := new(big.Float).Quo(floatPriceSQRTQ64, Q64)
	floatPrice := new(big.Float).Mul(floatPriceSQRT, floatPriceSQRT)
	price, _ := floatPrice.Float64()
	return price
}

func GetTicksForPrice(price *big.Float, tick, tickSpacing *big.Int, spreadPCT *big.Float) (*big.Int, *big.Int) {
	higherPCT := new(big.Float).Quo(new(big.Float).Add(big.NewFloat(100), spreadPCT), big.NewFloat(100))
	lowerPCT := new(big.Float).Quo(new(big.Float).Sub(big.NewFloat(100), spreadPCT), big.NewFloat(100))
	// log.Println("pct", lowerPCT.String(), higherPCT.String())
	// log.Println(higherPCT, lowerPCT)
	priceHigher := new(big.Float).Mul(price, higherPCT)
	priceLower := new(big.Float).Mul(price, lowerPCT)
	log.Println("pl", "ph", priceLower, priceHigher)
	tickNormalizedFee := new(big.Int).Mul(new(big.Int).Div(tick, tickSpacing), tickSpacing)
	// log.Println("normalized tick", tickNormalizedFee.String())
	lower := tickNormalizedFee
	// log.Println("lower", tickNormalizedFee)
	for {
		TICK := lower.Int64()
		TICKfloat := float64(TICK)
		// time.Sleep(5 * time.Second)
		price := math.Pow(1.0001, TICKfloat)
		priceCalculated := big.NewFloat(price)
		if priceLower.Cmp(priceCalculated) == 1 {
			break
		}
		lower.Sub(lower, tickSpacing)
	}
	tickNormalizedFee = new(big.Int).Mul(new(big.Int).Div(tick, tickSpacing), tickSpacing)
	higher := tickNormalizedFee
	// log.Println("higher", tickNormalizedFee)
	for {
		TICK := higher.Int64()
		TICKfloat := float64(TICK)
		price := math.Pow(1.0001, TICKfloat)
		priceCalculated := big.NewFloat(price)
		// log.Println("ph", "pc", priceHigher, priceCalculated)
		if priceHigher.Cmp(priceCalculated) == -1 {
			break
		}
		// log.Println(higher, priceHigher, priceCalculated, price)
		higher.Add(higher, tickSpacing)
	}
	return lower, higher
}

func BigIntToBinUint128(value *big.Int) (bin.Uint128, error) {
	var returnValue bin.Uint128
	bytes := value.Bytes()
	if len(bytes) > 16 {
		return returnValue, errors.New("Overflow, too many bytes from big.Int")
	}
	switch {
	case len(bytes) <= 8:
		low := make([]byte, 8)
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		for idx, byt := range bytes {
			low[idx] = byt
		}
		returnValue.Lo = binary.LittleEndian.Uint64(low)
	case len(bytes) > 8 && len(bytes) < 17:
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		low := make([]byte, 8)
		high := make([]byte, 8)
		highBytes := bytes[8:]
		lowBytes := bytes[:8]
		for idx, byt := range highBytes {
			high[idx] = byt
		}
		for idx, byt := range lowBytes {
			low[idx] = byt
		}
		returnValue.Lo = binary.LittleEndian.Uint64(low)
		returnValue.Hi = binary.LittleEndian.Uint64(high)
	}
	//low := make([]byte, 8)
	//high := make([]byte, 8)
	//log.Println(low, high)
	return returnValue, nil
}
