package main

import (
	"example/plutonke-server/api"
	"example/plutonke-server/storage"
)

const PORT = ":8080"

func main() {
	// testingDb := storage.NewMemoryStorage()
	// server := api.NewServer(PORT, testingDb)
	
	sqlDb := storage.NewSQLiteStorage()
	server := api.NewServer(PORT, sqlDb)


	if err := server.Start(); err != nil {
		panic(err)
	}
	server.Close()
}


