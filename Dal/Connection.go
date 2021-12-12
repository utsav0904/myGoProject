package Dal

import (
	"Interface/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv"
)

type CatalogDatabase struct{

	Productdb *gorm.DB

}

func NewCatalogDataBase() *CatalogDatabase{

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=catalog password=Dunzo@123 sslmode=disable ")

	if err != nil{
		panic(err)
	}else {
		fmt.Println("Database Connection Successful !!!")
	}

	db.AutoMigrate(&Models.Product{})
	db.AutoMigrate(&Models.Bills{})

	return &CatalogDatabase{db}

}


