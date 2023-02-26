package repositories

import (
	"github.com/moodrets/go-tickets/database"
	"github.com/moodrets/go-tickets/models"
)

type TicketRepository struct{}

func (t *TicketRepository) GetList(params map[string]string) ([]models.Ticket, error) {
	client := database.Connection()
	defer client.Close()

	rows, err := client.Query("select * from tickets")

	if err != nil {
		return nil, err
	}

	var tickets []models.Ticket

	for rows.Next() {
		var ticket models.Ticket
		rows.Scan(
			&ticket.Id,
			&ticket.CreatorId,
			&ticket.InWorkUserId,
			&ticket.Title,
			&ticket.Status,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (t *TicketRepository) Create() (bool, error) {
	return true, nil
}

func (t *TicketRepository) Update() (bool, error) {
	return true, nil
}

func (t *TicketRepository) Delete() (bool, error) {
	return true, nil
}
