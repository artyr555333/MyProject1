package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Products []Product

// Продукт
type Product struct {
	Id          int    `json:id`
	Name        string `json:name`
	Price       int    `json:price`
	Category    string `json:category`
	Weight      string `json:weight`
	Description string `json:description`
}

// Покупатель
type Customer struct {
	Name  string
	Phone string
	Email string
}

// Адрес покупателя
type Address struct {
	Country string
	City    string
	Street  string
	House   string
}

// Заказ
type Order struct {
	Items         Products
	Customer      Customer
	PaymentMethod string
	Address       Address
}

func main() {
	products := getProducts()

	var customer Customer

	fmt.Println("--Заполните данные покупателя--")

	fmt.Println("**Введите имя**")
	fmt.Scanf("%s\n", &customer.Name)

	fmt.Println("**Введите номер телефона**")
	fmt.Scanf("%s\n", &customer.Phone)

	fmt.Println("**Введите email**")
	fmt.Scanf("%s\n", &customer.Email)

	var address Address

	fmt.Println("--Введите свой адрес--")

	fmt.Println("**Введите страну**")
	fmt.Scanf("%s\n", &address.Country)

	fmt.Println("**Введите город**")
	fmt.Scanf("%s\n", &address.City)

	fmt.Println("**Введите улицу**")
	fmt.Scanf("%s\n", &address.Street)

	fmt.Println("**Введите номер дома**")
	fmt.Scanf("%s\n", &address.House)

	fmt.Println("--Выберите номера продуктов для заказа--")

	for _, product := range products {
		fmt.Printf("[%d] - %s - цена:  %d\n", product.Id, product.Name, product.Price)
	}

	var prodId int
	var basket []Product

	fmt.Println("Укажите номера | Для выхода из корзины введите -1")

	for prodId != -1 {
		fmt.Scanf("%d\n", &prodId)
		if prodId == -1 {
			continue
		}
		basket = append(basket, searchProductId(prodId))

	}

	var order Order

	fmt.Println("**Введите способ оплаты**")
	fmt.Scanf("%s", &order.PaymentMethod)

	//

	fmt.Println("--ИНФОРМАЦИЯ О ЗАКАЗЕ--")

	fmt.Println("--Покупатель--")

	fmt.Printf("Имя: %s\n", customer.Name)
	fmt.Printf("Телефон: %s\n", customer.Phone)
	fmt.Printf("Почта: %s\n", customer.Email)
	fmt.Printf("Способ оплаты: %s\n", order.PaymentMethod)

	fmt.Println("--Адрес Покупателя--")

	fmt.Printf("Страна : %s\n", address.Country)
	fmt.Printf("Город : %s\n", address.City)
	fmt.Printf("Улица : %s\n", address.Street)
	fmt.Printf("Номер дома : %s\n", address.House)

	fmt.Println("--Корзина--")

	var summ int
	fmt.Scanf("%d", &summ)
	for _, v := range basket {
		fmt.Printf("[%d] - %s - цена: %d\n", v.Id, v.Name, v.Price)
		summ += v.Price
	}

	fmt.Printf("Общая сумма заказа: %d", summ)

}

func getProducts() Products {

	var products Products

	byteValue, _ := os.ReadFile("products.json")

	json.Unmarshal(byteValue, &products)

	return products

}

func searchProductId(productId int) Product {
	products := getProducts()
	var fineProduct Product
	for _, v := range products {
		if productId == v.Id {
			fineProduct = v
		}
	}
	return fineProduct
}
