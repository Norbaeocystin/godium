// Code Generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package amm_v3

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetRewardParams is the `setRewardParams` instruction.
type SetRewardParams struct {
	RewardIndex           *uint8
	EmissionsPerSecondX64 *ag_binary.Uint128
	OpenTime              *uint64
	EndTime               *uint64

	// [0] = [SIGNER] authority
	//
	// [1] = [] ammConfig
	//
	// [2] = [WRITE] poolState
	//
	// [3] = [] operationState
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetRewardParamsInstructionBuilder creates a new `SetRewardParams` instruction builder.
func NewSetRewardParamsInstructionBuilder() *SetRewardParams {
	nd := &SetRewardParams{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetRewardIndex sets the "rewardIndex" parameter.
func (inst *SetRewardParams) SetRewardIndex(rewardIndex uint8) *SetRewardParams {
	inst.RewardIndex = &rewardIndex
	return inst
}

// SetEmissionsPerSecondX64 sets the "emissionsPerSecondX64" parameter.
func (inst *SetRewardParams) SetEmissionsPerSecondX64(emissionsPerSecondX64 ag_binary.Uint128) *SetRewardParams {
	inst.EmissionsPerSecondX64 = &emissionsPerSecondX64
	return inst
}

// SetOpenTime sets the "openTime" parameter.
func (inst *SetRewardParams) SetOpenTime(openTime uint64) *SetRewardParams {
	inst.OpenTime = &openTime
	return inst
}

// SetEndTime sets the "endTime" parameter.
func (inst *SetRewardParams) SetEndTime(endTime uint64) *SetRewardParams {
	inst.EndTime = &endTime
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetRewardParams) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetRewardParams {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetRewardParams) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmConfigAccount sets the "ammConfig" account.
func (inst *SetRewardParams) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *SetRewardParams {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ammConfig)
	return inst
}

// GetAmmConfigAccount gets the "ammConfig" account.
func (inst *SetRewardParams) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolStateAccount sets the "poolState" account.
func (inst *SetRewardParams) SetPoolStateAccount(poolState ag_solanago.PublicKey) *SetRewardParams {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
func (inst *SetRewardParams) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOperationStateAccount sets the "operationState" account.
func (inst *SetRewardParams) SetOperationStateAccount(operationState ag_solanago.PublicKey) *SetRewardParams {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(operationState)
	return inst
}

// GetOperationStateAccount gets the "operationState" account.
func (inst *SetRewardParams) GetOperationStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SetRewardParams) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetRewardParams,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetRewardParams) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetRewardParams) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardIndex == nil {
			return errors.New("RewardIndex parameter is not set")
		}
		if inst.EmissionsPerSecondX64 == nil {
			return errors.New("EmissionsPerSecondX64 parameter is not set")
		}
		if inst.OpenTime == nil {
			return errors.New("OpenTime parameter is not set")
		}
		if inst.EndTime == nil {
			return errors.New("EndTime parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OperationState is not set")
		}
	}
	return nil
}

func (inst *SetRewardParams) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetRewardParams")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          RewardIndex", *inst.RewardIndex))
						paramsBranch.Child(ag_format.Param("EmissionsPerSecondX64", *inst.EmissionsPerSecondX64))
						paramsBranch.Child(ag_format.Param("             OpenTime", *inst.OpenTime))
						paramsBranch.Child(ag_format.Param("              EndTime", *inst.EndTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     ammConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     poolState", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("operationState", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SetRewardParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardIndex` param:
	err = encoder.Encode(obj.RewardIndex)
	if err != nil {
		return err
	}
	// Serialize `EmissionsPerSecondX64` param:
	err = encoder.Encode(obj.EmissionsPerSecondX64)
	if err != nil {
		return err
	}
	// Serialize `OpenTime` param:
	err = encoder.Encode(obj.OpenTime)
	if err != nil {
		return err
	}
	// Serialize `EndTime` param:
	err = encoder.Encode(obj.EndTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetRewardParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardIndex`:
	err = decoder.Decode(&obj.RewardIndex)
	if err != nil {
		return err
	}
	// Deserialize `EmissionsPerSecondX64`:
	err = decoder.Decode(&obj.EmissionsPerSecondX64)
	if err != nil {
		return err
	}
	// Deserialize `OpenTime`:
	err = decoder.Decode(&obj.OpenTime)
	if err != nil {
		return err
	}
	// Deserialize `EndTime`:
	err = decoder.Decode(&obj.EndTime)
	if err != nil {
		return err
	}
	return nil
}

// NewSetRewardParamsInstruction declares a new SetRewardParams instruction with the provided parameters and accounts.
func NewSetRewardParamsInstruction(
	// Parameters:
	rewardIndex uint8,
	emissionsPerSecondX64 ag_binary.Uint128,
	openTime uint64,
	endTime uint64,
	// Accounts:
	authority ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	operationState ag_solanago.PublicKey) *SetRewardParams {
	return NewSetRewardParamsInstructionBuilder().
		SetRewardIndex(rewardIndex).
		SetEmissionsPerSecondX64(emissionsPerSecondX64).
		SetOpenTime(openTime).
		SetEndTime(endTime).
		SetAuthorityAccount(authority).
		SetAmmConfigAccount(ammConfig).
		SetPoolStateAccount(poolState).
		SetOperationStateAccount(operationState)
}
