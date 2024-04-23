package repository

import (
	"github.com/omidfth/testish/example/exampleModels"
	"gorm.io/gorm"
)

type testRepo struct {
	db *gorm.DB
}

func (r *testRepo) GetFirst() exampleModels.TestModel {
	var exampleModel exampleModels.TestModel
	r.db.First(&exampleModel)
	return exampleModel
}
