package main

import (
	"sirclo-test/weight_analytic/database"
	"sirclo-test/weight_analytic/httphandler"
	repository "sirclo-test/weight_analytic/repository/weight_history"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
)

func main() {
	db, err := memdb.NewMemDB(database.Schema())
	if err != nil {
		panic(err)
	}

	weightHistoryRepo := repository.NewWeightHistoryRepository(db)
	weightHistoryHandler := httphandler.WeightHistoryHandler{weightHistoryRepo}

	router := gin.Default()

	router.POST("/weight-histories", weightHistoryHandler.Store)
	router.GET("/weight-histories", weightHistoryHandler.Index)
	router.GET("/weight-histories/:id", weightHistoryHandler.Show)
	router.PUT("/weight-histories/:id", weightHistoryHandler.Update)
	router.DELETE("/weight-histories/:id", weightHistoryHandler.Delete)

	router.Run(":8080")
}
