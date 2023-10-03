package database

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strings"
)

func Seed(db *sql.DB){
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS electric_cars (
			vin TEXT,
			county TEXT,
			city TEXT,
			state TEXT,
			postalCode INT,
			modelYear INT,
			make TEXT,
			model TEXT,
			vehicleType TEXT,
			electricRange INT,
			vehicleLocation TEXT
		)
	`)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("./../../db_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records[1:] {
		_, vehicleType, _ := strings.Cut(record[8], "(")
		vehicleType, _ = strings.CutSuffix(vehicleType, ")")
		_, err = db.Exec("INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], vehicleType, record[10], record[14])
		if err != nil {
			panic(err)
		}
	}
}