package godium

import (
	"context"
	"github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
	"math/big"
)

type LiquidityData struct {
	PoolData            amm_v3.PoolState
	Token0              token.Mint
	Token1              token.Mint
	LiquiditiesForTicks []LiquidityForTick
}

type LiquidityForTick struct {
	Current        bool
	Tick           int32
	TickPrice      float64
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
	Token0         *big.Int
	Token1         *big.Int
}

func GetAllLiquidities(client *rpc.Client, market solana.PublicKey) LiquidityData {
	var liqData LiquidityData
	poolData := GetPoolState(*client, market)
	liqData.PoolData = poolData
	account0, _ := client.GetAccountInfo(context.TODO(), poolData.TokenMint0)
	var baseData token.Mint
	decoder := bin.NewBinDecoder(account0.GetBinary())
	decoder.Decode(&baseData)
	// log.Println(baseData)
	token0Decimals := math.Pow(10, float64(baseData.Decimals))
	account1, _ := client.GetAccountInfo(context.TODO(), poolData.TokenMint1)
	var quoteData token.Mint
	decoder = bin.NewBinDecoder(account1.GetBinary())
	decoder.Decode(&quoteData)
	token1Decimals := math.Pow(10, float64(quoteData.Decimals))
	liqData.Token0 = baseData
	liqData.Token1 = quoteData
	ktas := GetTickArrays(client, market)
	lickForTicks := make([]LiquidityForTick, 0)
	for _, kta := range ktas {
		for _, tickData := range kta.TickArrayState.Ticks {
			tick := tickData.Tick
			if tickData.LiquidityNet.BigInt().Cmp(big.NewInt(0)) != 0 {
				var liqForTick LiquidityForTick
				price := (math.Pow(1.0001, float64(tick)) * token0Decimals) / token1Decimals
				liqForTick.LiquidityNet = tickData.LiquidityNet.BigInt()
				liqForTick.LiquidityGross = tickData.LiquidityGross.BigInt()
				liqForTick.Tick = tick
				liqForTick.TickPrice = price
				if tick == poolData.TickCurrent {
					liqForTick.Current = true
				} else {
					liqForTick.Current = false
				}
				if tickData.LiquidityNet.BigInt().Cmp(big.NewInt(0)) == 1 {
					token0, token1 := CalculateAmounts(tickData.LiquidityNet.BigInt(), int64(tick-int32(poolData.TickSpacing)), int64(tick+int32(poolData.TickSpacing)), int64(poolData.TickCurrent))
					liqForTick.Token0 = token0
					liqForTick.Token1 = token1
				}
				lickForTicks = append(lickForTicks, liqForTick)
			}
		}
	}
	liqData.LiquiditiesForTicks = lickForTicks
	return liqData
}
