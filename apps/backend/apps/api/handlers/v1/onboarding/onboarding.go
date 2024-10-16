package onboarding_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type Config struct {
	UsersRepo users.Repository
	Config    *config.OpenpulseAPIConfig
}

type handler struct {
	usersRepo users.Repository
	config    *config.OpenpulseAPIConfig
}

func New(cfg Config) *handler {
	return &handler{config: cfg.Config, usersRepo: cfg.UsersRepo}
}

func (h *handler) CreateOrganization(w http.ResponseWriter, r *http.Request) error {
	var payload CreateOrganizationInput
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	claims := auth.GetClaims(r.Context())
	orgId, err := h.usersRepo.CreateOrganization(r.Context(), users.NewOrganization{
		AdminId:        claims.UserId,
		Name:           payload.Name,
		Description:    payload.Description,
		LogoURL:        payload.LogoURL,
		Designation:    payload.Designation,
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
		CreateOrganizationResponse{OrgId: orgId},
	)
}

func (h *handler) CreateTeam(w http.ResponseWriter, r *http.Request) error {
	var input CreateTeamInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	claims := auth.GetClaims(r.Context())
	code, err := nanoid.New()
	if err != nil {
		return err
	}

	teamId, err := h.usersRepo.CreateTeam(
		r.Context(),
		users.NewTeam{
			InvitationCode: code,
			OrgId:          input.OrgId,
			CreatorId:      claims.UserId,
			CreatorRoleId:  claims.RoleId,
			Name:           input.TeamName,
			Description:    input.TeamDescription,
			UserRBAC:       []users.UserRBAC{},
		},
	)
	if err != nil {
		return err
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Details saved successfully.",
		CreateTeamResponse{TeamId: teamId},
	)
}
