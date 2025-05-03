package repository

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
)

type AuthRepository struct {
	db           *pgxpool.Pool
	publicKeyPEM string // เพิ่ม field สำหรับเก็บ Public Key
}

// NewAuthRepository constructor สำหรับสร้าง authRepository
func NewAuthRepository(db *pgxpool.Pool, pubKey string) *AuthRepository {
	return &AuthRepository{
		db:           db,
		publicKeyPEM: pubKey, // รับค่า Public Key
	}
}

// ValidateAPIKey ใช้ในการตรวจสอบ API Key
func (r *AuthRepository) ValidateAPIKey(ctx context.Context, apiKey string) (int, error) {
	var clientID int
	err := r.db.QueryRow(ctx, "SELECT id FROM clients WHERE api_key = $1", apiKey).Scan(&clientID)
	return clientID, err
}

// LogRequest ใช้ในการบันทึกการเรียก API
func (r *AuthRepository) LogRequest(ctx context.Context, clientID int, endpoint string, method string) error {
	_, err := r.db.Exec(ctx, "INSERT INTO logs (client_id, endpoint, method) VALUES ($1, $2, $3)",
		clientID, endpoint, method)
	return err
}

// ValidateJWTToken ใช้ในการตรวจสอบ JWT Token โดยการใช้ Public Key
func (r *AuthRepository) ValidateJWTToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	// แปลง Public Key จาก PEM format
	block, _ := pem.Decode([]byte(r.publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Cast public key to rsa.PublicKey
	rsaPublicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not RSA public key")
	}

	// Parse JWT Token using the RSA public key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่า signing method เป็น RSA หรือไม่
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not parse JWT token: %v", err)
	}

	// ตรวจสอบว่า token ถูกต้องและ extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
