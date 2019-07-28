package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountId(t *testing.T) {
	acct1 := NewAccount("veera")
	acct2 := NewAccount("veera")

	if acct1.Id == acct2.Id {
		t.Errorf("two accounts should not have same id")
	}
}

func TestAccountMoneyAfterCredit(t *testing.T) {
	account := NewAccount("veera")
	CreditTo(&account, Money(500))

	if account.Balance != Money(500) {
		t.Errorf("money is not credited to account")
	}
}

func TestAccountMoneyAfterDebit(t *testing.T) {
	account := NewAccount("veera")
	CreditTo(&account, Money(500))

	DebitFrom(&account, Money(400))

	if account.Balance != Money(100) {
		t.Errorf("money is not debited from account")
	}
}

func TestAccountTransactions(t *testing.T) {
	account := NewAccount("veera")
	CreditTo(&account, Money(500))

	DebitFrom(&account, Money(400))
	expectedTransactions := make([]Transaction, 2)

	expectedTransactions[0] = Transaction{Type: Credit, Amount: 500}
	expectedTransactions[1] = Transaction{Type: Debit, Amount: 400}

	assert.Equal(t, expectedTransactions, account.Transactions)
}
