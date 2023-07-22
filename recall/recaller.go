package recall

import "rec_sys/common"

type Recaller interface {
	Recall(AllProducts []*common.Product, n int) []*common.Product
	Name() string
}
