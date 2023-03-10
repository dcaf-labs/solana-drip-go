// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package drip

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DepositWithMetadata is the `depositWithMetadata` instruction.
type DepositWithMetadata struct {
	Params *DepositParams

	// [0] = [WRITE] vault
	//
	// [1] = [WRITE] vaultPeriodEnd
	//
	// [2] = [WRITE] userPosition
	//
	// [3] = [] tokenAMint
	//
	// [4] = [WRITE, SIGNER] userPositionNftMint
	//
	// [5] = [WRITE] vaultTokenAAccount
	//
	// [6] = [WRITE] userTokenAAccount
	//
	// [7] = [WRITE] userPositionNftAccount
	//
	// [8] = [WRITE, SIGNER] depositor
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] associatedTokenProgram
	//
	// [11] = [] rent
	//
	// [12] = [] systemProgram
	//
	// [13] = [WRITE] positionMetadataAccount
	//
	// [14] = [] metadataProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type DepositWithMetadataAccounts struct {
	Vault                   ag_solanago.PublicKey
	VaultPeriodEnd          ag_solanago.PublicKey
	UserPosition            ag_solanago.PublicKey
	TokenAMint              ag_solanago.PublicKey
	UserPositionNftMint     ag_solanago.PublicKey
	VaultTokenAAccount      ag_solanago.PublicKey
	UserTokenAAccount       ag_solanago.PublicKey
	UserPositionNftAccount  ag_solanago.PublicKey
	Depositor               ag_solanago.PublicKey
	TokenProgram            ag_solanago.PublicKey
	AssociatedTokenProgram  ag_solanago.PublicKey
	Rent                    ag_solanago.PublicKey
	SystemProgram           ag_solanago.PublicKey
	PositionMetadataAccount ag_solanago.PublicKey
	MetadataProgram         ag_solanago.PublicKey
}

// NewDepositWithMetadataInstructionBuilder creates a new `DepositWithMetadata` instruction builder.
func NewDepositWithMetadataInstructionBuilder() *DepositWithMetadata {
	nd := &DepositWithMetadata{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

func (inst *DepositWithMetadata) GetDepositWithMetadataAccounts() *DepositWithMetadataAccounts {
	res := &DepositWithMetadataAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.Vault = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.VaultPeriodEnd = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.UserPosition = inst.AccountMetaSlice[2].PublicKey
	}
	if inst.AccountMetaSlice[3] != nil {
		res.TokenAMint = inst.AccountMetaSlice[3].PublicKey
	}
	if inst.AccountMetaSlice[4] != nil {
		res.UserPositionNftMint = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.VaultTokenAAccount = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.UserTokenAAccount = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.UserPositionNftAccount = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.Depositor = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.TokenProgram = inst.AccountMetaSlice[9].PublicKey
	}
	if inst.AccountMetaSlice[10] != nil {
		res.AssociatedTokenProgram = inst.AccountMetaSlice[10].PublicKey
	}
	if inst.AccountMetaSlice[11] != nil {
		res.Rent = inst.AccountMetaSlice[11].PublicKey
	}
	if inst.AccountMetaSlice[12] != nil {
		res.SystemProgram = inst.AccountMetaSlice[12].PublicKey
	}
	if inst.AccountMetaSlice[13] != nil {
		res.PositionMetadataAccount = inst.AccountMetaSlice[13].PublicKey
	}
	if inst.AccountMetaSlice[14] != nil {
		res.MetadataProgram = inst.AccountMetaSlice[14].PublicKey
	}
	return res
}

// SetParams sets the "params" parameter.
func (inst *DepositWithMetadata) SetParams(params DepositParams) *DepositWithMetadata {
	inst.Params = &params
	return inst
}

// SetVaultAccount sets the "vault" account.
func (inst *DepositWithMetadata) SetVaultAccount(vault ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *DepositWithMetadata) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetVaultPeriodEndAccount sets the "vaultPeriodEnd" account.
func (inst *DepositWithMetadata) SetVaultPeriodEndAccount(vaultPeriodEnd ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(vaultPeriodEnd).WRITE()
	return inst
}

// GetVaultPeriodEndAccount gets the "vaultPeriodEnd" account.
func (inst *DepositWithMetadata) GetVaultPeriodEndAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserPositionAccount sets the "userPosition" account.
func (inst *DepositWithMetadata) SetUserPositionAccount(userPosition ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userPosition).WRITE()
	return inst
}

// GetUserPositionAccount gets the "userPosition" account.
func (inst *DepositWithMetadata) GetUserPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenAMintAccount sets the "tokenAMint" account.
func (inst *DepositWithMetadata) SetTokenAMintAccount(tokenAMint ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenAMint)
	return inst
}

// GetTokenAMintAccount gets the "tokenAMint" account.
func (inst *DepositWithMetadata) GetTokenAMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserPositionNftMintAccount sets the "userPositionNftMint" account.
func (inst *DepositWithMetadata) SetUserPositionNftMintAccount(userPositionNftMint ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userPositionNftMint).WRITE().SIGNER()
	return inst
}

