package onboarding_handlers

import (
	"database/sql"
	stdErrors "errors"
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/organizations"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
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
	Config           *config.APIConfig
	RoleMap          auth.RoleMappings
	Users            users.Repository
	Organizations    organizations.Repository
	AccessControlMap auth.RoleNameToAccessControlMap
}

type handler struct {
	users            users.Repository
	config           *config.APIConfig
	roleMap          auth.RoleMappings
	organizations    organizations.Repository
	accessControlMap auth.RoleNameToAccessControlMap
}

func New(cfg Config) *handler {
	return &handler{
		config:           cfg.Config,
		roleMap:          cfg.RoleMap,
		accessControlMap: cfg.AccessControlMap,
		users:            cfg.Users,
		organizations:    cfg.Organizations,
	}
}

func (h *handler) CreateOrganization(w http.ResponseWriter, r *http.Request) error {
	var payload CreateOrganizationInput
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	user := auth.GetUser(r.Context())
	orgId, err := h.users.CreateOrganization(
		r.Context(),
		users.NewOrganization{
			AdminId:        user.Id,
			Name:           payload.Name,
			Description:    payload.Description,
			LogoURL:        payload.LogoURL,
			Designation:    payload.Designation,
			TotalEmployees: payload.MembersCount,
		},
	)

	if err != nil {
		if err := database.CheckPQError(
			err,
			func(err *pq.Error) error {
				if err.Column == "admin_id" && err.Code == pgerrcode.UniqueViolation {
					return errors.NewRequestError(
						"One user can create only one organization.",
						http.StatusBadRequest,
						errors.DuplicateValue,
					)
				}
				return nil
			},
		); err != nil {
			return err
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

	user := auth.GetUser(r.Context())
	code, err := nanoid.New()
	if err != nil {
		return err
	}

	org, err := h.organizations.QueryById(r.Context(), input.OrgId)
	if err != nil {
		if stdErrors.Is(err, sql.ErrNoRows) {
			return errors.NewRequestError("Organization not found.", http.StatusNotFound, errors.NotFound)
		}
		return err
	}

	if org.Admin.Id != user.Id {
		return errors.NewRequestError("Organization not found.", http.StatusForbidden, errors.Forbidden)
	}

	userRBAC := make([]users.UserRBAC, 0)
	resources := h.accessControlMap[roles.RoleOrgAdmin]
	admin := h.roleMap.ByName[roles.RoleOrgAdmin]

	for _, resPerms := range resources {
		for _, permission := range resPerms.Permissions {
			userRBAC = append(
				userRBAC,
				users.UserRBAC{
					UserId:       user.Id,
					RoleId:       admin.Id,
					PermissionId: permission.Id,
					ResourceId:   resPerms.Resource.Id,
				},
			)
		}
	}

	teamId, err := h.users.CreateTeam(
		r.Context(),
		users.NewTeam{
			InvitationCode: code,
			CreatorId:      user.Id,
			OrgId:          input.OrgId,
			CreatorRoleId:  user.Role.Id,
			Name:           input.TeamName,
			Description:    input.TeamDescription,
			CreatorRBAC:    userRBAC,
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
