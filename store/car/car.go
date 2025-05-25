package car

import (
	"context"
	"database/sql"

	"github.com/yuvrajinbhakti/cardock---car-managment-system/models"
)

type Store struct {	//create a struct for the store
	db *sql.DB
}

//creating store instance with db connection
func new(db *sql.DB) Store{ //create a new store
	return Store{db: db}
}

// functions with db

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	// TODO: Implement car retrieval logic
	return models.Car{}, nil
}


func(s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) (models.Car, error){
	// TODO: Implement car retrieval logic
	return models.Car{}, nil
}

func(s Store) CreateCar(ctx context.Context, carReq *models.CarRequest)(models.Car, error){
	// TODO: Implement car creation logic
	return models.Car{}, nil
}

func(s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest)(models.Car, error){
	// TODO: Implement car update logic
	return models.Car{}, nil
}

func(s Store) DeleteCar(ctx context.Context, id string)(error){
	// TODO: Implement car deletion logic
	return nil
}

