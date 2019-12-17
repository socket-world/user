package ledger

import (
	"sync"
	"github.com/google/uuid"

)

type User struct {
	Id uuid.UUID	`json:"id"`
	Name string		`json:"name"`

}

type Users map[string]User

var (
	users_once sync.Once
	users_instance Users
)

func FetchUsers() Users {

	users_once.Do(func() {

		users_instance = make(Users)

	})

	return users_instance

}
