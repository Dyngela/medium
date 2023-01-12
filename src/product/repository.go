package product

import (
	"database/sql"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github/Dyngela/medium/src/utils"
	"time"
)

type Product struct {
	// In some cases the binding between the db and the model can fail
	// To prevent it we use the db:"" tag to tell the exact name of our field
	ProductId uint    `db:"product_id"`
	Name      string  `db:"name"`
	Price     float64 `db:"price"`
	// Golang doesn't handle null value, so if a field of our struct can be null
	// we have to use this kind of type
	Image     sql.NullString `db:"image"`
	IsInStock bool           `db:"is_in_stock"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`

	CategoryId uint `db:"category_id"`
}

func findProductById(id uint, c *gin.Context) (GetProductByIdResponse, error) {
	// We acquire a connection from our pool initialized in the init function
	conn, err := utils.DbPool.Acquire(c)
	// here we allow the connection to be free and use else where at the end of our function
	defer conn.Release()
	if err != nil {
		return GetProductByIdResponse{}, err
	}
	// The actual SQL request to postgres, the $1 at the end is used to replace something is your statement, here the id
	query := `select product_id, name, price, image, is_in_stock from product where product_id = $1`
	var resp GetProductByIdResponse
	// this method comes from scany, a little library who helps you getting rid of some verbosity which are source for
	// common mistake with the standard sql implementation of go
	if err = pgxscan.Get(c, conn, &resp, query, id); err != nil {
		return GetProductByIdResponse{}, err
	}
	// if everything ran smooth we can just return the row queried.
	return resp, nil
}

func createProductRepo(product Product, c *gin.Context) (string, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return "Error while creating your product", err
	}

	query := `insert into product values (null, $1, $2, $3, $4, now(), null, null, $5);`
	_, err = conn.Query(c, query,
		product.Name,
		product.Price,
		product.Image,
		product.IsInStock,
		product.CategoryId)
	if err != nil {
		return "Error while creating your product", err
	}
	return "Product successfully created", nil
}

func findAllProduct(c *gin.Context) ([]GetAllProductResponse, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return []GetAllProductResponse{}, err
	}
	query := `select product_id, name, price, image, is_in_stock from product`
	var resp []GetAllProductResponse
	if err = pgxscan.Select(c, conn, &resp, query); err != nil {
		return []GetAllProductResponse{}, err
	}
	return resp, nil
}

func updateProductRepo(product Product, c *gin.Context) (string, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return "Error while updating your product", err
	}

	query := `update product
				SET name = $2
				SET price = $3
				SET image = $4
				SET is_in_stock = $5
				SET category_id = $6
				SET updated_at = now()
				where product_id = $7`
	_, err = conn.Query(c, query,
		product.Name,
		product.Price,
		product.Image,
		product.IsInStock,
		product.CategoryId,
		product.ProductId)
	if err != nil {
		return "Error while updating your product", err
	}
	return "Product successfully updated", nil
}

func deleteProductRepo(id uint, c *gin.Context) (string, error) {
	conn, err := utils.DbPool.Acquire(c)
	defer conn.Release()
	if err != nil {
		return "Product successfully deleted", err
	}
	// We can either update deletedAt value and modify other methods to
	// not return the row marked as deleted, or we can just delete it just like here
	query := `delete from product where product_id = $1`
	_, err = conn.Query(c, query, id)
	if err != nil {
		return "Error while deleting your product", err
	}
	return "Product successfully deleted", nil
}
