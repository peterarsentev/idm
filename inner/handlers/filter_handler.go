package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"idm/inner/models"
	"idm/inner/services"
	"strconv"
)

type FilterHandler struct {
	filterService *services.FilterService
}

func NewFilterHandler(filterService *services.FilterService) *FilterHandler {
	return &FilterHandler{
		filterService: filterService,
	}
}

func (h *FilterHandler) Filters(c *fiber.Ctx) error {
	filters, err := h.filterService.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch filters")
	}
	return c.Render("filters/view", fiber.Map{
		"filters": filters,
	})
}

func (h *FilterHandler) CreateView(c *fiber.Ctx) error {
	return c.Render("filter/create", struct{}{})
}

func (h *FilterHandler) Save(c *fiber.Ctx) error {
	filterName := c.FormValue("name")
	if filterName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Filter name is required")
	}

	filter := models.Filter{Name: filterName}
	filterID, err := h.filterService.Add(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save filter")
	}
	return c.Redirect(fmt.Sprintf("/filter/view/%d", filterID))
}

func (h *FilterHandler) View(c *fiber.Ctx) error {
	filterIDStr := c.Params("id")
	filterID, err := strconv.ParseInt(filterIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid filter ID")
	}
	filter, err := h.filterService.FindById(c.Context(), filterID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch filter")
	}
	keys, err := h.filterService.FindKeyByFilterID(filter.ID)
	IDS := make([]int64, len(keys))
	for i, key := range keys {
		IDS[i] = key.ID
	}
	kvalues, err := h.filterService.FindVKeyInKeys(IDS)
	return c.Render("filter/view", fiber.Map{
		"Filter":  filter,
		"Keys":    keys,
		"KValues": kvalues,
	})
}
