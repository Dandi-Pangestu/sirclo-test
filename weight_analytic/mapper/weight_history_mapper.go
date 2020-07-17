package mapper

import (
	"sirclo-test/weight_analytic/dto/response"
	"sirclo-test/weight_analytic/models"
	"sort"
)

func ToWeightHistoryDTO(m models.WeightHistory) response.WeightHistoryDTO {
	diff := m.Max - m.Min
	return response.WeightHistoryDTO{
		ID:   m.ID,
		Date: m.Date.Format("2006-01-02"),
		Max:  m.Max,
		Min:  m.Min,
		Diff: diff,
	}
}

func ListWeightHistoryDTO(ms []models.WeightHistory) response.ListWeightHistoryDTO {
	var histories []response.WeightHistoryDTO
	var totalMax int
	var totalMin int
	var totalDiff int
	totalRow := len(ms)

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Date.Unix() <= ms[j].Date.Unix()
	})

	for _, v := range ms {
		history := ToWeightHistoryDTO(v)
		histories = append(histories, history)

		totalMax += history.Max
		totalMin += history.Min
		totalDiff += history.Diff
	}

	maxAverage := float64(totalMax) / float64(totalRow)
	minAverage := float64(totalMin) / float64(totalRow)
	diffAverage := float64(totalDiff) / float64(totalRow)

	return response.ListWeightHistoryDTO{
		Histories:   histories,
		MaxAverage:  maxAverage,
		MinAverage:  minAverage,
		DiffAverage: diffAverage,
	}
}
