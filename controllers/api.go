package controllers

import (
	"encoding/xml"

	"github.com/bouzourene/xz6-ip/tools"
	"github.com/gofiber/fiber/v2"
)

func ApiJson(c *fiber.Ctx) error {
	ip := c.IP()

	return c.JSON(fiber.Map{
		"ip":      ip,
		"reverse": tools.DnsReverse(ip),
		"version": tools.IpVersion(ip),
	})
}

func ApiXml(c *fiber.Ctx) error {
	ip := c.IP()

	values := struct {
		XMLName xml.Name `xml:"Response"`
		Ip      string   `xml:"Ip"`
		Reverse string   `xml:"Reverse"`
		Version string   `xml:"Version"`
	}{
		Ip:      ip,
		Reverse: tools.DnsReverse(ip),
		Version: tools.IpVersion(ip),
	}

	return c.XML(values)
}

func ApiTextIp(c *fiber.Ctx) error {
	return c.SendString(c.IP())
}

func ApiTextReverse(c *fiber.Ctx) error {
	return c.SendString(
		tools.DnsReverse(c.IP()),
	)
}

func ApiTextVersion(c *fiber.Ctx) error {
	return c.SendString(
		tools.IpVersion(c.IP()),
	)
}
