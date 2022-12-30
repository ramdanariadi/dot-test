package product

type ProductService interface {
	FindById(id string) *Product
	FindByIds(id []string) []*Product
	FindAll() []*Product
	Save(request ProductDTO)
	Update(request ProductDTO, id string)
	Delete(id string)
}
