package models

import (
	"net/url"
	"time"
)

type Link struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	Enabled    bool   `json:"enabled"`
	OriginURL  string `json:"origin_url"`
	ShortCode  string `json:"short_code"`
	ClickCount uint   `json:"click_count"`
}

func (link *Link) ValidateURL() error {
	url, err := url.ParseRequestURI(link.OriginURL)

	if err != nil {
		return err
	}

	link.OriginURL = url.String()
	return nil
}
