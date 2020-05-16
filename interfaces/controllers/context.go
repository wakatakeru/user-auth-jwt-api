package controllers

type Context interface {
	Param(string) string
	Bind(interface{})
	Status(int)
	JSON(int, interface{})
}
