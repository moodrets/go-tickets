package models

import "time"

type Ticket struct {
	Id           uint            `json:"id"`
	CreatorId    uint            `json:"creator_id"`
	InWorkUserId uint            `json:"in_work_user_id"`
	Title        string          `json:"title"`
	Status       string          `json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Comments     []TicketComment `json:"comments,omitempty"`
}
