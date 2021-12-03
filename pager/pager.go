package pager

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pager struct {
	Items        int
	Page         int
	ItemsPerPage int
	Pages        int
}

func NewPager(c echo.Context, itemsPerPage int) Pager {
	p := Pager{
		ItemsPerPage: itemsPerPage,
		Page:         1,
	}

	if page := c.QueryParam("page"); page != "" {
		if pageInt, err := strconv.Atoi(page); err != nil {
			if pageInt > 0 {
				p.Page = pageInt
			}
		}
	}

	return p
}

func (p *Pager) SetItems(items int) {
	p.Items = items
	p.Pages = int(math.Ceil(float64(items) / float64(p.ItemsPerPage)))

	if p.Page > p.Pages {
		p.Page = p.Pages
	}
}

func (p *Pager) IsBeginning() bool {
	return p.Page == 1
}

func (p *Pager) IsEnd() bool {
	return p.Page >= p.Pages
}

func (p *Pager) GetOffset() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.ItemsPerPage
}
