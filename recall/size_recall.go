package recall

import "rec_sys/common"

type SizeRecall struct {
	Tag string
}

func (size SizeRecall) Name() string {
	return size.Tag
}

func (size SizeRecall) Recall(n int) (rect []*common.Product) {
	rect = make([]*common.Product, 0, n)

	for _, product := range allProducts {
		if product.Size < 200 {
			rect = append(rect, product)
		}
		if len(rect) >= n {
			break
		}
	}
	return
}
