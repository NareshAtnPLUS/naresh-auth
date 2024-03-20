package ports

import (
	"naresh/m/auth/internal/models"
	"naresh/m/auth/internal/types"

	"github.com/golang-jwt/jwt/v4"
)

/*
	This layer sits between framework and domain layer. This layer orchestrates the request
	This layer is connected to the core of the current system and the also connected
	to the external entities of the system such as db and email service etc.
*/

type APIPort interface {
	SignIn(*types.SigInBody) (*types.SignedUser, error)
	SignOut(email string) (types.User, error)
	SignUp(*types.SignUpBody) (*models.User, error)
	ResetPassword(*types.UserEmail) (types.User, error)
	ChangePassword(email, code, newPassword string) (types.User, error)
	ValidateToken(tokenString *string) (*jwt.Token, error)
}