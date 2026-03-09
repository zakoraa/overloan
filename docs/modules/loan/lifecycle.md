<!DOCTYPE html>
<html>

<body>

<h1>Loan Lifecycle</h1>

<p>
Loans follow a controlled lifecycle managed by borrower, policy authorities,
and omnibus authorities.
</p>

<h2>Step 1 — Create Loan</h2>

<p>
A borrower submits a loan request using <b>MsgCreateLoan</b>.
</p>

<p>
The loan is created with status:
</p>

<p><b>LOAN_STATUS_PENDING</b></p>


<h2>Step 2 — Loan Approval</h2>

<p>
An authorized LAZ authority reviews the loan request.
</p>

<p>
Possible outcomes:
</p>

<ul>

<li>
<b>Approve Loan</b><br>
MsgApproveLoan<br>
Status becomes <b>LOAN_STATUS_APPROVED</b>
</li>

<li>
<b>Reject Loan</b><br>
MsgRejectLoan<br>
Status becomes <b>LOAN_STATUS_REJECTED</b>
</li>

</ul>


<h2>Step 3 — Disbursement</h2>

<p>
After approval, an authorized omnibus authority processes disbursement.
</p>

<ul>

<li>
<b>Confirm Disbursement</b><br>
MsgConfirmDisbursement<br>
Status becomes <b>LOAN_STATUS_DISBURSED</b>
</li>

<li>
<b>Reject Disbursement</b><br>
MsgRejectDisbursement<br>
Status becomes <b>LOAN_STATUS_CANCELLED</b>
</li>

</ul>


<h2>Step 4 — Repayment</h2>

<p>
Repayments are recorded using <b>MsgRepayLoan</b>.
</p>

<p>
Each repayment reduces the <b>outstanding</b> loan balance.
</p>

<ul>

<li>
Outstanding balance decreases after each repayment.
</li>

<li>
When outstanding reaches zero, the loan status becomes:
<b>LOAN_STATUS_REPAID</b>
</li>

</ul>


<h2>Lifecycle Summary</h2>

<table border="1">

<tr>
<th>State</th>
<th>Description</th>
</tr>

<tr>
<td>PENDING</td>
<td>Loan request submitted</td>
</tr>

<tr>
<td>APPROVED</td>
<td>Loan approved by LAZ authority</td>
</tr>

<tr>
<td>REJECTED</td>
<td>Loan rejected by LAZ authority</td>
</tr>

<tr>
<td>DISBURSED</td>
<td>Loan funds disbursed by omnibus</td>
</tr>

<tr>
<td>CANCELLED</td>
<td>Disbursement rejected</td>
</tr>

<tr>
<td>REPAID</td>
<td>Loan fully repaid</td>
</tr>

</table>

</body>
</html>
