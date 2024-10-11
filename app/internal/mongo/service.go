package mongo

import "context"

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(ctx context.Context, uid UID) (*Group, error) {
	return s.repo.Get(ctx, uid)
}

func (s *Service) Create(ctx context.Context, group *Group) error {
	return s.repo.Save(ctx, group)
}

func (s *Service) Update(ctx context.Context, group *Group) error {
	return s.repo.Save(ctx, group)
}

func (s *Service) Delete(ctx context.Context, uid UID) error {
	return s.repo.Delete(ctx, uid)
}
