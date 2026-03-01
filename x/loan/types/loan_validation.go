package types

import loanv1 "github.com/cosmos/cosmos-sdk/api/overloan/loan/v1"

// CanApprove memastikan loan dalam status yang valid untuk disetujui
func CanApprove(l *loanv1.Loan) error {

	// Loan hanya boleh disetujui jika masih berstatus PENDING
	if l.Status != loanv1.LoanStatus_LOAN_STATUS_PENDING {
		return ErrInvalidStateTransition.Wrap("loan must be pending")
	}

	return nil
}

func CanReject(l *loanv1.Loan) error {
	if l.Status != loanv1.LoanStatus_LOAN_STATUS_PENDING {
		return ErrInvalidStateTransition.
			Wrap("loan must be pending to reject")
	}
	return nil
}

// CanDisburse memastikan loan dalam status yang valid untuk dicairkan
func CanDisburse(l *loanv1.Loan) error {

	// Loan hanya boleh dicairkan jika sudah berstatus APPROVED
	if l.Status != loanv1.LoanStatus_LOAN_STATUS_APPROVED {
		return ErrInvalidStateTransition.Wrap("loan must be approved")
	}

	return nil
}

// IsActiveStatus menentukan apakah loan masih dianggap aktif secara bisnis
func IsActiveStatus(status loanv1.LoanStatus) bool {

	switch status {
	case loanv1.LoanStatus_LOAN_STATUS_PENDING,
		loanv1.LoanStatus_LOAN_STATUS_APPROVED,
		loanv1.LoanStatus_LOAN_STATUS_DISBURSED:
		return true
	default:
		return false
	}
}
