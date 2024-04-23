package main

import (
	"github.com/omidfth/testish"
	"github.com/omidfth/testish/internal/types/serviceNames"
)

func main() {
	testish.NewTestish(
		testish.NewOption(
			serviceNames.MYSQL,
			3309,
			"./mysql_dump.sql",
		),
	)
}
