package api

import (
	"api-users/domain"

	"github.com/google/uuid"
)

type application struct {
	data map[domain.ID]domain.User
}

func NewApplication() application {
	return application{
		data: make(map[domain.ID]domain.User),
	}
}

func (a application) insert(user domain.User) domain.ID {
	id := domain.ID(uuid.New())

	a.data[id] = user

	return id
}

func (a application) findAll() []domain.UserWithID {
	users := make([]domain.UserWithID, 0, len(a.data))

	for id, user := range a.data {
		users = append(users, domain.NewUserWithID(id, user))
	}

	return users
}

func (a application) findByID(id domain.ID) *domain.UserWithID {
	user, ok := a.data[id]

	if !ok {
		return nil
	}

	userWithID := domain.NewUserWithID(id, user)

	return &userWithID
}

func (a application) deleteByID(id domain.ID) bool {
	if _, ok := a.data[id]; !ok {
		return false
	}

	delete(a.data, id)

	return true
}

func (a application) updateByID(id domain.ID) *domain.UserWithID {
	user, ok := a.data[id]

	if !ok {
		return nil
	}

	userWithID := domain.NewUserWithID(id, user)

	return &userWithID
}
