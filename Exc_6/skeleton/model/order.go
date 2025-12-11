package model

import (
	"fmt"
	"time"
)

const (
	orderFilename = "order_%d.md"

	// markdown template for the order receipt
	// placeholders: orderID, createdAt, drinkID, amount
	markdownTemplate = `# Order Receipt

- **Order ID:** %d
- **Created At:** %s
- **Drink ID:** %d
- **Amount:** %d
`
)

type Order struct {
	Base
	Amount uint64 `json:"amount"`
	// Relationships
	// foreign key
	DrinkID uint  `json:"drink_id" gorm:"not null"`
	Drink   Drink `json:"drink"`
}

func (o *Order) ToMarkdown() string {
	return fmt.Sprintf(markdownTemplate, o.ID, o.CreatedAt.Format(time.Stamp), o.DrinkID, o.Amount)
}

func (o *Order) GetFilename() string {
	return fmt.Sprintf(orderFilename, o.ID)
}
