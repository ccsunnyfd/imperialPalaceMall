package middleware

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
)

func JWTAuth(ckTokenFunc CheckToken) middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("cannot get context")
			}
			tr.ReplyHeader().Set("Vary", "Authorization")

			authorizationHeader := tr.RequestHeader().Get("Authorization")
			if authorizationHeader == "" {
				return nil, errors.New("no authorization header")
			}

			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				tr.ReplyHeader().Set("WWW-Authenticate", "Bearer")
				message := "invalid or missing authentication token"
				return nil, errors.New(message)
			}

			// Validate the token to make sure it is in a sensible format.
			err = validateTokenPlaintext(headerParts[1])
			if err != nil {
				return nil, err
			}

			// 请求用户服务去获取userId塞到context中
			user, err := ckTokenFunc(ctx, headerParts[1])
			if err != nil {
				return nil, err
			}
			return next(ContextSetUser(ctx, user), req)
		}
	}
}

func validateTokenPlaintext(tokenPlaintext string) error {
	if tokenPlaintext == "" {
		return errors.New("token must be provided")
	}
	if len(tokenPlaintext) != 26 {
		return errors.New("token must be 26 bytes long")
	}
	return nil
}

type User struct {
	UserId int64
}

var AnonymousUser = &User{}

// IsAnonymous Check if a User instance is the AnonymousUser.
func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type CheckToken func(ctx context.Context, token string) (*User, error)

type contextKey string

const userContextKey = contextKey("user")

func ContextSetUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func ContextGetUser(ctx context.Context) *User {
	user, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
