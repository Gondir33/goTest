package storage

import (
	"goTest/internal/models"
	"strconv"
)

func buildSqlForUpdateCar(id int, car models.Car) string {
	comma := false
	sql := "UPDATE car SET "
	if car.RegNum != "" {
		sql += "regNum = " + car.RegNum
		comma = true
	}
	if car.Mark != "" {
		if comma {
			sql += ", "
		}
		comma = true
		sql += "mark = " + car.Mark
	}
	if car.Model != "" {
		if comma {
			sql += ", "
		}
		comma = true
		sql += "model = " + car.Model
	}
	if car.Year != 0 {
		if comma {
			sql += ", "
		}
		comma = true
		sql += "year = " + strconv.Itoa(car.Year)
	}
	sql += " WHERE id = " + strconv.Itoa(id) + " RETURNING people_id"
	return sql
}

func buildSqlForUpdatePeople(people_id int, car models.Car) string {
	comma := false
	sql := "UPDATE people SET "
	if car.Owner.Name != "" {
		sql += "name = " + car.Owner.Name
		comma = true
	}
	if car.Owner.Surname != "" {
		if comma {
			sql += ", "
		}
		comma = true
		sql += "surname = " + car.Owner.Surname
	}
	if car.Owner.Patronymic != "" {
		if comma {
			sql += ", "
		}
		comma = true
		sql += "patronymic = " + car.Owner.Patronymic
	}
	sql += " WHERE id = " + strconv.Itoa(people_id)
	return sql
}

func buildSqlForGetCars(filters map[string]string) string {
	sql := "SELECT * FROM car JOIN people ON car.people_id = people.id"

	first := false
	for key, value := range filters {
		if !first {
			sql += " WHERE (" + key + " = " + value + ")"
			first = true
		} else {
			sql += " AND (" + key + " = " + value + ")"
		}
	}
	sql += " ODRER BY id LIMIT $1 OFFSET $2"
	return sql
}
