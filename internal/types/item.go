package types

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
