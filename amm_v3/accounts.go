// Code Generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package amm_v3

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AmmConfig struct {
	Bump            uint8
	Index           uint16
	Owner           ag_solanago.PublicKey
	ProtocolFeeRate uint32
	TradeFeeRate    uint32
	TickSpacing     uint16
	FundFeeRate     uint32
	PaddingU32      uint32
	FundOwner       ag_solanago.PublicKey
	Padding         [3]uint64
}

var AmmConfigDiscriminator = [8]byte{218, 244, 33, 104, 203, 203, 43, 111}

func (obj AmmConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AmmConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeRate` param:
	err = encoder.Encode(obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeRate` param:
	err = encoder.Encode(obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Serialize `TickSpacing` param:
	err = encoder.Encode(obj.TickSpacing)
	if err != nil {
		return err
	}
	// Serialize `FundFeeRate` param:
	err = encoder.Encode(obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Serialize `PaddingU32` param:
	err = encoder.Encode(obj.PaddingU32)
	if err != nil {
		return err
	}
	// Serialize `FundOwner` param:
	err = encoder.Encode(obj.FundOwner)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmmConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AmmConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[218 244 33 104 203 203 43 111]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeRate`:
	err = decoder.Decode(&obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `TradeFeeRate`:
	err = decoder.Decode(&obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `TickSpacing`:
	err = decoder.Decode(&obj.TickSpacing)
	if err != nil {
		return err
	}
	// Deserialize `FundFeeRate`:
	err = decoder.Decode(&obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `PaddingU32`:
	err = decoder.Decode(&obj.PaddingU32)
	if err != nil {
		return err
	}
	// Deserialize `FundOwner`:
	err = decoder.Decode(&obj.FundOwner)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type OperationState struct {
	Bump            uint8
	OperationOwners [10]ag_solanago.PublicKey
	WhitelistMints  [100]ag_solanago.PublicKey
}

var OperationStateDiscriminator = [8]byte{19, 236, 58, 237, 81, 222, 183, 252}

func (obj OperationState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(OperationStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `OperationOwners` param:
	err = encoder.Encode(obj.OperationOwners)
	if err != nil {
		return err
	}
	// Serialize `WhitelistMints` param:
	err = encoder.Encode(obj.WhitelistMints)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OperationState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(OperationStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[19 236 58 237 81 222 183 252]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `OperationOwners`:
	err = decoder.Decode(&obj.OperationOwners)
	if err != nil {
		return err
	}
	// Deserialize `WhitelistMints`:
	err = decoder.Decode(&obj.WhitelistMints)
	if err != nil {
		return err
	}
	return nil
}

type ObservationState struct {
	Initialized  bool
	PoolId       ag_solanago.PublicKey
	Observations [1000]Observation
	Padding      [5]ag_binary.Uint128
}

var ObservationStateDiscriminator = [8]byte{122, 174, 197, 53, 129, 9, 165, 132}

func (obj ObservationState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ObservationStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized` param:
	err = encoder.Encode(obj.Initialized)
	if err != nil {
		return err
	}
	// Serialize `PoolId` param:
	err = encoder.Encode(obj.PoolId)
	if err != nil {
		return err
	}
	// Serialize `Observations` param:
	err = encoder.Encode(obj.Observations)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ObservationState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ObservationStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[122 174 197 53 129 9 165 132]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized`:
	err = decoder.Decode(&obj.Initialized)
	if err != nil {
		return err
	}
	// Deserialize `PoolId`:
	err = decoder.Decode(&obj.PoolId)
	if err != nil {
		return err
	}
	// Deserialize `Observations`:
	err = decoder.Decode(&obj.Observations)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type PersonalPositionState struct {
	Bump                    uint8
	NftMint                 ag_solanago.PublicKey
	PoolId                  ag_solanago.PublicKey
	TickLowerIndex          int32
	TickUpperIndex          int32
	Liquidity               ag_binary.Uint128
	FeeGrowthInside0LastX64 ag_binary.Uint128
	FeeGrowthInside1LastX64 ag_binary.Uint128
	TokenFeesOwed0          uint64
	TokenFeesOwed1          uint64
	RewardInfos             [3]PositionRewardInfo
	Padding                 [8]uint64
}

var PersonalPositionStateDiscriminator = [8]byte{70, 111, 150, 126, 230, 15, 25, 117}

func (obj PersonalPositionState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PersonalPositionStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `NftMint` param:
	err = encoder.Encode(obj.NftMint)
	if err != nil {
		return err
	}
	// Serialize `PoolId` param:
	err = encoder.Encode(obj.PoolId)
	if err != nil {
		return err
	}
	// Serialize `TickLowerIndex` param:
	err = encoder.Encode(obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Serialize `TickUpperIndex` param:
	err = encoder.Encode(obj.TickUpperIndex)
	if err != nil {
		return err
	}
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthInside0LastX64` param:
	err = encoder.Encode(obj.FeeGrowthInside0LastX64)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthInside1LastX64` param:
	err = encoder.Encode(obj.FeeGrowthInside1LastX64)
	if err != nil {
		return err
	}
	// Serialize `TokenFeesOwed0` param:
	err = encoder.Encode(obj.TokenFeesOwed0)
	if err != nil {
		return err
	}
	// Serialize `TokenFeesOwed1` param:
	err = encoder.Encode(obj.TokenFeesOwed1)
	if err != nil {
		return err
	}
	// Serialize `RewardInfos` param:
	err = encoder.Encode(obj.RewardInfos)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PersonalPositionState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PersonalPositionStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[70 111 150 126 230 15 25 117]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `NftMint`:
	err = decoder.Decode(&obj.NftMint)
	if err != nil {
		return err
	}
	// Deserialize `PoolId`:
	err = decoder.Decode(&obj.PoolId)
	if err != nil {
		return err
	}
	// Deserialize `TickLowerIndex`:
	err = decoder.Decode(&obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Deserialize `TickUpperIndex`:
	err = decoder.Decode(&obj.TickUpperIndex)
	if err != nil {
		return err
	}
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthInside0LastX64`:
	err = decoder.Decode(&obj.FeeGrowthInside0LastX64)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthInside1LastX64`:
	err = decoder.Decode(&obj.FeeGrowthInside1LastX64)
	if err != nil {
		return err
	}
	// Deserialize `TokenFeesOwed0`:
	err = decoder.Decode(&obj.TokenFeesOwed0)
	if err != nil {
		return err
	}
	// Deserialize `TokenFeesOwed1`:
	err = decoder.Decode(&obj.TokenFeesOwed1)
	if err != nil {
		return err
	}
	// Deserialize `RewardInfos`:
	err = decoder.Decode(&obj.RewardInfos)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type PoolState struct {
	Bump                      uint8                 // 1
	AmmConfig                 ag_solanago.PublicKey // 32
	Owner                     ag_solanago.PublicKey // 32
	TokenMint0                ag_solanago.PublicKey // 32
	TokenMint1                ag_solanago.PublicKey //
	TokenVault0               ag_solanago.PublicKey
	TokenVault1               ag_solanago.PublicKey
	ObservationKey            ag_solanago.PublicKey
	MintDecimals0             uint8
	MintDecimals1             uint8
	TickSpacing               uint16
	Liquidity                 ag_binary.Uint128
	SqrtPriceX64              ag_binary.Uint128
	TickCurrent               int32
	ObservationIndex          uint16
	ObservationUpdateDuration uint16
	FeeGrowthGlobal0X64       ag_binary.Uint128
	FeeGrowthGlobal1X64       ag_binary.Uint128
	ProtocolFeesToken0        uint64
	ProtocolFeesToken1        uint64
	SwapInAmountToken0        ag_binary.Uint128
	SwapOutAmountToken1       ag_binary.Uint128
	SwapInAmountToken1        ag_binary.Uint128
	SwapOutAmountToken0       ag_binary.Uint128
	Status                    uint8
	Padding                   [7]uint8
	RewardInfos               [3]RewardInfo
	TickArrayBitmap           [16]uint64
	TotalFeesToken0           uint64
	TotalFeesClaimedToken0    uint64
	TotalFeesToken1           uint64
	TotalFeesClaimedToken1    uint64
	FundFeesToken0            uint64
	FundFeesToken1            uint64
	Padding1                  [26]uint64
	Padding2                  [32]uint64
}

var PoolStateDiscriminator = [8]byte{247, 237, 227, 245, 215, 195, 222, 70}

func (obj PoolState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PoolStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `AmmConfig` param:
	err = encoder.Encode(obj.AmmConfig)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `TokenMint0` param:
	err = encoder.Encode(obj.TokenMint0)
	if err != nil {
		return err
	}
	// Serialize `TokenMint1` param:
	err = encoder.Encode(obj.TokenMint1)
	if err != nil {
		return err
	}
	// Serialize `TokenVault0` param:
	err = encoder.Encode(obj.TokenVault0)
	if err != nil {
		return err
	}
	// Serialize `TokenVault1` param:
	err = encoder.Encode(obj.TokenVault1)
	if err != nil {
		return err
	}
	// Serialize `ObservationKey` param:
	err = encoder.Encode(obj.ObservationKey)
	if err != nil {
		return err
	}
	// Serialize `MintDecimals0` param:
	err = encoder.Encode(obj.MintDecimals0)
	if err != nil {
		return err
	}
	// Serialize `MintDecimals1` param:
	err = encoder.Encode(obj.MintDecimals1)
	if err != nil {
		return err
	}
	// Serialize `TickSpacing` param:
	err = encoder.Encode(obj.TickSpacing)
	if err != nil {
		return err
	}
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `SqrtPriceX64` param:
	err = encoder.Encode(obj.SqrtPriceX64)
	if err != nil {
		return err
	}
	// Serialize `TickCurrent` param:
	err = encoder.Encode(obj.TickCurrent)
	if err != nil {
		return err
	}
	// Serialize `ObservationIndex` param:
	err = encoder.Encode(obj.ObservationIndex)
	if err != nil {
		return err
	}
	// Serialize `ObservationUpdateDuration` param:
	err = encoder.Encode(obj.ObservationUpdateDuration)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthGlobal0X64` param:
	err = encoder.Encode(obj.FeeGrowthGlobal0X64)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthGlobal1X64` param:
	err = encoder.Encode(obj.FeeGrowthGlobal1X64)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeesToken0` param:
	err = encoder.Encode(obj.ProtocolFeesToken0)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeesToken1` param:
	err = encoder.Encode(obj.ProtocolFeesToken1)
	if err != nil {
		return err
	}
	// Serialize `SwapInAmountToken0` param:
	err = encoder.Encode(obj.SwapInAmountToken0)
	if err != nil {
		return err
	}
	// Serialize `SwapOutAmountToken1` param:
	err = encoder.Encode(obj.SwapOutAmountToken1)
	if err != nil {
		return err
	}
	// Serialize `SwapInAmountToken1` param:
	err = encoder.Encode(obj.SwapInAmountToken1)
	if err != nil {
		return err
	}
	// Serialize `SwapOutAmountToken0` param:
	err = encoder.Encode(obj.SwapOutAmountToken0)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	// Serialize `RewardInfos` param:
	err = encoder.Encode(obj.RewardInfos)
	if err != nil {
		return err
	}
	// Serialize `TickArrayBitmap` param:
	err = encoder.Encode(obj.TickArrayBitmap)
	if err != nil {
		return err
	}
	// Serialize `TotalFeesToken0` param:
	err = encoder.Encode(obj.TotalFeesToken0)
	if err != nil {
		return err
	}
	// Serialize `TotalFeesClaimedToken0` param:
	err = encoder.Encode(obj.TotalFeesClaimedToken0)
	if err != nil {
		return err
	}
	// Serialize `TotalFeesToken1` param:
	err = encoder.Encode(obj.TotalFeesToken1)
	if err != nil {
		return err
	}
	// Serialize `TotalFeesClaimedToken1` param:
	err = encoder.Encode(obj.TotalFeesClaimedToken1)
	if err != nil {
		return err
	}
	// Serialize `FundFeesToken0` param:
	err = encoder.Encode(obj.FundFeesToken0)
	if err != nil {
		return err
	}
	// Serialize `FundFeesToken1` param:
	err = encoder.Encode(obj.FundFeesToken1)
	if err != nil {
		return err
	}
	// Serialize `Padding1` param:
	err = encoder.Encode(obj.Padding1)
	if err != nil {
		return err
	}
	// Serialize `Padding2` param:
	err = encoder.Encode(obj.Padding2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PoolState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PoolStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[247 237 227 245 215 195 222 70]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `AmmConfig`:
	err = decoder.Decode(&obj.AmmConfig)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `TokenMint0`:
	err = decoder.Decode(&obj.TokenMint0)
	if err != nil {
		return err
	}
	// Deserialize `TokenMint1`:
	err = decoder.Decode(&obj.TokenMint1)
	if err != nil {
		return err
	}
	// Deserialize `TokenVault0`:
	err = decoder.Decode(&obj.TokenVault0)
	if err != nil {
		return err
	}
	// Deserialize `TokenVault1`:
	err = decoder.Decode(&obj.TokenVault1)
	if err != nil {
		return err
	}
	// Deserialize `ObservationKey`:
	err = decoder.Decode(&obj.ObservationKey)
	if err != nil {
		return err
	}
	// Deserialize `MintDecimals0`:
	err = decoder.Decode(&obj.MintDecimals0)
	if err != nil {
		return err
	}
	// Deserialize `MintDecimals1`:
	err = decoder.Decode(&obj.MintDecimals1)
	if err != nil {
		return err
	}
	// Deserialize `TickSpacing`:
	err = decoder.Decode(&obj.TickSpacing)
	if err != nil {
		return err
	}
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `SqrtPriceX64`:
	err = decoder.Decode(&obj.SqrtPriceX64)
	if err != nil {
		return err
	}
	// Deserialize `TickCurrent`:
	err = decoder.Decode(&obj.TickCurrent)
	if err != nil {
		return err
	}
	// Deserialize `ObservationIndex`:
	err = decoder.Decode(&obj.ObservationIndex)
	if err != nil {
		return err
	}
	// Deserialize `ObservationUpdateDuration`:
	err = decoder.Decode(&obj.ObservationUpdateDuration)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthGlobal0X64`:
	err = decoder.Decode(&obj.FeeGrowthGlobal0X64)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthGlobal1X64`:
	err = decoder.Decode(&obj.FeeGrowthGlobal1X64)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeesToken0`:
	err = decoder.Decode(&obj.ProtocolFeesToken0)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeesToken1`:
	err = decoder.Decode(&obj.ProtocolFeesToken1)
	if err != nil {
		return err
	}
	// Deserialize `SwapInAmountToken0`:
	err = decoder.Decode(&obj.SwapInAmountToken0)
	if err != nil {
		return err
	}
	// Deserialize `SwapOutAmountToken1`:
	err = decoder.Decode(&obj.SwapOutAmountToken1)
	if err != nil {
		return err
	}
	// Deserialize `SwapInAmountToken1`:
	err = decoder.Decode(&obj.SwapInAmountToken1)
	if err != nil {
		return err
	}
	// Deserialize `SwapOutAmountToken0`:
	err = decoder.Decode(&obj.SwapOutAmountToken0)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	// Deserialize `RewardInfos`:
	err = decoder.Decode(&obj.RewardInfos)
	if err != nil {
		return err
	}
	// Deserialize `TickArrayBitmap`:
	err = decoder.Decode(&obj.TickArrayBitmap)
	if err != nil {
		return err
	}
	// Deserialize `TotalFeesToken0`:
	err = decoder.Decode(&obj.TotalFeesToken0)
	if err != nil {
		return err
	}
	// Deserialize `TotalFeesClaimedToken0`:
	err = decoder.Decode(&obj.TotalFeesClaimedToken0)
	if err != nil {
		return err
	}
	// Deserialize `TotalFeesToken1`:
	err = decoder.Decode(&obj.TotalFeesToken1)
	if err != nil {
		return err
	}
	// Deserialize `TotalFeesClaimedToken1`:
	err = decoder.Decode(&obj.TotalFeesClaimedToken1)
	if err != nil {
		return err
	}
	// Deserialize `FundFeesToken0`:
	err = decoder.Decode(&obj.FundFeesToken0)
	if err != nil {
		return err
	}
	// Deserialize `FundFeesToken1`:
	err = decoder.Decode(&obj.FundFeesToken1)
	if err != nil {
		return err
	}
	// Deserialize `Padding1`:
	err = decoder.Decode(&obj.Padding1)
	if err != nil {
		return err
	}
	// Deserialize `Padding2`:
	err = decoder.Decode(&obj.Padding2)
	if err != nil {
		return err
	}
	return nil
}

type ProtocolPositionState struct {
	Bump                    uint8
	PoolId                  ag_solanago.PublicKey
	TickLowerIndex          int32
	TickUpperIndex          int32
	Liquidity               ag_binary.Uint128
	FeeGrowthInside0LastX64 ag_binary.Uint128
	FeeGrowthInside1LastX64 ag_binary.Uint128
	TokenFeesOwed0          uint64
	TokenFeesOwed1          uint64
	RewardGrowthInside      [3]ag_binary.Uint128
	Padding                 [8]uint64
}

var ProtocolPositionStateDiscriminator = [8]byte{100, 226, 145, 99, 146, 218, 160, 106}

func (obj ProtocolPositionState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ProtocolPositionStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `PoolId` param:
	err = encoder.Encode(obj.PoolId)
	if err != nil {
		return err
	}
	// Serialize `TickLowerIndex` param:
	err = encoder.Encode(obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Serialize `TickUpperIndex` param:
	err = encoder.Encode(obj.TickUpperIndex)
	if err != nil {
		return err
	}
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthInside0LastX64` param:
	err = encoder.Encode(obj.FeeGrowthInside0LastX64)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthInside1LastX64` param:
	err = encoder.Encode(obj.FeeGrowthInside1LastX64)
	if err != nil {
		return err
	}
	// Serialize `TokenFeesOwed0` param:
	err = encoder.Encode(obj.TokenFeesOwed0)
	if err != nil {
		return err
	}
	// Serialize `TokenFeesOwed1` param:
	err = encoder.Encode(obj.TokenFeesOwed1)
	if err != nil {
		return err
	}
	// Serialize `RewardGrowthInside` param:
	err = encoder.Encode(obj.RewardGrowthInside)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProtocolPositionState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ProtocolPositionStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[100 226 145 99 146 218 160 106]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `PoolId`:
	err = decoder.Decode(&obj.PoolId)
	if err != nil {
		return err
	}
	// Deserialize `TickLowerIndex`:
	err = decoder.Decode(&obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Deserialize `TickUpperIndex`:
	err = decoder.Decode(&obj.TickUpperIndex)
	if err != nil {
		return err
	}
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthInside0LastX64`:
	err = decoder.Decode(&obj.FeeGrowthInside0LastX64)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthInside1LastX64`:
	err = decoder.Decode(&obj.FeeGrowthInside1LastX64)
	if err != nil {
		return err
	}
	// Deserialize `TokenFeesOwed0`:
	err = decoder.Decode(&obj.TokenFeesOwed0)
	if err != nil {
		return err
	}
	// Deserialize `TokenFeesOwed1`:
	err = decoder.Decode(&obj.TokenFeesOwed1)
	if err != nil {
		return err
	}
	// Deserialize `RewardGrowthInside`:
	err = decoder.Decode(&obj.RewardGrowthInside)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type TickArrayState struct {
	PoolId               ag_solanago.PublicKey // 32
	StartTickIndex       int32                 // 4
	Ticks                [60]TickState         // 10080 + 36 + 116 = 10232
	InitializedTickCount uint8                 // 1
	Padding              [115]uint8            //  115
}

var TickArrayStateDiscriminator = [8]byte{192, 155, 85, 205, 49, 249, 129, 42}

func (obj TickArrayState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TickArrayStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolId` param:
	err = encoder.Encode(obj.PoolId)
	if err != nil {
		return err
	}
	// Serialize `StartTickIndex` param:
	err = encoder.Encode(obj.StartTickIndex)
	if err != nil {
		return err
	}
	// Serialize `Ticks` param:
	err = encoder.Encode(obj.Ticks)
	if err != nil {
		return err
	}
	// Serialize `InitializedTickCount` param:
	err = encoder.Encode(obj.InitializedTickCount)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TickArrayState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TickArrayStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[192 155 85 205 49 249 129 42]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolId`:
	err = decoder.Decode(&obj.PoolId)
	if err != nil {
		return err
	}
	// Deserialize `StartTickIndex`:
	err = decoder.Decode(&obj.StartTickIndex)
	if err != nil {
		return err
	}
	// Deserialize `Ticks`:
	err = decoder.Decode(&obj.Ticks)
	if err != nil {
		return err
	}
	// Deserialize `InitializedTickCount`:
	err = decoder.Decode(&obj.InitializedTickCount)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}
