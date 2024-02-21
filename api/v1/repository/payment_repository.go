package repository

import "shopping-chart/api/v1/model"

type PaymentRepository interface {
	Create(payment model.Payment) model.Payment
}
