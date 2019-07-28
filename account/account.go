package account

type AccountId int
type Money int

type TransactionType int

const (
	Debit TransactionType = iota
	Credit
)

var accountId AccountId = 1

type Account struct {
	name         string
	Id           AccountId
	Balance      Money
	Transactions []Transaction
}

type Transaction struct {
	Type   TransactionType
	Amount Money
}

func NewAccount(name string) Account {
	account := Account{name: name}
	account.Id = accountId
	accountId += 1
	return account
}

func CreditTo(account *Account, money Money) {
	account.Balance += money
	account.Transactions = append(account.Transactions, Transaction{Type: Credit, Amount: money})
}

func DebitFrom(account *Account, money Money) {
	account.Balance -= money
	account.Transactions = append(account.Transactions, Transaction{Type: Debit, Amount: money})
}
