package loan

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

// AutoCLIOptions mengatur konfigurasi CLI otomatis
// berdasarkan service Query dan Msg di protobuf.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{

		// ---------------------------
		// Query Commands (read-only)
		// ---------------------------
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{

				// Query parameter modul
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Menampilkan parameter modul loan",
				},

				// Query satu loan berdasarkan ID
				{
					RpcMethod: "Loan",
					Use:       "loan [loan-id]",
					Short:     "Menampilkan detail satu loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},

				// Query semua loan (dengan pagination)
				{
					RpcMethod: "Loans",
					Use:       "loans",
					Short:     "Menampilkan daftar semua loan",
				},

				// Query loan berdasarkan borrower
				{
					RpcMethod: "LoansByBorrower",
					Use:       "loans-by-borrower [borrower]",
					Short:     "Menampilkan daftar loan milik borrower",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "borrower"},
					},
				},
			},
		},

		// ---------------------------
		// Transaction Commands (state changing)
		// ---------------------------
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,

			// true agar tetap bisa gabung dengan custom command kalau ada
			EnhanceCustomCommand: true,

			RpcCommandOptions: []*autocliv1.RpcCommandOptions{

				// Update parameter (biasanya authority/gov only)
				{
					RpcMethod: "UpdateParams",
					Use:       "update-params",
					Short:     "Mengubah parameter modul loan (authority only)",
				},

				// Membuat loan baru
				{
					RpcMethod: "CreateLoan",
					Use:       "create-loan",
					Short:     "Mengajukan permohonan loan baru",
				},

				// Menyetujui loan
				{
					RpcMethod: "ApproveLoan",
					Use:       "approve-loan [loan-id]",
					Short:     "Menyetujui loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},

				// Menolak loan
				{
					RpcMethod: "RejectLoan",
					Use:       "reject-loan [loan-id]",
					Short:     "Menolak loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},

				// Membayar loan
				{
					RpcMethod: "RepayLoan",
					Use:       "repay-loan [loan-id]",
					Short:     "Melakukan pembayaran loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},

				// Konfirmasi pencairan dana
				{
					RpcMethod: "ConfirmDisbursement",
					Use:       "confirm-disbursement [loan-id]",
					Short:     "Konfirmasi pencairan loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
			},
		},
	}
}
