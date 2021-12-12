package Models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

//type Items struct{
//	Products map[int]Product
//}
//
//func InitialiseItems() *Items {
//	products:=make(map[int]Product)
//	products[1]= Product{
//		Id: 1, ProductName: "Apple",
//		Description: "Fruit",
//		Price: 50,
//		Quantity: 10,
//	}
//	products[2]= Product{Id: 2,
//		ProductName: "Chess",
//		Description: "Sports",
//		Price: 1050,
//		Quantity: 100,
//	}
//
//
//	store:= Items{products}
//	return &store
//}
