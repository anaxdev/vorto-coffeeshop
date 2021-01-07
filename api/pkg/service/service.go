package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vorto-coffeeshop/api/models"
)

var ErrNotFound = errors.New("not found")

type DeliveryRequest struct {
	BeanId     int64
	CarrierId  int64
	SupplierId int64
}

type DeliveryResponse struct {
	Status  int64
	Message string
}

type StatisticsRequest struct {
	Reason string
}

type StatisticsResponse struct {
	Percent int64
}

// Service defines the interface exposed by this package.
type Service interface {
	Delivery(request *DeliveryRequest) (*DeliveryResponse, error)
	Statistics(request *StatisticsRequest) (*StatisticsResponse, error)
}

type service struct {
	Conn *gorm.DB
}

func NewService(conn *gorm.DB) Service {
	return &service{Conn: conn}
}

func (s *service) Delivery(request *DeliveryRequest) (*DeliveryResponse, error) {
	type Exist struct {
		CarrierId  int64 `json:"cb"`
		SupplierId int64 `json:"sb"`
	}
	var exists []Exist
	sql := `select c.id as cb, s.id as sb from carrier_bean_type c, supplier_bean_type s where c.bean_type_id=%d and s.bean_type_id=%d and c.carrier_id=%d and s.supplier_id=%d and (select id from driver where carrier_id=%d) is not NULL`
	query := fmt.Sprintf(sql, request.BeanId, request.BeanId, request.CarrierId, request.SupplierId, request.CarrierId)
	s.Conn.Raw(query).Scan(&exists)
	if len(exists) == 0 {
		return &DeliveryResponse{Status: -1, Message: "Invalid delivery"}, errors.New("Invalid delivery")
	}
	var driver models.Driver
	s.Conn.Table("driver").Where("carrier_id = ?", request.CarrierId).First(&driver)
	entry := models.Delivery{
		SupplierId: request.SupplierId,
		DriverId:   driver.Id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	s.Conn.Table("delivery").Create(&entry)
	return &DeliveryResponse{Status: 0, Message: "success"}, nil
}

func (s *service) Statistics(request *StatisticsRequest) (*StatisticsResponse, error) {
	type ValidCount struct {
		ValidDeliveriesCount  int64 `gorm:"column:valid_deliveries_count; type:integer;" json:"valid_deliveries_count"`
	}

	type DeliveryCount struct {
		Deliveries  int64 `gorm:"column:all_deliveries_count; type:integer;" json:"all_deliveries_count"`
	}

	var beans []models.BeanType
	s.Conn.Table("bean_type").Scan(&beans)

	var delivery DeliveryCount
	sql := `select (select count(*) as c from supplier) * (select count(*) from driver) * (select count(*) from bean_type) as all_deliveries_count`
	s.Conn.Raw(sql).Scan(&delivery)
	if delivery.Deliveries == 0 {
		return &StatisticsResponse{}, nil
	}

	var validCount int64
	for _, bean := range beans {
		sql := `select (select count(*) from supplier_bean_type where bean_type_id=%d) * (select count(d.id) from carrier_bean_type cb, driver d where cb.bean_type_id=%d and cb.carrier_id = d.carrier_id) as valid_deliveries_count`
		query := fmt.Sprintf(sql, bean.Id, bean.Id)
		var count ValidCount
		s.Conn.Raw(query).Scan(&count)
		validCount += count.ValidDeliveriesCount
	}

	return &StatisticsResponse{Percent: (validCount * 100) / delivery.Deliveries}, nil
}
