package ent

// var (
// 	serviceInstance *Service
// 	lock            = &sync.Mutex{}
// )

// // Service ...
// type Service struct {
// 	repo *entities.EntityRepository
// }

// // NewService ...
// func NewService(configs ...func(s *Service) error) *Service {
// 	if serviceInstance == nil {
// 		svc := &Service{}

// 		for _, cfg := range configs {
// 			if err := cfg(svc); err != nil {
// 				return nil
// 			}
// 		}

// 		serviceInstance = svc
// 	}

// 	return serviceInstance
// }

// // WithRepository ...
// func WithRepository(repo entities.EntityRepository) func(s *Service) error {
// 	return func(svc *Service) error {
// 		svc.repo = &repo

// 		return nil
// 	}
// }

// // FindEntity ...
// func (s Service) FindEntity(context.Context, string) (entities.Entity, error) {
// 	return entities.Entity{}, nil
// }

// // ListEntities ...
// func (s Service) ListEntities(context.Context, entities.EntityFilter) ([]entities.ListEntity, error) {
// 	return nil, nil
// }

// // Resource ...
// type Resource struct {
// 	svc *Service
// }

// // New ...
// func New(svc *Service) *Resource {
// 	return &Resource{svc}
// }

// // TestMe ...
// func (r *Resource) TestMe() (entities.Entity, error) {
// 	return r.svc.FindEntity(context.Background(), "1")
// }
