package repository

import "github.com/jmoiron/sqlx"

type CustomRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomRepositoryDB(db *sqlx.DB) CustomRepositoryDB {
	return CustomRepositoryDB{db: db}
}

func (r CustomRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r CustomRepositoryDB) GetById(id int) (*Customer, error) {

	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
