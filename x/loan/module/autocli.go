package loan

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{

		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Displaying loan module parameters",
				},
				{
					RpcMethod: "Loan",
					Use:       "loan [loan-id]",
					Short:     "Displaying details of one loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
				{
					RpcMethod: "Loans",
					Use:       "loans",
					Short:     "Displaying a list of all loans",
				},
				{
					RpcMethod: "LoansByBorrower",
					Use:       "loans-by-borrower [borrower]",
					Short:     "Displaying a list of loans belonging to the borrower",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "borrower"},
					},
				},
			},
		},

		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Use:       "update-params",
					Short:     "Changing loan module parameters (authority only)",
				},
				{
					RpcMethod: "CreateLoan",
					Use:       "create-loan",
					Short:     "Apply for a new loan",
				},
				{
					RpcMethod: "ApproveLoan",
					Use:       "approve-loan [loan-id]",
					Short:     "Approve the loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
				{
					RpcMethod: "RejectLoan",
					Use:       "reject-loan [loan-id]",
					Short:     "Rejecting a loan",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
				{
					RpcMethod: "RepayLoan",
					Use:       "repay-loan [loan-id]",
					Short:     "Making loan payments",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
				{
					RpcMethod: "ConfirmDisbursement",
					Use:       "confirm-disbursement [loan-id]",
					Short:     "Confirmation of loan disbursement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "loan_id"},
					},
				},
			},
		},
	}
}
