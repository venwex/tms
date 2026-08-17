package auth

/*
type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
*/

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
}
