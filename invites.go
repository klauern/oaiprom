package openaiorgs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Invite struct {
	ObjectType string       `json:"object"`
	ID         string       `json:"id"`
	Email      string       `json:"email"`
	Role       string       `json:"role"`
	Status     string       `json:"status"`
	CreatedAt  UnixSeconds  `json:"created_at"`
	ExpiresAt  UnixSeconds  `json:"expires_at"`
	AcceptedAt *UnixSeconds `json:"accepted_at,omitempty"`
}

const InviteListEndpoint = "/organization/invites"

func (c *Client) ListInvites() ([]Invite, error) {
	// Get the raw response
	rawResp, err := c.httpClient.R().Get(InviteListEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get invites: %w", err)
	}

	// Read and log the raw response body
	body, err := io.ReadAll(bytes.NewReader(rawResp.Body()))
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Attempt to parse the response
	var resp ListResponse[Invite]
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return resp.Data, nil
}

func (c *Client) CreateInvite(email string, role RoleType) (*Invite, error) {
	body := map[string]string{
		"email": email,
		"role":  string(role),
	}

	invite, err := Post[Invite](c.httpClient, InviteListEndpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create invite: %w", err)
	}

	return invite, nil
}

func (c *Client) RetrieveInvite(id string) (*Invite, error) {
	resp, err := GetSingle[Invite](c.httpClient, InviteListEndpoint+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve invite: %w", err)
	}

	return resp, nil
}

func (c *Client) DeleteInvite(id string) error {
	err := Delete[Invite](c.httpClient, InviteListEndpoint+"/"+id)
	if err != nil {
		return fmt.Errorf("failed to delete invite: %w", err)
	}

	return nil
}
