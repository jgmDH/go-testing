package main

import (
	"fmt"

	"github.com/bootcamp-go/cmd/server/handler"
	"github.com/bootcamp-go/internal/transactions"
	"github.com/bootcamp-go/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Para staticcheck

	p "github.com/bootcamp-go/pruebastaticcheck"
)

// Detectado por staticcheck, golangci-lint
//var variablenousada string

func main() {
	loadEnv()

	// Detectado por vet, staticcheck, golangci-lint
	fmt.Printf("%d prueba con vet\n\n\n", 2)

	// Detectado por staticcheck
	fmt.Println(p.Print())
	fmt.Println(p.Print())

	router := gin.Default()
	db := store.NewFileStore(store.FileType, "./transactions.json")
	r := transactions.NewRepository(db)
	s := transactions.NewService(r)
	h := handler.NewTransaction(s)

	rTransaction := router.Group("transactions")
	rTransaction.GET("/", h.GetAll())
	rTransaction.POST("/", h.Store())

	// Detectado por golangci-lint
	if err := router.Run(); err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
