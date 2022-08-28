package types

import "fmt"

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
