package controllers

import (
	"github.com/bouzourene/xz6-ip/tools"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	ip := c.IP()

	return c.Render("views/home", fiber.Map{
		"ip":      ip,
		"reverse": tools.DnsReverse(ip),
		"version": tools.IpVersion(ip),
	})
}

func AltProto(c *fiber.Ctx) error {
	version := tools.IpVersion(c.IP())

	address := tools.ConfigValue("ADDRESS_IPV6")
	if version == "v6" {
		address = tools.ConfigValue("ADDRESS_IPV4")
	}

	return c.JSON(fiber.Map{
		"address": address,
	})
}
