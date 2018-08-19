package apis

import (
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/yawlhead91/nfc_scanner_prototype/app"
	"github.com/yawlhead91/nfc_scanner_prototype/models"
)

type (
	// supplementService specifies the interface for the supplent service needed by supplementResource.
	supplementService interface {
		Get(rs app.RequestScope, id int) (*models.Artist, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Artist) (*models.Artist, error)
		Update(rs app.RequestScope, id int, model *models.Artist) (*models.Artist, error)
		Delete(rs app.RequestScope, id int) (*models.Artist, error)
	}

	// supplementResource defines the handlers for the CRUD APIs.
	supplementResource struct {
		service artistService
	}
)

// ServeSupplementResource sets up the routing of supplements endpoints and the corresponding handlers.
func ServeSupplementResource(rg *routing.RouteGroup, service supplementService) {
	r := &supplementResource{service}
	rg.Get("/supplements/<id>", r.get)
	rg.Get("/supplements", r.query)
	rg.Post("/supplements", r.create)
	rg.Put("/supplements/<id>", r.update)
	rg.Delete("/supplements/<id>", r.delete)
}

func (r *supplementResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *supplementResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *supplementResource) create(c *routing.Context) error {
	var model models.Artist
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *supplementResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *supplementResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
