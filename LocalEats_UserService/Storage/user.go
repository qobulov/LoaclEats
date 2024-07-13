package postgres

import (
	pb "AuthService/genproto/proto"
	"AuthService/models"
	"database/sql"
	"log"
	"time"

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
			(username, email, password_hash, full_name, user_type, address, phone_number, specialties, years_of_experience, bio, is_verified)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`
	_, err := r.Db.Exec(query, req.Username, req.Email, req.Password, req.FullName, req.UserType, req.Address, req.PhoneNumber, req.Specialties, req.YearsOfExperience, req.Bio, req.IsVerified)
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
		&user.UpdatedAt,)
		
		if err != nil {
			return nil, err
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
			phone_number = $4
			address = $5
			full_name = $6
			user_type = $7
		WHERE
			id = $8

	`

	_, err := r.Db.Exec(query, req.Username, req.Email, req.FullName, req.Phone, req.Address, req.UserType, req.UserId)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil

}

func (r *UserRepo) ResetPassword(req *pb.ResetPasswordRequest) (*pb.Status, error) {
	query := `
		UPDATE
			users
		SET
			password_hash = $1
		WHERE
			email = $2
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
	_, err := r.Db.Exec(query, req.Token, req.UserId, req.Email, req.ExpiresAt,req.CreatedAt)
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
			id, username, email, password_hash, full_name, user_type, address, phone_number
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
		&user.Address,
		&user.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) DeleteRefreshToken(email string) (string,error) {
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
	_, err := r.Db.Exec(query, email)
	if err != nil {
		return "Logout Unsuccessful, You already logged out!",err
	}
	return "Logout Successful",nil
}
