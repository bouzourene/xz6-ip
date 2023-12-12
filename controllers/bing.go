package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func BingProxy(c *fiber.Ctx) error {
	url := "https://www.bing.com/HPImageArchive.aspx?format=js&cacheDir=false&date=0&locale=en-us&n=1&orientation=landscape&quality=1980x1080"
	if err := proxy.Do(c, url); err != nil {
		return err
	}

	// Remove headers
	c.Response().Header.Del(fiber.HeaderServer)
	c.Response().Header.Del(fiber.HeaderSetCookie)

	return nil
}
