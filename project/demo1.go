package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryID  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {

	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	// bodyString := string(bodyBytes)

	// fmt.Println(bodyString)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	// fmt.Println(products)
	return products, nil

}

func AddProduct() (Product, error) {
	product := Product{ID: 5,ProductName: "M3 Macbook", CategoryID: 1, UnitPrice: 500}

	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()
	// fmt.Println("Başarılı Bir Şekilde Eklendi Yapıştırrr")

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	// // bodyString := string(bodyBytes)

	// // fmt.Println(bodyString)

	var productresponse Product
	json.Unmarshal(bodyBytes, &productresponse)
	// fmt.Println(product)

	return productresponse, nil
}
