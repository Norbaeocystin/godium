package godium

import (
	"context"
	"fmt"
	amm_v32 "github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func OpenNewPositionAndAddLiquidity(tickLower,
	tickUpper int32, client rpc.Client, tokenAMax,
	tokenBMax uint64, liquidity bin.Uint128,
	tokenA, tokenB, tokenVault0, tokenVault1, poolStateAddress solana.PublicKey, wallet solana.PrivateKey) (solana.Signature, error) {
	owner := wallet
	amm_v32.ProgramID = RAYDIUM_PROGRAM_ID
	nftMint := solana.NewWallet()
	protocolPosition, _ := GetProtocolPositionAddress(poolStateAddress, tickLower, tickUpper)
	personalPosition, _ := GetPersonalPositionAddress(nftMint.PublicKey())
	metadataAccount, _ := GetNFTMetadaAddress(nftMint.PublicKey())
	positionNFTAccount, _ := GetPositionNFTAccount(owner.PublicKey(), nftMint.PublicKey())
	//  TickArrayLowerStartIndex -39600
	ktas := GetTickArrays(&client, poolStateAddress)
	ktaLower := GetTickArray(tickLower, ktas)
	ktaUpper := GetTickArray(tickUpper, ktas)
	i0 := solana.NewInstruction(COMPUTE_BUDGET,
		[]*solana.AccountMeta{},
		// fee 1, u
		[]uint8{0, 32, 161, 7, 0, 1, 0, 0, 0},
	)
	i1 := amm_v32.NewOpenPositionInstruction(
		tickLower,
		tickUpper,
		ktaLower.TickArrayState.StartTickIndex,
		ktaUpper.TickArrayState.StartTickIndex,
		liquidity,
		tokenAMax,
		tokenBMax,
		owner.PublicKey(),
		owner.PublicKey(),
		nftMint.PublicKey(),
		positionNFTAccount, // ata pda
		metadataAccount,    //pda
		poolStateAddress,
		protocolPosition, // pda
		ktaLower.Account, // pda
		ktaUpper.Account, //pda
		personalPosition, //pda
		tokenA,
		tokenB,
		tokenVault0,
		tokenVault1,
		solana.SysVarRentPubkey, // const
		solana.SystemProgramID,  // const
		solana.TokenProgramID,   // const
		solana.SPLAssociatedTokenAccountProgramID, // const
		METAPLEX, // const
	).Build()
	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			i0, i1,
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
			if nftMint.PublicKey().Equals(key) {
				return &nftMint.PrivateKey
			}
			return nil

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

func OpenNewPositionAndAddLiquidityIx(tickLower,
	tickUpper int32, client rpc.Client, tokenAMax,
	tokenBMax uint64, liquidity bin.Uint128,
	tokenA, tokenB, tokenVault0, tokenVault1, poolStateAddress, nftMint solana.PublicKey, wallet solana.PrivateKey) ([]solana.Instruction, error) {
	owner := wallet
	amm_v32.ProgramID = RAYDIUM_PROGRAM_ID
	// nftMint := solana.NewWallet()
	protocolPosition, _ := GetProtocolPositionAddress(poolStateAddress, tickLower, tickUpper)
	personalPosition, _ := GetPersonalPositionAddress(nftMint)
	metadataAccount, _ := GetNFTMetadaAddress(nftMint)
	positionNFTAccount, _ := GetPositionNFTAccount(owner.PublicKey(), nftMint)
	//  TickArrayLowerStartIndex -39600
	ktas := GetTickArrays(&client, poolStateAddress)
	ktaLower := GetTickArray(tickLower, ktas)
	ktaUpper := GetTickArray(tickUpper, ktas)
	i1 := amm_v32.NewOpenPositionInstruction(
		tickLower,
		tickUpper,
		ktaLower.TickArrayState.StartTickIndex,
		ktaUpper.TickArrayState.StartTickIndex,
		liquidity,
		tokenAMax,
		tokenBMax,
		owner.PublicKey(),
		owner.PublicKey(),
		nftMint,
		positionNFTAccount, // ata pda
		metadataAccount,    //pda
		poolStateAddress,
		protocolPosition, // pda
		ktaLower.Account, // pda
		ktaUpper.Account, //pda
		personalPosition, //pda
		tokenA,
		tokenB,
		tokenVault0,
		tokenVault1,
		solana.SysVarRentPubkey, // const
		solana.SystemProgramID,  // const
		solana.TokenProgramID,   // const
		solana.SPLAssociatedTokenAccountProgramID, // const
		METAPLEX, // const
	).Build()
	return []solana.Instruction{i1}, nil
}
