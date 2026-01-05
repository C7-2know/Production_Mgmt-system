package entities

import (
	"time"

)

type WeeklyReport struct {
	WeekStart    time.Time
	WeekEnd      time.Time
	TotalProduction int
	TotalWastage    int
	ItemsProduced  map[string]int
}

type MonthlyReport struct {
	Month         time.Time
	TotalProduction int
	TotalWastage    int
	ItemsProduced  map[string]int
}

type WastageReport struct {
	StartDate    time.Time
	EndDate      time.Time
	TotalWastage int
	WastageByType map[string]int
}

type StockReport struct {
	ItemsInStock map[string]int
	LastUpdated  time.Time
}

type CostReport struct {
	StartDate    time.Time
	EndDate      time.Time
	TotalCost    float64
	CostByCategory map[string]float64
}