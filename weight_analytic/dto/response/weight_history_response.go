package response

type WeightHistoryDTO struct {
	ID   string `json:"id"`
	Date string `json:"date"`
	Max  int    `json:"max"`
	Min  int    `json:"min"`
	Diff int    `json:"diff"`
}

type ListWeightHistoryDTO struct {
	Histories   []WeightHistoryDTO `json:"histories"`
	MaxAverage  float64            `json:"max_average"`
	MinAverage  float64            `json:"min_average"`
	DiffAverage float64            `json:"diff_average"`
}
