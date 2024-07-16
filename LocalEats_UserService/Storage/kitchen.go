package postgres

import (
	pb "AuthService/genproto/users"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func (u *UserRepo) CreateKitchen(kitchen *pb.CreateKitchenRequest) (*pb.Status, error) {

	query := `
		INSERT INTO 
			kitchens_profile
			(name, address, phone_number,cuisine_type,description)
		VALUES
			($1, $2, $3, $4, $5)
		`
	_, err := u.Db.Exec(query, kitchen.Name, kitchen.Address, kitchen.PhoneNumber, kitchen.CuisineType, kitchen.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (u *UserRepo) GetKitchenByID(kitchenID *pb.GetKitchenByIDRequest) (*pb.GetKitchenByIDResponse, error) {

	query := `
		SELECT 
			id, name, address, phone_number, cuisine_type, description,rating,total_orders
		FROM 
			kitchens_profile
		WHERE 
			id = $1
		AND
			deleted_at IS NULL
	`
	var kitchen pb.Kitchen
	err := u.Db.QueryRow(query, kitchenID.Id).Scan(
		&kitchen.Id,
		&kitchen.Name,
		&kitchen.Address,
		&kitchen.PhoneNumber,
		&kitchen.CuisineType,
		&kitchen.Description,
		&kitchen.Rating,
		&kitchen.TotalOrders,
	)
	if err != nil {
		return nil, err
	}
	if kitchen.Description == "" {
		kitchen.Description = "No description"
	}
	if kitchen.CuisineType == "" {
		kitchen.CuisineType = "No cuisine type"
	}

	return &pb.GetKitchenByIDResponse{Kitchen: &kitchen}, nil
}

func (u *UserRepo) UpdateKitchen(kitchen *pb.UpdateKitchenRequest) (*pb.Status, error) {

	query := `
		UPDATE
			kitchens_profile
		SET
			name = $1,
			address = $2,
			phone_number = $3,
			cuisine_type = $4,
			description = $5
		WHERE
			id = $6
		AND
			deleted_at IS NULL
	`
	_, err := u.Db.Exec(query, kitchen.Name, kitchen.Address, kitchen.PhoneNumber, kitchen.CuisineType, kitchen.Description, kitchen.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (u *UserRepo) DeleteKitchen(kitchenID *pb.GetKitchenByIDRequest) (*pb.Status, error) {

	query := `
		UPDATE
			kitchens_profile
		SET
			deleted_at = $1
		WHERE
			id = $2
	`
	_, err := u.Db.Exec(query, time.Now(), kitchenID.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (u *UserRepo) GetAllKitchens(req *pb.GetKitchensRequest) (*pb.GetKitchensResponse, error) {
	query := `
		SELECT 
			id,  name, address, phone_number, cuisine_type, description,rating,total_orders
		FROM 
			kitchens_profile
		WHERE
			deleted_at IS NULL
		LIMIT
			$1
		OFFSET
			$2
	`
	rows, err := u.Db.Query(query, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var kitchens []*pb.Kitchen
	for rows.Next() {
		var kitchen pb.Kitchen
		err := rows.Scan(
			&kitchen.Id,
			&kitchen.Name,
			&kitchen.Address,
			&kitchen.PhoneNumber,
			&kitchen.CuisineType,
			&kitchen.Description,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			return nil, err
		}
		kitchens = append(kitchens, &kitchen)
	}

	return &pb.GetKitchensResponse{Kitchens: kitchens}, nil
}
func (u *UserRepo) SearchKitchens(filter *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	params := make(map[string]interface{})
	var args []interface{}
	filterStr := "WHERE deleted_at IS NULL" // Initial condition to exclude deleted kitchens

	if filter.Name != "" {
		params["name"] = "%" + filter.Name + "%"
		filterStr += " AND name ILIKE :name"
	}

	if filter.CuisineType != "" {
		params["cuisine_type"] = filter.CuisineType
		filterStr += " AND cuisine_type = :cuisine_type"
	}

	if filter.Rating > 0 {
		params["rating"] = filter.Rating
		filterStr += " AND rating >= :rating"
	}

	limit := ""
	if filter.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(int(filter.Limit))
	}

	offset := ""
	if filter.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(int(filter.Offset))
	}

	query := "SELECT id, owner_id, name, address, phone_number, cuisine_type, description, rating, total_orders FROM kitchens_profile " + filterStr + limit + offset
	query, args = ReplaceQueryParams(query, params)

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kitchens []*pb.Kitchen
	for rows.Next() {
		var kitchen pb.Kitchen
		var ownerId sql.NullString

		err := rows.Scan(
			&kitchen.Id,
			&ownerId,
			&kitchen.Name,
			&kitchen.Address,
			&kitchen.PhoneNumber,
			&kitchen.CuisineType,
			&kitchen.Description,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			return nil, err
		}

		// Handle nullable fields
		if ownerId.Valid {
			kitchen.OwnerId = ownerId.String
		}

		kitchens = append(kitchens, &kitchen)
	}

	return &pb.SearchKitchensResponse{Kitchens: kitchens}, nil
}

// ReplaceQueryParams replaces named parameters in the SQL query with positional arguments.
func ReplaceQueryParams(query string, params map[string]interface{}) (string, []interface{}) {
	args := make([]interface{}, 0, len(params))
	for key, val := range params {
		query = strings.ReplaceAll(query, ":"+key, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, val)
	}
	return query, args
}

