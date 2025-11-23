package product

type Service interface {
	GetProducts() ([]Product, error)
	CreateProduct(p *Product) error
}

type productService struct {
	repo Repository
}

func NewProductService(r Repository) Service {
	return &productService{r}
}

func (s *productService) GetProducts() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *productService) CreateProduct(p *Product) error {
	return s.repo.Create(p)
}
