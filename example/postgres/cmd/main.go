package main

import (
	"github.com/omidfth/testish"
	"github.com/omidfth/testish/internal/types/serviceNames"
	"gorm.io/gorm"
	"log"
)

func main() {
	t := testish.NewTestish(
		testish.NewOption(
			serviceNames.POSTGRESQL,
			5432,
			"./postgres_dump.sql",
		),
	)
	var mm []struct {
		gorm.Model
		Name string
	}
	t.GetDB(serviceNames.POSTGRESQL).Select("*").Table("test_models").Scan(&mm)
	log.Println(mm)
}
