package httphandler

import (
	"fmt"
	"net/http"
	"sirclo-test/weight_analytic/dto/request"
	"sirclo-test/weight_analytic/mapper"
	"sirclo-test/weight_analytic/models"
	repository "sirclo-test/weight_analytic/repository/weight_history"
	"time"

	"github.com/gin-gonic/gin"
)

type WeightHistoryHandler struct {
	WeightHistoryRepo repository.Repository
}

func (h *WeightHistoryHandler) Index(c *gin.Context) {
	weightHistories, err := h.WeightHistoryRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if weightHistories == nil || len(*weightHistories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}

	c.JSON(http.StatusOK, mapper.ListWeightHistoryDTO(*weightHistories))
}

func (h *WeightHistoryHandler) Show(c *gin.Context) {
	id := c.Param("id")

	weightHistory, err := h.WeightHistoryRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if weightHistory == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}

	c.JSON(http.StatusOK, mapper.ToWeightHistoryDTO(*weightHistory))
}

func (h *WeightHistoryHandler) Store(c *gin.Context) {
	var request request.WeightHistoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var weightHistory models.WeightHistory
	weightHistory.Date, _ = time.Parse("2006-01-02", request.Date)
	weightHistory.Max = request.Max
	weightHistory.Min = request.Min
	if err := h.WeightHistoryRepo.SaveOrUpdate(&weightHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success"})
}

func (h *WeightHistoryHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var request request.WeightHistoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var weightHistory models.WeightHistory
	weightHistory.ID = id
	weightHistory.Date, _ = time.Parse("2006-01-02", request.Date)
	weightHistory.Max = request.Max
	weightHistory.Min = request.Min
	fmt.Println(weightHistory)
	if err := h.WeightHistoryRepo.SaveOrUpdate(&weightHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (h *WeightHistoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	weightHistory, err := h.WeightHistoryRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = h.WeightHistoryRepo.Delete(weightHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
