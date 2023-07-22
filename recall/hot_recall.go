package recall

import (
	"rec_sys/common"
	"sort"
)

type HotRecall struct {
	Tag string
}

func (r HotRecall) Name() string {
	return r.Tag
}

func (r HotRecall) Recall(AllProducts []*common.Product, n int) []*common.Product {

	sort.Slice(AllProducts, func(i, j int) bool {
		return AllProducts[i].Sale > AllProducts[j].Sale // 销量从大到小
	})
	rect := make([]*common.Product, 0, n)
	for _, product := range AllProducts {
		rect = append(rect, product)
		if len(rect) >= n {
			break
		}
	}
	return rect
}
