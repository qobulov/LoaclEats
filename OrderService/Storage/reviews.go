package postgres

import (
	pb "Orders/genproto/orders"
	"fmt"
	"time"
)

func (o *OrderService) CreateReview(req *pb.CreateReviewRequest) (*pb.Review, error) {
	query := `
		INSERT INTO
			reviews
			(user_id, order_id, kitchen_id, rating, comment, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING
			id, user_id, order_id, kitchen_id, rating, comment, created_at
	`
	
	update := `
		SELECT 
			count(*)
		FROM
			reviews
		WHERE
			kitchen_id = $1
		AND
			deleted_at IS NULL`
	var count int
	err := o.Db.QueryRow(update, req.KitchenId).Scan(&count)
	if err != nil {
		return nil, err
	}

	updateRating := `
		UPDATE 
			kitchens
		SET
			rating = (rating * $2 + $1) / ($2 + 1)
		WHERE
			id = $3`

	var review pb.Review
	err = o.Db.QueryRow(
		query,
		req.UserId,
		req.OrderId,
		req.KitchenId,
		req.Rating,
		req.Comment,
		time.Now(),
	).Scan(
		&review.Id,
		&review.UserId,
		&review.OrderId,
		&review.KitchenId,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	_, err = o.Db.Exec(updateRating, req.Rating, count, req.KitchenId)
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (o *OrderService) GetReviews(req *pb.GetReviewsRequest) (*pb.GetReviewsResponse, error) {
	query := `
		SELECT
			id, user_id, order_id, rating, comment, created_at
		FROM
			reviews
		WHERE
			kitchen_id = $1
		LIMIT
			$2
		OFFSET
			$3
	`

	rows, err := o.Db.Query(query, req.KitchenId, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var reviews []*pb.Review
	for rows.Next() {
		var review pb.Review
		if err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.OrderId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		reviews = append(reviews, &review)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &pb.GetReviewsResponse{Reviews: reviews}, nil
}

func (o *OrderService) DeleteReview(req *pb.Review) (*pb.Status, error) {
	query := `
		UPDATE
			reviews
		SET
			deleted_at = $1
		WHERE
			id = $2
	`
	_, err := o.Db.Exec(query, time.Now(), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (o *OrderService) UpdateReview(req *pb.UpdateReviewRequest) (*pb.Status, error) {
	query := `
		UPDATE
			reviews
		SET
			comment = $1,
			updated_at = $2
		WHERE
			id = $3
	`
	_, err := o.Db.Exec(query, req.Comment, time.Now(), req.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}