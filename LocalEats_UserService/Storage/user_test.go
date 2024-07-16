package postgres

import (
	pb "AuthService/genproto/users"
	"AuthService/models"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestRegister(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	req := &pb.RegisterRequest{
		Username:    "test",
		Email:       "test",
		Password:    "test",
		FullName:    "test",
		UserType:    "test",
		PhoneNumber: "test",
	}
	res, err := repo.Register(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	id := &pb.UserId{
		Id: "test",
	}
	res, err := repo.GetProfile(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestUpdateProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	req := &pb.UpdateProfileRequest{
		UserId:   "test",
		Username: "test",
		Email:    "test",
		FullName: "test",
		Phone:    "test",
	}
	res, err := repo.UpdateProfile(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteProfile(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	id := "test"
	res, err := repo.DeleteProfile(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetUserByEmail(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	email := "test"
	res, err := repo.GetUserByEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteRefreshToken(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	id := "test"
	res, err := repo.DeleteRefreshToken(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRefreshToken(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	email := "test"
	res, err := repo.RefreshToken(email)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestStoreRefreshToken(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	token := &models.RefreshToken{
		Token:     "test",
		UserId:    "test",
		Email:     "test",
		ExpiresAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	err = repo.StoreRefreshToken(token)
	if err != nil {
		t.Fatal(err)
	}
}
