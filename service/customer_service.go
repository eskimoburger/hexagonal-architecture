package service

import (
	"bank/repository"
	"database/sql"
	"fmt"
	"log"
)

type customerService struct {
	custRepo repository.CustomRepository
}

func NewCustomerService(custRepo repository.CustomRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		cusResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}

		custResponses = append(custResponses, cusResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {

	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s", "Customer not found")
		}
		log.Println(err)
		return nil, err
	}
	cusResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &cusResponse, nil
}
