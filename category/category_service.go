package category

type CategoryService interface {
	FindById(id string) *Category
	FindAll() []*Category
	Save(request *CategoryDTO)
	Update(request *CategoryDTO, id string)
	Delete(id string)
}
