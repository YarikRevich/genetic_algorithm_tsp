package history

import (
	"sort"
	"university/generic_algorithm_project/internal/entity"
)

type DistanceHistoryTracker struct {
	records []entity.DistanceHistoryRecord
}

func (d *DistanceHistoryTracker) AddRecord(record entity.DistanceHistoryRecord) {
	d.records = append(d.records, record)
}

func (d *DistanceHistoryTracker) GetAllOrdered() []entity.DistanceHistoryRecord {
	sort.Slice(d.records, func(i, j int) bool {
		return d.records[i].Distance < d.records[j].Distance
	})

	return d.records
}

func NewDistanceHistoryTracker() *DistanceHistoryTracker {
	return new(DistanceHistoryTracker)
}
