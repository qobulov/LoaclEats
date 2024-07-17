package postgres

import (
	pb "AuthService/genproto/users"
	"AuthService/models"
	"AuthService/redis"
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"math/rand"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (r *UserRepo) Register(req *pb.RegisterRequest) (*pb.Status, error) {
	query := `
		INSERT INTO 
			users
			(username, email, password_hash, full_name, user_type, phone_number)
		VALUES
			($1, $2, $3, $4, $5, $6)
		`
	_, err := r.Db.Exec(query, req.Username, req.Email, req.Password, req.FullName, req.UserType, req.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *UserRepo) GetProfile(UserId *pb.UserId) (*pb.GetProfileResponse, error) {
	query := `
		SELECT 
			id, username, email, full_name, user_type, address, bio, specialties,
			years_of_experience, is_verified, phone_number, created_at, updated_at
		FROM 
			users
		WHERE 
			id = $1
		AND
			deleted_at IS NULL
	`
	
	var user pb.User
	err := r.Db.QueryRow(query, UserId.Id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.UserType,
		&user.Address,
		&user.Bio,
		&user.Specialties,
		&user.YearsOfExperience,
		&user.IsVerified,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	// nullbio := sql.NullString{}
	if user.Bio == "" {
		user.Bio = "No bio"
	}
	if user.Specialties == "" {
		user.Specialties = "No Specialties"
	}
	if user.Address == "" {
		user.Address = "No address"
	}

	return &pb.GetProfileResponse{User: &user}, nil
}

func (r *UserRepo) UpdateProfile(req *pb.UpdateProfileRequest) (*pb.Status, error) {
	query := `
		UPDATE
			users
		SET
			username = $1,
			email = $2,
			full_name = $3,
			phone_number = $4,
			address = $5,
			user_type = $6,
			bio = $7,
			specialties = $8,
			years_of_experience = $9,
			is_verified = $10
		WHERE
			id = $11

	`

	_, err := r.Db.Exec(query, req.Username, req.Email, req.FullName, req.Phone, req.Address, req.UserType, req.Bio, req.Specialties, req.YearsOfExperience, req.IsVerified, req.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil

}

var Verification int32

func (r *UserRepo) ResetPassword(req *pb.ResetPasswordRequest) (*pb.Status, error) {
	if req.Verification != Verification {
		return &pb.Status{Status: false}, nil
	}
	query := `
		UPDATE
			users
		SET
			password_hash = $1
		WHERE
			email = $2
		AND
			deleted_at IS NULL
		`
	_, err := r.Db.Exec(query, req.NewPassword, req.Email)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *UserRepo) StoreRefreshToken(req *models.RefreshToken) error {
	query := `
		INSERT INTO 
			refresh_token
			(token, user_id,email, expires_at, created_at)
		VALUES
			($1, $2, $3, $4, $5)
		`
	_, err := r.Db.Exec(query, req.Token, req.UserId, req.Email, req.ExpiresAt, req.CreatedAt)
	if err != nil {
		log.Fatalf("Error with inserting refresh_token: %v", err)
		return err
	}
	return nil
}

func (r *UserRepo) DeleteProfile(id string) (*pb.Status, error) {
	query := `
		UPDATE
			users
		SET
			deleted_at = $1
		WHERE
			id = $2
		AND
			deleted_at IS NULL
	`
	_, err := r.Db.Exec(query, time.Now(), id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*pb.User, error) {
	query := `
		SELECT
			id, username, email, password_hash, full_name, user_type, phone_number
		FROM
			users
		WHERE
			email = $1
		AND
			deleted_at IS NULL
	`
	row := r.Db.QueryRow(query, email)
	var user pb.User
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.UserType,
		&user.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) DeleteRefreshToken(email string) (string, error) {
	query := `
		UPDATE 
			refresh_token
		SET
			deleted_at = $1
		WHERE
			email = $2
		AND
			deleted_at IS NULL
	`
	res, err := r.Db.Exec(query, time.Now().Unix(),email)
	if err != nil {
		return "Logout Unsuccessful, You already logged out!", err
	}
	isNull ,_ := res.RowsAffected()
	if  isNull == 0 {
		return "Logout Unsuccessful, You already logged out!", nil
	}
	return "Logout Successful", nil
}

func (r *UserRepo) RefreshToken(email string) (string, error) {
	query := `
		SELECT
			token
		FROM
			refresh_token
		WHERE
			email = $1
		AND
			deleted_at IS NULL`

	row1 := r.Db.QueryRow(query, email)
	var Token string
	err := row1.Scan(&Token)
	if err != nil {
		return "", err
	}

	return Token, nil
}

func (s *UserRepo) ForgotPassword(req *pb.ResetPasswordRequest) (*pb.Status, error) {
	query := `
		SELECT
			full_name
		FROM
			users
		WHERE
			email = $1
		AND
			deleted_at IS NULL`

	row := s.Db.QueryRow(query, req.Email)
	var name string
	err := row.Scan(&name)
	if err != nil {
		return nil, err
	}


	reds, err := redis.ConnectRedis()
	if err != nil {
		return nil, err
	}
	defer reds.Close()

	ran := rand.Intn(9999)
	err = redis.SendVerificationToEmail(req.Email, strconv.Itoa(ran))
	if err != nil {
		return nil, err
	}

	err = reds.Set(context.Background(),req.Email, strconv.Itoa(ran), 0).Err()
	if err != nil {
		return nil, err
	}
	Verification = int32(ran)
	return &pb.Status{Status: true}, nil
}
