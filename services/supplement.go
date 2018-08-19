package services

import (
	"github.com/yawlhead91/nfc_scanner_prototype/app"
	"github.com/yawlhead91/nfc_scanner_prototype/models"
)

// supplementDAO specifies the interface of the supplement
// DAO needed by SupplementService.
type supplementDAO interface {
	// Get returns the supplement with the specified supplement ID.
	Get(rs app.RequestScope, id int) (*models.Supplement, error)
	// Count returns the number of artists.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of supplement with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Supplement, error)
	// Create saves a new supplement in the storage.
	Create(rs app.RequestScope, artist *models.Supplement) error
	// Update updates the supplement with given ID in the storage.
	Update(rs app.RequestScope, id int, artist *models.Supplement) error
	// Delete removes the supplement with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// SupplementService provides services related with supplement.
type SupplementService struct {
	dao supplementDAO
}

// NewSupplementService creates a new SupplementService with the given artist DAO.
func NewSupplementService(dao supplementDAO) *SupplementService {
	return &SupplementService{dao}
}

// Get returns the supplement with the specified the artist ID.
func (s *SupplementService) Get(rs app.RequestScope, id int) (*models.Supplement, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new supplement.
func (s *SupplementService) Create(rs app.RequestScope, model *models.Supplement) (*models.Supplement, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}

// Update updates the artist with the specified ID.
func (s *SupplementService) Update(rs app.RequestScope, id int, model *models.Supplement) (*models.Supplement, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the Supplement with the specified ID.
func (s *SupplementService) Delete(rs app.RequestScope, id int) (*models.Supplement, error) {
	supplement, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return supplement, err
}

// Count returns the number of supplements.
func (s *SupplementService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the supplements with the specified offset and limit.
func (s *SupplementService) Query(rs app.RequestScope, offset, limit int) ([]models.Supplement, error) {
	return s.dao.Query(rs, offset, limit)
}
