package repository

import (
	"context"
	"errors"
	"otus/internal/domain"
	"otus/internal/repository/dao"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = `public."user"`
)

var (
	ErrUpdate = errors.New("error updating or no changes")
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	rawQuery := r.Builder.Insert(tableName).Columns(dao.UserColumns...).Values(user.Id(), user.UserName(), user.FirstName(), user.LastName(), user.Email(), user.Phone(), user.CreatedAt(), user.ModifiedAt())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id uuid.UUID, upFn func(oldUser *domain.User) (user *domain.User, err error)) (*domain.User, error) {
	oldUser, err := r.oneUser(ctx, id)
	if err != nil {
		return nil, err
	}

	newUser, err := upFn(oldUser)
	if err != nil {
		return nil, err
	}

	rawQuery := r.Builder.Update(tableName).Set("user_name", newUser.UserName()).Set("first_name", newUser.FirstName()).Set("last_name", newUser.LastName()).Set("email", newUser.Email()).Set("phone", newUser.Phone()).Set("modified_at", newUser.ModifiedAt())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return nil, err
	}

	res, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if res.RowsAffected() == 0 {
		return nil, ErrUpdate
	}

	return newUser, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	rawQuery := r.Builder.Delete(tableName).Where("id = ?", id)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ReadUserById(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	return r.oneUser(ctx, id)
}

func (r *Repository) oneUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	rawQuery := r.Builder.Select(dao.UserColumns...).From(tableName).Where("id = ?", id)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return nil, err
	}
	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoUser, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.User])
	if err != nil {
		return nil, err
	}

	return r.toDomainUser(&daoUser)
}
