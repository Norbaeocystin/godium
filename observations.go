package godium

import (
	"context"
	amm_v32 "github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"sort"
	"time"
)

type Observations []amm_v32.Observation

func (a Observations) Len() int           { return len(a) }
func (a Observations) Less(i, j int) bool { return a[i].BlockTimestamp < a[j].BlockTimestamp }
func (a Observations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetObservations(client *rpc.Client, poolObservationKey solana.PublicKey) {
	observations, _ := client.GetAccountInfo(context.TODO(), poolObservationKey)
	decoder := bin.NewBorshDecoder(observations.GetBinary())
	var observe amm_v32.ObservationState
	decoder.Decode(&observe)
	for _, observation := range observe.Observations {
		price := CalculatePriceFromSQRPriceQ64(observation.SqrtPriceX64.BigInt())
		// time.Unix(int64(observation.BlockTimestamp), 0)
		log.Println(price, time.Unix(int64(observation.BlockTimestamp), 0))
	}
}

func GetLRFromObservations(client *rpc.Client, poolObservationKey solana.PublicKey) {
	observations, _ := client.GetAccountInfo(context.TODO(), poolObservationKey)
	decoder := bin.NewBorshDecoder(observations.GetBinary())
	var observe amm_v32.ObservationState
	decoder.Decode(&observe)
	observationsDecoded := Observations{}
	for _, observation := range observe.Observations {
		// price := CalculatePriceFromSQRPriceQ64(observation.SqrtPriceX64.BigInt())
		// time.Unix(int64(observation.BlockTimestamp), 0)
		observationsDecoded = append(observationsDecoded, observation)
		// log.Println(price, time.Unix(int64(observation.BlockTimestamp), 0))
	}
	sort.Sort(observationsDecoded)
	series := make(Series, 0)
	for idx, observation := range observationsDecoded {
		price := CalculatePriceFromSQRPriceQ64(observation.SqrtPriceX64.BigInt())
		series = append(series, Coordinate{float64(idx), price})
	}
	log.Println(observationsDecoded[len(observationsDecoded)-1].BlockTimestamp - observationsDecoded[0].BlockTimestamp)
	gradient, intercept, angle, err := LinearRegression(series)
	log.Println(gradient, intercept, angle, err)
}
