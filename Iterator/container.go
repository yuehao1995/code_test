/**
 * @author zhangyuehao
 * @date 2019-01-04 10:41
 */

package iterator

import "container/list"

type Container interface {
	Iterator() Iterator
}

type List struct {
	list list.List
}

func (this *List) Iterator() Iterator {
	return &ListIterator{this.list.Front(), this.list.Back()}
}

func (this *List) Add(value interface{}) {
	this.list.PushBack(value)
}
