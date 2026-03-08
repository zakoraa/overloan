package types

// CanApprove memastikan loan dalam status yang valid untuk disetujui
func CanApprove(l *Loan) error {

	// Loan hanya boleh disetujui jika masih berstatus PENDING
	if l.Status != LoanStatus_LOAN_STATUS_PENDING {
		return ErrInvalidStateTransition.Wrap("loan must be pending")
	}

	return nil
}

// CanReject memastikan loan dalam status yang valid untuk ditolak
func CanReject(l *Loan) error {
	if l.Status != LoanStatus_LOAN_STATUS_PENDING {
		return ErrInvalidStateTransition.
			Wrap("loan must be pending to reject")
	}
	return nil
}

// CanDisburse memastikan loan dalam status yang valid untuk dicairkan
func CanDisburse(l *Loan) error {

	// Loan hanya boleh dicairkan jika sudah berstatus APPROVED
	if l.Status != LoanStatus_LOAN_STATUS_APPROVED {
		return ErrInvalidStateTransition.Wrap("loan must be approved")
	}

	return nil
}

// CanRejectDisburse memastikan loan dalam status yang valid
// untuk menolak pencairan dana oleh omnibus
func CanRejectDisburse(l *Loan) error {

	// Penolakan pencairan hanya boleh jika loan masih APPROVED
	if l.Status != LoanStatus_LOAN_STATUS_APPROVED {
		return ErrInvalidStateTransition.
			Wrap("loan must be approved to reject disbursement")
	}

	return nil
}

// CanRepay memastikan loan dalam status yang valid untuk borrower melakukan pembayaran
func CanRepay(l *Loan) error {

	// Repayment hanya boleh jika dana sudah dicairkan
	if l.Status != LoanStatus_LOAN_STATUS_DISBURSED {
		return ErrInvalidStateTransition.
			Wrap("loan must be disbursed to repay")
	}

	// Tidak boleh repay jika outstanding sudah habis
	if l.Outstanding.IsZero() {
		return ErrInvalidStateTransition.
			Wrap("loan already fully repaid")
	}

	return nil
}

// IsActiveStatus menentukan apakah loan masih dianggap aktif secara bisnis
func IsActiveStatus(status LoanStatus) bool {

	switch status {
	case LoanStatus_LOAN_STATUS_PENDING,
		LoanStatus_LOAN_STATUS_APPROVED,
		LoanStatus_LOAN_STATUS_DISBURSED:
		return true
	default:
		return false
	}
}
