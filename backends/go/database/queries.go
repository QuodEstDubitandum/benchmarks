package database

import (
	"benchmarks/utils"
)

// /select
func SelectQuery(opt *utils.Options){ 
	db := opt.DB

	for i:=0; i<opt.N; i++{
		stmt, _ := db.Prepare(`
			SELECT COUNT(*) FROM electric_cars 
			WHERE vehicletype = ? AND model = ?
		`)
		stmt.Exec("BEV", "MODEL 3")
	}
}

// /insert
func InsertQuery(opt *utils.Options){
	db := opt.DB
	for i:=0; i<opt.N; i++{
		stmt, _ := db.Prepare(`
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`)
		stmt.Exec("test", "test", "test", "test", 1, 1, "test", "test", "test", 1, "test")
	}
}

// /update
func UpdateQuery(opt *utils.Options){
	db := opt.DB
	for i:=0; i<opt.N; i++{
		stmt, _ := db.Prepare(`
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`)
		stmt.Exec("test", "test", "test", "test", 1, 1, "test", "test", "test", 1, "test")
	}
}

// /delete
func DeleteQuery(opt *utils.Options){
	db := opt.DB
	for i:=0; i<opt.N; i++{
		stmt, _ := db.Prepare(`
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`)
		stmt.Exec("test", "test", "test", "test", 1, 1, "test", "test", "test", 1, "test")
	}
}