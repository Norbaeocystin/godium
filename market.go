package godium

import (
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
	m.SetKtas()
	m.SetData()
	return m
}

type Market struct {
	ProgramId solana.PublicKey
	MarketId  solana.PublicKey
	PoolState amm_v3.PoolState
	KTAS      []KeyedTickArray
	Client    *rpc.Client
}

func (m Market) FetchData() amm_v3.PoolState {
	data := GetPoolState(*m.Client, m.MarketId)
	return data
}

func (m *Market) SetData() {
	m.PoolState = m.FetchData()
}

func (m Market) FetchKtas() []KeyedTickArray {
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
	kta := GetTickArray(m.PoolState.TickCurrent, m.KTAS)
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
		kta.Account,
	).Build()
}

func (m Market) SwapAtoBExactInputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price - (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) * priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := GetTickArray(tick, m.KTAS)
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
		kta.Account,
	).Build()
}

func (m Market) SwapBToAExactInputInstructionWithSlippageUseState(amount uint64, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	price := CalculatePriceFromSQRPriceQ64(m.PoolState.SqrtPriceX64.BigInt())
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) / priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := GetTickArray(m.PoolState.TickCurrent, m.KTAS)
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
		kta.Account,
	).Build()
}

func (m Market) SwapBToAExactInputInstruction(amount, otherAmountThreshold uint64, sqrtPriceLimit bin.Uint128, owner, ownerTokenAAddress, ownerTokenBAddress, kta solana.PublicKey) solana.Instruction {
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

func (m Market) SwapBToAExactInputInstructionWithSlippageUsePrice(amount uint64, price, slippagePCT float64, owner, ownerTokenAAddress, ownerTokenBAddress solana.PublicKey) solana.Instruction {
	amm_v3.ProgramID = m.ProgramId
	tick := (PriceToTick(price) / int32(m.PoolState.TickSpacing)) * int32(m.PoolState.TickSpacing)
	priceWithSlippage := price + (price * (slippagePCT / 100))
	otherAmountThreshold := uint64(float64(amount) / priceWithSlippage)
	sqrtPriceLimit, _ := BigIntToBinUint128(CalculateSqrtPriceQ64(big.NewFloat(priceWithSlippage)))
	kta := GetTickArray(tick, m.KTAS)
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
		kta.Account,
	).Build()
}
