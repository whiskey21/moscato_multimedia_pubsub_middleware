package modules

import (
	"errors"
	"fmt"
	"strconv"
)

// Structure for managing subscriptions
type sub_manager struct {
	list       topicList          // Data Structure for Manage TopicNode
	count_sub  int                // Subscription#
	emptylist  []int              // For administrate Subscription#(Deleted)
	ip2sub     map[string][]int   // For mapping {ip : Sub#s List}
	sub2ip     map[int]string     // For mapping {Sub# : ip}
	sub2node   map[int][]nodeInfo // For mapping {Sub# : NodeInfo List}
	israngesub map[int]bool       // To manage when deleted

	subnum2ip   map[int]string
	subnum2conv map[int][]float64
}

type nodeInfo struct {
	valNodeList []*valueNode
	topic       []int64
}

type convNode struct {
	subNum int
}

func (manager *sub_manager) Initialize() {
	// Some initialize
	manager.ip2sub = make(map[string][]int)
	manager.sub2ip = make(map[int]string)
	manager.sub2node = make(map[int][]nodeInfo)
	manager.israngesub = make(map[int]bool)

	manager.subnum2ip = make(map[int]string)
	manager.subnum2conv = make(map[int][]float64)
}

func newSubmng() *sub_manager {
	subMng := &sub_manager{}
	subMng.Initialize()
	return subMng
}

func (manager *sub_manager) isDuplicated(msg MsgUnit) bool {
	//from := msg.(SubscriptionImage).From
	//topic := msg.(SubscriptionImage).Topic
	//value := msg.(SubscriptionMsg).Value
	//operator := msg.(SubscriptionMsg).Operator
	//subList := manager.ip2sub[from]
	canFind := false

	//for i := 0; i < len(subList) && canFind == false; i++ {
	//	sub := subList[i]
	//	nodeinfoList := manager.sub2node[sub]
	//
	//	for j := 0; j < len(nodeinfoList); j++ {
	//		node := nodeinfoList[j]

	//if Compare(node.topic, topic) == 0 {
	//	cnt := 0
	//	if len(operator) == 1 {
	//		valPtr := node.valNodeList[0]
	//		if Compare(valPtr.val, value) == 0 {
	//			op := operator[0]
	//			cnt += valPtr.findOperatorList(op, sub, true)
	//		}
	//	} else {
	//		leftop := operator[0]
	//		logicalop := operator[1]
	//		rightop := operator[2]
	//
	//		leftValuePtr := node.valNodeList[0]
	//		rightValuePtr := node.valNodeList[1]
	//
	//		nodeValList := []int64{leftValuePtr.val[0], rightValuePtr.val[0]}
	//
	//		if Compare(nodeValList, value) == 0 {
	//			if logicalop == "&&" {
	//				cnt += leftValuePtr.findOperatorList(leftop, sub, false)
	//				cnt += rightValuePtr.findOperatorList(rightop, sub, false)
	//			} else {
	//				cnt += leftValuePtr.findOperatorList(leftop, sub, true)
	//				cnt += rightValuePtr.findOperatorList(rightop, sub, true)
	//			}
	//		}
	//	}
	//	if cnt == len(node.valNodeList) {
	//		canFind = true
	//		break
	//	}
	//}
	//	}
	//}
	return canFind
}

