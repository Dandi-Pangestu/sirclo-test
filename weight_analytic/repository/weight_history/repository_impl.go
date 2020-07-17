package repository

import (
	"sirclo-test/weight_analytic/models"

	"github.com/hashicorp/go-memdb"
	uuid "github.com/satori/go.uuid"
)

type weightHistoryRepository struct {
	TableName string
	DB        *memdb.MemDB
}

func NewWeightHistoryRepository(db *memdb.MemDB) Repository {
	return &weightHistoryRepository{
		TableName: "weight_histories",
		DB:        db,
	}
}

func (r *weightHistoryRepository) FindByID(id string) (*models.WeightHistory, error) {
	txn := r.DB.Txn(false)
	defer txn.Abort()

	raw, err := txn.First(r.TableName, "id", id)
	if err != nil {
		return nil, err
	}

	return raw.(*models.WeightHistory), nil
}

func (r *weightHistoryRepository) FindAll() (*[]models.WeightHistory, error) {
	txn := r.DB.Txn(false)
	defer txn.Abort()

	it, err := txn.Get(r.TableName, "id")
	if err != nil {
		return nil, err
	}

	results := make([]models.WeightHistory, 0)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		result := obj.(*models.WeightHistory)
		results = append(results, *result)
	}

	return &results, nil
}

func (r *weightHistoryRepository) SaveOrUpdate(m *models.WeightHistory) error {
	txn := r.DB.Txn(true)
	defer txn.Abort()

	uuid4, _ := uuid.NewV4()
	m.ID = uuid4.String()
	if err := txn.Insert(r.TableName, m); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (r *weightHistoryRepository) Delete(m *models.WeightHistory) error {
	txn := r.DB.Txn(true)
	defer txn.Abort()

	if err := txn.Delete(r.TableName, m); err != nil {
		return err
	}

	txn.Commit()

	return nil
}
