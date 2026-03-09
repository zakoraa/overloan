<!DOCTYPE html>
<html>

<body>

<h1>Loan Module Developer Guide</h1>

<p>
This guide explains how developers can run a local chain,
configure module parameters through governance,
and test the complete lifecycle of a loan using the CLI.
</p>

<p>
For deeper understanding of the module architecture and data model, see:
</p>

<ul>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/concepts.md">
Loan Module Concepts
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/lifecycle.md">
Loan Lifecycle
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/params.md">
Module Parameters
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/messages.md">
Module Messages
</a>
</li>

<li>
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/queries.md">
Module Queries
</a>
</li>

</ul>


<h2>Prerequisites</h2>

<ul>
<li>Go installed</li>
<li>Cosmos SDK application built</li>
<li>CLI binary <code>simd</code> available</li>
</ul>



<h2>Step 1 — Initialize Development Chain</h2>

<pre>
make init-dev
</pre>

<p>
This command will:
</p>

<ul>

<li>Remove previous chain state</li>
<li>Initialize a new chain</li>
<li>Create the following accounts:
<ul>
<li>validator</li>
<li>borrower</li>
<li>laz</li>
<li>omnibus</li>
</ul>
</li>

<li>Update governance voting duration</li>
<li>Start the node</li>

</ul>



<h2>Step 2 — Submit Governance Proposal</h2>

<p>
The Loan module parameters must be configured through governance.
</p>

<pre>
make create-proposal
</pre>

<p>
This proposal updates module parameters such as:
</p>

<ul>
<li>settlement_denom</li>
<li>min_loan_amount</li>
<li>max_loan_amount</li>
<li>max_tenor_months</li>
<li>laz_authorities</li>
<li>omnibus_authorities</li>
</ul>

<p>
For a full description of parameters see:
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/params.md">
Module Parameters
</a>
</p>



<h2>Step 3 — Check Proposal Status</h2>

<pre>
./build/simd query gov proposal 1
</pre>



<h2>Step 4 — Create Loan</h2>

<pre>
./build/simd tx loan create-loan \
--laz $(./build/simd keys show laz -a --keyring-backend test --home ./private/.overloand) \
--principal 1000stake \
--tenor-months 6 \
--metadata-hash QmHash123 \
--from borrower \
--chain-id overloan-1 \
--keyring-backend test \
--home ./private/.overloand \
--gas auto --gas-adjustment 1.3 \
-y
</pre>

<p>
Loan status becomes:
</p>

<p><b>LOAN_STATUS_PENDING</b></p>

<p>
See message documentation:
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/messages.md">
Messages Documentation
</a>
</p>



<h2>Step 5 — Approve or Reject Loan</h2>

<h3>Approve Loan</h3>

<pre>
./build/simd tx loan approve-loan &lt;loan_id&gt; --from laz --chain-id overloan-1 --keyring-backend test --home ./private/.overloand --gas auto --gas-adjustment 1.3 -y
</pre>

<p>Status: <b>LOAN_STATUS_APPROVED</b></p>

<h3>Reject Loan</h3>

<pre>
./build/simd tx loan reject-loan &lt;loan_id&gt; --from laz --chain-id overloan-1 --keyring-backend test --home ./private/.overloand -y
</pre>

<p>Status: <b>LOAN_STATUS_REJECTED</b></p>



<h2>Step 6 — Confirm or Reject Disbursement</h2>

<h3>Confirm Disbursement</h3>

<pre>
./build/simd tx loan confirm-disbursement &lt;loan_id&gt; \
--from omnibus \
--chain-id overloan-1 \
--keyring-backend test \
--home ./private/.overloand \
--gas auto --gas-adjustment 1.3 \
-y
</pre>

<p>Status: <b>LOAN_STATUS_DISBURSED</b></p>

<h3>Reject Disbursement</h3>

<pre>
./build/simd tx loan reject-disbursement \
--loan-id &lt;loan_id&gt; \
--reason "bank transfer failed" \
--from omnibus \
--chain-id overloan-1 \
--keyring-backend test \
--home ./private/.overloand \
--gas auto --gas-adjustment 1.3 \
-y
</pre>

<p>Status: <b>LOAN_STATUS_CANCELLED</b></p>



<h2>Step 7 — Repay Loan</h2>

<p>
Repayments reduce the outstanding loan balance.
</p>

<pre>
./build/simd tx loan repay-loan &lt;loan_id&gt; \
--amount 500stake \
--from omnibus \
--chain-id overloan-1 \
--keyring-backend test \
--home ./private/.overloand \
--gas auto --gas-adjustment 1.3 \
-y
</pre>

<p>
If the outstanding balance becomes zero,
the loan status changes to:
</p>

<p><b>LOAN_STATUS_REPAID</b></p>



<h2>Loan Lifecycle Summary</h2>

<ol>

<li>Create loan request → <b>PENDING</b></li>

<li>Approve or reject by LAZ → <b>APPROVED / REJECTED</b></li>

<li>Disbursement confirmed by omnibus → <b>DISBURSED</b></li>

<li>Disbursement may be rejected → <b>CANCELLED</b></li>

<li>Repayments reduce outstanding balance</li>

<li>When outstanding reaches zero → <b>REPAID</b></li>

</ol>


<p>
For a detailed explanation of state transitions see:
<a href="https://github.com/zakoraa/overloan/blob/main/docs/modules/loan/lifecycle.md">
Loan Lifecycle Documentation
</a>
</p>


</body>
</html>
