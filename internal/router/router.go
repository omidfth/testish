package router

import (
	"github.com/omidfth/testish/internal/types/serviceNames"
)

type Router interface {
	Serve(key serviceNames.ServiceName, i interface{}) (interface{}, string)
	On(key serviceNames.ServiceName, f routerFunc) *route
}

func NewRouter() Router {
	return &router{events: make(map[serviceNames.ServiceName]*route)}
}

type route struct {
	handler handler
}

func (s *router) On(key serviceNames.ServiceName, f routerFunc) *route {
	return s.addHandler(key, f)
}

type handler interface {
	Serve(interface{}) (interface{}, string)
}

type routerFunc func(i interface{}) (interface{}, string)

func (f routerFunc) Serve(i interface{}) (interface{}, string) {
	return f(i)
}

type router struct {
	events map[serviceNames.ServiceName]*route
}

func (s *router) addHandler(key serviceNames.ServiceName, handler handler) *route {
	route := route{handler: handler}
	s.events[key] = &route
	return &route
}

func (s *router) Serve(key serviceNames.ServiceName, i interface{}) (interface{}, string) {
	return s.events[key].handler.Serve(i)
}
