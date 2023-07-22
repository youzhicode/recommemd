package source

import "rec_sys/common"

type LocalSource struct {
	SourceName string
}

func (ls LocalSource) Name() string {
	return ls.SourceName
}

func (ls LocalSource) Source() []*common.Product {
	allProducts := []*common.Product{
		{Id: 1, Sale: 1234, Name: "p1"},
		{Id: 2, Sale: 8273, Name: "p2"},
		{Id: 3, Sale: 123, Name: "p3"},
		{Id: 4, Sale: 9883, Name: "p4"},
		{Id: 5, Sale: 352, Name: "p5"},
		{Id: 6, Sale: 492, Name: "p6"},
		{Id: 7, Sale: 4123, Name: "p7"},
		{Id: 8, Sale: 9821, Name: "p8"},
	}

	return allProducts

}
