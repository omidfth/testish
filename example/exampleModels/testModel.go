package exampleModels

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name string
}
