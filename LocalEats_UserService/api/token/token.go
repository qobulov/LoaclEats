package token

import (
	"log"
	"AuthService/config"
	pb "AuthService/genproto/proto"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateJWT(user *pb.User) *pb.Token {
	accesstoken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	accessClaim := accesstoken.Claims.(jwt.MapClaims)
	accessClaim["user_id"] = user.Id
	accessClaim["username"] = user.Username
	accessClaim["email"] = user.Email
	accessClaim["phone_number"] = user.PhoneNumber
	accessClaim["user_type"] = user.UserType
	accessClaim["address"] = user.Address
	accessClaim["full_name"] = user.FullName
	accessClaim["years_of_experience"] = user.YearsOfExperience
	accessClaim["bio"] = user.Bio
	accessClaim["specialties"] = user.Specialties
	accessClaim["is_verified"] = user.IsVerified
	accessClaim["iat"] = time.Now().Unix()
	accessClaim["exp"] = time.Now().Add(time.Hour).Unix()

	con := config.Load()
	access, err := accesstoken.SignedString([]byte(con.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Error with generating access token: %s", err)
	}

	refreshClaim := refreshToken.Claims.(jwt.MapClaims)
	refreshClaim["user_id"] = user.Id
	refreshClaim["username"] = user.Username
	refreshClaim["email"] = user.Email
	refreshClaim["phone_number"] = user.PhoneNumber
	refreshClaim["user_type"] = user.UserType
	refreshClaim["address"] = user.Address
	refreshClaim["full_name"] = user.FullName
	refreshClaim["years_of_experience"] = user.YearsOfExperience
	refreshClaim["bio"] = user.Bio
	refreshClaim["specialties"] = user.Specialties
	refreshClaim["is_verified"] = user.IsVerified
	refreshClaim["iat"] = time.Now().Unix()
	refreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := refreshToken.SignedString([]byte(con.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Error with generating access token: %s", err)
	}

	return &pb.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}
