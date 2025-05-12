package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"sync"

	test "gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/pkg/api/api/test"
)

const (
	length = 10
)

type Service struct {
	test.OrderServiceServer
	data map[string]*test.Order
	mu   sync.Mutex
}

func New() *Service {
	return &Service{data: map[string]*test.Order{}}
}

func (s *Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	buf := make([]byte, length)
	rand.Read(buf)
	id := string(buf)
	s.data[id] = &test.Order{Id: id, Item: req.Item, Quantity: req.Quantity}
	return &test.CreateOrderResponse{Id: id}, nil
}

func (s *Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	id := req.Id
	s.mu.Lock()
	defer s.mu.Unlock()
	order, flag := s.data[id]
	if !flag {
		return nil, fmt.Errorf("there is no element with such id")
	}
	return &test.GetOrderResponse{Order: order}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	id := req.Id
	s.mu.Lock()
	defer s.mu.Unlock()
	_, flag := s.data[id]
	if !flag {
		return nil, fmt.Errorf("there is no element with such id")
	}
	s.data[id] = &test.Order{Id: id, Item: req.Item, Quantity: req.Quantity}
	return &test.UpdateOrderResponse{Order: s.data[id]}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	id := req.Id
	s.mu.Lock()
	defer s.mu.Unlock()
	_, flag := s.data[id]
	if !flag {
		return &test.DeleteOrderResponse{Success: false}, fmt.Errorf("there is no element with such id")
	}
	delete(s.data, id)
	return &test.DeleteOrderResponse{Success: true}, nil
}

func (s *Service) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]*test.Order, 0)
	for _, val := range s.data {
		list = append(list, val)
	}
	return &test.ListOrdersResponse{Orders: list}, nil
}
