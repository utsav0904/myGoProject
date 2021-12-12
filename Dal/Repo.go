package Dal

import (
	"Interface/Models"
	"github.com/jinzhu/gorm"
)


type ProductRepoInterface interface {

	//Transaction Functions
	AddTransactions			(SaleRecord Models.Bills) bool
	GetAllTransaction		() []Models.Bills
	GetTop5Products			() []Models.Bills

	//Products Functions
	GetProductById			(id int) Models.Product
	CheckAvailabilityById	(id int) bool
	GetProductQuantityById	(id int) int
	ReduceQuantity			(quantity Models.Product,id int)
	IncreaseQuantity		(quantity Models.Product,id int)
	GetAllProduct			() []Models.Product
	InsertProduct			(product Models.Product) (int,int)
	//GetAvailableProducts	() []models.Product


}

type ProductRepo struct{
	Database *CatalogDatabase
}

func NewProductRepo(datastore *CatalogDatabase) ProductRepoInterface {
	return &ProductRepo{datastore}
}

func (r ProductRepo) GetProductById(id int) Models.Product{

	var pro Models.Product
	r.Database.Productdb.First(&pro,id)
	return pro
}

func (r ProductRepo) CheckAvailabilityById(id int) bool {

	err:= r.Database.Productdb.Where("Id=?",id).Find(&Models.Product{}).Error;

	if gorm.IsRecordNotFoundError(err){
		return false
	}
	return true
}

func (r ProductRepo) GetProductQuantityById(id int) int{
	var pro Models.Product
	r.Database.Productdb.First(&pro,id)
	return pro.Quantity
}

func (r ProductRepo) ReduceQuantity	(quantity Models.Product,id int) {
	Product:= r.GetProductById(id)
	Product.Quantity-=quantity.Quantity
	r.Database.Productdb.Save(Product)
}

func (r ProductRepo) IncreaseQuantity(quantity Models.Product,id int) {
	Product:= r.GetProductById(id)
	Product.Quantity+=quantity.Quantity
	r.Database.Productdb.Save(Product)
}

func (r ProductRepo) GetAllProduct() []Models.Product{
	AllProduct :=make([]Models.Product,0)
	r.Database.Productdb.Find(&AllProduct)
	return AllProduct
}

func (r ProductRepo) GetAllTransaction() []Models.Bills{
	AllProduct :=make([]Models.Bills,0)
	r.Database.Productdb.Find(&AllProduct)
	return AllProduct
}



func (r ProductRepo) InsertProduct(product Models.Product) (id int , err int){

	e:=r.Database.Productdb.NewRecord(product)
	if e!=true {
		return product.Id,1
	}
	r.Database.Productdb.Create(&product)
	return product.Id,0
}

func (r ProductRepo) AddTransactions(SaleRecord Models.Bills)bool {

	err:=r.Database.Productdb.NewRecord(SaleRecord)
	if err!=true{
		return false
	}
	r.Database.Productdb.Create(&SaleRecord)
	return true
}

func (r ProductRepo) GetTop5Products() []Models.Bills{
	var SalesRecord []Models.Bills

	sqlStr:="SELECT product_id ,SUM(quantity)AS Total FROM (SELECT product_id,product_name ,quantity FROM bills WHERE transaction_time > NOW() - INTERVAL '1 hour') Temp GROUP BY product_id ORDER BY Total DESC LIMIT 5"

	//sqlStr :="SELECT Pid AS Product ID, ProductName , Total As Billable Amount FROM (SELECT product_id,quantity_sold FROM sales_records WHERE sales_time> NOW() - INTERVAL '1 hour') Temp GROUP BY product_id ORDER BY Total DESC LIMIT 5"
	rows, err := r.Database.Productdb.Raw(sqlStr).Rows()
	if err!=nil{
		panic(err)
	}
	for rows.Next(){
		var p Models.Bills
		var Id int
		var QuantitySold int
		err = rows.Scan(&Id,&QuantitySold)
		p.ProductId = Id
		p.Quantity = QuantitySold
		SalesRecord = append(SalesRecord, p)
	}
	return SalesRecord
}
//
//
//
//
//
//func (r ProductRepo) GetAvailableProducts() []models.Product{
//	TempProduct:=make([]models.Product,0)
//	for _,Product :=range r.GetAllProduct(){
//		if Product.Quantity>0{
//			TempProduct = append(TempProduct, Product)
//		}
//	}
//	return TempProduct
//}
//
//func (r ProductRepo) DeleteProduct(id int) {
//	var product models.Product
//	r.Database.Productdb.First(&product,id)
//	r.Database.Productdb.Delete(&product)
//
//}
