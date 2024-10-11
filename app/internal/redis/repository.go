package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	c *redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{client}
}

func (r *Repository) Save(ctx context.Context, dto Student) error {
	data, err := json.Marshal(dto)

	if err != nil {
		return err
	}

	r.c.Set(ctx, string(dto.UID), data, 0)

	return nil
}

func (r *Repository) Get(ctx context.Context, uid UID) (*Student, error) {
	res := r.c.Get(ctx, string(uid))

	if res.Err() != nil {
		return nil, res.Err()
	}

	var student Student

	err := json.Unmarshal([]byte(res.Val()), &student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *Repository) Delete(ctx context.Context, uid UID) {
	r.c.Del(ctx, string(uid))
}
