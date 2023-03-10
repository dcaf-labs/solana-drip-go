// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package drip

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type InitializeVaultPeriodParams struct {
	PeriodId uint64
}

func (obj InitializeVaultPeriodParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PeriodId` param:
	err = encoder.Encode(obj.PeriodId)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeVaultPeriodParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PeriodId`:
	err = decoder.Decode(&obj.PeriodId)
	if err != nil {
		return err
	}
	return nil
}

type InitVaultProtoConfigParams struct {
	Granularity             uint64
	TokenADripTriggerSpread uint16
	TokenBWithdrawalSpread  uint16
	Admin                   ag_solanago.PublicKey
}

func (obj InitVaultProtoConfigParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Granularity` param:
	err = encoder.Encode(obj.Granularity)
	if err != nil {
		return err
	}
	// Serialize `TokenADripTriggerSpread` param:
	err = encoder.Encode(obj.TokenADripTriggerSpread)
	if err != nil {
		return err
	}
	// Serialize `TokenBWithdrawalSpread` param:
	err = encoder.Encode(obj.TokenBWithdrawalSpread)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitVaultProtoConfigParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Granularity`:
	err = decoder.Decode(&obj.Granularity)
	if err != nil {
		return err
	}
	// Deserialize `TokenADripTriggerSpread`:
	err = decoder.Decode(&obj.TokenADripTriggerSpread)
	if err != nil {
		return err
	}
	// Deserialize `TokenBWithdrawalSpread`:
	err = decoder.Decode(&obj.TokenBWithdrawalSpread)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

type InitializeVaultParams struct {
	MaxSlippageBps   uint16
	WhitelistedSwaps []ag_solanago.PublicKey
}

func (obj InitializeVaultParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MaxSlippageBps` param:
	err = encoder.Encode(obj.MaxSlippageBps)
	if err != nil {
		return err
	}
	// Serialize `WhitelistedSwaps` param:
	err = encoder.Encode(obj.WhitelistedSwaps)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeVaultParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MaxSlippageBps`:
	err = decoder.Decode(&obj.MaxSlippageBps)
	if err != nil {
		return err
	}
	// Deserialize `WhitelistedSwaps`:
	err = decoder.Decode(&obj.WhitelistedSwaps)
	if err != nil {
		return err
	}
	return nil
}

type DepositParams struct {
	TokenADepositAmount uint64
	NumberOfSwaps       uint64
}

func (obj DepositParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenADepositAmount` param:
	err = encoder.Encode(obj.TokenADepositAmount)
	if err != nil {
		return err
	}
	// Serialize `NumberOfSwaps` param:
	err = encoder.Encode(obj.NumberOfSwaps)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DepositParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenADepositAmount`:
	err = decoder.Decode(&obj.TokenADepositAmount)
	if err != nil {
		return err
	}
	// Deserialize `NumberOfSwaps`:
	err = decoder.Decode(&obj.NumberOfSwaps)
	if err != nil {
		return err
	}
	return nil
}
