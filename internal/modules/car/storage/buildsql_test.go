package storage

import (
	"goTest/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSqlForUpdateCar(t *testing.T) {
	assert := assert.New(t)

	car := models.Car{RegNum: "XYZ", Mark: "Ford", Model: "Focus", Year: 2020}
	want := "UPDATE car SET regNum = XYZ, mark = Ford, model = Focus, year = 2020 WHERE id = 1 RETURNING people_id"

	got := buildSqlForUpdateCar(1, car)

	assert.Equal(want, got, "Should result in SQL query for updating car")

	car = models.Car{Model: "Focus", Year: 2020}
	want = "UPDATE car SET model = Focus, year = 2020 WHERE id = 1 RETURNING people_id"

	got = buildSqlForUpdateCar(1, car)

	assert.Equal(want, got, "Should result in SQL query for updating car")
}

func TestBuildSqlForUpdatePeople(t *testing.T) {
	assert := assert.New(t)

	car := models.Car{Owner: models.Owner{Name: "John", Surname: "Doe", Patronymic: "Patronym"}}
	// Примечание - судя по исходному коду, поля внутри структуры car.Owner присваиваются некорректно
	// Связать поля models.Car с даннутми в models.People могут быть нарушением SRP

	want := "UPDATE people SET name = John, surname = Doe, patronymic = Patronym WHERE id = 1" // Установите ожидаемую строку SQL здесь
	got := buildSqlForUpdatePeople(1, car)

	assert.Equal(want, got, "Should result in SQL query for updating people")

	// Добавьте больше тестов, если необходимо
	car = models.Car{Owner: models.Owner{Patronymic: "Patronym"}}
	// Примечание - судя по исходному коду, поля внутри структуры car.Owner присваиваются некорректно
	// Связать поля models.Car с даннутми в models.People могут быть нарушением SRP

	want = "UPDATE people SET patronymic = Patronym WHERE id = 1" // Установите ожидаемую строку SQL здесь
	got = buildSqlForUpdatePeople(1, car)

	assert.Equal(want, got, "Should result in SQL query for updating people")
}

func TestBuildSqlForGetCarsWithTestify(t *testing.T) {
	assert := assert.New(t)
	filters := map[string]string{
		"regNum":     "123",
		"mark":       "BMW",
		"model":      "X5",
		"year":       "2018",
		"name":       "John",
		"surname":    "Doe",
		"patronymic": "Jr",
	}
	expected := "SELECT * FROM car JOIN people ON car.people_id = people.id WHERE (regNum = 123) AND (mark = BMW) AND (model = X5) AND (year = 2018) AND (name = John) AND (surname = Doe) AND (patronymic = Jr) ODRER BY id LIMIT $1 OFFSET $2"
	result := buildSqlForGetCars(filters)
	assert.Equal(expected, result, "The two SQL strings should be the same.")
}
