package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stm := `
	INSERT INTO snippets (title, content, created, expires)
	VALUES ($1, $2, CURRENT_TIMESTAMP AT TIME ZONE 'UTC', CURRENT_TIMESTAMP AT TIME ZONE 'UTC' + ($3 || ' days')::interval)
	RETURNING id;`

	var id int

	err := m.DB.QueryRow(stm, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stm := `
	SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AND id = $1;`

	var s Snippet

	row := m.DB.QueryRow(stm, id)

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stm := `
	SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
	ORDER BY id DESC
	LIMIT 10;`

	rows, err := m.DB.Query(stm)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
