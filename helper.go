package testish

import (
	"github.com/omidfth/testish/internal/router"
	"github.com/omidfth/testish/internal/types/serviceNames"
	"gorm.io/gorm"
)

type Service interface {
	GetEntity(name serviceNames.ServiceName) interface{}
	GetDB(name serviceNames.ServiceName) *gorm.DB
	Close()
}

type testishService struct {
	entities    map[serviceNames.ServiceName]interface{}
	serviceList []string
}

func NewTestish(options ...*Option) Service {
	r := router.NewRouter()
	addListeners(r)
	models := make(map[serviceNames.ServiceName]interface{})
	var serviceList []string
	for _, option := range options {
		name := option.ServiceName
		entity, srv := r.Serve(name, option)
		models[name] = entity
		serviceList = append(serviceList, srv)
	}

	s := testishService{entities: models, serviceList: serviceList}
	return &s
}

func (t *testishService) Close() {
	for _, service := range t.serviceList {
		stopDockerCompose(service)
	}
}

func (t *testishService) GetEntity(name serviceNames.ServiceName) interface{} {
	return t.entities[name]
}

func (t *testishService) GetDB(name serviceNames.ServiceName) *gorm.DB {
	switch t.entities[name].(type) {
	case *gorm.DB:
		return t.entities[name].(*gorm.DB)
	}
	panic("db not found!")
	return nil
}
