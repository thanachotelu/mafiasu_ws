package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository struct {
	db   *pgxpool.Pool
	jwks *keyfunc.JWKS
}

// NewAuthRepository constructor สำหรับสร้าง authRepository
func NewAuthRepository(db *pgxpool.Pool, jwksURL string) *AuthRepository {
	option := keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			fmt.Printf("Error refreshing JWKS: %v\n", err)
		},
	}
	jwks, err := keyfunc.Get(jwksURL, option)
	if err != nil {
		log.Fatalf("❌ Failed to get JWKS from URL: %v", err)
		return nil
	}
	if db == nil {
		log.Println("🚨 DB pool is nil!")
	}
	return &AuthRepository{
		db:   db,
		jwks: jwks,
	}
}

// LogRequest ใช้ในการบันทึกการเรียก API
func (r *AuthRepository) LogRequest(ctx context.Context, clientID *int, userID *string, endpoint string, method string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO logs (client_id, user_id, endpoint, method) 
		 VALUES ($1, $2, $3, $4)`,
		clientID, userID, endpoint, method)
	return err
}

// ValidateJWTToken ใช้ในการตรวจสอบ JWT Token โดยการใช้ Public Key
func (r *AuthRepository) ValidateJWTToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	// // แปลง Public Key จาก PEM format
	// block, _ := pem.Decode([]byte(r.publicKeyPEM))
	// if block == nil {
	// 	return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	// }

	// pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse public key: %v", err)
	// }

	// // Cast public key to rsa.PublicKey
	// rsaPublicKey, ok := pub.(*rsa.PublicKey)
	// if !ok {
	// 	return nil, fmt.Errorf("not RSA public key")
	// }

	// // Parse JWT Token using the RSA public key
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// ตรวจสอบว่า signing method เป็น RSA หรือไม่
	// 	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	return rsaPublicKey, nil
	// })

	// if err != nil {
	// 	return nil, fmt.Errorf("could not parse JWT token: %v", err)
	// }

	// // ตรวจสอบว่า token ถูกต้องและ extract claims
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	log.Println("Raw token:", tokenString)
	// 	log.Printf("Decoded claims: %+v\n", claims)
	// 	log.Println("Token valid:", token.Valid)
	// 	return claims, nil
	// }

	// return nil, fmt.Errorf("invalid token")
	token, err := jwt.Parse(tokenString, r.jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("could not parse JWT Token %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (r *AuthRepository) ValidateAPIKey(ctx context.Context, apikey string) (int, error) {
	var clientID int
	err := r.db.QueryRow(ctx, "select id from clients where api_key = $1", apikey).Scan(&clientID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("API key not found")
		}
		return 0, err
	}
	return clientID, nil
}
