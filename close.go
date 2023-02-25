package godium

import (
	"context"
	"fmt"
	amm_v32 "github.com/Norbaeocystin/godium/amm_v3"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
)

func DecreaseLiquidityAndClose(position amm_v32.PersonalPositionState, client rpc.Client, token0, token1, tokenVault0, tokenVault1,
	poolStateAddress solana.PublicKey, wallet solana.PrivateKey, rewardsAccounts []*solana.AccountMeta) (solana.Signature, error) {
	owner := wallet
	amm_v32.ProgramID = RAYDIUM_PROGRAM_ID
	ktas := GetTickArrays(&client, poolStateAddress)
	protocolPosition, _ := GetProtocolPositionAddress(poolStateAddress, position.TickLowerIndex, position.TickUpperIndex)
	personalPosition, _ := GetPersonalPositionAddress(position.NftMint)
	// metadataAccount, _ := GetNFTMetadaAddress(position.NftMint)
	positionNFTAccount, _ := GetPositionNFTAccount(owner.PublicKey(), position.NftMint)
	ktaLower := GetTickArray(position.TickLowerIndex, ktas)
	ktaUpper := GetTickArray(position.TickUpperIndex, ktas)
	i0 := solana.NewInstruction(COMPUTE_BUDGET,
		[]*solana.AccountMeta{},
		// fee 1, u
		[]uint8{0, 32, 161, 7, 0, 1, 0, 0, 0},
	)
	i := amm_v32.NewDecreaseLiquidityInstruction(
		position.Liquidity,
		0,
		0,
		wallet.PublicKey(),
		positionNFTAccount,
		personalPosition,
		poolStateAddress,
		protocolPosition,
		tokenVault0,
		tokenVault1,
		ktaLower.Account,
		ktaUpper.Account,
		token0,
		token1,
		solana.TokenProgramID,
	)
	if len(rewardsAccounts) > 0 {
		log.Println(len(i.AccountMetaSlice))
		i.AccountMetaSlice = append(i.AccountMetaSlice, rewardsAccounts...)
	}
	i1 := i.Build()
	log.Println(len(i1.Accounts()))
	i2 := amm_v32.NewClosePositionInstruction(
		owner.PublicKey(),
		position.NftMint,
		positionNFTAccount, // ata pda
		personalPosition,
		solana.SystemProgramID, // const
		solana.TokenProgramID,  // const
	).Build()
	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			i0, i1, i2,
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
	// log.Println(position, positionMint.PublicKey(), positionTokenAccount, positionMA)
	return sig, err
}
