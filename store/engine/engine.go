package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/yuvrajinbhakti/cardock---car-managment-system/models"
)
type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e *EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine 
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}
	defer func(){
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v", rbErr)
			} else {
				if cmErr := tx.Commit(); cmErr != nil {
					fmt.Printf("Error committing transaction: %v", cmErr)
				}
			}
		} 
	}()
	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1", id).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, errors.New("engine not found")
		}
		return engine, err
	}
	return engine, nil
}



func (e *EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func(){	
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v", rbErr)
			} else {
				if cmErr := tx.Commit(); cmErr != nil {
					fmt.Printf("Error committing transaction: %v", cmErr)
				}
			}
		}
	}()

	engineId := uuid.New()
	


	query := `INSERT INTO engine (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)`
	_,err = tx.ExecContext(ctx, query, engineId, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange)
	if err != nil {
		return models.Engine{}, err
	}
	engine := models.Engine{
		EngineID: engineId,
		Displacement: engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange: engineReq.CarRange,
	}
	return engine, nil
}

func (e *EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	engineId, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("Invalid engine ID: %v", err)
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func(){
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v", rbErr)
			} else {
				if cmErr := tx.Commit(); cmErr != nil {
					fmt.Printf("Error committing transaction: %v", cmErr)
				}
			}
		}
	}()

	results, err := tx.ExecContext(ctx, "UPDATE engine SET displacement = $1, no_of_cylinders = $2, car_range = $3 WHERE id = $4", engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, engineId)
	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No new rows updated")
	}
	engine := models.Engine{
		EngineID: engineId,
		Displacement: engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange: engineReq.CarRange,
	}
	return engine, nil
}


func (e *EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	engineId, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("Invalid engine ID: %v", err)
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func(){
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v", rbErr)
			} else {
				if cmErr := tx.Commit(); cmErr != nil {
					fmt.Printf("Error committing transaction: %v", cmErr)
				}
			}
		}
	}()
	results, err := tx.ExecContext(ctx, "DELETE FROM engine WHERE id = $1", engineId)
	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No rows deleted")
	}
	engine := models.Engine{
		EngineID: engineId,
	}
	return engine, nil
}
