package storage

import (
	"context"
	"goTest/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Carer interface {
	CreateCars(ctx context.Context, cars []models.Car) error
	DeleteCar(ctx context.Context, id int) error
	UpdateCar(ctx context.Context, id int, car models.Car) error
	GetCars(ctx context.Context, filters map[string]string, limit, offset int) ([]models.Car, error)
}
type CarStorage struct {
	pool *pgxpool.Pool
}

func NewCarStorage(pool *pgxpool.Pool) Carer {
	return &CarStorage{pool}
}

func (c *CarStorage) CreateCars(ctx context.Context, cars []models.Car) error {
	var id int
	// Создаем транзакцию
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}
	// Добавляем defer, чтобы была возможность откатить транзакцию в случае ошибки
	defer tx.Rollback(ctx)

	for _, car := range cars {
		// Выполняем первый запрос
		sql := "INSERT INTO people (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id"
		err = tx.QueryRow(ctx, sql, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic).Scan(&id)
		if err != nil {
			return err
		}

		// Выполняем второй запрос
		sql = "INSERT INTO car (regNum, mark, model, year, people_id) VALUES ($1, $2, $3, $4, $5)"
		_, err = tx.Exec(ctx, sql, car.RegNum, car.Mark, car.Model, car.Year, id)
		if err != nil {
			return err
		}
	}
	// Если все запросы выполнены успешно, утверждаем транзакцию
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *CarStorage) DeleteCar(ctx context.Context, id int) error {
	var people_id int

	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	sql := "DELETE FROM car WHERE id = $1 RETURNING people_id"
	err = tx.QueryRow(ctx, sql, id).Scan(&people_id)
	if err != nil {
		return err
	}

	sql = "DELETE FROM people WHERE id = $1"
	_, err = tx.Exec(ctx, sql, people_id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *CarStorage) UpdateCar(ctx context.Context, id int, car models.Car) error {
	var people_id int

	sql := buildSqlForUpdateCar(id, car)
	if err := c.pool.QueryRow(ctx, sql).Scan(&people_id); err != nil {
		return err
	}
	sql = buildSqlForUpdatePeople(people_id, car)
	if _, err := c.pool.Exec(ctx, sql); err != nil {
		return err
	}
	return nil
}

func (c *CarStorage) GetCars(ctx context.Context, filters map[string]string, limit, offset int) ([]models.Car, error) {

	sql := buildSqlForGetCars(filters)

	rows, err := c.pool.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]models.Car, 0)

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.ID, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}
