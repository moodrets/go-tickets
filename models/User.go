package models

import "time"

type User struct {
	Id        uint            `json:"id"`
	Login     string          `json:"login"`
	Name      string          `json:"name"`
	Comments  []TicketComment `json:"comments,omitempty"`
	Tickets   []Ticket        `json:"tickets,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
