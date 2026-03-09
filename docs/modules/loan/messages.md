<!DOCTYPE html>
<html>

<body>

<h1>Loan Module Messages</h1>

<p>
The Loan module defines a set of transaction messages used to manage
the lifecycle of loans.
</p>

<h2>MsgUpdateParams</h2>

<p>
Updates module parameters through governance.
Only the configured authority (typically the governance module)
can execute this message.
</p>


<h2>MsgCreateLoan</h2>

<p>
Creates a new loan request with status <b>LOAN_STATUS_PENDING</b>.
This message is signed by the borrower.
</p>


<h2>MsgApproveLoan</h2>

<p>
Approves a loan that is currently in <b>PENDING</b> status.
When approved, the loan status becomes <b>LOAN_STATUS_APPROVED</b>.
This message is signed by an authorized LAZ account.
</p>


<h2>MsgRejectLoan</h2>

<p>
Rejects a loan request that is currently in <b>PENDING</b> status.
The loan status becomes <b>LOAN_STATUS_REJECTED</b>.
This message is signed by an authorized LAZ account.
</p>


<h2>MsgConfirmDisbursement</h2>

<p>
Confirms that loan funds have been disbursed to the borrower.
The loan status becomes <b>LOAN_STATUS_DISBURSED</b>.
This message is signed by an authorized omnibus account.
</p>


<h2>MsgRejectDisbursement</h2>

<p>
Rejects a previously approved loan before funds are disbursed.
The loan status becomes <b>LOAN_STATUS_CANCELLED</b>.
This message is signed by an authorized omnibus account.
</p>


<h2>MsgRepayLoan</h2>

<p>
Records a repayment toward the loan.
The repayment amount reduces the <b>outstanding</b> balance.
When the outstanding amount reaches zero, the loan status becomes
<b>LOAN_STATUS_REPAID</b>.
This message is signed by an authorized omnibus account.
</p>

</body>
</html>
