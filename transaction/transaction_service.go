package transaction

type TransactionService interface {
	FindByTransactionId(id string) *Transaction
	FindByUserId(id string) []*Transaction
	Save(transaction TransactionDTO)
}