// GetUserPositionNftMintAccount gets the "userPositionNftMint" account.
func (inst *DepositWithMetadata) GetUserPositionNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultTokenAAccountAccount sets the "vaultTokenAAccount" account.
func (inst *DepositWithMetadata) SetVaultTokenAAccountAccount(vaultTokenAAccount ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultTokenAAccount).WRITE()
	return inst
}

// GetVaultTokenAAccountAccount gets the "vaultTokenAAccount" account.
func (inst *DepositWithMetadata) GetVaultTokenAAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUserTokenAAccountAccount sets the "userTokenAAccount" account.
func (inst *DepositWithMetadata) SetUserTokenAAccountAccount(userTokenAAccount ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(userTokenAAccount).WRITE()
	return inst
}

// GetUserTokenAAccountAccount gets the "userTokenAAccount" account.
func (inst *DepositWithMetadata) GetUserTokenAAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetUserPositionNftAccountAccount sets the "userPositionNftAccount" account.
func (inst *DepositWithMetadata) SetUserPositionNftAccountAccount(userPositionNftAccount ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(userPositionNftAccount).WRITE()
	return inst
}

// GetUserPositionNftAccountAccount gets the "userPositionNftAccount" account.
func (inst *DepositWithMetadata) GetUserPositionNftAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDepositorAccount sets the "depositor" account.
func (inst *DepositWithMetadata) SetDepositorAccount(depositor ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(depositor).WRITE().SIGNER()
	return inst
}

// GetDepositorAccount gets the "depositor" account.
func (inst *DepositWithMetadata) GetDepositorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DepositWithMetadata) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DepositWithMetadata) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *DepositWithMetadata) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *DepositWithMetadata) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
func (inst *DepositWithMetadata) SetRentAccount(rent ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *DepositWithMetadata) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *DepositWithMetadata) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *DepositWithMetadata) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetPositionMetadataAccountAccount sets the "positionMetadataAccount" account.
func (inst *DepositWithMetadata) SetPositionMetadataAccountAccount(positionMetadataAccount ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(positionMetadataAccount).WRITE()
	return inst
}

// GetPositionMetadataAccountAccount gets the "positionMetadataAccount" account.
func (inst *DepositWithMetadata) GetPositionMetadataAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetMetadataProgramAccount sets the "metadataProgram" account.
func (inst *DepositWithMetadata) SetMetadataProgramAccount(metadataProgram ag_solanago.PublicKey) *DepositWithMetadata {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(metadataProgram)
	return inst
}

// GetMetadataProgramAccount gets the "metadataProgram" account.
func (inst *DepositWithMetadata) GetMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst DepositWithMetadata) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DepositWithMetadata,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DepositWithMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositWithMetadata) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.VaultPeriodEnd is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserPosition is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenAMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserPositionNftMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultTokenAAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.UserTokenAAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.UserPositionNftAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Depositor is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.PositionMetadataAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.MetadataProgram is not set")
		}
	}
	return nil
}

func (inst *DepositWithMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositWithMetadata")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 vault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        vaultPeriodEnd", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          userPosition", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            tokenAMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   userPositionNftMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           vaultTokenA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            userTokenA", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       userPositionNft", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("             depositor", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("      positionMetadata", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("       metadataProgram", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj DepositWithMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DepositWithMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewDepositWithMetadataInstruction declares a new DepositWithMetadata instruction with the provided parameters and accounts.
func NewDepositWithMetadataInstruction(
	// Parameters:
	params DepositParams,
	// Accounts:
	vault ag_solanago.PublicKey,
	vaultPeriodEnd ag_solanago.PublicKey,
	userPosition ag_solanago.PublicKey,
	tokenAMint ag_solanago.PublicKey,
	userPositionNftMint ag_solanago.PublicKey,
	vaultTokenAAccount ag_solanago.PublicKey,
	userTokenAAccount ag_solanago.PublicKey,
	userPositionNftAccount ag_solanago.PublicKey,
	depositor ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	positionMetadataAccount ag_solanago.PublicKey,
	metadataProgram ag_solanago.PublicKey) *DepositWithMetadata {
	return NewDepositWithMetadataInstructionBuilder().
		SetParams(params).
		SetVaultAccount(vault).
		SetVaultPeriodEndAccount(vaultPeriodEnd).
		SetUserPositionAccount(userPosition).
		SetTokenAMintAccount(tokenAMint).
		SetUserPositionNftMintAccount(userPositionNftMint).
		SetVaultTokenAAccountAccount(vaultTokenAAccount).
		SetUserTokenAAccountAccount(userTokenAAccount).
		SetUserPositionNftAccountAccount(userPositionNftAccount).
		SetDepositorAccount(depositor).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram).
		SetPositionMetadataAccountAccount(positionMetadataAccount).
		SetMetadataProgramAccount(metadataProgram)
}
