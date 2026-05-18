package dto

import (
	"go-simple-template/internal/core/domain/entity"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Auth Request DTOs

type AuthRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthLogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Auth Response DTOs

type AuthRegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthProfileResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfileResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *AuthRegisterRequest) ToEntity() entity.User {
	return entity.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}

func (d *AuthLoginRequest) ToEntity() entity.User {
	return entity.User{
		Email:    d.Email,
		Password: d.Password,
	}
}

func (d *AuthRefreshRequest) ToEntity() entity.AuthToken {
	return entity.AuthToken{
		RefreshToken: d.RefreshToken,
	}
}

func (d *AuthLogoutRequest) ToEntity() entity.AuthToken {
	return entity.AuthToken{
		RefreshToken: d.RefreshToken,
	}
}

func ToRegisterResponse(user *entity.User) *AuthRegisterResponse {
	return &AuthRegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToAuthTokenResponse(token *entity.AuthToken) *AuthTokenResponse {
	return &AuthTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}

func ToProfileResponse(user *entity.User) *UserProfileResponse {
	return &UserProfileResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (d *AuthRegisterRequest) ValidateRegister() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required, validation.Length(3, 20)),
		validation.Field(&d.Email,
			validation.Required,
			validation.Length(5, 50),
			is.EmailFormat.Error("email is not valid"),
		),
		validation.Field(&d.Password, validation.Required, validation.Length(6, 100)),
	)
}

func (d *AuthLoginRequest) ValidateLogin() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Email,
			validation.Required,
			validation.Length(5, 100),
		),
		validation.Field(&d.Password, validation.Required, validation.Length(6, 100)),
	)
}

func (d *AuthRefreshRequest) ValidateRefreshToken() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.RefreshToken, validation.Required),
	)
}

func (d *AuthLogoutRequest) ValidateLogout() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.RefreshToken, validation.Required),
	)
}
