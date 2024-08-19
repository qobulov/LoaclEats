package postgres

import (
	"testing"

	_ "github.com/lib/pq"

	pb "Orders/genproto/orders"
)

func TestGetKitchenStatistics(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	service := NewOrdersRepo(db)

	req := &pb.GetKitchenStatisticsRequest{
		KitchenId: "45e86e04-4f57-4f46-a330-df4e02ed281f",
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	_, err = service.GetKitchenStatistics(req)
	if err != nil {
		t.Fatal(err)
	}

}
