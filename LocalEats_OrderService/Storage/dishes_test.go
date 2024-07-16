package postgres

import (
	"testing"

	pb "Orders/genproto/orders"

	_ "github.com/lib/pq"
)

func TestCreateDish(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	ingd := []string{"Cheese", "Tomato"}
	orderService := NewOrdersRepo(db)
	_, err = orderService.CreateDish(&pb.CreateDishRequest{
		KitchenId:   "0ec9754d-10a3-47c6-83e1-17683faf7876",
		Name:        "Shawarma",
		Description: "Shawarma with cheese",
		Price:       10.0,
		Category:    "Fast Food",
		Ingredients: ingd,
		Available:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDishes(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	dish, err := orderService.GetDishes(&pb.GetDishesRequest{
		KitchenId: "0ec9754d-10a3-47c6-83e1-17683faf7876",
		Limit:     10,
		Offset:    0,
	})
	if len(dish.Dishes) == 0 {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateDish(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.UpdateDish(&pb.UpdateDishRequest{
		Id:          "26d96315-95d1-4171-b751-f69dcd62bd8a",
		Name:        "Shawarma",
		Description: "Shawarma with cheese",
		Price:       10.0,
		Category:    "Fast Food",
		Ingredients: []string{"Cheese", "Tomato"},
		Available:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteDish(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	orderService := NewOrdersRepo(db)
	_, err = orderService.DeleteDish(&pb.DeleteDishRequest{
		Id: "26d96315-95d1-4171-b751-f69dcd62bd8a",
	})
	if err != nil {
		t.Fatal(err)
	}
}