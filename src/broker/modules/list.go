package modules

import "fmt"

/* List to manage Sub information */

type topicList struct {
	head *topicNode
	tail *topicNode
	size int
}

type topicNode struct {
	topic []float64 // Encrypted topic
	next  *topicNode
	prev  *topicNode
	list  valueList
}

type valueList struct {
	head *valueNode
	tail *valueNode
	size int
}

type valueNode struct {
	val  []int64 // Encrypted value
	next *valueNode
	prev *valueNode

	// single //
	single2sub_s  []int // (sub_val < pub_val) sub#s
	single2sub_es []int // (sub_val <= pub_val) sub#s
	single2sub_b  []int // (sub_val > pub_val) sub#s
	single2sub_eb []int // (sub_val >= pub_val) sub#s
	single2sub_e  []int // (sub_val == pub_val) sub#s

	// range //
	range2sub_s  []int // (sub_val < pub_val) sub#s
	range2sub_es []int // (sub_val <= pub_val) sub#s
	range2sub_b  []int // (sub_val > pub_val) sub#s
	range2sub_eb []int // (sub_val >= pub_val) sub#s
}

// To delete element in slice Array
func remove(ary []int, i int) []int {
	return append(ary[:i], ary[i+1:]...)
}

// To find a specific sub# in a Value node
func findSub(ary []int, sub int) int {
	for i := 0; i < len(ary); i++ {
		if ary[i] == sub {
			return i
		}
	}
	return -1
}

func ExactCompare(a []float64, b []float64) int {
	if len(a) != len(b) {
		return 0
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return 0
			}
		}
	}
	return 1
}

func ConvCompare(a []float64, b []float64) int {
	if len(a) < len(b) {
		return 1
	} else if len(a) > len(b) {
		return -1
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return 1
			} else if a[i] > b[i] {
				return -1
			}
		}
		return 0
	}
}

// Compare -> To compare two encrypted arrays
func Compare(a []int64, b []int64) int {
	if len(a) < len(b) {
		return 1
	} else if len(a) > len(b) {
		return -1
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return 1
			} else if a[i] > b[i] {
				return -1
			}
		}
		return 0
	}
}

func (l *topicList) getTopicBySimilarity(topic []float64) ([]*topicNode, []float64) {

	var threshold = float64(0.9)
	var ptr []*topicNode
	var cosineIdx []float64
	topicPtr := l.head

	for topicPtr != nil {
		if len(topicPtr.topic) == 0 {
			topicPtr = topicPtr.next
			continue
		}
		fmt.Println("Compare", FloatSlice2String(topicPtr.topic), "///", FloatSlice2String(topic))
		cosine, _ := Cosine(topicPtr.topic, topic)
		fmt.Println("Similarity %", cosine)
		if cosine > threshold {
			ptr = append(ptr, topicPtr)
			cosineIdx = append(cosineIdx, cosine)
		}
		topicPtr = topicPtr.next
	}

	return ptr, cosineIdx
}

// To get the position of a Topic node with a specific Topic
func (l *topicList) getTopicNodePos(topic []float64) *topicNode {
	topicPtr := l.head
	for topicPtr != nil {
		if len(topicPtr.topic) == 0 {
			topicPtr = topicPtr.next
			continue
		}
		if ConvCompare(topicPtr.topic, topic) == 0 {
			return topicPtr
		}
		topicPtr = topicPtr.next
	}
	return nil
}

func (l *topicList) addConvTopic(topic []float64) {
	newNode := &topicNode{topic, nil, nil, valueList{}}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// To add a Topic node to the Topic list
func (l *topicList) addTopicNode(topic []float64) {
	newNode := &topicNode{topic, nil, nil, valueList{}}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// To add a Value node to the Value list
func (l *valueList) addValueNode(value []int64) {
	newValNode := &valueNode{value, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	if l.head == nil {
		l.head = newValNode
		l.tail = l.head
	} else {
		newValNode.prev = l.tail
		l.tail.next = newValNode
		l.tail = newValNode
	}
	l.size++
}

// To get the position of a Value node with a specific Value
func (l *valueList) getValueNodePos(value []int64) *valueNode {
	valPtr := l.head
	for valPtr != nil {
		if len(valPtr.val) == 0 {
			valPtr = valPtr.next
			continue
		}
		if Compare(value, valPtr.val) == 0 {
			return valPtr
		}
		valPtr = valPtr.next
	}
	return nil
}

// To insert sub# into the Operator list of the Value node
func (node *valueNode) insertSub(op string, sub int, issingle bool) {
	if issingle == true {
		switch op {
		case "<":
			node.single2sub_s = append(node.single2sub_s, sub)
		case "<=":
			node.single2sub_es = append(node.single2sub_es, sub)
		case ">":
			node.single2sub_b = append(node.single2sub_b, sub)
		case ">=":
			node.single2sub_eb = append(node.single2sub_eb, sub)
		case "==":
			node.single2sub_e = append(node.single2sub_e, sub)
		}
	} else {
		switch op {
		case "<":
			node.range2sub_s = append(node.range2sub_s, sub)
		case "<=":
			node.range2sub_es = append(node.range2sub_es, sub)
		case ">":
			node.range2sub_b = append(node.range2sub_b, sub)
		case ">=":
			node.range2sub_eb = append(node.range2sub_eb, sub)
		}
	}
}

// To check if sub# is in the Value node's operator list
func (node *valueNode) findOperatorList(op string, sub int, issingle bool) int {
	ret := -1
	if issingle == true {
		switch op {
		case "<":
			ret = findSub(node.single2sub_s, sub)
		case "<=":
			ret = findSub(node.single2sub_es, sub)
		case ">":
			ret = findSub(node.single2sub_b, sub)
		case ">=":
			ret = findSub(node.single2sub_eb, sub)
		case "==":
			ret = findSub(node.single2sub_e, sub)
		}
	} else {
		switch op {
		case "<":
			ret = findSub(node.range2sub_s, sub)
		case "<=":
			ret = findSub(node.range2sub_es, sub)
		case ">":
			ret = findSub(node.range2sub_b, sub)
		case ">=":
			ret = findSub(node.range2sub_eb, sub)
		}
	}
	if ret < 0 {
		return 0
	} else {
		return 1
	}
}

// To check if a specific Value node is empty
func (node *valueNode) isEmpty() bool {
	empty := true
	if len(node.single2sub_s) != 0 {
		empty = false
	}
	if len(node.single2sub_es) != 0 {
		empty = false
	}
	if len(node.single2sub_b) != 0 {
		empty = false
	}
	if len(node.single2sub_eb) != 0 {
		empty = false
	}
	if len(node.single2sub_e) != 0 {
		empty = false
	}
	if len(node.range2sub_s) != 0 {
		empty = false
	}
	if len(node.range2sub_es) != 0 {
		empty = false
	}
	if len(node.range2sub_b) != 0 {
		empty = false
	}
	if len(node.range2sub_eb) != 0 {
		empty = false
	}
	return empty
}
