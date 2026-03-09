<!DOCTYPE html>
<html>

<body>

<h1>Loan Module Queries</h1>

<p>
The Loan module exposes gRPC query endpoints that allow clients
to retrieve loan data and module configuration.
</p>


<h2>Query Params</h2>

<p>
Returns the current module configuration parameters.
</p>

<pre>
GET /cosmos/loan/v1/params
</pre>


<h2>Query Loan</h2>

<p>
Returns detailed information about a specific loan by its ID.
</p>

<pre>
GET /cosmos/loan/v1/loan/{loan_id}
</pre>


<h2>Query Loans</h2>

<p>
Returns a list of all loans stored in the module.
Supports pagination.
</p>

<pre>
GET /cosmos/loan/v1/loans
</pre>


<h2>Query Loans By Borrower</h2>

<p>
Returns all loans belonging to a specific borrower address.
Supports pagination.
</p>

<pre>
GET /cosmos/loan/v1/borrowers/{borrower}/loans
</pre>


</body>
</html>
