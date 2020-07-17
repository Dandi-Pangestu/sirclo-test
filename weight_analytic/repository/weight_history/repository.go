package repository

import "sirclo-test/weight_analytic/models"

type Repository interface {
	FindByID(id string) (*models.WeightHistory, error)
	FindAll() (*[]models.WeightHistory, error)
	SaveOrUpdate(m *models.WeightHistory) error
	Delete(m *models.WeightHistory) error
}
