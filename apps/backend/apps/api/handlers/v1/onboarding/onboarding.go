package onboarding_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/organizations"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	OrgRepo organizations.Repository
}

type handler struct {
	orgRepo organizations.Repository
}

func New(cfg Config) *handler {
	return &handler{orgRepo: cfg.OrgRepo}
}

func (h *handler) CreateOrganization(w http.ResponseWriter, r *http.Request) error {
	var payload OnboardingOrganizationPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	claims := auth.GetClaims(r.Context())
	orgId, err := h.orgRepo.Create(r.Context(), organizations.NewOrganization{
		Name:           payload.Name,
		AdminId:        claims.UserId,
		Description:    payload.Description,
		LogoURL:        payload.LogoURL,
		TotalEmployees: payload.MembersCount,
	})

	if err != nil {
		if ok := database.CheckPQError(
			err,
			func(err *pq.Error) bool {
				return err.Column == "admin_id" && err.Code == pgerrcode.UniqueViolation
			},
		); ok {
			return errors.NewRequestError(
				"One user can create only one organization.",
				http.StatusBadRequest,
				errors.DuplicateValue,
			)
		}

		return errors.NewRequestError(
			"Error while creating organization.",
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Organization created successfully.",
		OnboardingOrganizationResponse{OrgId: orgId},
	)
}

func (h *handler) CreateTeam(w http.ResponseWriter, r *http.Request) error {
	var payload OnboardingTeamPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Details saved successfully.",
		OnboardingOrganizationResponse{OrgId: 0},
	)
}
