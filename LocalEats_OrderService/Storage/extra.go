package postgres

import (
	"context"
	"fmt"
	"time"

	pb "Orders/genproto/orders"

	_ "github.com/lib/pq"
)

func (e *OrderService) GetKitchenStatistics(req *pb.GetKitchenStatisticsRequest) (*pb.GetKitchenStatisticsResponse, error) {
	kitchenStatsQuery := `
		SELECT
			COUNT(o.id) AS total_orders,
			k.rating AS average_rating,
			SUM(o.total_amount) AS revenue
		FROM
			kitchens k
		LEFT JOIN
			orders o ON k.id = o.kitchen_id
		WHERE
			k.id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			k.id
	`

	topDishesQuery := `
		SELECT
			d.id, d.name, COUNT(o.id) AS order_count, SUM(d.price) AS revenue
		FROM
			dishes d
		JOIN
			orders o ON o.items_id = d.id
		WHERE
			d.kitchen_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			d.id, d.name
		ORDER BY
			order_count DESC
	`

	busiestHoursQuery := `
		SELECT
			EXTRACT(HOUR FROM o.created_at) AS hour,
			COUNT(o.id) AS order_count
		FROM
			orders o
		WHERE
			o.kitchen_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			hour
		ORDER BY
			order_count DESC
	`

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		fmt.Println(1111111111)
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		fmt.Println(2222222222)
		return nil, err
	}

	// Kitchen statistics
	var totalOrders int32
	var averageRating float64
	var revenue float64
	err = e.Db.QueryRow(kitchenStatsQuery, req.KitchenId, startDate, endDate).Scan(
		&totalOrders,
		&averageRating,
		&revenue,
	)
	if err != nil {
		fmt.Println(3333333333)
		return nil, err
	}

	// Top dishes
	topDishesRows, err := e.Db.Query(topDishesQuery, req.KitchenId, startDate, endDate)
	if err != nil {
		fmt.Println(4444444444)
		return nil, err
	}
	defer topDishesRows.Close()

	var topDishes []*pb.TopDish
	for topDishesRows.Next() {
		var topDish pb.TopDish
		if err := topDishesRows.Scan(
			&topDish.Id,
			&topDish.Name,
			&topDish.OrdersCount,
			&topDish.Revenue,
		); err != nil {
			fmt.Println(5555555555)
			return nil, err
		}
		topDishes = append(topDishes, &topDish)
	}

	// Busiest hours
	busiestHoursRows, err := e.Db.Query(busiestHoursQuery, req.KitchenId, startDate, endDate)
	if err != nil {
		fmt.Println(6666666666)
		return nil, err
	}
	defer busiestHoursRows.Close()

	var busiestHours []*pb.BusiestHour
	for busiestHoursRows.Next() {
		var busiestHour pb.BusiestHour
		if err := busiestHoursRows.Scan(
			&busiestHour.Hour,
			&busiestHour.OrdersCount,
		); err != nil {
			fmt.Println(7777777777)
			return nil, err
		}
		busiestHours = append(busiestHours, &busiestHour)
	}

	return &pb.GetKitchenStatisticsResponse{
		TotalOrders:   totalOrders,
		AverageRating: averageRating,
		TopDishes:     topDishes,
		BusiestHours:  busiestHours,
		TotalRevenue:  revenue,
	}, nil
}

