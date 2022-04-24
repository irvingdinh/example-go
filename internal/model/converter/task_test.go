package converter

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/irvingdinh/example-go/dto"
	"github.com/irvingdinh/example-go/internal/model"
)

var _ = Describe("TaskToDTO", func() {
	Context("nil provided", func() {
		It("should returns", func() {
			output := TaskToDTO(nil)
			Expect(output).To(BeNil())
		})
	})

	Context("Task provided", func() {
		It("should returns", func() {
			currentTime := time.Now()

			output := TaskToDTO(&model.Task{
				Model: model.Model{
					ID:        1,
					CreatedAt: currentTime,
					UpdatedAt: currentTime,
					DeletedAt: invalidNullTime,
				},
				Summary:     "Lorem ipsum dolor sit amet",
				IsCompleted: true,
			})

			Expect(output).To(Equal(&dto.Task{
				ID:          1,
				Summary:     "Lorem ipsum dolor sit amet",
				IsCompleted: true,
				CreatedAt:   currentTime.Format(time.RFC3339),
				UpdatedAt:   currentTime.Format(time.RFC3339),
				DeletedAt:   invalidNullTimeAsString,
			}))
		})
	})
})
