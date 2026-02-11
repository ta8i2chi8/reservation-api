package jwt

import (
	"os"
	"testing"
)

func TestJWTTokenGeneration(t *testing.T) {
	// Clear the environment variable first
	os.Unsetenv("JWT_SECRET")

	// Test token generation without JWT_SECRET should fail
	_, err := GenerateToken(1, "test@example.com")
	if err == nil {
		t.Error("GenerateToken() should fail without JWT_SECRET environment variable")
	}

	// Test token validation without JWT_SECRET should fail
	_, err = ValidateToken("any-token")
	if err == nil {
		t.Error("ValidateToken() should fail without JWT_SECRET environment variable")
	}

	// Set JWT_SECRET and test again
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateToken(1, "test@example.com")
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
	}

	if token == "" {
		t.Error("GenerateToken() should return a valid token")
	}

	// Test token validation
	claims, err := ValidateToken(token)
	if err != nil {
		t.Errorf("ValidateToken() error = %v", err)
	}

	if claims.UserID != 1 {
		t.Errorf("Expected UserID 1, got %v", claims.UserID)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %v", claims.Email)
	}
}
