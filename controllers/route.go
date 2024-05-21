package controllers

import "github.com/gofiber/fiber/v2"

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}
