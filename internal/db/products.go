package db

import (
	"time"
)

type Product struct {
	ID          string    `gorm:"primaryKey;size:36" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int       `gorm:"not null" json:"price"`
	Category    string    `gorm:"size:100" json:"category"`
	Brand       string    `gorm:"size:100" json:"brand"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

var mockRecommendations = []Product{
	Product{ID: "1", Name: "product 1", Price: 100},
	Product{ID: "2", Name: "product 2", Price: 200, Category: "food"},
	Product{ID: "3", Name: "product 3", Price: 300},
	Product{ID: "4", Name: "product 4", Price: 400},
	Product{ID: "5", Name: "product 5", Price: 500, Description: "product 5"},
}

func (h *dbHelper) ListRecommendations() ([]Product, error) {
	time.Sleep(3 * time.Second) // Simulate the behavior of db
	return mockRecommendations, nil
}
