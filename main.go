package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Products struct {
	model   string
	company string
	price   float32
}

func (p Products) NewProducts() (NewPosition Products) {
	fmt.Print("\nВведите модель продукта")
	fmt.Scan(NewPosition.model)
	fmt.Print("\nВведите компанию изготовителя")
	fmt.Scan(NewPosition.company)
	fmt.Print("\nВведите цену прогукта:")
	fmt.Scan(NewPosition.price)
	return NewPosition
}

func main() {
	type Products struct {
		model   string
		company string
		price   float32
	}
	//Задаём коннект
	connStr := "postgres://postgres:123@localhost:5438/MyDb?sslmode=disable"

	//Открывкаем соединение с ДБ
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("\n\nошибка открытия: ", err)
	}

	// Создаём отложенную функцию закрытия нашего соединения
	defer db.Close()
	if err != nil {
		log.Fatal("\n\nошибка закрытия: ", err)
	}

	//Делаем проверку соединения
	if err = db.Ping(); err != nil {
		log.Fatal("Ping is err", err)
	}

	test := Products.NewProducts()
	fmt.Print(test)

	//createProductsTable(db)

	//DropTable(db)
}

// Удаление таблици products
func DropTable(db *sql.DB) {
	query := `DROP TABLE products;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("\n\nошибка выполнения запроса: ", err)
	}
}

// Создание таблици products
func createProductsTable(db *sql.DB) {
	/*Table: products
	- ID - первичный ключ который автоматически увеличивает своё значение
	- Model - строка, обязательно должно быть заполнено
	-Company - строка, обязательно должно быть заполнено
	-Price - число, обязательно должно быть заполнено
	*/

	//Запрос на создание таблицы который сначала делает проверку что такой таблицы ещё нет, после чего создаёт её с заданными полями
	query := `CREATE TABLE IF NOT EXISTS products (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	model 	VARCHAR(100) NOT NULL UNIQUE,
	company VARCHAR (100) NOT NULL,
	price NUMERIC (6,2) NOT NULL)`

	//Exec выполняет запрос без возврата значения
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("\n\nошибка выполнения запроса: ", err)
	}
}
