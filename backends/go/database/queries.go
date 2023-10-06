package database

import (
	"benchmarks/utils"
	"fmt"
	"time"
)

// /select
func SelectQuery(opt *utils.Options){ 
	db := opt.DB
	var count int

	for i:=0; i<opt.N; i++{
		db.QueryRow(`
			SELECT COUNT(*) FROM electric_cars 
			WHERE vehicletype = ? AND model = ?
		`, "BEV", "MODEL 3").Scan(&count)
	}
	fmt.Println(count)
}

// /insert
func InsertQuery(opt *utils.Options){
	db := opt.DB
	for i:=0; i<opt.N; i++{
		db.Exec(`
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, "test", "test", "test", "test", 1, 1, "test", "test", "test", 1, "test")
	}
}

// /update
func UpdateQuery(opt *utils.Options){
	db := opt.DB
	db.Exec(`
		UPDATE electric_cars
		SET model = ?
		WHERE model = "I3"
	`, "I4")
	db.Exec(`
		UPDATE electric_cars
		SET model = ?
		WHERE model = "I4"
	`, "I3")
}

// /delete
func DeleteQuery(opt *utils.Options) error {
	db := opt.DB
	var duration int64

	for i:=0; i<opt.N; i++{
		db.Exec(`
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,"test", "test", "test", "test", 1, 1, "test", "test", "test", 1, "test")

		start := time.Now()
		db.Exec(`
			DELETE FROM electric_cars
			WHERE vin = "test" AND county = "test"
		`)
		duration += time.Since(start).Milliseconds()
	}

	return opt.Context.SendString(fmt.Sprintf("%v", duration/5))
}