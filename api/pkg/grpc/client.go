package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/vorto-coffeeshop/api/pkg/service"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient DeliveryServiceClient
}

func NewGRPCService(connString string) (service.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: NewDeliveryServiceClient(conn)}, nil
}

func (s *grpcService) Delivery(request *service.DeliveryRequest) (*service.DeliveryResponse, error) {
	req := &DeliveryRequest{
		BeanId:     request.BeanId,
		CarrierId:  request.CarrierId,
		SupplierId: request.SupplierId,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.Delivery(ctx, req)
	if resp == nil {
		return nil, err
	}
	convert := unmarshalDeliveryResponse(resp)
	return &convert, nil
}

func (s *grpcService) Statistics(request *service.StatisticsRequest) (*service.StatisticsResponse, error) {
	req := &StatisticsRequest{
		Reason: request.Reason,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.Statistics(ctx, req)
	if resp == nil {
		return nil, err
	}
	convert := unmarshalStatisticsResponse(resp)
	return &convert, nil
}

func unmarshalDeliveryResponse(response *DeliveryResponse) (result service.DeliveryResponse) {
	result.Status = response.Status
	result.Message = response.Message
	return
}

func unmarshalStatisticsResponse(response *StatisticsResponse) (result service.StatisticsResponse) {
	result.Percent = response.Percent
	return
}
