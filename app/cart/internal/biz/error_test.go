package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	cartPb "imperialPalaceMall/api/cart/v1"
	"imperialPalaceMall/app/pkg/middleware"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	err := fun3()
	//t.Logf("%+v", err)

	e := errors.FromError(pkgErrors.Cause(err))
	//t.Logf("%+v", e)

	assert.Equal(t, true, e.Reason == cartPb.ErrorReason_CART_ITEM_NOT_FOUND.String() && e.Code == 404)

	logger := log.With(log.NewStdLogger(os.Stdout))
	_, _ = middleware.LogServer(logger)(next)(context.Background(), nil)
}

func fun1() error {
	err := cartPb.ErrorCartItemNotFound("FindOneBy %d", 2)
	return pkgErrors.Wrap(err, "kratos")
}

func fun2() error {
	err := fun1()
	return err
}

func fun3() error {
	err := fun2()
	return err
}

func next(ctx context.Context, req interface{}) (reply interface{}, err error) {
	return nil, fun3()
}
