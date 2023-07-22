package source

import "rec_sys/common"

type Sourcer interface {
	Name() string
	Source() []*common.Product
}
