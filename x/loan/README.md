
<h1>Loan Module</h1>

<p>
The <b>Loan module</b> implements an on-chain lending lifecycle within a Cosmos SDK application.
It allows borrowers to submit loan requests and enables authorized policy actors to approve,
reject, disburse, cancel, and record repayments of loans.
</p>

<p>
The module manages the full lifecycle of loans stored on-chain and ensures that operations follow
governance‑defined parameters and policy authorities.
</p>

<p>
More detailed documentation is available in the following pages:
</p>

<ul>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/overview.md">
Module Overview
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/concepts.md">
Concepts
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/lifecycle.md">
Loan Lifecycle
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/state.md">
State Layout
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/params.md">
Module Parameters
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/messages.md">
Messages
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/queries.md">
Queries
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/dev-guide.md">
Developer Guide
</a>
</li>

</ul>

<h2>Features</h2>

<ul>
<li>Borrowers can request loans on-chain</li>
<li>LAZ approve or reject loan requests</li>
<li>Loan disbursement confirmation by authorized omnibus accounts</li>
<li>Repayment tracking with outstanding balance updates</li>
<li>Governance‑controlled parameters</li>
<li>Queryable loan data through gRPC and REST APIs</li>
</ul>

<h2>Loan Lifecycle</h2>

<ol>
<li>Borrower submits a loan request (<b>LOAN_STATUS_PENDING</b>)</li>
<li>LAZ authority approves or rejects the request (<b>APPROVED</b> or <b>REJECTED</b>)</li>
<li>Omnibus confirms disbursement (<b>LOAN_STATUS_DISBURSED</b>)</li>
<li>Disbursement may also be rejected (<b>LOAN_STATUS_CANCELLED</b>)</li>
<li>Repayments reduce outstanding balance</li>
<li>When outstanding reaches zero the loan becomes <b>LOAN_STATUS_REPAID</b></li>
</ol>

<h2>Key Roles</h2>

<table border="1">
<tr>
<th>Role</th>
<th>Description</th>
</tr>

<tr>
<td>Borrower</td>
<td>User who submits a loan request</td>
</tr>

<tr>
<td>LAZ Authority</td>
<td>Policy authority responsible for approving or rejecting loans</td>
</tr>

<tr>
<td>Omnibus Authority</td>
<td>Confirms disbursement and records repayments</td>
</tr>

<tr>
<td>Governance</td>
<td>Controls module parameters</td>
</tr>
</table>

<h2>Module Parameters</h2>

<ul>
<li><b>settlement_denom</b> – token denomination used for loan settlement</li>
<li><b>min_loan_amount</b> – minimum allowed loan size</li>
<li><b>max_loan_amount</b> – maximum allowed loan size</li>
<li><b>max_tenor_months</b> – maximum loan duration</li>
<li><b>laz_authorities</b> – addresses authorized to approve or reject loans</li>
<li><b>omnibus_authorities</b> – addresses authorized to confirm disbursement and repayments</li>
</ul>

<h2>Messages</h2>

<ul>
<li>MsgCreateLoan – submit a new loan request</li>
<li>MsgApproveLoan – approve a pending loan</li>
<li>MsgRejectLoan – reject a pending loan</li>
<li>MsgConfirmDisbursement – confirm loan disbursement</li>
<li>MsgRejectDisbursement – reject disbursement</li>
<li>MsgRepayLoan – record loan repayment</li>
<li>MsgUpdateParams – update module parameters via governance</li>
</ul>

<h2>Queries</h2>

<ul>
<li>Params – returns module parameters</li>
<li>Loan – returns loan details by ID</li>
<li>Loans – list all loans</li>
<li>LoansByBorrower – list loans belonging to a borrower</li>
</ul>

<h2>Genesis State</h2>

<p>
The loan module defines a custom genesis state that allows a chain to initialize with predefined loan data.
</p>

<ul>
<li><b>params</b> – module configuration parameters</li>
<li><b>next_id</b> – next loan identifier</li>
<li><b>loans</b> – list of existing loans</li>
</ul>


