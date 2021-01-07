package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	rpcService Service
)

func delivery(w http.ResponseWriter, r *http.Request) {
	beanIdStr := r.FormValue("bean_type_id")
	carrierIdStr := r.FormValue("carrierIdStr")
	supplierIdStr := r.FormValue("supplierIdStr")
	beanId, _ := strconv.Atoi(beanIdStr)
	carrierId, _ := strconv.Atoi(carrierIdStr)
	supplierId, _ := strconv.Atoi(supplierIdStr)

	_, err := rpcService.Delivery(&DeliveryRequest{
		BeanId:     int64(beanId),
		SupplierId: int64(supplierId),
		CarrierId:  int64(carrierId)})
	ret := DeliveryResponse{
		Status:  0,
		Message: "success",
	}
	if err != nil {
		ret.Status = -1
		ret.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(ret)
}

func statistics(w http.ResponseWriter, r *http.Request) {
	response, _ := rpcService.Statistics(&StatisticsRequest{})
	ret := StatisticsResponse{
		Percent: 0,
	}
	if response != nil {
		ret.Percent = response.Percent
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode("CoffeeShop API")
}

func ConfigureHandlers(r *mux.Router, rpc Service) error {
	r.HandleFunc("/delivery", delivery).Methods("POST")
	r.HandleFunc("/statistics", statistics).Methods("GET")
	r.HandleFunc("/", hello).Methods("GET")
	rpcService = rpc
	return nil
}
