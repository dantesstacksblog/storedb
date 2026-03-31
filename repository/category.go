package repository

import (
	"database/sql"
	"fmt"
	"storedb/models"

)


type CategoryRepository struct {
	db *sql.DB
}


func NewCategoryRepository(db *sql.DB) *CategoryRepository {

	return &CategoryRepository{db: db}
}



