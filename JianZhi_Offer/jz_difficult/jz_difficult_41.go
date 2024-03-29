package jz_difficult

import "container/heap"

// MedianFinder 数据流中的中位数
type MedianFinder struct {
	large *IntHeap //这里的large 是大顶堆  是 如：6，7，8
	small *IntHeap //这里的small 是小顶堆  是 如 ：-5 -4，-3
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1] // 这里一定得写 *h 不能写 old
	return value
}

// Constructor /** initialize your data structure here. */
func Constructor() MedianFinder {
	a := &IntHeap{}
	b := &IntHeap{}
	heap.Init(a) //这里是 heap.Init(a) 不是a.Init()
	heap.Init(b)
	return MedianFinder{
		large: a,
		small: b,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(*this.small) == 0 || num <= -(*this.small)[0] {
		heap.Push(this.small, -num)

		if len(*this.small) > len(*this.large)+1 {
			heap.Push(this.large, -heap.Pop(this.small).(int))
		}
	} else {
		heap.Push(this.large, num)
		if len(*this.small) < len(*this.large) {
			heap.Push(this.small, -heap.Pop(this.large).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (*this.large).Len() < this.small.Len() { //指针可以直接调用Len方法 ,值也可以掉用
		return float64(-(*this.small)[0])
	} else if this.large.Len() > this.small.Len() {
		return float64((*this.large)[0])
	}
	return float64(-(*this.small)[0]+(*this.large)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
