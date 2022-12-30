package transaction

type TransactionService interface {
	FindByTransactionId(id string, userId string) *Transaction
	Find(userId string) []*Transaction
	Save(transaction TransactionDTO)
}
