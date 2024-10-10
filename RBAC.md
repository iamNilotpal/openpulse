# Role-Based Access Control (RBAC) Guidelines

This document outlines the Role-Based Access Control (RBAC) model implemented within the system, detailing the roles, permissions, actions, and resource access that define the organization’s security framework. Roles determine what actions users can perform on various resources, ensuring that access rights align with their responsibilities.

## Roles Overview

Roles define the access levels and permissions granted to users. Each role is designed to ensure that users can only perform actions necessary for their specific responsibilities within the organization.

### 1. **Org Admin**

The **Org Admin** role is the top-level administrator with full access and control over all aspects of the organization. They can manage everything across teams and organization-level settings.

- **Permissions**: Full access to create, view, update, delete, and manage all organizational resources.
- **Resources**: The Org Admin can control billing, dashboards, global API tokens, heartbeats, incidents, integrations, monitors, on-call calendars, escalation policies, sources, status pages, team members, and teams, as well as overall organizational settings.

**Responsibilities**:

- Full oversight and configuration of all organization and team resources.
- Setting up and managing the organization’s global infrastructure.
- Managing users, roles, teams, and their associated permissions.
- Ensuring proper billing and payment setup for the organization.
- Defining global API tokens and monitoring solutions.
- Managing security levels, escalations, and on-call scheduling for all teams.

### 2. **Admin**

The **Admin** role shares most responsibilities with the Org Admin, but with a focus on team-level management and more granular control over specific resources.

- **Permissions**: Similar to Org Admin, with full access to create, view, update, delete, and manage resources at both the organizational and team levels.
- **Resources**: Admins can control billing, dashboards, global API tokens, heartbeats, incidents, integrations, monitors, on-call calendars, escalation policies, security levels, sources, status pages, team members, teams, and overall organizational settings.

**Responsibilities**:

- Manage teams, including team composition, access controls, and team-specific settings.
- Oversee global and team-specific monitoring solutions, such as monitors, heartbeats, and incidents.
- Ensure proper billing and payment management.
- Set up integrations, status pages, and dashboards for visibility into operations and incident response.

### 3. **Billing Admin**

The **Billing Admin** role is focused exclusively on managing billing information and payments for the organization.

- **Permissions**: Limited to the billing resources.
- **Resources**: Can view, update, and manage billing and payment information, ensuring the organization’s financial settings are correct.

**Responsibilities**:

- Oversee and manage invoices, payment methods, and billing history for the organization.
- Ensure accurate and up-to-date financial records for the team and organization.

### 4. **Team Lead**

The **Team Lead** role is responsible for managing team-specific operations and resources. They have a broader range of permissions than most team members but are restricted to actions related to their team.

- **Permissions**: Full access to create, view, update, delete, and manage team-specific resources.
- **Resources**: Can manage billing, global API tokens, dashboards, heartbeats, incidents, integrations, monitors, on-call calendars, escalation policies, security levels, sources, status pages, and team members. The Team Lead can also manage teams at the organizational level.

**Responsibilities**:

- Manage team members and their roles within the team.
- Set up and maintain monitoring, incident response, and on-call scheduling for their team.
- Oversee team billing and payment processes.
- Ensure that team-specific integrations, dashboards, and status pages are configured and updated.

### 5. **Responder**

The **Responder** role is focused on handling incidents and team-related resources. This role has permissions to act quickly on alerts and issues, ensuring incidents are resolved.

- **Permissions**: Can view, update, and manage certain resources related to incidents and team operations.
- **Resources**: Access to billing, team members, teams, dashboards, heartbeats, incidents, integrations, monitors, on-call calendars, escalation policies, security levels, sources, and status pages.

**Responsibilities**:

- Respond to incidents and alerts within their team.
- Monitor the health of services and ensure proper incident management workflows are followed.
- Update dashboards and status pages to reflect the current state of incidents.
- Work with team members to ensure smooth operations and incident resolution.

### 6. **Member**

The **Member** role is designed for general team members who have a limited scope of permissions. They contribute to team activities but with restricted access to critical resources.

- **Permissions**: Limited access to view and update certain resources.
- **Resources**: Access to billing, team members, teams, dashboards, and sources.

**Responsibilities**:

- Participate in team tasks, working on assigned resources such as dashboards and sources.
- Assist in maintaining team visibility through dashboards and reporting tools.
- Follow team-specific workflows and updates.

---

Certainly! Here's a more detailed explanation of **Permission Actions** and **Resources** without the SQL.

---

## Permission Actions

Permission actions define what roles can do with resources. These actions are key to managing access control within the system. Each role is assigned one or more of these actions, dictating how they can interact with a given resource.

### 1. **View**

The `view` action allows users to **see** the details of a resource without making any changes. This is the most basic level of access and is typically used to ensure that users can monitor or inspect a resource without the ability to alter its state.

- **Example**: A user with the `view` permission on a **dashboard** can see reports and visualizations but cannot change the data sources or customize the layout.

### 2. **Create**

The `create` action allows users to **add new resources** to the system. This action gives users the ability to introduce new entities or data, but it doesn’t necessarily allow them to modify existing ones.

- **Example**: A **Team Lead** with `create` permissions on **monitors** can set up new monitoring configurations for services but cannot change or delete existing monitors unless they have additional permissions.

### 3. **Update**

The `update` action allows users to **modify existing resources**. This permission is crucial when users need to make adjustments to resources they manage. It typically includes the ability to change resource configurations, settings, or data.

