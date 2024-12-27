package models

import (
	"gin-api/database"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Begin       time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(
					name,
					description,
					location,
					begin,
					user_id
				) VALUES (?,?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Begin, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Begin, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	//When you know you will get back only 1 row.
	row := db.DB.QueryRow(query, id)
	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Begin, &event.UserId)

	if err != nil {
		// type return is *Event = null pointer for Event is nil,
		return nil, err
	}

	return &event, nil
}

func (e *Event) UpdateEvent() error {
	query := `UPDATE events
			SET name=?, description=?, location=?, begin=?, user_id=?
			WHERE id=?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Begin, e.UserId, e.ID)

	return err
}

func (e *Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id=?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}
