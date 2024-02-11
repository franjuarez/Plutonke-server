package main

import (
	"example/plutonke-server/api"
	"example/plutonke-server/storage"
)

const PORT = ":8080"

func main() {
	testingDb := storage.NewMemoryStorage()

	server := api.NewServer(PORT, testingDb)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
