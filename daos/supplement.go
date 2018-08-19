package daos

import (
	"github.com/yawlhead91/nfc_scanner_prototype/app"
	"github.com/yawlhead91/nfc_scanner_prototype/models"
)

// SupplementDAO persists supplement data in database
type SupplementDAO struct{}

// NewSupplementDAO creates a new SupplementDAO
func NewSupplementDAO() *SupplementDAO {
	return &SupplementDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *SupplementDAO) Get(rs app.RequestScope, id int) (*models.Supplement, error) {
	var supplement models.Supplement
	err := rs.Tx().Select().Model(id, &supplement)
	return &supplement, err
}

// Create saves a new supplement record in the database.
// The Supplement.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *SupplementDAO) Create(rs app.RequestScope, supplement *models.Supplement) error {
	supplement.ID = 0
	return rs.Tx().Model(supplement).Insert()
}

// Update saves the changes to an Supplement in the database.
func (dao *SupplementDAO) Update(rs app.RequestScope, id int, supplement *models.Supplement) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	supplement.ID = id
	return rs.Tx().Model(supplement).Exclude("id").Update()
}

// Delete deletes an supplement with the specified ID from the database.
func (dao *SupplementDAO) Delete(rs app.RequestScope, id int) error {
	supplement, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(supplement).Delete()
}

// Count returns the number of the supplement records in the database.
func (dao *SupplementDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("supplement").Row(&count)
	return count, err
}

// Query retrieves the supplement records with the specified offset and limit from the database.
func (dao *SupplementDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Supplement, error) {
	supplement := []models.Supplement{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&supplement)
	return supplement, err
}
