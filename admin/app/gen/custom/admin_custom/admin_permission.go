package admin_custom

import "time"

type AdminPermissionMenu struct {
	ID        int32     `json:"id"`
	MenuID    int32     `json:"menu_id"`
	MenuName  string    `json:"menu_name"`
	MenuPath  string    `json:"menu_path"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	Describe  string    `json:"describe"`
	Type      string    `json:"type"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