- **Example**: A **Responder** with `update` permissions on **incidents** can modify the status of an ongoing incident or update the details of how the incident is being handled.

### 4. **Delete**

The `delete` action grants the ability to **remove resources** from the system. This is a powerful permission that should be carefully assigned, as it involves the permanent removal of data or configurations.

- **Example**: An **Admin** with `delete` permissions on **team members** can remove individuals from the organization or a team, effectively revoking their access to all associated resources.

### 5. **Manage**

The `manage` action is the highest level of permission. It allows users not only to create, view, update, and delete a resource but also to **configure settings**, assign roles, and adjust access controls for other users. Managing a resource typically includes controlling how others can interact with it.

- **Example**: A **Team Admin** with `manage` permissions on **team members** can add new members, assign roles, modify their permissions, and remove them from the team. This gives them full control over team composition and governance.

---

## Resources

Resources represent the various components or entities within the organization that users interact with. These can range from teams and team members to more complex entities like dashboards, incidents, and monitoring configurations. Below is a detailed description of the different resource types:

### 1. **Teams**

Teams represent groups of users working together on specific projects or tasks. Managing teams involves adding members, setting roles, and organizing workflows. Teams are a foundational resource in the organization.

- **Example**: A **Team Lead** can manage team composition, adding or removing members as needed.

### 2. **Team Members**

Team members are the individuals who are part of a team. Permissions for this resource typically revolve around adding new members, modifying their roles, or removing them from the team.

- **Example**: A **Team Admin** can manage the roles of individual team members, changing their access rights based on their responsibilities.

### 3. **Billing**

Billing covers the financial aspects of the organization, including payments, invoices, and billing details. Only certain roles, like **Billing Admin** or **Org Admin**, typically have access to billing-related resources.

- **Example**: A **Billing Admin** can manage payment methods, review invoices, and handle financial disputes.

### 4. **Global API Tokens**

Global API tokens allow external applications or systems to interact with the organization’s services at a high level. These tokens often grant broad access, so managing them is critical for security.

- **Example**: An **Org Admin** can create or revoke global API tokens that grant access to key services across the organization.

### 5. **Team API Tokens**

Team API tokens are similar to global API tokens but are scoped to individual teams. These are used to control programmatic access to team-specific resources.

- **Example**: A **Team Lead** can generate a token that allows external systems to pull data from the team’s monitors or dashboards.

### 6. **Monitors**

Monitors are used to track the performance and availability of services or applications. Permissions for monitors typically involve creating, modifying, or deleting configurations that trigger alerts based on specific conditions.

- **Example**: A **Responder** can modify a monitor to adjust its sensitivity or add new services to be monitored.

### 7. **Heartbeats**

Heartbeats are periodic checks that ensure a system or service is running as expected. If a heartbeat check fails, it can trigger an alert, signaling potential downtime. Managing heartbeats is critical for ensuring system uptime.

- **Example**: An **Admin** can configure how often heartbeats check for system health and what actions to take if a service is unresponsive.

### 8. **Integrations**

Integrations are connections with third-party tools or services (e.g., Slack, PagerDuty) that help streamline workflows and communication. Permissions related to integrations allow users to set up, modify, or remove these connections.

- **Example**: A **Team Lead** can configure an integration to send notifications to a Slack channel when a new incident is created.

### 9. **Incidents**

Incidents represent issues or problems that need to be addressed, typically related to system performance or failures. Permissions for incidents involve creating, updating, managing, or resolving incidents as they occur.

- **Example**: A **Responder** can update the status of an incident, marking it as resolved once the underlying issue has been fixed.

### 10. **Invitations**

Invitations allow new users to join teams or the organization. Permissions to manage invitations typically include sending, resending, or revoking invitations.

- **Example**: A **Team Admin** can send invitations to prospective team members, granting them access to join the team.

### 11. **Status Pages**

Status pages provide real-time or historical updates on the health and availability of services. These are often publicly accessible and used to inform users of any ongoing issues.

- **Example**: An **Admin** can configure a status page to display the current state of the organization’s services and past incidents.

### 12. **Escalation Policies**

Escalation policies define how incidents are handled within the organization. They determine the order in which responders are notified and the actions taken if the issue escalates further.

- **Example**: A **Team Lead** can create an escalation policy that notifies responders in a specified order when a monitor fails.

### 13. **On-call Calendars**

On-call calendars are used to manage the scheduling of team members who are responsible for responding to incidents. Permissions here involve setting or modifying on-call rotations and schedules.

- **Example**: A **Team Lead** can configure an on-call schedule to ensure that someone is always available to respond to incidents.

### 14. **Sources**

Sources represent the origin of data used in monitoring or reporting. They could be external data sources or internal metrics. Managing sources involves configuring where the system pulls data from.

- **Example**: A **Member** can add a new source of monitoring data to track specific metrics for the team's dashboards.

### 15. **Dashboards**

Dashboards provide a visual representation of data, metrics, and alerts in real-time. Permissions for dashboards involve creating, modifying, or viewing the layouts and data sources presented on them.

- **Example**: A **Member** can modify a team’s dashboard to display relevant metrics for tracking system performance.

---

### Summary

By assigning specific permission actions to roles, the RBAC system ensures that users can interact with resources in a controlled and secure manner. Resources, such as teams, monitors, and incidents, are fundamental components of the system, and the appropriate assignment of permissions allows for efficient and secure management of these entities.
