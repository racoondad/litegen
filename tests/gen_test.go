package tests_test

import (
	"context"
	"sync"

	"github.com/racoondad/litegen/tests/.expect/dal_test/query"
)

var (
	useOnce sync.Once
	ctx     = context.Background()
)

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
}
