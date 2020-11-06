package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	key string
}

var ctx context.Context

func NewUser(context context.Context, username string, hash []byte) (*User, error) {
	ctx = context
	id, err := client.Incr(context, "user:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(context, key, "id", id)
	pipe.HSet(context, key, "username", username)
	pipe.HSet(context, key, "hash", hash)
	pipe.HSet(context, "user:by-username", username, id)
	_, err = pipe.Exec(context)
	if err != nil {
		return nil, err
	}
	return &User{key}, nil
}

func (user *User) GetUsername() (string,error) {
	return client.HGet(ctx, user.key, "username").Result()
}

func (user *User) GetHash() ([]byte,error) {
	return client.HGet(ctx, user.key, "hash").Bytes()
}

func GetUserByUsername(username string) (*User, error) {
	id, err := client.HGet(ctx, "user:by-username", username).Int64()
	if err == redis.Nil {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	return &User{key}, nil
}

func RegisterUser(ctx context.Context, username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	_, err = NewUser(ctx, username, hash)
	return err
}

func (user *User) Authenticate(password string) error {
	hash, err := user.GetHash()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrInvalidLogin
	}
	return err
}

func AuthenticateUser(ctx context.Context, username, password string) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return user.Authenticate(password)
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid login")
)