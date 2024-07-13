package token

import (
	"AuthService/config"
	pb "AuthService/genproto/proto"
	"fmt"
	"log"
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
func ExtractClaims(tokenstr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
func RefreshToken(token string) *pb.Token {
	refreshClaim := jwt.New(jwt.SigningMethodHS256)
	refClaim, err := ExtractClaims(token)
	if err != nil {
		log.Fatalf("Error with extracting claims: %s", err)
	}
	claims := refreshClaim.Claims.(jwt.MapClaims)
	claims["user_id"] = refClaim["user_id"]
	claims["username"] = refClaim["username"]
	claims["email"] = refClaim["email"]
	claims["phone_number"] = refClaim["phone_number"]
	claims["user_type"] = refClaim["user_type"]
	claims["address"] = refClaim["address"]
	claims["full_name"] = refClaim["full_name"]
	claims["years_of_experience"] = refClaim["years_of_experience"]
	claims["bio"] = refClaim["bio"]
	claims["specialties"] = refClaim["specialties"]
	claims["is_verified"] = refClaim["is_verified"]
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh, err := refreshClaim.SignedString([]byte(config.Load().SIGNING_KEY))
	if err != nil {
		log.Fatalf("Error with generating access token: %s", err)
	}
	return &pb.Token{
		RefreshToken: refresh,
	}
}
