package question_two


type skipListNode struct {
	id string
	name string
	score uint16
	pre *skipListNode //双向链表的前驱指针
	next *skipListNode //双向链表的后继指针
	hnext *skipListNode //处理hash冲突的拉链指针
	down *skipListNode //跳表的下沉指针
}

type skipList struct {
	level int
	headNodeArr []*skipListNode
}

type ranking struct {
	hashMap map[string]*skipListNode
}
