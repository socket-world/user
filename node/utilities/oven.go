package utilities

import (
	"sync"
	"gopkg.in/macaroon-bakery.v2/bakery"

)

var (
	oven_once sync.Once
	oven_instance *bakery.Oven
)

func FetchOven() *bakery.Oven {

	oven_once.Do(func() {

		oven_instance = bakery.NewOven(bakery.OvenParams{Location: `socket.world`})

	})

	return oven_instance

}
