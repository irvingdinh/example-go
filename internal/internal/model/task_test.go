package model

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	Describe("AsMap", func() {
		It("should returns", func() {
			currentTime := time.Now()

			record := Task{
				Model: Model{
					ID:        1,
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
					DeletedAt: invalidDeletedAt,
				},
				Summary:     "Lorem ipsum dolor sit amet",
				IsCompleted: true,
			}

			Expect(record.AsMap()).To(BeEquivalentTo(M{
				"id":           uint64(1),
				"summary":      "Lorem ipsum dolor sit amet",
				"is_completed": true,
				"created_at":   currentTime.Format(time.RFC3339),
				"updated_at":   currentTime.Format(time.RFC3339),
				"deleted_at":   invalidDeletedAtAsMapValue,
			}))
		})
	})
})
