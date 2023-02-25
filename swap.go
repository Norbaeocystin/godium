package godium

import (
	"context"
	"fmt"
	amm_v32 "github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
)

func Swap(client *rpc.Client, amountIn, amountOutMin uint64, sqrtPriceLimit bin.Uint128, isBaseInput bool, tokenIn,
	tokenOut, observationState, inputVault, outputVault, poolState, ammConfig, tickarray solana.PublicKey, owner solana.PrivateKey) (solana.Signature, error) {
	amm_v32.ProgramID = RAYDIUM_PROGRAM_ID
	i0 := solana.NewInstruction(COMPUTE_BUDGET,
		[]*solana.AccountMeta{},
		// fee 1, u
		[]uint8{0, 32, 161, 7, 0, 1, 0, 0, 0},
	)
	i := amm_v32.NewSwapInstruction(
		amountIn,
		amountOutMin,
		sqrtPriceLimit,
		isBaseInput,
		owner.PublicKey(),
		ammConfig,
		poolState,
		tokenIn,
		tokenOut,
		inputVault,
		outputVault,
		observationState,
		solana.TokenProgramID,
		solana.MustPublicKeyFromBase58("4Msu9Acwj2U9rZCsWCKFUqhXL5UWovPGUxi99VJ2aXYA"), // tickarray,
	).Build()
	accountsMeta := i.Accounts()
	for idx, accountMeta := range accountsMeta {
		log.Println(idx, accountMeta.PublicKey)
	}
	log.Println("tx prepared")
	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	log.Println("blockhash prepared")
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			i0, i,
		},
		recent.Value.Blockhash, //NONCE
		solana.TransactionPayer(owner.PublicKey()),
	)
	// log.Println(tx, err)
	// TODO intiliaze those 2 accounts
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if owner.PublicKey().Equals(key) {
				return &owner
			}
			return nil
		},
	)
	log.Println("tx signed")
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	sig, err := client.SendTransactionWithOpts(context.TODO(), tx,
		rpc.TransactionOpts{
			Encoding:            "",
			SkipPreflight:       false,
			PreflightCommitment: "",
			MaxRetries:          nil,
			MinContextSlot:      nil,
		},
	)
	// log.Println(sig)
	return sig, err
}
