package repository

import "shopping-chart/api/v1/model"

type CustomerRepository interface {
	Create(customer model.Customer) model.Customer
}
