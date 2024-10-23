package resources_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Resources resources.Repository
}

type handler struct {
	resources resources.Repository
}

func New(cfg Config) *handler {
	return &handler{resources: cfg.Resources}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var input NewResourceInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	resource, err := resources.ParseAppResourceString(input.Resource)
	if err != nil {
		return err
	}

	id, err := h.resources.Create(r.Context(), resources.NewResource{
		Resource:    resource,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		if err := database.CheckPQError(
			err,
			func(err *pq.Error) error {
				if err.Column == "name" && err.Code == pgerrcode.UniqueViolation {
					return errors.NewRequestError(
						"Resource with same name already exists.",
						http.StatusConflict,
						errors.DuplicateValue,
					)
				}
				if err.Column == "resource" && err.Code == pgerrcode.CheckViolation {
					return errors.NewRequestError("Invalid resource.", http.StatusConflict, errors.DuplicateValue)
				}
				return nil
			},
		); err != nil {
			return err
		}
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Resource created.",
		NewResourceResponse{Id: id, Name: input.Name, Resource: string(resource)},
	)
}
