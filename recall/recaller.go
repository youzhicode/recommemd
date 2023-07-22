package recall

import "rec_sys/common"

type Recaller interface {
	Recall(n int) []*common.Product
	Name() string
}
