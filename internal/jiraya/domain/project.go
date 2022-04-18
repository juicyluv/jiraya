package domain

import "time"

type Project struct {
	ProjectID   string     `json:"project_id"`
	Title       string     `json:"title"`
	CreatorID   string     `json:"creator_id"`
	Description *string    `json:"description,omitempty"`
	IconURL     *string    `json:"icon_url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	ClosedAt    *time.Time `json:"closed_at,omitempty"`
}

type GetProjectRequest struct {
	ProjectID string `json:"project_id"`
}

type GetProjectResponse struct {
	Project *Project `json:"project"`
}

type CreateProjectRequest struct {
	Title       string  `json:"title"`
	CreatorID   string  `json:"creator_id"`
	Description *string `json:"description,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
}
