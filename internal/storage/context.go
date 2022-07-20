package storage

import "context"

type storegeContextKey struct{}

func WithStorage(ctx context.Context, repo Repo) context.Context {
	return context.WithValue(ctx, storegeContextKey{}, repo)
}

func GetStorage(ctx context.Context) Repo {
	return ctx.Value(storegeContextKey{}).(Repo)
}
