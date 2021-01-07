package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/anaxdev/go-microservice/models"
	"github.com/jinzhu/gorm"
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

	// s.Conn.Debug().Table("delivery").Create(&entry)
	s.Conn.Table("delivery").Create(&entry)
	return &DeliveryResponse{Status: 0, Message: "success"}, nil
}

func (s *service) Statistics(request *StatisticsRequest) (*StatisticsResponse, error) {
	return &StatisticsResponse{}, nil
}
