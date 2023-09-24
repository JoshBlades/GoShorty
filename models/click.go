package models

type Click struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	LinkReferer int    `json:"url_id"`
	Link        Link   `gorm:"foreignKey:LinkReferer"`
	IPAddr      string `json:"ip_address"`
}
