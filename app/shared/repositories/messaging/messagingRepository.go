package owner

import (
	"backend/app/shared/models"
	"database/sql"

	// "log"

	_ "github.com/lib/pq"
)

type MessagingRepository struct {
	db *sql.DB
}

func NewMessagingRepository(db *sql.DB) *MessagingRepository {
	return &MessagingRepository{db}
}

// create chat
func (Repository *MessagingRepository) CreateNewChat(user1Id int, user2Id int) (*models.Chat, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO chat("user_1", "user_2") Values ($1, $2) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(user1Id, user2Id).Scan(&id)

	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.GetOwner(id) // <-- FIX THIS
}

// create message
func (Repository *MessagingRepository) CreateNewMessage(chatId int, sentFrom int, sentTo int, message, string, messageType, string, attatchment string) (*models.Message, error) {
	var id int
	stmt, err := Repository.db.Prepare(`INSERT INTO message("chat_id", "sent_from", "sentto", "message", "message_type", "attachment") Values ($1, $2, $3, $4, $5, $6) RETURNING id;`)

	if err != nil {
		return nil, err
	}

	stmtERR := stmt.QueryRow(chatId, sentFrom, sentTo, message, messageType, attatchment).Scan(&id)

	if stmtERR != nil {
		return nil, stmtERR
	}

	return Repository.GetOwner(id) // <-- FIX THIS
}
