package modules

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTopicNode(t *testing.T) {

	//list := &topicList{nil, nil, 0}
	//ary := []int64{123, 456, 789, 1234}
	//
	//for i := 0; i < 4; i++ {
	//	tmp := []int64{ary[i]}
	//	list.addTopicNode(tmp)
	//	exp := []int64{ary[i]}
	//
	//	assert.Equal(t, list.tail.topic, exp, "add TopicNode failed")
	//}
}

func Test_addValueNode(t *testing.T) {

	list := &valueList{nil, nil, 0}
	ary := []int64{123, 456, 789, 1234}

	for i := 0; i < 4; i++ {
		tmp := []int64{ary[i]}
		list.addValueNode(tmp)
		exp := []int64{ary[i]}

		assert.Equal(t, list.tail.val, exp, "add ValueNode failed")
	}
}

func Test_findSub(t *testing.T) {

	ary := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		assert.Equal(t, findSub(ary, i+1), i, "findSub is failed")
	}

	// Don't Exist
	assert.Equal(t, findSub(ary, 7), -1, "findSub is failed")

	// array = {1, 2, 3, 4, 5, 6} -> {1, 2, 4, 5, 6}
	ary = remove(ary, 2)
	assert.Equal(t, findSub(ary, 6), 4)
}

func Test_remove(t *testing.T) {
	l := make([]int, 4)
	var i int
	for i = 0; i < 4; i++ {
		l[i] = i
	}

	// array = {0, 1, 2, 3} -> {0, 1, 3}
	l = remove(l, 2)
	assert.Equal(t, l, []int{0, 1, 3}, "Array Delete is failed")

	// array = {0, 1, 3} -> {1, 3}
	l = remove(l, 0)
	assert.Equal(t, l, []int{1, 3}, "Array Delete is failed")

	// array = {1, 2} -> {1}
	l = remove(l, 1)
	assert.Equal(t, l, []int{1}, "Array Delete is failed")

	// array = {1} -> {}
	l = remove(l, 0)
	assert.Equal(t, l, []int{}, "Array Delete is failed")
}

func Test_isempty(t *testing.T) {
	node := &valueNode{next: nil, prev: nil}

	// 1. One element in Node
	node.single2sub_eb = append(node.single2sub_eb, 1)
	assert.Equal(t, node.isEmpty(), false, "isEmpty is failed")

	// 2. Empty
	node.single2sub_eb = remove(node.single2sub_eb, 0)
	assert.Equal(t, node.isEmpty(), true, "isEmpty is failed")

	// 3. Many elements in Node
	node.single2sub_b = append(node.single2sub_b, 1)
	node.single2sub_eb = append(node.single2sub_eb, 1)
	node.single2sub_s = append(node.single2sub_s, 1)
	node.single2sub_es = append(node.single2sub_es, 1)
	node.single2sub_e = append(node.single2sub_e, 1)
	assert.Equal(t, node.isEmpty(), false, "isEmpty is failed")
}

func Test_insertSub(t *testing.T) {
	node := &valueNode{next: nil, prev: nil}

	// 1. Single2sub

	// (1) <
	node.insertSub("<", 1, true)
	assert.Equal(t, 0, findSub(node.single2sub_s, 1), "singleInsertSub (<) is failed")

	// (2) <=
	node.insertSub("<=", 2, true)
	assert.Equal(t, 0, findSub(node.single2sub_es, 2), "singleInsertSub (<=) is failed")

	// (3) >
	node.insertSub(">", 3, true)
	assert.Equal(t, 0, findSub(node.single2sub_b, 3), "singleInsertSub (>) is failed")

	// (4) >=
	node.insertSub(">=", 4, true)
	assert.Equal(t, 0, findSub(node.single2sub_eb, 4), "singleInsertSub (>=) is failed")

	// (5) ==
	node.insertSub("==", 5, true)
	assert.Equal(t, 0, findSub(node.single2sub_e, 5), "singleInsertSub (==) is failed")

	// 2. range2sub

	// (1) <
	node.insertSub("<", 1, false)
	assert.Equal(t, 0, findSub(node.range2sub_s, 1), "rangeInsertSub (<) is failed")

	// (2) <=
	node.insertSub("<=", 2, false)
	assert.Equal(t, 0, findSub(node.range2sub_es, 2), "rangeInsertSub (<=) is failed")

	// (3) >
	node.insertSub(">", 3, false)
	assert.Equal(t, 0, findSub(node.range2sub_b, 3), "rangeInsertSub (>) is failed")

	// (4) >=
	node.insertSub(">=", 4, false)
	assert.Equal(t, 0, findSub(node.range2sub_eb, 4), "rangeInsertSub (>=) is failed")

}

func Test_getTopicPos(t *testing.T) {
	//l := topicList{}
	//l.addTopicNode(nil)
	//
	//// 1. find Topic in empty topiclist
	//if l.head.topic != nil {
	//	fmt.Println("addTopicNode is failed")
	//}
	//
	//// 2. find {{12},{34},{56},{78},{89},{90}} in topiclist
	//ary := [][]int64{{12}, {34}, {56}, {78}, {89}, {90}}
	//if l.getTopicNodePos(ary[0]) != nil {
	//	fmt.Println("getTopicNodePos is failed")
	//}
	//
	//for i := 0; i < 5; i++ {
	//	l.addTopicNode(ary[i])
	//	assert.Equal(t, ary[i], l.tail.topic)
	//}

}

func Test_getValuePos(t *testing.T) {
	l := valueList{}
	l.addValueNode(nil)

	ary := [][]int64{{12}, {34}, {56}, {78}, {89}, {90}}

	// 1. find {{12},{34},{56},{78},{89},{90}} in topiclist
	for i := 0; i < 5; i++ {
		l.addValueNode(ary[i])
		assert.Equal(t, ary[i], l.tail.val)
	}
}
