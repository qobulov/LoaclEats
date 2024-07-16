package postgres

import (
	"context"
	"fmt"
	"time"

	pb "Orders/genproto/orders"

	_ "github.com/lib/pq"
)

func (e *OrderService) GetKitchenStatistics(req *pb.GetKitchenStatisticsRequest) (*pb.GetKitchenStatisticsResponse, error) {
	query := `
		SELECT
			COUNT(*) AS total_orders,
			COALESCE(SUM(total_amount), 0) AS total_revenue,
			COALESCE(AVG(rating), 0) AS average_rating
		FROM
			orders o
		LEFT JOIN
			reviews r ON o.id = r.order_id
		WHERE
			o.kitchen_id = $1 AND
			o.created_at BETWEEN $2 AND $3
	`
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		fmt.Print(1111111111)
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		fmt.Print(2222222222)
		return nil, err
	}

	var totalOrders int32
	var totalRevenue float64
	var averageRating float64

	err = e.Db.QueryRow(query, req.KitchenId, startDate, endDate).Scan(
		&totalOrders,
		&totalRevenue,
		&averageRating,
	)
	if err != nil {
		fmt.Print(3333333333)
		return nil, err
	}

	topDishesQuery := `
		SELECT
			d.id, d.name, COUNT(*) AS orders_count, SUM(o.total_amount) AS revenue
		FROM
			dishes d
		JOIN
			orders o ON o.kitchen_id = d.kitchen_id
		WHERE
			d.kitchen_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			d.id, d.name
		ORDER BY
			orders_count DESC
	`

	topDishesRows, err := e.Db.Query(topDishesQuery, req.KitchenId, startDate, endDate)
	if err != nil {
		fmt.Print(4444444444)
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
			fmt.Print(5555555555)
			return nil, err
		}
		topDishes = append(topDishes, &topDish)
	}

	busiestHoursQuery := `
		SELECT
			EXTRACT(HOUR FROM o.created_at) AS hour, COUNT(*) AS orders_count
		FROM
			orders o
		WHERE
			o.kitchen_id = $1 AND
			o.created_at BETWEEN $2 AND $3
		GROUP BY
			hour
		ORDER BY
			orders_count DESC
	`

	busiestHoursRows, err := e.Db.Query(busiestHoursQuery, req.KitchenId, startDate, endDate)
	if err != nil {
		fmt.Print(6666666666)
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
			fmt.Print(7777777777)
			return nil, err
		}
		busiestHours = append(busiestHours, &busiestHour)
	}
	// fmt.Print(topDishes)

	return &pb.GetKitchenStatisticsResponse{
		TotalOrders:   totalOrders,
		TotalRevenue:  totalRevenue,
		AverageRating: averageRating,
		TopDishes:     topDishes,
		BusiestHours:  busiestHours,
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
			cuisine_type, COUNT(*) AS orders_count
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
			orders_count DESC
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
			k.id, k.name, COUNT(*) AS orders_count
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
			orders_count DESC
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
