package onboarding_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Config struct{}

type handler struct{}

func New(cfg Config) *handler {
	return &handler{}
}

func (h *handler) SaveOrganizationDetails(w http.ResponseWriter, r *http.Request) error {
	var payload OnboardingOrganizationPayload
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

func (h *handler) SaveTeamDetails(w http.ResponseWriter, r *http.Request) error {
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
