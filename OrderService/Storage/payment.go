package postgres

import (
	pb "Orders/genproto/orders"
	"context"
	"time"
)

func (o *OrderService) CreatePayment(req *pb.CreatePaymentRequest) (*pb.Payment, error) {
	check, err := o.ValidateUser(context.Background(), req.UserId)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, nil
	}
	
	query := `
		INSERT INTO
			payments
			(order_id, amount, payment_method, card_number, expiry_date, cvv, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id, order_id, amount, payment_method, status, created_at
	`
	var payment pb.Payment
	err = o.Db.QueryRow(
		query,
		req.OrderId,
		req.Amount,
		req.PaymentMethod,
		req.CardNumber,
		req.ExpiryDate,
		req.Cvv,
		time.Now(),
	).Scan(
		&payment.Id,
		&payment.OrderId,
		&payment.Amount,
		&payment.PaymentMethod,
		&payment.PaymentStatus,
		&payment.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (o *OrderService) GetPayments(req *pb.GetPaymentsRequest) (*pb.GetPaymentsResponse, error) {

	query := `
		SELECT
			id, order_id, amount, payment_method, status, created_at
		FROM
			payments
		WHERE
			order_id = $1
	`
	rows, err := o.Db.Query(query, req.OrderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*pb.Payment
	for rows.Next() {
		var payment pb.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.OrderId,
			&payment.Amount,
			&payment.PaymentMethod,
			&payment.PaymentStatus,
			&payment.CreatedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return &pb.GetPaymentsResponse{Payments: payments}, nil
}
