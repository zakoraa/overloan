<!DOCTYPE html>
<html>

<body>

<h1>Loan Module Concepts</h1>

<p>
The Loan module manages the lifecycle of loans issued to borrowers
under a controlled policy framework.
</p>

<p>
Loans are recorded on-chain and governed by policy authorities defined
through module parameters controlled by governance.
</p>

<h2>Loan Object</h2>

<p>
A <b>Loan</b> represents a borrowing agreement recorded on-chain.
</p>

<table border="1">

<tr>
<th>Field</th>
<th>Description</th>
</tr>

<tr>
<td>id</td>
<td>Unique identifier of the loan</td>
</tr>

<tr>
<td>borrower</td>
<td>Address requesting the loan</td>
</tr>

<tr>
<td>laz</td>
<td>Policy authority responsible for approving or rejecting the loan</td>
</tr>

<tr>
<td>omnibus</td>
<td>Authority responsible for confirming disbursement and recording repayments</td>
</tr>

<tr>
<td>principal</td>
<td>Initial loan amount</td>
</tr>

<tr>
<td>outstanding</td>
<td>Remaining loan balance</td>
</tr>

<tr>
<td>tenor_months</td>
<td>Loan duration in months</td>
</tr>

<tr>
<td>status</td>
<td>Current loan lifecycle state</td>
</tr>

<tr>
<td>created_at</td>
<td>Timestamp when the loan request was created</td>
</tr>

<tr>
<td>approved_at</td>
<td>Timestamp when the loan was approved</td>
</tr>

<tr>
<td>disbursed_at</td>
<td>Timestamp when the loan funds were disbursed</td>
</tr>

<tr>
<td>metadata_hash</td>
<td>Optional hash of external loan documents</td>
</tr>

</table>

<h2>Actors</h2>

<h3>Borrower</h3>

<p>
The borrower submits loan requests and receives loan funds.
</p>

<h3>LAZ Authority</h3>

<p>
Authorized policy actor responsible for approving or rejecting loan requests.
</p>

<h3>Omnibus Authority</h3>

<p>
Responsible for confirming disbursement and recording repayments.
</p>

<h3>Governance</h3>

<p>
Governance controls module parameters including authorized policy actors
and loan limits.
</p>

<h2>Loan Status</h2>

<p>
Loans transition through the following lifecycle states.
</p>

<table border="1">

<tr>
<th>Status</th>
<th>Description</th>
</tr>

<tr>
<td>LOAN_STATUS_PENDING</td>
<td>Loan request submitted by borrower</td>
</tr>

<tr>
<td>LOAN_STATUS_APPROVED</td>
<td>Loan approved by LAZ authority</td>
</tr>

<tr>
<td>LOAN_STATUS_DISBURSED</td>
<td>Loan funds disbursed to borrower</td>
</tr>

<tr>
<td>LOAN_STATUS_REPAID</td>
<td>Loan fully repaid</td>
</tr>

<tr>
<td>LOAN_STATUS_REJECTED</td>
<td>Loan rejected by LAZ authority</td>
</tr>

<tr>
<td>LOAN_STATUS_CANCELLED</td>
<td>Loan disbursement rejected or cancelled</td>
</tr>

</table>

</body>
</html>
