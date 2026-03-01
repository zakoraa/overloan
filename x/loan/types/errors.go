package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidPrincipal       = errorsmod.Register(ModuleName, 1, "invalid principal")
	ErrInvalidTenor           = errorsmod.Register(ModuleName, 2, "invalid tenor")
	ErrUnauthorized           = errorsmod.Register(ModuleName, 3, "unauthorized")
	ErrInvalidStateTransition = errorsmod.Register(ModuleName, 4, "invalid state transition")
	ErrLoanNotFound           = errorsmod.Register(ModuleName, 5, "loan not found")
	ErrActiveLoanExists       = errorsmod.Register(ModuleName, 6, "active loan already exists")
	ErrInvalidAuthority       = errorsmod.Register(ModuleName, 7, "invalid authority")
	ErrInvalidAddress         = errorsmod.Register(ModuleName, 8, "invalid address")
	ErrInvalidRequest         = errorsmod.Register(ModuleName, 9, "invalid request")
	ErrInvalidCoin            = errorsmod.Register(ModuleName, 10, "invalid coin")
)
