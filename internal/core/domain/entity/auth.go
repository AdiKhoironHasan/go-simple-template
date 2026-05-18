package entity

// AuthToken represents the JWT tokens returned after authentication
type AuthToken struct {
	Id           string
	AccessToken  string
	RefreshToken string
}
