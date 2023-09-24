package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joshblades/goshorty/inits"
	"github.com/joshblades/goshorty/models"
	"github.com/joshblades/goshorty/utils"
)

func ParseURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	if shortCode == "" {
		log.Println("No ShortCode provided in request. Redirecting to /.")
		return c.Redirect("/", fiber.StatusMovedPermanently)
	}

	var link models.Link

	inits.Database.First(&link, "short_code = ?", shortCode)

	if link.ID == 0 {
		return c.Status(404).JSON("A link with that URL does not exist.")
	}

	if !link.Enabled {
		return c.Status(400).JSON("Link disabled.")
	}

	return c.Redirect(link.OriginURL, fiber.StatusMovedPermanently)

}

func CreateLink(c *fiber.Ctx) error {
	var link models.Link

	if err := c.BodyParser(&link); err != nil {
		return c.Status(400).JSON("Invalid body in request. Unable to Parse.")
	}

	if link.ShortCode == "" {
		link.ShortCode = utils.RandString(12)
	}

	if err := link.ValidateURL(); err != nil {
		return c.Status(400).JSON("Invalid origin URL. Unable to Parse.")
	}

	link.Enabled = true

	inits.Database.Create(&link)

	return c.Status(201).JSON(link)

}

// GetLinks
// GetLinks

func GetLink(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	var link models.Link

	inits.Database.First(&link, "short_code = ?", shortCode)

	if link.ShortCode == "" {
		return c.Status(404).JSON("Link with provided short code does not exist.")
	}

	return c.Status(200).JSON(&link)

}

func GetLinks(c *fiber.Ctx) error {
	links := []models.Link{}
	inits.Database.Find(&links)

	return c.Status(200).JSON(links)
}

func DeleteLink(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	var link models.Link

	inits.Database.First(&link, "short_code = ?", shortCode)

	if link.ShortCode == "" {
		return c.Status(404).JSON("Link with provided short code does not exist.")
	}

	inits.Database.Delete(&link, 1)

	return c.Status(200).JSON("Link deleted successfully.")

}

func ToggleLinkStatus(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	var link models.Link

	inits.Database.First(&link, "short_code = ?", shortCode)

	if link.ShortCode == "" {
		return c.Status(404).JSON("Link with provided short code does not exist.")
	}

	inits.Database.Model(&link).Update("Enabled", !link.Enabled)

	return c.Status(200).JSON("Link updated successfully.")
}
