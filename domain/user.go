package domain

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

func (u User) ValidateUser() (bool, string) {
	if len(u.FirstName) < 2 || len(u.FirstName) > 20 {
		return false, "First name must be between 2 and 20 characters"
	}

	if len(u.LastName) < 2 || len(u.LastName) > 20 {
		return false, "Last name must be between 2 and 20 characters"
	}

	if len(u.Bio) < 20 || len(u.Bio) > 450 {
		return false, "Bio must be between 20 and 450 characters"
	}

	return true, ""
}

type UserWithID struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

func NewUserWithID(id ID, user User) UserWithID {
	return UserWithID{
		ID:        id.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Bio:       user.Bio,
	}
}

// outra forma de fazer
// func (u UserWithID) MarshalJSON() ([]byte, error) {
// 	type Alias UserWithID
//
// 	return json.Marshal(&struct {
// 		ID string `json:"id"`
// 		Alias
// 	}{
// 		ID:    u.ID.String(),
// 		Alias: (Alias)(u),
// 	})
// }
