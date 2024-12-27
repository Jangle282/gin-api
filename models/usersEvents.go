package models

import (
	"gin-api/database"
)

type UsersEvent struct {
	ID      int64
	UserId  int64 `binding:"required"`
	EventId int64 `binding:"required"`
}

func (uE *UsersEvent) Save() error {
	query := `INSERT INTO users_events (user_id, event_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uE.UserId, uE.EventId)

	if err != nil {
		return err
	}

	return err
}

func (uE *UsersEvent) Delete() error {
	query := `DELETE FROM users_events WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uE.UserId, uE.EventId)

	return err
}

func GetUsersEvent(eventId, userId int64) (*UsersEvent, error) {
	query := `SELECT * FROM users_events WHERE event_id=? AND user_id=?`

	row := db.DB.QueryRow(query, eventId, userId)

	var usersEvent UsersEvent
	err := row.Scan(&usersEvent.ID, &usersEvent.EventId, &usersEvent.UserId)

	if err != nil {
		// type return is *Event = null pointer for Event is nil,
		return nil, err
	}

	return &usersEvent, nil
}
