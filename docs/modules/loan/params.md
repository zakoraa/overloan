<!DOCTYPE html>
<html>

<body>

<h1>Loan Module Parameters</h1>

<p>
Module parameters define global configuration values that control
how the loan system behaves. These parameters are stored in module
state and can only be modified through governance.
</p>

<h2>Parameter List</h2>

<table border="1">

<tr>
<th>Parameter</th>
<th>Description</th>
</tr>

<tr>
<td>settlement_denom</td>
<td>Token denomination used for loan settlement</td>
</tr>

<tr>
<td>min_loan_amount</td>
<td>Minimum loan size allowed</td>
</tr>

<tr>
<td>max_loan_amount</td>
<td>Maximum loan size allowed</td>
</tr>

<tr>
<td>max_tenor_months</td>
<td>Maximum loan duration in months</td>
</tr>

<tr>
<td>laz_authorities</td>
<td>Addresses authorized to approve or reject loans</td>
</tr>

<tr>
<td>omnibus_authorities</td>
<td>Addresses authorized to confirm disbursement and record repayments</td>
</tr>

</table>


<h2>Parameter Governance</h2>

<p>
Parameters can only be modified through a governance proposal
that executes the <b>MsgUpdateParams</b> message.
</p>


<h3>Example Proposal</h3>

<pre>
{
  "title": "Update Loan Params",
  "summary": "Update module parameters",
  "deposit": "10000000stake",
  "messages": [
    {
      "@type": "/cosmos.loan.v1.MsgUpdateParams",
      "authority": "cosmos1govmoduleaddress",
      "params": {
        "settlement_denom": "stake",
        "min_loan_amount": "100",
        "max_loan_amount": "1000000",
        "max_tenor_months": "12",
        "laz_authorities": [
          "cosmos1..."
        ],
        "omnibus_authorities": [
          "cosmos1..."
        ]
      }
    }
  ]
}
</pre>

</body>
</html>
