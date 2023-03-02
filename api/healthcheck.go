package api

import (
	"encoding/json"
	"net/http"
)

type Healthcheck struct {
	Overall  ServiceHealthCheckStatus `json:"overall"`
	Services []ServiceHealthcheck     `json:"services,omitempty"`
}

type ServiceHealthcheck struct {
	Name    string                   `json:"name"`
	Status  ServiceHealthCheckStatus `json:"status"`
	handler func() ServiceHealthCheckStatus
}

type ServiceHealthCheckStatus string

const (
	ServiceHealthy   ServiceHealthCheckStatus = "HEALTHY"
	ServicePending   ServiceHealthCheckStatus = "PENDING"
	ServiceUnhealthy ServiceHealthCheckStatus = "UNHEALTHY"
)

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{
		Overall: ServiceHealthy,
	}
}

func (h *Healthcheck) handlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.check()
		healthBytes, marshalErr := json.Marshal(h)
		if marshalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("NOK"))
			return
		}
		if h.Overall == ServiceUnhealthy {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Write(healthBytes)
	}
}

func (h *Healthcheck) RegisterService(serviceName string, handler func() ServiceHealthCheckStatus) {
	h.Services = append(h.Services, ServiceHealthcheck{
		Name:    serviceName,
		Status:  ServicePending,
		handler: handler,
	})
}

func (h *Healthcheck) check() {
	for index, service := range h.Services {
		status := service.handler()
		h.Services[index].Status = status
		if status == ServiceUnhealthy {
			h.Overall = ServiceUnhealthy
		}
	}
}
