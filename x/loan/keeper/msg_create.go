package keeper

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

// CreateLoan menangani pembuatan loan baru oleh borrower
func (m msgServer) CreateLoan(
	ctx context.Context,
	msg *types.MsgCreateLoan,
) (*types.MsgCreateLoanResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Ambil params modul
	params, err := m.GetParams(sdkCtx)
	if err != nil {
		return nil, err
	}

	// Validasi laz yang dipilih borrower
	if err := m.ValidateLazAuthority(sdkCtx, msg.Laz); err != nil {
		return nil, err
	}

	// Validasi denom harus sesuai settlement denom
	if msg.Principal.Denom != params.SettlementDenom {
		return nil, types.ErrInvalidPrincipal.Wrap("invalid settlement denom")
	}

	// Validasi amount harus positif
	if !msg.Principal.Amount.IsPositive() {
		return nil, types.ErrInvalidPrincipal.Wrap("amount must be positive")
	}

	amount := msg.Principal.Amount

	// Validasi amount dalam range min dan max
	if amount.Uint64() < params.MinLoanAmount ||
		amount.Uint64() > params.MaxLoanAmount {
		return nil, types.ErrInvalidPrincipal.Wrap("amount out of range")
	}

	// Validasi tenor tidak boleh melebihi maksimum
	if msg.TenorMonths > params.MaxTenorMonths {
		return nil, types.ErrInvalidTenor.Wrap("tenor exceeds maximum")
	}

	// Parse alamat borrower dan pastikan valid
	borrowerAddr, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return nil, types.ErrInvalidAddress.Wrap(err.Error())
	}

	// Cegah borrower memiliki lebih dari satu loan aktif
	if m.HasActiveLoan(sdkCtx, borrowerAddr) {
		return nil, types.ErrActiveLoanExists
	}

	// Ambil ID loan berikutnya dari sequence
	loanID, err := m.GetNextLoanID(sdkCtx)
	if err != nil {
		return nil, err
	}

	// Ambil waktu block sebagai timestamp pembuatan
	now := sdkCtx.BlockTime()

	// Entity Loan baru dengan status awal PENDING
	loan := &types.Loan{
		Id:        loanID,
		Borrower:  msg.Borrower,
		Laz:       msg.Laz,
		Principal: msg.Principal,
		Outstanding: &sdk.Coin{
			Denom:  msg.Principal.Denom,
			Amount: sdkmath.ZeroInt(),
		},

		TenorMonths:  msg.TenorMonths,
		Status:       types.LoanStatus_LOAN_STATUS_PENDING,
		CreatedAt:    &now,
		MetadataHash: msg.MetadataHash,
	}

	// Simpan loan ke store utama
	m.SetLoan(sdkCtx, loan)

	// Index loan berdasarkan borrower
	m.SetLoanByBorrower(sdkCtx, borrowerAddr, loanID)

	// Emit event agar bisa ditangkap CLI / indexer
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLoanCreated,
			sdk.NewAttribute(types.AttributeKeyLoanID, fmt.Sprintf("%d", loanID)),
			sdk.NewAttribute(types.AttributeKeyBorrower, msg.Borrower),
			sdk.NewAttribute(types.AttributeKeyLaz, msg.Laz),
			sdk.NewAttribute(types.AttributeKeyPrincipal, msg.Principal.String()),
		),
	)

	return &types.MsgCreateLoanResponse{
		LoanId: loanID,
	}, nil
}
