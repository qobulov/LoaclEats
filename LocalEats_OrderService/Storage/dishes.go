package postgres

import (
	pb "Orders/genproto/orders"
	"time"

	"github.com/lib/pq"
)


func (o *OrderService) GetDishes(req *pb.GetDishesRequest) (*pb.GetDishesResponse, error) {
	query := `
		SELECT
			id, kitchen_id, name, description, price, category, ingredients, available, created_at, updated_at
		FROM
			dishes
		WHERE
			kitchen_id = $1
		ORDER BY
			name
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

	var dishes []*pb.Dish
	for rows.Next() {
		dish := &pb.Dish{}
		var ingredients pq.StringArray // Use pq.StringArray for PostgreSQL arrays
		err = rows.Scan(
			&dish.Id,
			&dish.KitchenId,
			&dish.Name,
			&dish.Description,
			&dish.Price,
			&dish.Category,
			&ingredients, // Scan into pq.StringArray
			&dish.Available,
			&dish.CreatedAt,
			&dish.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		dish.Ingredients = ingredients // Assign pq.StringArray to pb.Dish
		dishes = append(dishes, dish)
	}
	return &pb.GetDishesResponse{Dishes: dishes}, nil
}

func (o *OrderService) GetDish(req *pb.GetDishRequest) (*pb.Dish, error) {
	query := `
		SELECT
			id, kitchen_id, name, description, price, category, ingredients, available, created_at, updated_at
		FROM
			dishes
		WHERE
			id = $1
	`

	var dish pb.Dish
	var ingredients pq.StringArray // Use pq.StringArray for PostgreSQL arrays
	err := o.Db.QueryRow(query, req.DishId).Scan(
		&dish.Id,
		&dish.KitchenId,
		&dish.Name,
		&dish.Description,
		&dish.Price,
		&dish.Category,
		&ingredients, // Scan into pq.StringArray
		&dish.Available,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	dish.Ingredients = ingredients // Assign pq.StringArray to pb.Dish
	return &dish, nil
}

func (o *OrderService) CreateDish(req *pb.CreateDishRequest) (*pb.Dish, error) {
	query := `
		INSERT INTO
			dishes
			(kitchen_id, name, description, price, category, ingredients, available)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id, kitchen_id, name, description, price, category, ingredients, available, created_at, updated_at
	`

	var dish pb.Dish
	var ingredients pq.StringArray = req.Ingredients // Use pq.StringArray for PostgreSQL arrays
	err := o.Db.QueryRow(
		query,
		req.KitchenId,
		req.Name,
		req.Description,
		req.Price,
		req.Category,
		ingredients, // Pass pq.StringArray here
		req.Available,
	).Scan(
		&dish.Id,
		&dish.KitchenId,
		&dish.Name,
		&dish.Description,
		&dish.Price,
		&dish.Category,
		(*pq.StringArray)(&dish.Ingredients), // Scan into pq.StringArray
		&dish.Available,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &dish, nil
}


func (o *OrderService) UpdateDish(req *pb.UpdateDishRequest) (*pb.Dish, error) {
	query := `
		UPDATE
			dishes
		SET
			name = $1,
			price = $2,
			available = $3,
			description = $4,
			ingredients = $5,
			category = $6,
			updated_at = $7
		WHERE
			id = $8
		RETURNING
			id, kitchen_id, name, description, price, category, ingredients, available, created_at, updated_at
	`
	var dish pb.Dish
	var ingredients pq.StringArray = req.Ingredients   // Assign ingredients from req to pq.StringArray

	err := o.Db.QueryRow(
		query,
		req.Name,
		req.Price,
		req.Available,
		req.Description,
		ingredients, // Use pq.StringArray here
		req.Category,
		time.Now(),
		req.Id,
	).Scan(
		&dish.Id,
		&dish.KitchenId,
		&dish.Name,
		&dish.Description,
		&dish.Price,
		&dish.Category,
		&ingredients, // Scan into pq.StringArray
		&dish.Available,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	dish.Ingredients = ingredients // Assign pq.StringArray to pb.Dish
	return &dish, nil
}
func (o *OrderService) DeleteDish(req *pb.DeleteDishRequest) (*pb.Status, error) {
	query := `
		DELETE FROM
			dishes
		WHERE
			id = $1
	`
	_, err := o.Db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}
