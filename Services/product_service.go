package Services

import (
	"Interface/Dal"
	"Interface/Models"
	"time"
)

type ProductServiceInterface interface {

	//Transaction service function
	GetAllTransaction() []Models.Bills
	GetTop5Products() []Models.Bills

	//Product Services Function
	AddNewProduct(product Models.Product) (int, int)
	GetAllProduct() []Models.Product
	GetProductById(id int) (p Models.Product, err int)
	BuyProduct(quantity Models.Product, id int) (int, int, string)
	//GetAvailableProduct	() []Models.Product

}

type ProductService struct {
	Repos Dal.ProductRepoInterface
}

func NewProductServices(Repo Dal.ProductRepoInterface) ProductServiceInterface {

	return &ProductService{Repo}
}

func (s ProductService) GetAllTransaction() []Models.Bills {

	return s.Repos.GetAllTransaction()
}

func (s ProductService) GetTop5Products() []Models.Bills {

	return s.Repos.GetTop5Products()

}

func (s ProductService) AddNewProduct(product Models.Product) (int, int) {

	id, b := s.Repos.InsertProduct(product)
	if b == 1 {
		return id, 1
	}

	return id, 0

}

func (s ProductService) GetAllProduct() []Models.Product {

	return s.Repos.GetAllProduct()

}

func (s ProductService) GetProductById(id int) (p Models.Product, err int) {

	e := s.Repos.CheckAvailabilityById(id)

	if e == false {
		return Models.Product{}, 1
	}
	return s.Repos.GetProductById(id), 0

}

func (s ProductService) BuyProduct(quantity Models.Product, id int) (int, int, string) {

	e := s.Repos.CheckAvailabilityById(id)
	if e == false {
		return 1, 0, ""
	}
	q := s.Repos.GetProductQuantityById(id)
	if q < quantity.Quantity {
		return 2, 0, ""
	}
	s.Repos.ReduceQuantity(quantity, id)
	pro := s.Repos.GetProductById(id)

	var t Models.Bills
	t.ProductId = pro.Id
	t.Price = pro.Price
	t.Quantity = quantity.Quantity
	var tot int
	tot = quantity.Quantity * pro.Price
	t.TotalAmount = quantity.Quantity * pro.Price
	t.ProductName = pro.Name
	t.TransactionTime = time.Now()

	s.Repos.AddTransactions(t)
	return 0, tot, t.ProductName

}
