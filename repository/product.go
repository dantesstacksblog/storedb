package repository


import (
	"database/sql"
	"fmt"
	"storedb/models"
)


type ProductRepository struct {

	db *sql.DB
}


func NewProductRepository(db *sql.DB) *ProductRepository {

	return &ProductRepository{db: db}

}


// insert a new product adn return its id

func (r *ProductRepository) Create(p *models.Product) (int64, error) {
	result, err := r.db.Exec(`
		INSERT INTO products (category_id, name, price, stock, low_stock_threshold)
		VALUES (?, ?, ?, ?, ?)`,
		p.CategoryID, p.Name, p.Price, p.Stock, p.LowStockThreshold,
	
	)

	if err != nil {
		return 0, fmt.Errorf("creating product: %w", err)
	}

	return result.LastInsertId()
}

// getid fetches a single product by id

func (r *ProductRepository) GetByID(id int64) (*models.Product, error) {

	p := &models.Product{}
	err := r.db.QueryRow(`
		SELECT id, category_id, name, price, stock, low_stock_threshold, created_at
		FROM products WHERE id = ?`, id).
		Scan(&p.ID, &p.CategoryID, &p.Name, &p.Price, 
			&p.Stock, &p.LowStockThreshold, &p.CreatedAt)

	if err == sql.ErrNoRows {

		return nil, fmt.Errorf("product %d not found", id)
	}
	if err != nil {
		
		return  nil, fmt.Errorf("getting product: %w", err)
	}

	return p, nil
	
}


// updatestock sets the stock level for product 

func (r *ProductRepository) UpdateStock(id int64, stock int) error {

	_, err := r.db.Exec(`UPDATE products SET stock = ? WHERE id = ?`, stock, id)
	if err != nil {
		return fmt.Errorf("updating stock: %w", err)
	}

	return nil
}


// delete a product by id 

func (r *ProductRepository) Delete(id int64) error {

	_, err := r.db.Exec(`DELETE FROM produts WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("deleting product: %w", err)
	}

	return nil
}

// return all products below their threshold


func (r *ProductRepository) GetLowStock() ([]models.Product, error) {

	rows, err := r.db.Query(`
	

	`)
}
