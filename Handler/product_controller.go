package Handler

import (
	"Interface/Handler/ErrorHandler"
	"Interface/Models"
	"Interface/Services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ControllerInterface interface {
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
	BuyProductMany(w http.ResponseWriter, r *http.Request)
	GetAllTransactions(w http.ResponseWriter, r *http.Request)
	GetTop5Products(w http.ResponseWriter, r *http.Request)
}

type ProductController struct {
	ps Services.ProductServiceInterface
}

func Initialise(p Services.ProductServiceInterface) ControllerInterface {
	return &ProductController{p}
}

func (pc ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id int
	var e error
	id, e = strconv.Atoi(params["id"])
	if e != nil {
		ErrorHandler.Response(w, 400, "Product Id in url should be Integer", "Incorrect_Url")
		return
	}

	IdProduct, err := pc.ps.GetProductById(id)
	if err == 1 {
		ErrorHandler.Response(w, 400, "Product With Given Id is not present in the store", "Product_Not_Found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(IdProduct)
}

func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Models.Product
	e := json.NewDecoder(r.Body).Decode(&product)
	if e != nil {
		ErrorHandler.Response(w, 400, "Error occur while Decoding: Id , Price ,Quantity should be int and check  structure body", "DataType_Mismatch")
		return
	}
	if product.Quantity<=0{
		ErrorHandler.Response(w, 400, "Can't able to Insert Negative or NULL quantity", "Negative_Quantity")
		return
	}
	if product.Price<=0{
		ErrorHandler.Response(w, 400, "Can't able to Insert Negative or NULL price", "Negative_Price")
		return
	}
	i, err := pc.ps.AddNewProduct(product)

	if err == 1 {
		ErrorHandler.Response(w, 400, "Id is auto generated please dont provide Id.", "PrimaryKey_Error")
		return
	}
	//msg:="Product with id:"+(string)(i)+"is created  Successfully and inserted into the Store"
	//err=json.NewEncoder(w).Encode((string)(msg))
	str := "Product Successfully inserted into the Store with Product ID: " + strconv.Itoa(i)
	ErrorHandler.Response1(w, 200, str, product.Name)

}

func (pc ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := pc.ps.GetAllProduct()
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		ErrorHandler.Response(w, 400, "Error occur while Encoding", "Encoding_Error")
		return
	}
	w.WriteHeader(200)
}

func (pc ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var value Models.Product
	e := json.NewDecoder(r.Body).Decode(&value)
	if e != nil {
		ErrorHandler.Response(w, 400, "Error occur while Decoding: Quantity should be int and check structure body ", "DataType_Mismatch")
		return
	}
	var id int
	//var e error
	id, e = strconv.Atoi(params["id"])
	if e != nil {
		ErrorHandler.Response(w, 400, "Product Id in url should be Integer", "Incorrect_Url")
		return
	}
	if value.Quantity<=0{
		ErrorHandler.Response(w, 400, "Can't able to purchase Negative or NULL quantity", "Negative_Quantity")
		return
	}
	err, total, name := pc.ps.BuyProduct(value, id)
	if err == 1 {
		ErrorHandler.Response(w, 400, "Product With Given Id is not present in the store", "Product_Not_Found")
		return
	} else if err == 2 {
		ErrorHandler.Response(w, 400, "sufficient quantity not available", "Insufficient_Quantity")
		return
	} else {
		str := "Product with id: " + strconv.Itoa(id) + " purchase successful"
		ErrorHandler.Response2(w, 200, str, name, total)
		return
	}
}

func (pc ProductController) BuyProductMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var value []Models.Product
	e := json.NewDecoder(r.Body).Decode(&value)
	if e != nil {
		ErrorHandler.Response(w, 400, "Error occur while Decoding: Id ,Quantity should be int and check structure body ", "dataType_Mismatch")
		return
	}
	for i, _ := range value {
		p, err := pc.ps.GetProductById(value[i].Id)
		if err == 1 {
			str := "Product With Id: " + strconv.Itoa(value[i].Id) + " is not available in the store"
			ErrorHandler.Response(w, 400, str, "Product_Not_Found")
			return
		}
		if value[i].Quantity<=0{

				ErrorHandler.Response(w, 400, "Can't able to purchase Negative or NULL quantity", "Negative_Quantity")
				return

		}
		if p.Quantity < value[i].Quantity {
			str := "Product With Id: " + strconv.Itoa(value[i].Id) + " doesn't have sufficient quantity"
			ErrorHandler.Response(w, 400, str, "Insufficient_Quantity")
			return
		}


	}
	tot := 0
	for i, _ := range value {
		_, t, _ := pc.ps.BuyProduct(value[i], value[i].Id)
		tot += t

	}
	ErrorHandler.Response3(w, 200, "All the products purchased successfully", tot)
	return
}

func (pc ProductController) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := pc.ps.GetAllTransaction()
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		ErrorHandler.Response(w, 400, "Error occur while Encoding", "Encoding_Error")
		return
	}
	w.WriteHeader(200)
}

type top struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func (pc ProductController) GetTop5Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := pc.ps.GetTop5Products()
	var top5 []top
	for i, _ := range t {
		var temp top
		temp.Id = t[i].ProductId
		p, _ := pc.ps.GetProductById(t[i].ProductId)
		temp.Name = p.Name
		temp.Quantity = t[i].Quantity
		top5 = append(top5, temp)
	}
	err := json.NewEncoder(w).Encode(top5)
	if err != nil {
		ErrorHandler.Response(w, 400, "Error occur while Encoding", "Encoding_Error")
		return
	}
	w.WriteHeader(200)

}
