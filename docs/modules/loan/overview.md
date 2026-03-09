<!DOCTYPE html>

<html>

<body>

<h1>Loan Module</h1>

<p>
The Loan module provides a lending lifecycle management system implemented as a Cosmos SDK module.
It allows borrowers to request loans and authorized policy actors to approve, reject, disburse, cancel,
and record repayments of loans.
</p>

<p>
The module manages the lifecycle of loans stored on-chain and enforces governance controlled parameters
for lending operations.
</p>

<h2>Purpose</h2>

<ul>
<li>Enable controlled lending within a blockchain environment</li>
<li>Track outstanding loan obligations</li>
<li>Provide governance controlled parameters</li>
<li>Allow transparent lifecycle tracking</li>
<li>Allow loans to be queried through gRPC endpoints</li>
</ul>

<h2>Key Roles</h2>

<table border="1">

<tr>
<th>Role</th>
<th>Description</th>
</tr>

<tr>
<td>Borrower</td>
<td>User requesting a loan</td>
</tr>

<tr>
<td>LAZ Authority</td>
<td>Policy authority approving or rejecting loan requests</td>
</tr>

<tr>
<td>Omnibus Authority</td>
<td>Authority confirming disbursement and recording repayments</td>
</tr>

<tr>
<td>Governance</td>
<td>Controls module parameters</td>
</tr>

</table>

<h2>Loan Lifecycle</h2>

<ol>

<li>Borrower submits a loan request</li>

<li>
LAZ authority reviews the request and either:
<ul>
<li>Approves the loan</li>
<li>Rejects the loan</li>
</ul>
</li>

<li>
Omnibus authority confirms that funds are disbursed
</li>

<li>
Loan repayments are recorded by omnibus authority
</li>

<li>
Outstanding balance is reduced until fully repaid
</li>

</ol>

<h2>Loan Status</h2>

<table border="1">

<tr>
<th>Status</th>
<th>Description</th>
</tr>

<tr>
<td>PENDING</td>
<td>Loan request has been submitted</td>
</tr>

<tr>
<td>APPROVED</td>
<td>Loan approved by LAZ authority</td>
</tr>

<tr>
<td>DISBURSED</td>
<td>Funds have been disbursed to borrower</td>
</tr>

<tr>
<td>REPAID</td>
<td>Loan fully repaid</td>
</tr>

<tr>
<td>REJECTED</td>
<td>Loan rejected by LAZ authority</td>
</tr>

<tr>
<td>CANCELLED</td>
<td>Disbursement rejected or cancelled</td>
</tr>

</table>

<h2>Loan Data Structure</h2>

<ul>

<li>id — unique loan identifier</li>
<li>borrower — borrower address</li>
<li>laz — policy authority address</li>
<li>omnibus — disbursement authority address</li>
<li>principal — initial loan amount</li>
<li>outstanding — remaining loan balance</li>
<li>tenor_months — loan duration in months</li>
<li>metadata_hash — optional document hash</li>

</ul>

<h2>Genesis State</h2>

<p>
The loan module defines a custom genesis state allowing a chain to initialize
with existing loan data.
</p>

<ul>
<li>params — module parameters</li>
<li>next_id — next loan identifier</li>
<li>loans — existing loan records</li>
</ul>

<h2>Query Endpoints</h2>

<ul>

<li>/cosmos/loan/v1/params</li>
<li>/cosmos/loan/v1/loan/{loan_id}</li>
<li>/cosmos/loan/v1/loans</li>
<li>/cosmos/loan/v1/borrowers/{borrower}/loans</li>

</ul>

</body>
</html>
