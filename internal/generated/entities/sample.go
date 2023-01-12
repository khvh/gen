package entities

import (
	"context"
)

// Sample ...
type Sample struct {
	ID    string `json:"_id"`
	Key   string `json:"_key"`
	Title string `json:"title"`
}

// NewSample constructs a Sample
func NewSample() *Sample {
	return &Sample{}
}

// NewSampleWithValues constructs a Sample (using provided values)
func NewSampleWithValues(
	_id string,
	_key string,
	title string,
) *Sample {
	return NewSample().SetID(_id).SetKey(_key).SetTitle(title)
}

// NewSampleWithStruct constructs a Sample (using provided struct)
func NewSampleWithStruct(data *Sample) *Sample {
	o := NewSample()

	if data != nil {
		o.SetID(data.ID)
		o.SetKey(data.Key)
		o.SetTitle(data.Title)
	}

	return o
}

// SetID sets ID and returns Sample
func (e *Sample) SetID(_id string) *Sample {

	if _id != "" {
		e.ID = _id
	}

	return e
}

// SetKey sets Key and returns Sample
func (e *Sample) SetKey(_key string) *Sample {

	if _key != "" {
		e.Key = _key
	}

	return e
}

// SetTitle sets Title and returns Sample
func (e *Sample) SetTitle(title string) *Sample {

	if title != "" {
		e.Title = title
	}

	return e
}

// SampleService ...
type SampleService interface {
}

// SampleRepository ...
type SampleRepository interface {
	Create(context.Context, Sample) (Sample, error)
	ByID(context.Context, string) (Sample, error)
	UpdateByID(context.Context, string, Sample) (Sample, error)
	DeleteByID(context.Context, string) error
	ListAll(context.Context) ([]Sample, error)
	ListWithPagination(context.Context, int, int) ([]Sample, error)
	ListWithPaginationAndFilter(context.Context, int, int, interface{}) ([]Sample, error)
	ListWithStringQuery(context.Context, string) ([]Sample, error)
	ListWithQuery(context.Context, interface{}) ([]Sample, error)
}

// SampleMockRepository ...
type SampleMockRepository struct{}

// NewSampleMockRepository ...
func NewSampleMockRepository() *SampleMockRepository {
	return &SampleMockRepository{}
}

// Create ...
func (r *SampleMockRepository) Create(context.Context, Sample) (*Sample, error) {
	return mockSample(), nil
}

// ByID ...
func (r *SampleMockRepository) ByID(context.Context, string) (*Sample, error) {
	return mockSample(), nil
}

// UpdateByID ...
func (r *SampleMockRepository) UpdateByID(context.Context, string, Sample) (*Sample, error) {
	return mockSample(), nil
}

// DeleteByID ...
func (r *SampleMockRepository) DeleteByID(context.Context, string) error {
	return nil
}

// ListAll ...
func (r *SampleMockRepository) ListAll(context.Context) ([]*Sample, error) {
	return []*Sample{mockSample()}, nil
}

// ListWithPagination ...
func (r *SampleMockRepository) ListWithPagination(context.Context, int, int) ([]*Sample, error) {
	return []*Sample{mockSample()}, nil
}

// ListWithPaginationAndFilter ...
func (r *SampleMockRepository) ListWithPaginationAndFilter(context.Context, int, int, interface{}) ([]*Sample, error) {
	return []*Sample{mockSample()}, nil
}

// ListWithStringQuery ...
func (r *SampleMockRepository) ListWithStringQuery(context.Context, string) ([]*Sample, error) {
	return []*Sample{mockSample()}, nil
}

// ListWithQuery ...
func (r *SampleMockRepository) ListWithQuery(context.Context, interface{}) ([]*Sample, error) {
	return []*Sample{mockSample()}, nil
}

func mockSample() *Sample {
	return NewSample()
}
