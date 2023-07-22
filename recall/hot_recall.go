package recall

import (
	"rec_sys/common"
	"sort"
)

var allProducts = []*common.Product{
	{Id: 1, Sale: 1234, Name: "p1"},
	{Id: 2, Sale: 8273, Name: "p2"},
	{Id: 3, Sale: 123, Name: "p3"},
	{Id: 4, Sale: 9883, Name: "p4"},
	{Id: 5, Sale: 352, Name: "p5"},
	{Id: 6, Sale: 492, Name: "p6"},
	{Id: 7, Sale: 4123, Name: "p7"},
	{Id: 8, Sale: 9821, Name: "p8"},
}

type HotRecall struct {
	Tag string
}

func (r HotRecall) Name() string {
	return r.Tag
}

func (r HotRecall) Recall(n int) []*common.Product {

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Sale > allProducts[j].Sale // 销量从大到小
	})
	rect := make([]*common.Product, 0, n)
	for _, product := range allProducts {
		rect = append(rect, product)
		if len(rect) >= n {
			break
		}
	}
	return rect
}
