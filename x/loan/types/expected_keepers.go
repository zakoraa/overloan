package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper mendefinisikan interface minimal yang dibutuhkan
// dari modul x/bank oleh modul loan
type BankKeeper interface {

	// GetBalance mengembalikan saldo suatu denom untuk alamat tertentu
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin

	// GetSupply mengembalikan total supply suatu denom
	GetSupply(ctx context.Context, denom string) sdk.Coin

	// SendCoins melakukan transfer koin antar akun biasa
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error

	// MintCoins mencetak koin baru ke module account.
	// Module account harus memiliki permission "minter"
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	// BurnCoins membakar koin dari module account.
	// Module account harus memiliki permission "burner"
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	// SendCoinsFromModuleToAccount mentransfer koin
	// dari module account ke akun user.
	SendCoinsFromModuleToAccount(
		ctx context.Context,
		senderModule string,
		recipientAddr sdk.AccAddress,
		amt sdk.Coins,
	) error

	// SendCoinsFromAccountToModule mentransfer koin
	// dari akun user ke module account.
	SendCoinsFromAccountToModule(
		ctx context.Context,
		senderAddr sdk.AccAddress,
		recipientModule string,
		amt sdk.Coins,
	) error
}

// AccountKeeper mendefinisikan interface minimal yang dibutuhkan
// dari modul x/auth
type AccountKeeper interface {
	// GetModuleAddress mengembalikan alamat akun modul
	// berdasarkan nama modul
	GetModuleAddress(moduleName string) sdk.AccAddress
}