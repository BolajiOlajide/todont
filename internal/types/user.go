package types

import "fmt"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Active    bool   `json:"active"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (u *User) String() string {
	return fmt.Sprintf("User<ID=%d, Name=%s>", u.ID, u.Name)
}
