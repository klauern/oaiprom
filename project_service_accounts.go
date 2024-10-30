package openaiorgs

import "fmt"

const ProjectServiceAccountsListEndpoint = "/organization/projects/%s/service_accounts"

type ProjectServiceAccount struct {
	Object    string      `json:"object"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Role      string      `json:"role"`
	CreatedAt UnixSeconds `json:"created_at"`
}

func (c *Client) ListProjectServiceAccounts(projectID string, limit int, after string) (*ListResponse[ProjectServiceAccount], error) {
	queryParams := make(map[string]string)
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}
	if after != "" {
		queryParams["after"] = after
	}

	return Get[ProjectServiceAccount](c.httpClient, fmt.Sprintf(ProjectServiceAccountsListEndpoint, projectID), queryParams)
}

func (c *Client) CreateProjectServiceAccount(projectID string, name string) (*ProjectServiceAccount, error) {
	body := map[string]string{"name": name}
	return Post[ProjectServiceAccount](c.httpClient, fmt.Sprintf(ProjectServiceAccountsListEndpoint, projectID), body)
}

func (c *Client) RetrieveProjectServiceAccount(projectID string, serviceAccountID string) (*ProjectServiceAccount, error) {
	return GetSingle[ProjectServiceAccount](c.httpClient, fmt.Sprintf(ProjectServiceAccountsListEndpoint+"/%s", projectID, serviceAccountID))
}

func (c *Client) DeleteProjectServiceAccount(projectID string, serviceAccountID string) error {
	return Delete[ProjectServiceAccount](c.httpClient, fmt.Sprintf(ProjectServiceAccountsListEndpoint+"/%s", projectID, serviceAccountID))
}
