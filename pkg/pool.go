package raydium

import (
	"context"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"raydium/pkg/amm_v3"
)

func GetPoolState(client rpc.Client, poolAddress solana.PublicKey) amm_v3.PoolState {
	account, _ := client.GetAccountInfoWithOpts(context.TODO(),
		poolAddress,
		&rpc.GetAccountInfoOpts{
			Encoding:       solana.EncodingBase64,
			Commitment:     rpc.CommitmentFinalized,
			DataSlice:      nil,
			MinContextSlot: nil,
		},
	)
	var pool amm_v3.PoolState
	dataPos := account.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&pool)
	// log.Println(wpData)
	return pool
}
