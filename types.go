package main

import "fmt"

// Item represents an action you don't want to do anymore
type item struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (i *item) String() string {
	return fmt.Sprintf("Items<ID=%d, Content=%s>", i.ID, i.Content)
}

// User represents an entity that can log in to the Todont app.
type user struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *user) String() string {
	return fmt.Sprintf("User<ID=%d, Name=%s>", u.ID, u.Name)
}

type response struct {
	Status  status `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
