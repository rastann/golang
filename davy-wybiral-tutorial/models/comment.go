package models

import (
	"context"
)

func GetComments(ctx context.Context) ([]string, error) {
	return client.LRange(ctx, "comments", 0, 10).Result()
}

func PostComments(ctx context.Context, comment string) (error) {
	return client.LPush(ctx, "comments", comment).Err()
}