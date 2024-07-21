package user

import "context"

type UserService interface {
	All(ctx context.Context) ([]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	Search(ctx context.Context, filter *UserFilter, limit int64, offset int64) ([]User, int64, error)
}

func NewUserUseCase(repository UserRepository) *UserUseCase {
	return &UserUseCase{repository: repository}
}

type UserUseCase struct {
	repository UserRepository
}

func (s *UserUseCase) All(ctx context.Context) ([]User, error) {
	return s.repository.All(ctx)
}
func (s *UserUseCase) Load(ctx context.Context, id string) (*User, error) {
	return s.repository.Load(ctx, id)
}
func (s *UserUseCase) Create(ctx context.Context, user *User) (int64, error) {
	return s.repository.Create(ctx, user)
}
func (s *UserUseCase) Update(ctx context.Context, user *User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *UserUseCase) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *UserUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
func (s *UserUseCase) Search(ctx context.Context, filter *UserFilter, limit int64, offset int64) ([]User, int64, error) {
	return s.repository.Search(ctx, filter, limit, offset)
}
