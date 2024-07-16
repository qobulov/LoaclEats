package postgres

import (
	"testing"

	pb "Orders/genproto/orders"

	_ "github.com/lib/pq"
)

func TestCreateReview(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.CreateReview(&pb.CreateReviewRequest{
		OrderId: "d4a5ad7b-64f0-4072-8905-9e917291409a",
		UserId:  "fdcfd039-781d-4a6d-b025-810986269f6b",
		KitchenId: "0ec9754d-10a3-47c6-83e1-17683faf7876",
		Rating:  5,
		Comment: "Great food!",

	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetReviews(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.GetReviews(&pb.GetReviewsRequest{
		KitchenId: "0ec9754d-10a3-47c6-83e1-17683faf7876",
	})
	if err != nil {
		t.Fatal(err)
	}
}	

func TestUpdateReview(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.UpdateReview(&pb.UpdateReviewRequest{
		Id: "d4a5ad7b-64f0-4072-8905-9e917291409a",
		OrderId: "d4a5ad7b-64f0-4072-8905-9e917291409a",
		UserId:  "fdcfd039-781d-4a6d-b025-810986269f6b",
		KitchenId: "0ec9754d-10a3-47c6-83e1-17683faf7876",
		Rating:  5,
		Comment: "Great food!",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteReview(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.DeleteReview(&pb.Review{
		Id: "d4a5ad7b-64f0-4072-8905-9e917291409a",
	})	
	if err != nil {
		t.Fatal(err)
	}
}