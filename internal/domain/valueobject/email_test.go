package valueobject

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "Valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "Invalid email - no @",
			email:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "Invalid email - no domain",
			email:   "test@",
			wantErr: true,
		},
		{
			name:    "Invalid email - empty",
			email:   "",
			wantErr: true,
		},
		{
			name:    "Valid email with subdomain",
			email:   "test@mail.example.com",
			wantErr: false,
		},
		{
			name:    "Valid email with special chars",
			email:   "test.user+tag@example.com",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if email == nil && !tt.wantErr {
				t.Error("NewEmail() returned nil email")
			}
			if email != nil && email.Value() != tt.email {
				t.Errorf("NewEmail() = %v, want %v", email.Value(), tt.email)
			}
		})
	}
}

func TestEmailEquals(t *testing.T) {
	email1, _ := NewEmail("test@example.com")
	email2, _ := NewEmail("test@example.com")
	email3, _ := NewEmail("different@example.com")

	if !email1.Equals(email2) {
		t.Error("Email1 should equal email2")
	}

	if email1.Equals(email3) {
		t.Error("Email1 should not equal email3")
	}

	if email1.Equals(nil) {
		t.Error("Email1 should not equal nil")
	}
}
