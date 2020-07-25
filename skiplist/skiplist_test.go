package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T){
	l := newSkipList()
	l.print()
	l.Insert(0, "0")
	l.Insert(20, "20")
	l.Insert(40, "20")
	l.Insert(35, "20")
	l.Insert(25, "20")
	l.Insert(32, "20")
	l.Insert(15, "20")
	l.Insert(4, "20")
	l.Insert(5, "20")
	l.Insert(87, "20")
	l.Insert(123, "20")
	l.Insert(1, "20")
	l.print()
	fmt.Println(l.Search(35))
	l.Delete(87)
	l.Delete(35)
	l.Delete(40)
	l.Delete(25)
	l.Delete(32)
	l.print()
}