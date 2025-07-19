package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	lg := &log.Logger{}
	ms := server.CreateMyServer(lg)
	err := ms.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
