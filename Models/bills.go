package Models

import "time"

type Bills struct {
	Id              int       `json:"id""`
	ProductId       int       `json:"Pid"`
	ProductName     string    `json:"name"`
	Price           int       `json:"price"`
	Quantity        int       `json:"quantity"`
	TotalAmount     int       `json:"total"`
	TransactionTime time.Time `json:"SalesTime"`
}

//type Transactions struct{
//	Transaction map[int]Bills
//}
//
//func InitialiseTransactions() *Transactions {
//T:=make(map[int]Bills)
//
//T[1]= Bills{
//}
//
//tran:= Transactions{T}
//return &tran
//}
