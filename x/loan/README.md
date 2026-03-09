# Loan Module

The Loan module provides an on-chain lending lifecycle that allows users to request, approve, disburse, and repay loans in a controlled governance-enabled environment.

The module manages the lifecycle of loans between borrowers and authorized policy accounts.

## Features

- Borrower loan request
- Policy approval / rejection
- Disbursement confirmation
- Repayment tracking
- Governance-controlled parameters

## Loan Lifecycle

1. Borrower submits a loan request
2. Authorized LAZ approves or rejects
3. Omnibus confirms disbursement
4. Borrower repays the outstanding balance

## Key Roles

| Role | Responsibility |
|-----|-----|
| Borrower | Creates loan request |
| LAZ | Approves or rejects loan |
| Omnibus | Handles disbursement and repayment |
| Governance | Controls module parameters |

## Module Parameters

- Settlement denom
- Min / Max loan amount
- Max tenor
- Authorized LAZ accounts
- Authorized Omnibus accounts

## CLI Examples

Create loan
