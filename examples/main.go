package main

import (
	"context"
	"fmt"

	"github.com/racoondad/litegen/examples/biz"
	"github.com/racoondad/litegen/examples/conf"
	"github.com/racoondad/litegen/examples/dal"
	"github.com/racoondad/litegen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
