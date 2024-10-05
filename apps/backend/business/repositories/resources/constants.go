package resources

type ResourceType string

var (
	ResourceTeams            ResourceType = "teams"
	ResourceTeamMembers      ResourceType = "team_members"
	ResourceBillings         ResourceType = "billings"
	ResourceGlobalAPITokens  ResourceType = "global_api_tokens"
	ResourceTeamAPITokens    ResourceType = "team_api_tokens"
	ResourceMonitors         ResourceType = "monitors"
	ResourceHeartbeats       ResourceType = "heartbeats"
	ResourceIntegrations     ResourceType = "integrations"
	ResourceIncidents        ResourceType = "incidents"
	ResourceInvitations      ResourceType = "invitations"
	ResourceStatusPages      ResourceType = "status_pages"
	ResourceEscalationPolicy ResourceType = "escalation_policy"
	ResourceOnCallCalenders  ResourceType = "on_call_calenders"
	ResourceSources          ResourceType = "sources"
	ResourceDashboards       ResourceType = "dashboards"
)

func FromResourceType(s ResourceType) string {
	return string(s)
}

func ToResourceType(s string) ResourceType {
	return ResourceType(s)
}
