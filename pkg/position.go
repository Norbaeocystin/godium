package raydium

import (
	"context"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"raydium/pkg/amm_v3"
	"strings"
)

type PositionKeys struct {
	NFTMint            solana.PublicKey
	ProtocolPosition   solana.PublicKey
	PersonalPosition   solana.PublicKey
	PositionNFTAccount solana.PublicKey
	RaydiumPosition    bool
}

func GetPosition(client rpc.Client, personalPosition solana.PublicKey) amm_v3.PersonalPositionState {
	account, _ := client.GetAccountInfoWithOpts(context.TODO(),
		personalPosition,
		&rpc.GetAccountInfoOpts{
			Encoding:       solana.EncodingBase64,
			Commitment:     rpc.CommitmentFinalized,
			DataSlice:      nil,
			MinContextSlot: nil,
		},
	)
	var position amm_v3.PersonalPositionState
	dataPos := account.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&position)
	// log.Println(wpData)
	return position
}

func FindRaydiumPositionsForOwner(client *rpc.Client, owner, poolAddress solana.PublicKey) ([]PositionKeys, error) {
	positionsKeys := make([]PositionKeys, 0)
	tokens, err := client.GetTokenAccountsByOwner(context.TODO(), owner,
		&rpc.GetTokenAccountsConfig{
			Mint:      nil,
			ProgramId: solana.TokenProgramID.ToPointer(),
		},

		&rpc.GetTokenAccountsOpts{
			Commitment: "",
			Encoding:   solana.EncodingBase64,
			DataSlice:  nil,
		})
	if err != nil {
		log.Println("fetching positions error", err)
	}
	for _, tk := range tokens.Value {
		var ta token.Account
		borshDec := bin.NewBorshDecoder(tk.Account.Data.GetBinary())
		borshDec.Decode(&ta)
		if ta.Amount == 1 {
			log.Println()
			keys, err := FindPublicKeysForPositioMint(client, ta.Mint, poolAddress)
			if err != nil {
				return positionsKeys, err
			}
			// var pubKey solana.PublicKey
			if keys.RaydiumPosition == true {
				positionsKeys = append(positionsKeys, keys)
			}
		}
	}
	return positionsKeys, nil
}

// not working if use with SOL, position needs to be open with WSOL
func FindPublicKeysForPositioMint(client *rpc.Client, positionMint, poolAdddress solana.PublicKey) (PositionKeys, error) {
	var positionKeys PositionKeys
	positionKeys.RaydiumPosition = false
	one := 1
	signature, err := client.GetSignaturesForAddressWithOpts(context.TODO(), positionMint,
		&rpc.GetSignaturesForAddressOpts{
			&one,
			solana.Signature{},
			solana.Signature{},
			rpc.CommitmentFinalized,
			nil,
		})
	if err != nil {
		return positionKeys, err
	}
	version := uint64(0)
	opts := rpc.GetTransactionOpts{
		Encoding:                       solana.EncodingBase64,
		Commitment:                     rpc.CommitmentFinalized,
		MaxSupportedTransactionVersion: &version,
	}
	// log.Println("sign", signature[0].Signature)

	txs, err := client.GetTransaction(context.TODO(), signature[0].Signature, &opts)
	if err != nil {
		log.Println(err)
		return positionKeys, err
	}
	if strings.Contains(strings.Join(txs.Meta.LogMessages, ","), RAYDIUM_PROGRAM_ID.String()) {
		tx, _ := txs.Transaction.GetTransaction()
		accounts, _ := tx.AccountMetaList()
		if len(accounts) > 13 {
			// log.Println("Match tx with signature", signature[0].Signature)
			positionKeys.NFTMint = accounts[1].PublicKey
			positionKeys.PositionNFTAccount = accounts[2].PublicKey
			positionKeys.PersonalPosition = accounts[7].PublicKey
			positionKeys.ProtocolPosition = accounts[5].PublicKey
			positionKeys.RaydiumPosition = true
		}
	}
	return positionKeys, nil
}
