package test

import (
	"Interface/Handler"
	"Interface/Models"
	"Interface/Services"
	mock_Dal "Interface/mock/Dal"
	mock_services "Interface/mock/Sevices"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct:=Models.Product{Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	//OutputProduct:=Models.Product{Id:1,ProductName:"wickets",Description :"wood",Price: 256,Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("POST", "/products/insert", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductServiceInterface(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(1,0)
	ProductController:=Handler.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/insert",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"Message    ":"Product Successfully inserted into the Store with Product ID: 1","ProductName":"wickets"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestCreateProduct2(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct:=Models.Product{Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	//OutputProduct:=Models.Product{Id:1,ProductName:"wickets",Description :"wood",Price: 256,Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("POST", "/products/insert", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductServiceInterface(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(1,1)
	ProductController:=Handler.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/insert",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"error":"Id is auto generated please dont provide Id.","code":"PrimaryKey_Error"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 400,rr.Code)
}

func TestGetProductById_AvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo :=mock_Dal.NewMockProductRepoInterface(ctrl)
	//MockRepo := mock_Databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().GetProductById(id).Return(Models.Product{
		Id : 1,
		Price: 22,
		Name : "Balls",
		Quantity: 10,
		Description: "adidas",
	})
	//MockRepo.EXPECT().CheckProductAvailableById(id).Return(true)
	MockRepo.EXPECT().CheckAvailabilityById(id).Return(true)
	//ProductService:=Services.NewProductService(MockRepo)
	ProductService:=Services.NewProductServices(MockRepo)
	product ,error :=ProductService.GetProductById(id)
	assert.NotNil(t, product)
	assert.Equal(t, 0,error)
}

func TestGetProductById_UnAvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo := mock_Dal.NewMockProductRepoInterface(ctrl)

	MockRepo.EXPECT().CheckAvailabilityById(id).Return(false)

	ProductService:=Services.NewProductServices(MockRepo)

	product ,error :=ProductService.GetProductById(id)
	assert.NotNil(t, product)
	assert.Equal(t, 1,error)
}