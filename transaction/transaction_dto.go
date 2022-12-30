package transaction

import "errors"

type TransactionDTO struct {
	UserId            string                 `json:"userId"`
	DetailTransaction []DetailTransactionDTO `json:"detailTransaction"`
}

type DetailTransactionDTO struct {
	ProductId string `json:"productId"`
	Total     uint   `json:"total"`
}

func (t *TransactionDTO) GetAllProductIds() []string {
	var ids []string
	for _, dt := range t.DetailTransaction {
		ids = append(ids, dt.ProductId)
	}
	return ids
}

func (t *TransactionDTO) GetProductTotal(productId string) (uint, error) {
	for _, dt := range t.DetailTransaction {
		if dt.ProductId == productId {
			return dt.Total, nil
		}
	}
	return 0, errors.New("NOT_FOUND")
}
