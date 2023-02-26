package models

import "time"

type TicketComment struct {
	Id        uint      `json:"id"`
	TicketId  uint      `json:"ticket_id"`
	AuthorId  uint      `json:"author_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
