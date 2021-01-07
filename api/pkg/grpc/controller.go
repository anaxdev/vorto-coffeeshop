package grpc

import (
	"context"

	"github.com/anaxdev/go-microservice/pkg/service"
)

type deliveryController struct {
	deliveryService service.Service
}

func (ctlr *deliveryController) mustEmbedUnimplementedDeliveryServiceServer() {
	panic("implement me")
}

func NewDeliveryController(deliveryService service.Service) DeliveryServiceServer {
	return &deliveryController{
		deliveryService: deliveryService,
	}
}

func (ctlr *deliveryController) Delivery(ctx context.Context, request *DeliveryRequest) (*DeliveryResponse, error) {
	req := &service.DeliveryRequest{
		BeanId:     request.BeanId,
		CarrierId:  request.CarrierId,
		SupplierId: request.SupplierId,
	}
	resp, err := ctlr.deliveryService.Delivery(req)
	if resp == nil {
		return nil, err
	}
	response := &DeliveryResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}
	return response, err
}

func (ctlr *deliveryController) Statistics(ctx context.Context, request *StatisticsRequest) (*StatisticsResponse, error) {
	req := &service.StatisticsRequest{
		Reason: request.Reason,
	}
	resp, err := ctlr.deliveryService.Statistics(req)
	if resp == nil {
		return nil, err
	}
	response := &StatisticsResponse{
		Percent: resp.Percent,
	}
	return response, nil
}
