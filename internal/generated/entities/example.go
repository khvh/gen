// Code generated by "gen"; DO NOT EDIT.

package entities

import (
	"context"

	"time"
)

// Example ...
type Example struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewExample constructs a Example
func NewExample() *Example {
	return &Example{}
}

// NewExampleWithValues constructs a Example (using provided values)
func NewExampleWithValues(
	id string,
	name string,
	createdAt time.Time,
) *Example {
	return NewExample().SetID(id).SetName(name).SetCreatedAt(createdAt)
}

// NewExampleWithStruct constructs a Example (using provided struct)
func NewExampleWithStruct(data *Example) *Example {
	o := NewExample()

	if data != nil {
		o.SetID(data.ID)
		o.SetName(data.Name)
		o.SetCreatedAt(data.CreatedAt)
	}

	return o
}

// SetID sets ID and returns Example
func (e *Example) SetID(id string) *Example {

	if id != "" {
		e.ID = id
	}

	return e
}

// SetName sets Name and returns Example
func (e *Example) SetName(name string) *Example {

	if name != "" {
		e.Name = name
	}

	return e
}

// SetCreatedAt sets CreatedAt and returns Example
func (e *Example) SetCreatedAt(createdAt time.Time) *Example {

	return e
}

// ExampleFilter ...
type ExampleFilter struct {
	ID string `json:"id"`
}

// NewExampleFilter constructs a ExampleFilter
func NewExampleFilter() *ExampleFilter {
	return &ExampleFilter{}
}

// NewExampleFilterWithValues constructs a ExampleFilter (using provided values)
func NewExampleFilterWithValues(
	id string,
) *ExampleFilter {
	return NewExampleFilter().SetID(id)
}

// NewExampleFilterWithStruct constructs a ExampleFilter (using provided struct)
func NewExampleFilterWithStruct(data *ExampleFilter) *ExampleFilter {
	o := NewExampleFilter()

	if data != nil {
		o.SetID(data.ID)
	}

	return o
}

// MapExampleFilterToExample maps ExampleFilter to Example
func MapExampleFilterToExample(data *ExampleFilter) *Example {
	return NewExampleWithStruct(&Example{
		ID: data.ID,
	})
}

// MapExampleToExampleFilter maps Example to ExampleFilter
func MapExampleToExampleFilter(data *Example) *ExampleFilter {
	return NewExampleFilterWithStruct(&ExampleFilter{
		ID: data.ID,
	})
}

// SetID sets ID and returns ExampleFilter
func (e *ExampleFilter) SetID(id string) *ExampleFilter {

	if id != "" {
		e.ID = id
	}

	return e
}

func mockExampleFilter() ExampleFilter {
	return ExampleFilter{}
}

// ListExample ...
type ListExample struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewListExample constructs a ListExample
func NewListExample() *ListExample {
	return &ListExample{}
}

// NewListExampleWithValues constructs a ListExample (using provided values)
func NewListExampleWithValues(
	id string,
	name string,
) *ListExample {
	return NewListExample().SetID(id).SetName(name)
}

// NewListExampleWithStruct constructs a ListExample (using provided struct)
func NewListExampleWithStruct(data *ListExample) *ListExample {
	o := NewListExample()

	if data != nil {
		o.SetID(data.ID)
		o.SetName(data.Name)
	}

	return o
}

// MapListExampleToExample maps ListExample to Example
func MapListExampleToExample(data *ListExample) *Example {
	return NewExampleWithStruct(&Example{
		ID:   data.ID,
		Name: data.Name,
	})
}

// MapExampleToListExample maps Example to ListExample
func MapExampleToListExample(data *Example) *ListExample {
	return NewListExampleWithStruct(&ListExample{
		ID:   data.ID,
		Name: data.Name,
	})
}

// SetID sets ID and returns ListExample
func (e *ListExample) SetID(id string) *ListExample {

	if id != "" {
		e.ID = id
	}

	return e
}

// SetName sets Name and returns ListExample
func (e *ListExample) SetName(name string) *ListExample {

	if name != "" {
		e.Name = name
	}

	return e
}

func mockListExample() ListExample {
	return ListExample{}
}

// ExampleService ...
type ExampleService interface {
	FindExample(context.Context, string) (Example, error)
	ListEntities(context.Context, *ExampleFilter) ([]*ListExample, error)
}

// ExampleRepository ...
type ExampleRepository interface {
	Create(context.Context, Example) (Example, error)
	ByID(context.Context, string) (Example, error)
	UpdateByID(context.Context, string, Example) (Example, error)
	DeleteByID(context.Context, string) error
	ListAll(context.Context) ([]Example, error)
	ListWithPagination(context.Context, int, int) ([]Example, error)
	ListWithPaginationAndFilter(context.Context, int, int, interface{}) ([]Example, error)
	ListWithStringQuery(context.Context, string) ([]Example, error)
	ListWithQuery(context.Context, interface{}) ([]Example, error)
}

// ExampleMockRepository ...
type ExampleMockRepository struct{}

// NewExampleMockRepository ...
func NewExampleMockRepository() *ExampleMockRepository {
	return &ExampleMockRepository{}
}

// Create ...
func (r *ExampleMockRepository) Create(context.Context, *Example) (*Example, error) {
	return mockExample(), nil
}

// ByID ...
func (r *ExampleMockRepository) ByID(context.Context, string) (*Example, error) {
	return mockExample(), nil
}

// UpdateByID ...
func (r *ExampleMockRepository) UpdateByID(context.Context, string, *Example) (*Example, error) {
	return mockExample(), nil
}

// DeleteByID ...
func (r *ExampleMockRepository) DeleteByID(context.Context, string) error {
	return nil
}

// ListAll ...
func (r *ExampleMockRepository) ListAll(context.Context) ([]*Example, error) {
	return []*Example{mockExample()}, nil
}

// ListWithPagination ...
func (r *ExampleMockRepository) ListWithPagination(context.Context, int, int) ([]*Example, error) {
	return []*Example{mockExample()}, nil
}

// ListWithPaginationAndFilter ...
func (r *ExampleMockRepository) ListWithPaginationAndFilter(context.Context, int, int, interface{}) ([]*Example, error) {
	return []*Example{mockExample()}, nil
}

// ListWithStringQuery ...
func (r *ExampleMockRepository) ListWithStringQuery(context.Context, string) ([]*Example, error) {
	return []*Example{mockExample()}, nil
}

// ListWithQuery ...
func (r *ExampleMockRepository) ListWithQuery(context.Context, interface{}) ([]*Example, error) {
	return []*Example{mockExample()}, nil
}

func mockExample() *Example {
	return NewExample()
}
