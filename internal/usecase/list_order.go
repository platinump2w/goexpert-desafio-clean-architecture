package usecase

import (
	"github.com/platinump2w/goexpert-desafio-clean-architecture/internal/entity"
	"github.com/platinump2w/goexpert-desafio-clean-architecture/pkg/events"
)

type OrderListOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (listOrdersUseCase *ListOrdersUseCase) Execute() (OrderListOutputDTO, error) {
	orders, err := listOrdersUseCase.OrderRepository.List()
	if err != nil {
		return OrderListOutputDTO{}, err
	}

	var orderOutputDtos = make([]OrderOutputDTO, 0)

	for _, order := range orders {
		orderOutputDto := buildOrderOutputDto(order)
		orderOutputDtos = append(orderOutputDtos, orderOutputDto)
	}

	orderListOutputDto := OrderListOutputDTO{
		Orders: orderOutputDtos,
	}

	listOrdersUseCase.OrdersListed.SetPayload(orderListOutputDto)
	listOrdersUseCase.EventDispatcher.Dispatch(listOrdersUseCase.OrdersListed)

	return orderListOutputDto, nil
}

func buildOrderOutputDto(order entity.Order) OrderOutputDTO {
	return OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
}
