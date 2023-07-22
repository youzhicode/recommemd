package sort

import (
	"rec_sys/common"
	"sort"
)

type SizeSorter struct {
	Tag string
}

func (ss SizeSorter) Name() string {
	return ss.Tag
}

func (ss SizeSorter) Sort(products []*common.Product) []*common.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Size > products[j].Size
	})

	return products
}
