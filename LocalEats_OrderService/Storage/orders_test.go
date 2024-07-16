package postgres

import (
	"Orders/genproto/orders"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Initialize a new OrdersRepo instance
	orderService := NewOrdersRepo(db)

	// Prepare a CreateOrderRequest
	request := &orders.CreateOrderRequest{
		UserId:          "7a110220-8a11-4dfd-a2b7-6d1af6585c53",
		KitchenId:       "0ec9754d-10a3-47c6-83e1-17683faf7876",
		Items:           "item1",
		TotalAmount:     10.0,
		DeliveryTime:    "2020-12-12",
		DeliveryAddress: "City",
	}

	// Create the order using CreateOrder method
	response, err := orderService.CreateOrder(request)
	if err != nil {
		t.Fatalf("error creating order: %v", err)
	}

	// Assert that the response is not nil
	if response == nil {
		t.Fatal("expected non-nil response, got nil")
	}

	// Assert that the created order ID is not empty
	if response.Order.Id == "" {
		t.Fatal("expected non-empty order ID, got empty")
	}
}

func TestGetUserOrders(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Initialize a new OrdersRepo instance
	orderService := NewOrdersRepo(db)

	// Prepare a GetUserOrdersRequest
	request := &orders.GetUserOrdersRequest{
		UserId: "7a110220-8a11-4dfd-a2b7-6d1af6585c53",
	}

	// Get the orders using GetUserOrders method
	response, err := orderService.GetUserOrders(request)
	if err != nil {
		t.Fatalf("error getting orders: %v", err)
	}

	// Assert that the response is not nil
	if response == nil {
		t.Fatal("expected non-nil response, got nil")
	}
}


func TestUpdateOrderStatus(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Initialize a new OrdersRepo instance
	orderService := NewOrdersRepo(db)

	// Prepare a UpdateOrderStatusRequest
	request := &orders.UpdateOrderStatusRequest{
		Id:     "7a110220-8a11-4dfd-a2b7-6d1af6585c53",
		Status: "pending",
	}

	// Update the order status using UpdateOrderStatus method
	_, err = orderService.UpdateOrderStatus(request)
	if err != nil {
		t.Fatalf("error updating order status: %v", err)
	}
}

func TestGetKitchenOrders(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Initialize a new OrdersRepo instance
	orderService := NewOrdersRepo(db)

	// Prepare a GetKitchenOrdersRequest
	request := &orders.GetKitchenOrdersRequest{
		KitchenId: "0ec9754d-10a3-47c6-83e1-17683faf7876",
	}	

	// Get the orders using GetKitchenOrders method
	response, err := orderService.GetKitchenOrders(request)
	if err != nil {
		t.Fatalf("error getting orders: %v", err)
	}

	// Assert that the response is not nil
	if response == nil {
		t.Fatal("expected non-nil response, got nil")
	}
}

func TestGetOrderByID(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()	

	// Initialize a new OrdersRepo instance
	orderService := NewOrdersRepo(db)

	// Prepare a GetOrderByIDRequest
	request := &orders.GetOrderByIDRequest{
		Id: "d4a5ad7b-64f0-4072-8905-9e917291409a",
	}

	// Get the order using GetOrderByID method
	response, err := orderService.GetOrderByID(request)
	if err != nil {
		t.Fatalf("error getting order: %v", err)
	}

	// Assert that the response is not nil
	if response == nil {
		t.Fatal("expected non-nil response, got nil")
	}
}