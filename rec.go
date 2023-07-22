package main

import (
	"fmt"
	"log"
	"rec_sys/common"
	"rec_sys/filter"
	"rec_sys/recall"
	"rec_sys/sort"
	"time"
)

type Recommender struct {
	Recallers []recall.Recaller // 召回
	Sorter    sort.Sorter       // 排序
	Filters   []filter.Filter   // 过滤
}

func (rec *Recommender) Rec() []*common.Product {
	RecallMap := make(map[int]*common.Product, 100)
	// 顺序遍历，每一路召回
	for _, recaller := range rec.Recallers {
		begin := time.Now()
		products := recaller.Recall(10)

		seconds := time.Since(begin).Nanoseconds()
		log.Printf("召回%s用时%dns, 召回了%d个商品", recaller.Name(), seconds, len(products))
		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("一共召回了%d个商品\n", len(RecallMap))
	// map to slice
	RecallSlice := make([]*common.Product, 0, len(RecallMap))

	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}

	// 对 召回结果作排序
	begin := time.Now()
	SortedResult := rec.Sorter.Sort(RecallSlice)
	log.Printf("排序耗时%d\n", time.Since(begin).Nanoseconds())

	// 顺序执行多个过滤规则
	FilteredResult := SortedResult
	prevCount := len(FilteredResult)
	for _, filter := range rec.Filters {
		begin := time.Now()
		FilteredResult = filter.Filter(FilteredResult)

		log.Printf("过滤规则%s耗时%dns, 过滤掉了%d个商品\n", filter.Name(),
			time.Since(begin).Nanoseconds(), prevCount-len(FilteredResult))
		prevCount = len(FilteredResult)
	}
	return FilteredResult
}

func main() {
	rec := Recommender{
		Recallers: []recall.Recaller{
			recall.HotRecall{Tag: "hot"},
			recall.SizeRecall{Tag: "size"},
		},
		Sorter: sort.RatioSorter{Tag: "ratio"},
		Filters: []filter.Filter{
			filter.RatioFilter{Tag: "ratio"},
		},
	}

	result := rec.Rec()

	for i, p := range result {
		fmt.Printf("第%d名: %d, %s\n", i, p.Id, p.Name)
	}
}
