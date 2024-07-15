package app

import (
	"main/internal/database"
	"main/internal/transport/rest"
)

func App() {
	database.Connect()
	rest.Handler()
}
