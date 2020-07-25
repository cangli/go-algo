package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

// 固定最大层数和概率因子，该跳表可容纳的元素数量为(1/p)^MaxLevel = 2^32
const (
	MaxLevel = 32
	Probability = 0.5
)


type skipListNode struct {
	key int
	val string
	level int	//标识该跳表节点的层高
	forward [MaxLevel]*skipListNode
}

func NewNode(level, key int, val string)*skipListNode{
	return &skipListNode{
		key:     key,
		val:     val,
		level: level,
	}
}

type skipList struct {
	header *skipListNode // header节点不储存数据
	r *rand.Rand	// 随机数生成器
	maxLevel int	// 跳表最大层高
}

func newSkipList()*skipList{
	return &skipList{
		header:   &skipListNode{
			level:   MaxLevel,
		},
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
		maxLevel: 0,
	}
}

func (l *skipList)print(){
	for i := 0; i < l.maxLevel; i++ {
		node := l.header
		for node != nil{
			if i >= l.maxLevel - node.level{
				fmt.Printf("%4d\t", node.key)
			}else{
				fmt.Print("\t")
			}
			node = node.forward[0]
		}
		fmt.Print("\n")
	}
}

// 搜索给定key的值，O(lgN)
func (l *skipList)Search(searchKey int)(string, bool){
	x := l.header
	for i:=l.maxLevel-1;i>=0;i--{
		for x.forward[i] != nil && x.forward[i].key < searchKey{
			x = x.forward[i]
			fmt.Println(x.key)
		}
		if x.forward[i] != nil && x.forward[i].key == searchKey{
			return x.val, true
		}
	}
	return "", false
}

func (l *skipList)randomLevel()int{
	lv := 1
	for l.r.Float64() < Probability && lv < MaxLevel-1{
		lv++
	}
	return lv
}

// 插入新元素,O(lgN)
func (l *skipList)Insert(key int, val string){
	var update [MaxLevel]*skipListNode
	x := l.header
	for i:=l.maxLevel-1;i>=0;i--{
		for x.forward[i] != nil && x.forward[i].key < key{
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key{
		x.val = val
	}else{
		level := l.randomLevel()
		if level > l.maxLevel{
			for i:=l.maxLevel;i< level;i++{
				update[i] = l.header
			}
			l.maxLevel = level
		}
		newNode := NewNode(level, key, val)
		for i:=0;i<level;i++{
			newNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = newNode
		}
	}
}

// 删除元素，O(lgN)
func (l *skipList)Delete(key int)bool{
	var update [MaxLevel]*skipListNode
	x := l.header
	for i:=l.maxLevel-1;i>=0;i--{
		for x.forward[i] != nil && x.forward[i].key < key{
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x.key == key {
		for i := 0; i < l.maxLevel; i++ {
			if update[i].forward[i] != x {
				break
			} else {
				update[i].forward[i] = x.forward[i]
			}
		}
		var i = l.maxLevel - 1
		for ; l.header.forward[i] == nil; i-- {
		}
		l.maxLevel = i + 1
		return true
	}
	return false
}

