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
	ParamsKey = []byte("loan_params")

	EventTypeLoanCreated          = "loan_created"
	EventTypeLoanApproved         = "loan_approved"
	EventTypeLoanRejected         = "loan_rejected"
	EventTypeLoanDisbursed        = "loan_disbursed"
	EventTypeLoanDisburseRejected = "loan_disburse_rejected"
	EventTypeLoanRepaid           = "loan_repaid"
	EventTypeUpdateParams         = "loan_update_params"

	AttributeKeyLoanID    = "loan_id"
	AttributeKeyBorrower  = "borrower"
	AttributeKeyLaz       = "laz"
	AttributeKeyOmnibus   = "omnibus"
	AttributeKeyAuthority = "authority"
	AttributeKeyPrincipal = "principal"
	AttributeKeyAmount    = "amount"
	AttributeKeyReason    = "reason"

	LoanKeyPrefix        = []byte{0x11}
	LoanByBorrowerPrefix = []byte{0x12}
	LoanIDKey            = []byte{0x13}
)