func (e *OrderService) GetUserActivity(req *pb.GetUserActivityRequest) (*pb.GetUserActivityResponse, error) {
	query := `
		SELECT
			COUNT(*) AS total_orders,
			COALESCE(SUM(total_amount), 0) AS total_spent
		FROM
			orders
		WHERE
			user_id = $1 AND
			created_at BETWEEN $2 AND $3
	`

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, err
	}

	var totalOrders int32
	var totalSpent float64

	err = e.Db.QueryRow(query, req.UserId, startDate, endDate).Scan(
		&totalOrders,
		&totalSpent,
	)
	if err != nil {
		return nil, err
	}

	favoriteCuisinesQuery := `
		SELECT
			cuisine_type, COUNT(*) AS order_count
		FROM
			kitchens k
		JOIN
			orders o ON o.kitchen_id = k.id
		WHERE  
			o.user_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			cuisine_type
		ORDER BY
			order_count DESC
	`

	favoriteCuisinesRows, err := e.Db.Query(favoriteCuisinesQuery, req.UserId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer favoriteCuisinesRows.Close()

	var favoriteCuisines []*pb.FavoriteCuisine
	for favoriteCuisinesRows.Next() {
		var favoriteCuisine pb.FavoriteCuisine
		if err := favoriteCuisinesRows.Scan(
			&favoriteCuisine.CuisineType,
			&favoriteCuisine.OrdersCount,
		); err != nil {
			return nil, err
		}
		favoriteCuisines = append(favoriteCuisines, &favoriteCuisine)
	}

	favoriteKitchensQuery := `
		SELECT
			k.id, k.name, COUNT(*) AS order_count
		FROM
			kitchens k
		JOIN
			orders o ON o.kitchen_id = k.id
		WHERE
			o.user_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			k.id, k.name
		ORDER BY
			order_count DESC
	`

	favoriteKitchensRows, err := e.Db.Query(favoriteKitchensQuery, req.UserId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer favoriteKitchensRows.Close()

	var favoriteKitchens []*pb.FavoriteKitchen
	for favoriteKitchensRows.Next() {
		var favoriteKitchen pb.FavoriteKitchen
		if err := favoriteKitchensRows.Scan(
			&favoriteKitchen.Id,
			&favoriteKitchen.Name,
			&favoriteKitchen.OrdersCount,
		); err != nil {
			return nil, err
		}
		favoriteKitchens = append(favoriteKitchens, &favoriteKitchen)
	}

	return &pb.GetUserActivityResponse{
		TotalOrders:      totalOrders,
		TotalSpent:       totalSpent,
		FavoriteCuisines: favoriteCuisines,
		FavoriteKitchens: favoriteKitchens,
	}, nil
}

func (e *OrderService) CreateWorkingHours(ctx context.Context, req *pb.CreateWorkingHoursRequest) (*pb.CreateWorkingHoursResponse, error) {
	query := `
		INSERT INTO working_hours (kitchen_id, day_of_week, open_time, close_time)
		VALUES ($1, $2, $3, $4)
	`

	for dayOfWeek, hours := range req.WorkingHours {
		openTime, err := time.Parse("15:04", hours.Open)
		if err != nil {
			return nil, err
		}
		closeTime, err := time.Parse("15:04", hours.Close)
		if err != nil {
			return nil, err
		}
		_, err = e.Db.Exec(query, req.KitchenId, dayOfWeek, openTime, closeTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateWorkingHoursResponse{
		KitchenId:    req.KitchenId,
		WorkingHours: req.WorkingHours,
		CreatedAt:    time.Now().Format(time.RFC3339),
	}, nil
}

func (e *OrderService) UpdateWorkingHours(ctx context.Context, req *pb.UpdateWorkingHoursRequest) (*pb.UpdateWorkingHoursResponse, error) {
	tx, err := e.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	deleteQuery := "DELETE FROM working_hours WHERE kitchen_id = $1"
	_, err = tx.Exec(deleteQuery, req.KitchenId)
	if err != nil {
		return nil, err
	}

	insertQuery := `
		INSERT INTO working_hours (kitchen_id, day_of_week, open_time, close_time)
		VALUES ($1, $2, $3, $4)
	`
	for dayOfWeek, hours := range req.WorkingHours {
		openTime, err := time.Parse("15:04", hours.Open)
		if err != nil {
			return nil, err
		}
		closeTime, err := time.Parse("15:04", hours.Close)
		if err != nil {
			return nil, err
		}
		_, err = tx.Exec(insertQuery, req.KitchenId, dayOfWeek, openTime, closeTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.UpdateWorkingHoursResponse{
		KitchenId:    req.KitchenId,
		WorkingHours: req.WorkingHours,
		UpdatedAt:    time.Now().Format(time.RFC3339),
	}, nil
}
