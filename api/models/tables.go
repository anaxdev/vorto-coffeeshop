package models

import (
	"time"
)

// BeanType model
type BeanType struct {
	Id        int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	name      string    `json:"name"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// Carrier model
type Carrier struct {
	Id        int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	name      string    `json:"name"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// CarrierBeanType model
type CarrierBeanType struct {
	Id         int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	BeanTypeId int64     `json:"bean_type_id"`
	CarrierId  int64     `json:"carrier_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// Delivery model
type Delivery struct {
	Id         int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	SupplierId int64     `json:"supplier_id"`
	DriverId   int64     `json:"driver_id"`
	CreatedAt  time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// Driver model
type Driver struct {
	Id        int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	name      string    `json:"name"`
	CarrierId int64     `json:"carrier_id"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// Supplier model
type Supplier struct {
	Id        int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	name      string    `json:"name"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}

// SupplierBeanType model
type SupplierBeanType struct {
	Id         int64     `gorm:"column:id; type:integer; primary_key=yes;" json:"id"`
	SupplierId int64     `json:"supplier_id"`
	BeanTypeId int64     `json:"bean_type_id"`
	CreatedAt  time.Time `gorm:"column:created_at; type:timestamp without time zone;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at; type:timestamp without time zone;" json:"updated_at"`
}
