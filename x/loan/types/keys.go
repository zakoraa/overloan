package types

const (
	// ModuleName mendefinisikan nama module
	ModuleName = "loan"

	// StoreKey nilai default dari store key
	StoreKey = ModuleName

	// RouterKey pesan route
	RouterKey = ModuleName
)

var (
	EventTypeLoanCreated   = "loan_created"
	EventTypeLoanApproved  = "loan_approved"
	EventTypeLoanRejected  = "loan_rejected"
	EventTypeLoanDisbursed = "loan_disbursed"
	EventTypeLoanRepaid    = "loan_repaid"

	AttributeKeyLoanID    = "loan_id"
	AttributeKeyBorrower  = "borrower"
	AttributeKeyPrincipal = "principal"
	AttributeKeyAuthority = "authority"
	AttributeKeyAmount    = "amount"

	LoanKeyPrefix        = []byte{0x11}
	LoanByBorrowerPrefix = []byte{0x12}
	LoanIDKey            = []byte{0x13}
)
