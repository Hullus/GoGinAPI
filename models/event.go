package models

import (
	"MaxRestAPI/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" ,binding:"required"`
	Description string    `json:"description" ,binding:"required"`
	Location    string    `json:"location" ,binding:"required"`
	DateTime    time.Time `json:"date_time" ,binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, err
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEventByID(e Event, eventID int64) error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ? `

	_, err := db.DB.Exec(query, e.Title, e.Description, e.Location, e.DateTime, e.UserID, eventID)
	if err != nil {
		return err
	}

	return err
}

func (e Event) DeleteByID() error {
	query := `DELETE FROM events WHERE id = ?`

	_, err := db.DB.Exec(query, e.ID)
	if err != nil {
		return err
	}
	return err
}
