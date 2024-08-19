package postgres

import (
	pb "Orders/genproto/orders"
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type OrderService struct {
	Db *sql.DB
}

func NewOrdersRepo(db *sql.DB) *OrderService {
	return &OrderService{Db: db}
}

func (o *OrderService) CreateOrder(req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	check,err := o.ValidateUser(context.Background(),req.UserId)
	if err != nil {
		return nil, err
	}
	if !check{
		return nil, nil
	}
	query := `
		INSERT INTO
			orders
			(user_id, kitchen_id, items_id, total_amount, delivery_address, delivery_time)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING
			id, user_id, kitchen_id, items_id, total_amount, delivery_address, delivery_time, created_at, updated_at
	`
	addOrder := `
		UPDATE
			kitchens
		SET
			total_orders = total_orders + 1
		WHERE
			id = $1`

	_, err = o.Db.Exec(addOrder, req.KitchenId)
	if err != nil {
		return nil, err
	}

	Time, _ := time.Parse("2006-01-02", req.DeliveryTime)
	var order pb.Order
	err = o.Db.QueryRow(
		query,
		req.UserId,
		req.KitchenId,
		req.ItemsId,
		req.TotalAmount,
		req.DeliveryAddress,
		Time,
	).Scan(
		&order.Id,
		&order.UserId,
		&order.KitchenId,
		&order.ItemsId,
		&order.TotalAmount,
		&order.DeliveryAddress,
		&order.DeliveryTime,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	dishOrder := `
		UPDATE
			dishes
		SET	
			order_count = order_count + 1
		WHERE
			id = $1`

	_, err = o.Db.Exec(dishOrder, req.ItemsId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{Order: &order}, nil
}

func (o *OrderService) GetUserOrders(req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersResponse, error) {
	query := `
		SELECT
			id, user_id, kitchen_id, items_id, total_amount, delivery_address, delivery_time, created_at, updated_at
		FROM
			orders
		WHERE
			user_id = $1
	`

	rows, err := o.Db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var order pb.Order
		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.KitchenId,
			&order.ItemsId,
			&order.TotalAmount,
			&order.DeliveryAddress,
			&order.DeliveryTime,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.GetUserOrdersResponse{Orders: orders}, nil
}

func (o *OrderService) UpdateOrderStatus(req *pb.UpdateOrderStatusRequest) (*pb.Status, error) {
	query := `
		UPDATE
			orders
		SET
			status = $1
		WHERE
			id = $2
	`
	_, err := o.Db.Exec(query, req.Status, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (o *OrderService) GetKitchenOrders(req *pb.GetKitchenOrdersRequest) (*pb.GetKitchenOrdersResponse, error) {
	query := `
		SELECT
			id, user_id, kitchen_id, items_id, total_amount, delivery_address, delivery_time, created_at, updated_at
		FROM
			orders
		WHERE
			kitchen_id = $1
		LIMIT
			$2
		OFFSET
			$3
	`
	rows, err := o.Db.Query(query, req.KitchenId, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.Order
	for rows.Next() {
		var order pb.Order
		if err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.KitchenId,
			&order.ItemsId,
			&order.TotalAmount,
			&order.DeliveryAddress,
			&order.DeliveryTime,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return &pb.GetKitchenOrdersResponse{Orders: orders}, nil
}

func (o *OrderService) GetOrderByID(req *pb.GetOrderByIDRequest) (*pb.GetOrderByIDResponse, error) {
	query := `
		SELECT
			id, user_id, kitchen_id, items_id, total_amount, delivery_address, delivery_time, created_at, updated_at
		FROM
			orders
		WHERE
			id = $1
	`
	var order pb.Order
	err := o.Db.QueryRow(query, req.Id).Scan(
		&order.Id,
		&order.UserId,
		&order.KitchenId,
		&order.ItemsId,
		&order.TotalAmount,
		&order.DeliveryAddress,
		&order.DeliveryTime,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderByIDResponse{Order: &order}, nil
}

func (u *OrderService) ValidateUser(ctx context.Context, id string) (bool, error) {
	query := `
	  select EXISTS (
	  select 1
	  from users
	  where deleted_at is null and id = $1
	)`

	var status bool
	err := u.Db.QueryRowContext(ctx, query, id).Scan(&status)
	if err != nil {
		return false, err
	}

	if !status {
		return false, err
	}
	return status, nil
}
