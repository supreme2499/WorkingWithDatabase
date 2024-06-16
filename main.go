package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:210101@localhost:5432/MyTestDb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("\n\nошибка открытия: ", err)
	}
	defer db.Close()
	if err != nil {
		log.Fatal("\n\nошибка закрытия: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("\n\nошибка пинга: ", err)

	}

	createProductsTable(db)

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
	id SERIA; PRIMARY KEY,
	model 	VARCAR(100) NOT NULL,
	company VARCHAR (100) NOT NULL,
	price NUMERIC (6,2) NOT NULL)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("\n\nошибка выполнения запроса: ", err)
	}
}
