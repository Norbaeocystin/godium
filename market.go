package godium

import (
	"context"
	"github.com/Norbaeocystin/godium/amm_v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math/big"
)

func NewMarket(client *rpc.Client, marketId solana.PublicKey) Market {
	var m Market
	m.Client = client
	m.ProgramId = RAYDIUM_PROGRAM_ID
	m.MarketId = marketId
	// too costly
	// m.SetKtas()
	m.SetData()
	return m
}

type Market struct {
	ProgramId solana.PublicKey
	MarketId  solana.PublicKey
	PoolState amm_v3.PoolState
	AmmConfig amm_v3.AmmConfig
	KTAS      KTAS
	Client    *rpc.Client
}

func (m Market) MintA() solana.PublicKey {
	return m.PoolState.TokenMint0
}

func (m Market) MintB() solana.PublicKey {
	return m.PoolState.TokenMint1
}

func (m Market) Fee() float64 {
	return float64(m.AmmConfig.TradeFeeRate) / 10_000_00.0
}

func (m Market) FetchData() amm_v3.PoolState {
	data := GetPoolState(*m.Client, m.MarketId)
	return data
}

func (m Market) FetchAmmConfig() amm_v3.AmmConfig {
	var ammConfig amm_v3.AmmConfig
	account, _ := m.Client.GetAccountInfo(context.TODO(), m.PoolState.AmmConfig)
	decoder := bin.NewBorshDecoder(account.Bytes())
	decoder.Decode(&ammConfig)
	return ammConfig
}

func (m *Market) SetData() {
	m.PoolState = m.FetchData()
	m.AmmConfig = m.FetchAmmConfig()
}

func (m Market) FetchKtas() KTAS {
	ktas := GetTickArrays(m.Client, m.MarketId)
	return ktas
}

func (m *Market) SetKtas() {
	m.KTAS = m.FetchKtas()
}

func (m Market) GetKtasForTicks(currentTick, tickForSqrtPriceLimit int32) (solana.PublicKey, solana.PublicKey, solana.PublicKey) {
	keyTickForSqrtPriceLimit := GetTickArray(tickForSqrtPriceLimit, m.KTAS).Account
	return GetTickArray(currentTick, m.KTAS).Account, keyTickForSqrtPriceLimit, keyTickForSqrtPriceLimit
}

// calculate derived key
func (m Market) GetTickAccount(tick int32) solana.PublicKey {
	startTick := GetStartTickIndex(tick, m.PoolState.TickSpacing)
	key, _ := GetTickArrayAddress(m.MarketId, startTick)
	return key
}

func (m Market) SwapAtoBExactInputInstruction(amount, otherAmountThreshold uint64, sqrtPriceLimit bin.Uint128, owner, ownerTokenAAddress, ownerTokenBAddress, kta solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

func (m Market) SwapAtoBExactInputInstructionWithSlippageUseState(amount uint64, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	price := CalculatePriceFromSQRPriceQ64(m.PoolState.SqrtPriceX64.BigInt())
	priceWithSlippage := price - (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) * priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := m.GetTickAccount(m.PoolState.TickCurrent)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

func (m Market) SwapAtoBExactInputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	// tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price - (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) * priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := m.GetTickAccount(PriceToTick(price))
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

func (m Market) SwapBtoAExactInputInstructionWithSlippageUseState(amount uint64, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	price := CalculatePriceFromSQRPriceQ64(m.PoolState.SqrtPriceX64.BigInt())
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) / priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := m.GetTickAccount(m.PoolState.TickCurrent)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

func (m Market) SwapBtoAExactInputInstruction(amount, otherAmountThreshold uint64, sqrtPriceLimit bin.Uint128, owner, ownerTokenAAddress, ownerTokenBAddress, kta solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

func (m Market) SwapBtoAExactInputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) / priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := m.GetTickAccount(tick)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		true,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output amount and otherAmount.. max amount of token to be used
func (m Market) SwapAtoBExactOutputInstruction(amount, otherAmountThreshold uint64, sqrtPriceLimit bin.Uint128, owner, ownerTokenAAddress, ownerTokenBAddress, kta solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output amount and otherAmount.. max amount of token to be used
func (m Market) SwapBtoAExactOutputInstruction(amount, otherAmountThreshold uint64, sqrtPriceLimit bin.Uint128, owner, ownerTokenAAddress, ownerTokenBAddress, kta solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output
func (m Market) SwapAtoBExactOutputInstructionWithSlippageUseState(amount uint64, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	price := CalculatePriceFromSQRPriceQ64(m.PoolState.SqrtPriceX64.BigInt())
	priceWithSlippage := price - (price * (slippagePCT / 100))
	otherAmountThreshold := uint64((float64(amount) / priceWithSlippage) * 1.01)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	tickp := PriceToTick(priceWithSlippage)
	kta := m.GetTickAccount(tickp) //)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output
func (m Market) SwapAtoBExactOutputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	// tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price - (price * (slippagePCT / 100))
	otherAmountThreshold := uint64((float64(amount) / priceWithSlippage) * 1.01)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	tickp := PriceToTick(priceWithSlippage)
	kta := m.GetTickAccount(tickp)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenAAddress,
		ownerTokenBAddress,
		m.PoolState.TokenVault0,
		m.PoolState.TokenVault1,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output
func (m Market) SwapBtoAExactOutputInstructionWithSlippageUseState(amount uint64, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	price := CalculatePriceFromSQRPriceQ64(m.PoolState.SqrtPriceX64.BigInt())
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64((float64(amount) * priceWithSlippage) * 1.01)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := m.GetTickAccount(m.PoolState.TickCurrent)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}

// amount is output amount - what means amount of token A - how much you want
func (m Market) SwapBtoAExactOutputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	// tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64((float64(amount) * priceWithSlippage) * 1.01)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	tickp := PriceToTick(priceWithSlippage)
	kta := m.GetTickAccount(tickp)
	return amm_v3.NewSwapInstruction(
		amount,
		otherAmountThreshold,
		sqrtPriceLimit,
		false,
		owner,
		m.PoolState.AmmConfig,
		m.MarketId,
		ownerTokenBAddress,
		ownerTokenAAddress,
		m.PoolState.TokenVault1,
		m.PoolState.TokenVault0,
		m.PoolState.ObservationKey,
		solana.TokenProgramID,
		kta,
	).Build()
}
