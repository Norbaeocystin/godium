package godium

import (
	"context"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"math/big"
)

func FindATA(wallet solana.PrivateKey, mint solana.PublicKey) solana.PublicKey {
	ataAccount, _, _ := solana.FindAssociatedTokenAddress(wallet.PublicKey(), mint)
	return ataAccount
}

func FindOrCreateATA(client *rpc.Client, wallet solana.PrivateKey, mint solana.PublicKey) solana.PublicKey {
	// cSol turbo sol AVxnqyCameKsKTCGVKeyJMA7vjHnxJit6afC8AM9MdMj
	// cUSDC turbo usdc HKijBKC2zKcV2BXA9CuNemmWUhTuFkPLLgvQBP7zrQjL
	// mint = solana.MustPublicKeyFromBase58("AVxnqyCameKsKTCGVKeyJMA7vjHnxJit6afC8AM9MdMj")
	ataAccount, _, _ := solana.FindAssociatedTokenAddress(wallet.PublicKey(), mint)
	tokens, err := client.GetTokenAccountsByOwner(context.TODO(), wallet.PublicKey(),
		&rpc.GetTokenAccountsConfig{
			Mint: &mint,
			// ProgramId: solana.TokenProgramID.ToPointer(),
		},

		&rpc.GetTokenAccountsOpts{
			Commitment: "",
			Encoding:   solana.EncodingBase64,
			DataSlice:  nil,
		})
	if err != nil {
		log.Println("got error during searching", err)
	}
	for _, token := range tokens.Value {
		log.Println("found ata", token.Pubkey)
		return token.Pubkey
	}
	log.Println("not found ata, creating ata")
	i := associatedtokenaccount.NewCreateInstruction(
		wallet.PublicKey(),
		wallet.PublicKey(),
		mint,
	).Build()
	for _, token := range tokens.Value {
		log.Println(token.Pubkey)
	}
	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			i,
		},
		recent.Value.Blockhash, //NONCE
		solana.TransactionPayer(wallet.PublicKey()),
	)
	// log.Println(tx, err)
	// TODO intiliaze those 2 accounts
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if wallet.PublicKey().Equals(key) {
				return &wallet
			}
			return nil
		},
	)
	if err != nil {
		log.Println("unable to sign transaction: %w", err)
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
	if err != nil {
		panic(err)
	}
	log.Println("tx for creating ata:", sig)
	// return sig
	return ataAccount
}

func GetBalance(client *rpc.Client, tokenAddress solana.PublicKey) (uint64, error) {
	result, err := client.GetTokenAccountBalance(context.TODO(), tokenAddress, rpc.CommitmentConfirmed)
	if err != nil {
		log.Println("fetching balance returned", err)
		return 1, err
	}
	amount, ok := new(big.Int).SetString(result.Value.Amount, 10)
	log.Println("got amount which can be used for liquidation", amount, ok)
	liquidityAmount := amount.Uint64()
	return liquidityAmount, nil
}

func GetTokenBalance(client *rpc.Client, tokenPublicMint, owner solana.PublicKey) (solana.PublicKey, solana.PublicKey,
	uint64) {
	tokens, _ := client.GetTokenAccountsByOwner(context.TODO(), owner,
		&rpc.GetTokenAccountsConfig{
			Mint:      nil,
			ProgramId: solana.TokenProgramID.ToPointer(),
		},
		&rpc.GetTokenAccountsOpts{
			Commitment: "",
			Encoding:   solana.EncodingBase64,
			DataSlice:  nil,
		})
	for _, tk := range tokens.Value {
		var ta token.Account
		borshDec := bin.NewBorshDecoder(tk.Account.Data.GetBinary())
		// log.Println(tk.Pubkey.String(), ta.Mint)
		borshDec.Decode(&ta)
		if ta.Mint == tokenPublicMint {
			return tk.Pubkey, ta.Mint, ta.Amount
		}
		return solana.PublicKey{}, solana.PublicKey{}, 0
	}
	return solana.PublicKey{}, solana.PublicKey{}, 0
}
