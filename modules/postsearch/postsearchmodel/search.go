package postsearchmodel

import (
	"time"
)

type DataSearch struct {
	From       *time.Time `json:"from,omitempty" form:"from"`
	To         *time.Time `json:"to,omitempty" form:"to"`
	Caption    *string    `json:"caption,omitempty" form:"caption"`
	SearchName *string    `json:"search_name,omitempty" form:"search_name"`
}
