package mygo_pagination

import (
	"math"

	"github.com/jinzhu/gorm"
)

type Pagging struct {
	//number of reacord's limit on each page
	Limit int `json:"limit" form:"limit"`
	//pagging
	Page int `json:"page" form:"page"`
}

type PaggingInfo struct {
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
	Offset      int `json:"-"`
	Limit       int `json:"-"`
	Page        int `json:"page:"`
}

func (inputs Pagging) GenOffset(query *gorm.DB) (pg PaggingInfo, err error) {
	var total int
	err = query.Count(&total).Error
	if err != nil {
		return
	}

	pg = inputs.PageInfoGenerator(total)
	return
}

func (inputs Pagging) PageInfoGenerator(total int) (pg PaggingInfo) {
	// if page set -1, means get all
	if inputs.Page == -1 {
		pg = PaggingInfo{
			TotalRecord: total,
			TotalPage:   1,
			Offset:      0,
			Limit:       total,
			Page:        1,
		}
	} else {
		pg = PaggingInfo{
			TotalRecord: total,
			TotalPage:   int(math.Ceil(float64(total) / float64(inputs.Limit))),
		}
		if inputs.Page <= 0 {
			inputs.Page = 1
		}
		pg.Page = inputs.Page
		if inputs.Page == 1 {
			pg.Offset = 0
		} else {
			pg.Offset = (inputs.Page - 1) * inputs.Limit
		}
		pg.Limit = (inputs.Page * inputs.Limit) - 1

		if pg.Limit > total {
			pg.Limit = total
		}
	}
	return
}
