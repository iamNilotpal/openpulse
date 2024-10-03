package resources

type ResourceType string

var (
	ResourceTeams            = "teams"
	ResourceTeamMembers      = "team_members"
	ResourceBillings         = "billings"
	ResourceGlobalAPITokens  = "global_api_tokens"
	ResourceTeamAPITokens    = "team_api_tokens"
	ResourceMonitors         = "monitors"
	ResourceHeartbeats       = "heartbeats"
	ResourceIntegrations     = "integrations"
	ResourceIncidents        = "incidents"
	ResourceInvitations      = "invitations"
	ResourceStatusPages      = "status_pages"
	ResourceEscalationPolicy = "escalation_policy"
	ResourceOnCallCalenders  = "on_call_calenders"
	ResourceSources          = "sources"
	ResourceDashboards       = "dashboards"
)

func FromResourceType(s ResourceType) string {
	return string(s)
}

func ToResourceType(s string) ResourceType {
	return ResourceType(s)
}
