package types

import "fmt"

type Items struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (i *Items) String() string {
	return fmt.Sprintf("Items<ID=%d, Content=%s>", i.ID, i.Content)
}
