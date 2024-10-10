package resources

import (
	"net/http"
	"slices"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type AppResource string

const (
	ResourceTeams            AppResource = "teams"
	ResourceTeamMembers      AppResource = "team_members"
	ResourceBillings         AppResource = "billings"
	ResourceGlobalAPITokens  AppResource = "global_api_tokens"
	ResourceTeamAPITokens    AppResource = "team_api_tokens"
	ResourceMonitors         AppResource = "monitors"
	ResourceHeartbeats       AppResource = "heartbeats"
	ResourceIntegrations     AppResource = "integrations"
	ResourceIncidents        AppResource = "incidents"
	ResourceInvitations      AppResource = "invitations"
	ResourceStatusPages      AppResource = "status_pages"
	ResourceEscalationPolicy AppResource = "escalation_policy"
	ResourceOnCallCalenders  AppResource = "on_call_calenders"
	ResourceSources          AppResource = "sources"
	ResourceDashboards       AppResource = "dashboards"
)

const (
	ResourceIntTeams int = iota + 1
	ResourceIntTeamMembers
	ResourceIntBillings
	ResourceIntGlobalAPITokens
	ResourceIntTeamAPITokens
	ResourceIntMonitors
	ResourceIntHeartbeats
	ResourceIntIntegrations
	ResourceIntIncidents
	ResourceIntInvitations
	ResourceIntStatusPages
	ResourceIntEscalationPolicy
	ResourceIntOnCallCalenders
	ResourceIntSources
	ResourceIntDashboards
)

var resources = []AppResource{
	ResourceTeams,
	ResourceTeamMembers,
	ResourceBillings,
	ResourceGlobalAPITokens,
	ResourceTeamAPITokens,
	ResourceMonitors,
	ResourceHeartbeats,
	ResourceIntegrations,
	ResourceIncidents,
	ResourceInvitations,
	ResourceStatusPages,
	ResourceEscalationPolicy,
	ResourceOnCallCalenders,
	ResourceSources,
	ResourceDashboards,
}

var resourcesMap = map[AppResource]int{
	ResourceTeams:            ResourceIntTeams,
	ResourceTeamMembers:      ResourceIntTeamMembers,
	ResourceBillings:         ResourceIntBillings,
	ResourceGlobalAPITokens:  ResourceIntGlobalAPITokens,
	ResourceTeamAPITokens:    ResourceIntTeamAPITokens,
	ResourceMonitors:         ResourceIntMonitors,
	ResourceHeartbeats:       ResourceIntHeartbeats,
	ResourceIntegrations:     ResourceIntIntegrations,
	ResourceIncidents:        ResourceIntIncidents,
	ResourceInvitations:      ResourceIntInvitations,
	ResourceStatusPages:      ResourceIntStatusPages,
	ResourceEscalationPolicy: ResourceIntEscalationPolicy,
	ResourceOnCallCalenders:  ResourceIntOnCallCalenders,
	ResourceSources:          ResourceIntSources,
	ResourceDashboards:       ResourceIntDashboards,
}

var resourcesReverseMap = map[int]AppResource{
	ResourceIntTeams:            ResourceTeams,
	ResourceIntTeamMembers:      ResourceTeamMembers,
	ResourceIntBillings:         ResourceBillings,
	ResourceIntGlobalAPITokens:  ResourceGlobalAPITokens,
	ResourceIntTeamAPITokens:    ResourceTeamAPITokens,
	ResourceIntMonitors:         ResourceMonitors,
	ResourceIntHeartbeats:       ResourceHeartbeats,
	ResourceIntIntegrations:     ResourceIntegrations,
	ResourceIntIncidents:        ResourceIncidents,
	ResourceIntInvitations:      ResourceInvitations,
	ResourceIntStatusPages:      ResourceStatusPages,
	ResourceIntEscalationPolicy: ResourceEscalationPolicy,
	ResourceIntOnCallCalenders:  ResourceOnCallCalenders,
	ResourceIntSources:          ResourceSources,
	ResourceIntDashboards:       ResourceDashboards,
}

func ParseAppResourceString(s string) (AppResource, error) {
	if contains := slices.Contains(resources, AppResource(s)); contains {
		return AppResource(s), nil
	}
	return "", errors.NewRequestError("Invalid resource.", http.StatusBadRequest, errors.BadRequest)
}

func ParseAppResource(s AppResource) int {
	return resourcesMap[s]
}

func ParseAppResourceInt(s int) AppResource {
	return resourcesReverseMap[s]
}
