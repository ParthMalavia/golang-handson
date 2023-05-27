package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USERNAME = "parth"
	PASSWORD = "parth123"
	DBNAME   = "exercise"
)

// List of phone numbers that will be added to database
var phonenums []string = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

type Phone struct {
	Id     int
	Number string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", HOST, PORT, USERNAME, PASSWORD)

	db, err := sql.Open("postgres", psqlInfo)
	ErrorHandler(err)

	// Uncomment while creating new database
	err = DropCreateDB(db, DBNAME)
	ErrorHandler(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s database=%s", psqlInfo, DBNAME)
	db, err = sql.Open("postgres", psqlInfo)
	ErrorHandler(err)
	defer db.Close()

	// Uncomment while creating new table
	ErrorHandler(CreatePhonenumTable(db))

	for _, n := range phonenums {
		_, err := InsertPhoneNo(db, n)
		if err != nil {
			ErrorHandler(err)
		}
		// fmt.Println(id, n)
	}

	ph, err := GetPhone(db, 5)
	ErrorHandler(err)
	fmt.Println("Get the phone number of id 5:", ph)

	phones, err := AllPhones(db)
	ErrorHandler(err)

	for _, ph := range phones {
		number := normalize(ph.Number)

		if number != ph.Number {
			fmt.Println("Updating phone with ID:", ph.Id)

			// Check if there exist a phone-number same as normalized number
			existing, err := FindPhone(db, number)
			ErrorHandler(err)

			if existing != nil {
				fmt.Println("Already exists number", number, "Deleting this one...")

				// DELETE given number
				ErrorHandler(DeletePhone(db, ph.Id))
				continue
			}

			// Update phone on this ID
			ph.Number = number
			ErrorHandler(UpdatePhone(db, ph))

		} else {
			fmt.Println("Nothing to change.")
		}
	}
}

func UpdatePhone(db *sql.DB, p Phone) error {
	statement := `
	UPDATE Phone_numbers 
	SET value=$2 
	WHERE id=$1
	`

	_, err := db.Exec(statement, p.Id, p.Number)
	return err
}

func DeletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM Phone_numbers WHERE id=$1`

	_, err := db.Exec(statement, id)
	return err
}

func FindPhone(db *sql.DB, number string) (*Phone, error) {
	var p Phone

	err := db.QueryRow("SELECT id FROM Phone_numbers WHERE value=$1;", number).Scan(&p.Id)
	if err != nil {
		// Skip error for no rows exists
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func GetPhone(db *sql.DB, id int) (string, error) {
	var number string

	err := db.QueryRow("SELECT value FROM Phone_numbers WHERE id=$1;", id).Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func AllPhones(db *sql.DB) ([]Phone, error) {
	data := []Phone{}

	rows, err := db.Query("SELECT id, value from Phone_numbers")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ph Phone
		rows.Scan(&ph.Id, &ph.Number)
		data = append(data, ph)
	}
	return data, nil
}

func InsertPhoneNo(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO Phone_numbers(value) VALUES($1) RETURNING id`
	var id int

	err := db.QueryRow(statement, phone).Scan(&id)
	ErrorHandler(err)
	return id, nil
}

func CreateDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE  DATABASE " + name)
	return err
}

func CreatePhonenumTable(db *sql.DB) error {
	statement := `
	CREATE TABLE Phone_numbers (
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
}

// TODO: TRY without deleting database
func DropCreateDB(db *sql.DB, name string) error {

	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		panic(err)
		// return err
	}
	return CreateDB(db, name)
}

func normalize2(number string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(number, "")
}

func normalize(number string) string {
	var t bytes.Buffer

	for _, c := range number {
		if c >= '0' && c <= '9' {
			t.WriteRune(c)
		}
	}
	return t.String()
}

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
		// fmt.Println("Something went wrong", err)
		// os.Exit(1)
	}
}
