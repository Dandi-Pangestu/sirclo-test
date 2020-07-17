package request

type WeightHistoryRequest struct {
	Date string `json:"date" binding:"required"`
	Max  int    `json:"max" binding:"required"`
	Min  int    `json:"min" binding:"required"`
}