// To Insert sub#
func (manager *sub_manager) addSubscription(msg MsgUnit) error {
	topic := msg.(SubscriptionImage).Topic
	//value := msg.(SubscriptionMsg).Value
	//operator := msg.(SubscriptionMsg).Operator
	subnumber := 0

	// 0. Check if same IP & same <Topic, Value> exists
	if manager.isDuplicated(msg) == true {
		return errors.New("Duplicater Subscription")
	}

	// 1. Mapping incoming IP address to sub #
	if len(manager.emptylist) == 0 {
		subnumber = manager.count_sub
		manager.ip2sub[msg.(SubscriptionImage).From] = append(manager.ip2sub[msg.(SubscriptionImage).From], subnumber)
		manager.sub2ip[subnumber] = msg.(SubscriptionImage).From

		manager.subnum2conv[subnumber] = msg.(SubscriptionImage).Topic //conv용

		manager.count_sub++

	} else {
		subnumber := manager.emptylist[len(manager.emptylist)-1]
		manager.emptylist = manager.emptylist[:len(manager.emptylist)-1]
		manager.ip2sub[msg.(SubscriptionImage).From] = append(manager.ip2sub[msg.(SubscriptionImage).From], subnumber)
		manager.sub2ip[subnumber] = msg.(SubscriptionImage).From

		manager.subnum2conv[subnumber] = msg.(SubscriptionImage).Topic //conv용
	}

	topicptr := manager.list.head
	existTopic := false

	// 2. Add Subscription

	// Find topic in topiclist, add if not found
	for topicptr != nil {
		if ConvCompare(topicptr.topic, topic) == 0 {
			existTopic = true
			break
		}
		topicptr = topicptr.next
	}

	if !existTopic {
		manager.list.addConvTopic(topic)
		topicptr = manager.list.tail
	}
	fmt.Println("SAVED SUB SM ")
	fmt.Println(FloatSlice2String(topic))

	// 안씀 오류시 날려
	//var addValNodeList []*valueNode

	// if single expression
	//if len(operator) == 1 {
	//	valptr := topicptr.list.getValueNodePos(value)
	//	op := operator[0]
	//
	//	if valptr == nil {
	//		topicptr.list.addValueNode(value)
	//		valptr = topicptr.list.tail
	//	}
	//
	//	addValNodeList = append(addValNodeList, valptr)
	//	manager.sub2node[subnumber] = append(manager.sub2node[subnumber], nodeInfo{addValNodeList, topic})
	//	manager.israngesub[subnumber] = false
	//	valptr.insertSub(op, subnumber, true)
	//
	//	return nil // AddSubscription ok
	//} else {
	//// if Multi expression
	//// (ex) : { (234 < x) && (x <= 1293) } , { (234 < x) || (x < 1293) }
	//
	//leftOperator := operator[0]
	//logicalOperator := operator[1] // For compound expressions bounded by '&&' and '||'
	//rightOperator := operator[2]
	//
	//// Find ValueNode = (topiclist[topic].list.val == value)
	//valptr1 := topicptr.list.getValueNodePos([]int64{value[0]})
	//valptr2 := topicptr.list.getValueNodePos([]int64{value[1]})
	//
	//if valptr1 == nil {
	//	topicptr.list.addValueNode([]int64{value[0]})
	//	valptr1 = topicptr.list.tail
	//}
	//if valptr2 == nil {
	//	topicptr.list.addValueNode([]int64{value[1]})
	//	valptr2 = topicptr.list.tail
	//}
	//
	//addValNodeList = append(addValNodeList, valptr1)
	//addValNodeList = append(addValNodeList, valptr2)
	//manager.sub2node[subnumber] = append(manager.sub2node[subnumber], nodeInfo{addValNodeList, topic})
	//
	//if logicalOperator == "&&" {
	//	// If they are enclosed in '&&' -> Insert Value to range_operator_list
	//	manager.israngesub[subnumber] = true
	//	valptr1.insertSub(leftOperator, subnumber, false)
	//	valptr2.insertSub(rightOperator, subnumber, false)
	//
	//} else {
	//	// if they are enclosed in '||' -> Insert Value to single_operator_list
	//	manager.israngesub[subnumber] = false
	//	valptr1.insertSub(leftOperator, subnumber, true)
	//	valptr2.insertSub(rightOperator, subnumber, true)
	//}
	//	return nil // addSubscription ok
	//}
	return nil
	//return errors.New("Can't addSubscription")
}

// To delete subscriptions
func (manager *sub_manager) delete(from string) error {
	ip := from
	cand := manager.ip2sub[ip]

	for i := 0; i < len(cand); i++ {
		sub := cand[i]

		for j := 0; j < len(manager.sub2node[sub]); j++ {
			nodeinfo := manager.sub2node[sub][j]
			node := nodeinfo.valNodeList

			if manager.israngesub[sub] {
				for k := 0; k < len(node); k++ {
					pos := findSub(node[k].range2sub_s, sub)
					if pos != -1 {
						node[k].range2sub_s = remove(node[k].range2sub_s, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].range2sub_es, sub)
					if pos != -1 {
						node[k].range2sub_es = remove(node[k].range2sub_es, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].range2sub_b, sub)
					if pos != -1 {
						node[k].range2sub_b = remove(node[k].range2sub_b, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].range2sub_eb, sub)
					if pos != -1 {
						node[k].range2sub_eb = remove(node[k].range2sub_eb, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}
				}
			} else {
				for k := 0; k < len(node); k++ {
					pos := findSub(node[k].single2sub_s, sub)
					if pos != -1 {
						node[k].single2sub_s = remove(node[k].single2sub_s, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].single2sub_es, sub)
					if pos != -1 {
						node[k].single2sub_es = remove(node[k].single2sub_es, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].single2sub_b, sub)
					if pos != -1 {
						node[k].single2sub_b = remove(node[k].single2sub_b, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].single2sub_eb, sub)
					if pos != -1 {
						node[k].single2sub_eb = remove(node[k].single2sub_eb, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}

					pos = findSub(node[k].single2sub_e, sub)
					if pos != -1 {
						node[k].single2sub_e = remove(node[k].single2sub_e, pos)
						manager.emptylist = append(manager.emptylist, sub)
					}
				}
			}
		}
		manager.ip2sub[ip] = nil // Delete sub#s mapped to Ip address
		return nil
	}
	return errors.New("Don't Exist Subscription to delete")
}

func FloatSlice2String(target []float64) string {
	var targetString string
	targetString = ""
	for _, value := range target {
		targetString += strconv.FormatFloat(float64(value), 'f', -1, 64) + " "
	}
	return targetString
}
