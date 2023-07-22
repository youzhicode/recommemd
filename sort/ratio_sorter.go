package sort

import (
	"rec_sys/common"
	"sort"
)

type RatioSorter struct {
	Tag string
}

func (rs RatioSorter) Name() string {
	return rs.Tag
}

func (rs RatioSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].PositiveRatio > products[j].PositiveRatio
	})

	return products
}
