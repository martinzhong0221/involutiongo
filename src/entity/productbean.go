package entity

import "time"

type Product struct {
	Id          int
	ProductName string
	ProductDesc string
	UtcUpdate   time.Time
	UtcInsert   time.Time
}
