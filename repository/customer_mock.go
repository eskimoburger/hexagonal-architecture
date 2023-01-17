package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{1001, "Soontorn", "Bangkok", "10110", "23051998", "Active"},
		{1002, "John", "New York", "50000", "23051998", "Active"},
	}
	return customerRepositoryMock{
		customers: customers,
	}

}
func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}
func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("Customer not found")
}
