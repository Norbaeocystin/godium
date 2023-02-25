package raydium

import (
	"bytes"
	"encoding/binary"
	"github.com/gagliardetto/solana-go"
)

func GetAmmConfigAddress() {

}

func GetPoolAddress(ammConfig, tokenMint0, tokenMint1 solana.PublicKey) (solana.PublicKey, uint8) {
	position, bump, _ := solana.FindProgramAddress([][]byte{POOL_SEED, ammConfig.Bytes(), tokenMint0.Bytes(),
		tokenMint1.Bytes()}, RAYDIUM_PROGRAM_ID)
	return position, bump
}

func GetPoolVaultAddress(pool, vaultTokenMint solana.PublicKey) (solana.PublicKey, uint8) {
	position, bump, _ := solana.FindProgramAddress([][]byte{POOL_VAULT_SEED, pool.Bytes(), vaultTokenMint.Bytes()}, RAYDIUM_PROGRAM_ID)
	return position, bump
}
func GetPoolRewardVaultAddress(pool, rewardTokenMint solana.PublicKey) (solana.PublicKey, uint8) {
	position, bump, _ := solana.FindProgramAddress([][]byte{POOL_VAULT_SEED, pool.Bytes(), rewardTokenMint.Bytes()}, RAYDIUM_PROGRAM_ID)
	return position, bump
}

func GetTickArrayAddress(pool solana.PublicKey, startIndex int32) (solana.PublicKey, uint8) {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, startIndex)
	taa, bump, _ := solana.FindProgramAddress([][]byte{TICK_ARRAY_SEED, pool.Bytes(),
		buff.Bytes()}, RAYDIUM_PROGRAM_ID)
	return taa, bump
}

func GetProtocolPositionAddress(pool solana.PublicKey, tickLower, tickUpper int32) (solana.PublicKey, uint8) {
	buffL := new(bytes.Buffer)
	binary.Write(buffL, binary.BigEndian, tickLower)
	buffU := new(bytes.Buffer)
	binary.Write(buffU, binary.BigEndian, tickUpper)
	position, bump, _ := solana.FindProgramAddress([][]byte{POSITION_SEED, pool.Bytes(),
		buffL.Bytes(), buffU.Bytes()}, RAYDIUM_PROGRAM_ID)
	return position, bump
}

func GetPersonalPositionAddress(nftMint solana.PublicKey) (solana.PublicKey, uint8) {
	position, bump, _ := solana.FindProgramAddress([][]byte{POSITION_SEED, nftMint.Bytes()}, RAYDIUM_PROGRAM_ID)
	return position, bump
}

func GetNFTMetadaAddress(nftMint solana.PublicKey) (solana.PublicKey, uint8) {
	// metadata account
	nftMeta, bump, _ := solana.FindProgramAddress([][]byte{[]byte("metadata"), METAPLEX.Bytes(), nftMint.Bytes()}, METAPLEX)
	return nftMeta, bump
}

func GetOperationAddress() {

}

func GetPositionNFTAccount(owner, mint solana.PublicKey) (solana.PublicKey, uint8) {
	positionNFTAccount, bump, _ := solana.FindProgramAddress([][]byte{owner.Bytes(), solana.TokenProgramID.Bytes(), mint.Bytes()}, solana.SPLAssociatedTokenAccountProgramID)
	return positionNFTAccount, bump
}
