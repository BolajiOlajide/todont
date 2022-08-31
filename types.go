package main

import "fmt"

// Item represents an action you don't want to do anymore
type Item struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (i *Item) String() string {
	return fmt.Sprintf("Items<ID=%d, Content=%s>", i.ID, i.Content)
}

// User represents an entity that can log in to the Todont app.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("User<ID=%d, Name=%s>", u.ID, u.Name)
}
