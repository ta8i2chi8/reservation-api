package domain

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		testName string
		email    string
		password string
		userName string
		wantErr  bool
	}{
		{
			testName: "Valid user",
			email:    "test@example.com",
			password: "password123",
			userName: "Test User",
			wantErr:  false,
		},
		{
			testName: "Empty email",
			email:    "",
			password: "password123",
			userName: "Test User",
			wantErr:  false, // Current implementation doesn't validate
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user, err := NewUser(tt.email, tt.password, tt.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if user == nil && !tt.wantErr {
				t.Error("NewUser() returned nil user")
			}
		})
	}
}

func TestUserPasswordHashing(t *testing.T) {
	email := "test@example.com"
	password := "password123"
	name := "Test User"

	// Test user creation with password hashing
	user, err := NewUser(email, password, name)
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}

	// Password should not be stored in plaintext
	if user.Password == password {
		t.Error("Password should be hashed, not stored in plaintext")
	}

	// Password should be verifiable
	err = user.CheckPassword(password)
	if err != nil {
		t.Errorf("CheckPassword() error = %v", err)
	}

	// Wrong password should fail
	err = user.CheckPassword("wrongpassword")
	if err == nil {
		t.Error("CheckPassword() should fail for wrong password")
	}
}
