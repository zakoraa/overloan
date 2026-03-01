package keeper

import (
	"context"
	"fmt"

	basev1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateLoan menangani pembuatan loan baru oleh borrower
func (m msgServer) CreateLoan(
	ctx context.Context,
	msg *loanv1.MsgCreateLoan,
) (*loanv1.MsgCreateLoanResponse, error) {

	// Unwrap context ke SDK context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Ambil parameter module untuk validasi stateful
	params := m.GetParams(sdkCtx)

	// Validasi denom harus sesuai settlement denom
	if msg.Principal.Denom != params.SettlementDenom {
		return nil, types.ErrInvalidPrincipal.Wrap("invalid settlement denom")
	}

	// Parse amount string menjadi sdkmath.Int
	amount, ok := sdkmath.NewIntFromString(msg.Principal.Amount)
	if !ok {
		return nil, types.ErrInvalidPrincipal.Wrap("invalid amount format")
	}

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
	loanID := m.GetNextLoanID(sdkCtx)

	// Ambil waktu block sebagai timestamp pembuatan
	now := sdkCtx.BlockTime()

	// entity Loan baru dengan status awal PENDING
	loan := &loanv1.Loan{
		Id:       loanID,
		Borrower: msg.Borrower,
		Principal: &basev1beta1.Coin{
			Denom:  msg.Principal.Denom,
			Amount: msg.Principal.Amount,
		},
		Outstanding: &basev1beta1.Coin{
			Denom:  msg.Principal.Denom,
			Amount: "0",
		},
		TenorMonths:  msg.TenorMonths,
		Status:       loanv1.LoanStatus_LOAN_STATUS_PENDING,
		CreatedAt:    timestamppb.New(now),
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
			sdk.NewAttribute(types.AttributeKeyPrincipal, msg.Principal.Amount),
		),
	)

	// Kembalikan response berisi ID loan
	return &loanv1.MsgCreateLoanResponse{
		LoanId: loanID,
	}, nil
}
