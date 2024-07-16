package postgres

import (
	"testing"

	pb "Orders/genproto/orders"

	_ "github.com/lib/pq"
)

func TestCreatePayment(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)

	_, err = orderService.CreatePayment(&pb.CreatePaymentRequest{
		OrderId: "d4a5ad7b-64f0-4072-8905-9e917291409a",
		UserId:  "fdcfd039-781d-4a6d-b025-810986269f6b",
		Amount:  10.0,
		PaymentMethod: "Visa",
		CardNumber:    "1234567890123456",
		ExpiryDate:    "2020-12-12",
		Cvv:           "123",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPayments(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.GetPayments(&pb.GetPaymentsRequest{
		OrderId: "d4a5ad7b-64f0-4072-8905-9e917291409a",
	})
	if err != nil {
		t.Fatal(err)
	}
}