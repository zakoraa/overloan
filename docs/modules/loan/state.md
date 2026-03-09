<!DOCTYPE html>
<html>

<body>

<h1>Loan Module State Storage</h1>

<p>
The Loan module stores loan records and module configuration
in the application state using key-value storage.
</p>

<h2>Primary Storage</h2>

<table border="1">

<tr>
<th>Key</th>
<th>Value</th>
<th>Description</th>
</tr>

<tr>
<td>loan_id</td>
<td>Loan</td>
<td>Stores the full loan object</td>
</tr>

</table>

<h2>Loan Object Structure</h2>

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
<td>Address of the borrower</td>
</tr>

<tr>
<td>laz</td>
<td>Authority responsible for approving or rejecting the loan</td>
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
<td>Remaining amount to be repaid</td>
</tr>

<tr>
<td>tenor_months</td>
<td>Loan duration in months</td>
</tr>

<tr>
<td>status</td>
<td>Current lifecycle state of the loan</td>
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
<td>Timestamp when the loan was disbursed</td>
</tr>

<tr>
<td>metadata_hash</td>
<td>Optional hash of external loan documentation</td>
</tr>

</table>


<h2>Secondary Index</h2>

<p>
The module maintains a secondary index to allow efficient lookup
of loans by borrower.
</p>

<table border="1">

<tr>
<th>Key</th>
<th>Value</th>
<th>Description</th>
</tr>

<tr>
<td>(borrower, loan_id)</td>
<td>loan_id</td>
<td>Index allowing queries of loans owned by a borrower</td>
</tr>

</table>


<h2>Module Parameters</h2>

<p>
The module stores configuration parameters which can be updated through governance.
</p>

<ul>

<li>settlement_denom</li>
<li>max_loan_amount</li>
<li>min_loan_amount</li>
<li>max_tenor_months</li>
<li>laz_authorities</li>
<li>omnibus_authorities</li>

</ul>


<h2>Genesis State</h2>

<p>
The module defines a custom genesis state allowing the chain
to start with predefined loan data.
</p>

<ul>

<li>params — module configuration parameters</li>
<li>next_id — next loan identifier</li>
<li>loans — existing loan records</li>

</ul>

</body>
</html>
