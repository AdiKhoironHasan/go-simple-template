package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		assertFn func(t *testing.T, hash string, err error)
	}{
		{
			name:     "success with default cost",
			password: "mysecretpassword",
			assertFn: func(t *testing.T, hash string, err error) {
				require.NoError(t, err)
				assert.NotEmpty(t, hash)
				assert.NotEqual(t, "mysecretpassword", hash)
			},
		},
		{
			name:     "empty password still hashes",
			password: "",
			assertFn: func(t *testing.T, hash string, err error) {
				require.NoError(t, err)
				assert.NotEmpty(t, hash)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			tt.assertFn(t, hash, err)
		})
	}
}

func TestHashPasswordWithCustomCost(t *testing.T) {
	hash, err := HashPassword("password", 4)
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password, 4) // low cost for fast tests
	require.NoError(t, err)

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "correct password matches",
			password: password,
			hash:     hash,
			want:     true,
		},
		{
			name:     "wrong password does not match",
			password: "wrongpassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "empty password does not match",
			password: "",
			hash:     hash,
			want:     false,
		},
		{
			name:     "invalid hash returns false",
			password: password,
			hash:     "not-a-valid-hash",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPasswordHash(tt.password, tt.hash)
			assert.Equal(t, tt.want, got)
		})
	}
}
