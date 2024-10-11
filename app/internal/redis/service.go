package redis

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) Get(ctx context.Context, uid UID) (*Student, error) {
	return s.repository.Get(ctx, uid)
}

func (s *Service) Create(ctx context.Context, dto CreateStudentDTO) (*Student, error) {
	student := dto.Student

	err := s.repository.Save(ctx, student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) Update(ctx context.Context, dto UpdateStudentDTO) (*Student, error) {
	student := dto.Student

	err := s.repository.Save(ctx, student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) Delete(ctx context.Context, uid UID) {
	s.repository.Delete(ctx, uid)
}
