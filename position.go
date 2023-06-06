package godium

import (
	"context"
	"github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
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
	//
	var position amm_v3.PersonalPositionState
	dataPos := account.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&position)
	// log.Println(wpData)
	return position
}

func FindRaydiumPositionsForOwner(client *rpc.Client, owner, poolAddress solana.PublicKey) ([]amm_v3.PersonalPositionState, error) {
	positions := make([]amm_v3.PersonalPositionState, 0)
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
			// tk.Pubkey = positionNFTAccount
			pak, _ := GetPersonalPositionAddress(ta.Mint)
			m := GetPosition(*client, pak)
			// var pubKey solana.PublicKey
			if m.PoolId.String() == poolAddress.String() {
				positions = append(positions, m)
			}
		}
	}
	return positions, nil
}
