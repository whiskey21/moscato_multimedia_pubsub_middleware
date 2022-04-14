package modules

//
//import (
//	"fmt"
//	_ "fmt"
//	"github.com/stretchr/testify/assert"
//	"math/rand"
//	"strconv"
//	"testing"
//	"time"
//)
//
//func makeData(IsAlpha bool) MsgUnit{
//	rand.Seed(time.Now().UnixNano())
//	// Set Ipaddr
//	msg := Message{"", "1.0", "", SM}
//	for i := 0; i < 4; i++{
//		itoa := strconv.Itoa(rand.Int() % 256)
//		msg.from += itoa
//		if i != 3{
//			msg.from += "."
//		}
//	}
//
//	// Set Time
//	msg.time += strconv.Itoa(rand.Int()%24) + ":"
//	msg.time += strconv.Itoa(rand.Int()%60)
//
//	// Set Topic, Value, Operator
//	Topic := []int64{}
//	Value := []int64{}
//	Operator := []string{}
//	candOp := []string{">", ">=", "<=" ,"<", "=="}
//	logicalOp := []string{"&&", "||"}
//
//	if IsAlpha {
//		topicLen := rand.Int() % 10 + 1
//		for i := 0 ; i < topicLen; i++{
//			Topic = append(Topic, rand.Int63())
//		}
//
//		valueLen := rand.Int() % 10 + 1
//		for i := 0 ; i < valueLen; i++{
//			Value = append(Value, rand.Int63())
//		}
//
//		Operator = append(Operator, "==")
//
//	} else{
//		valueLen := rand.Int() % 2 + 1
//
//		Topic = append(Topic, rand.Int63())
//		Value = append(Value, rand.Int63())
//
//		if valueLen == 1{
//			Operator = append(Operator, candOp[rand.Int()%5])
//		} else{
//			Value = append(Value, rand.Int63())
//			op := rand.Int()%2
//			Operator = append(Operator, candOp[rand.Int()%2])
//			if op == 0 {
//				Operator = append(Operator, logicalOp[0])
//			} else{
//				Operator = append(Operator, logicalOp[1])
//			}
//			Operator = append(Operator, candOp[rand.Int()%2 + 2])
//		}
//	}
//
//	return &SubscriptionMsg{msg, Topic, Value, Operator, IsAlpha}
//}
//
//func makeMsgList(dataLen int, IsAlpha bool) []MsgUnit{
//	rand.Seed(time.Now().UnixNano())
//	var ret []MsgUnit
//	for i := 0 ; i < dataLen ; i++{
//		ret = append(ret, makeData(IsAlpha))
//	}
//	return ret
//}
//
//func checkOperatorList(isSingle bool, sub int, Operator string, l *valueNode) int{
//	ret := -1
//	if isSingle{
//		switch Operator {
//		case "<":
//			ret = findSub(l.single2sub_s, sub)
//		case "<=":
//			ret = findSub(l.single2sub_es, sub)
//		case ">":
//			ret = findSub(l.single2sub_b, sub)
//		case ">=":
//			ret = findSub(l.single2sub_eb, sub)
//		case "==":
//			ret = findSub(l.single2sub_e, sub)
//		}
//
//	} else{
//		switch Operator {
//		case "<":
//			ret = findSub(l.range2sub_s, sub)
//		case "<=":
//			ret = findSub(l.range2sub_es, sub)
//		case ">":
//			ret = findSub(l.range2sub_b, sub)
//		case ">=":
//			ret = findSub(l.range2sub_eb, sub)
//		}
//	}
//	return ret
//}
//
//func watchData(msgList []MsgUnit, dataLen int, isSubscription bool){
//	for i := 0; i < dataLen; i++{
//		msg := msgList[i]
//		if isSubscription {
//			fmt.Println(
//				"\nfrom = ", msg.(*SubscriptionMsg).Message.from,
//				"\ntime = ", msg.(*SubscriptionMsg).Message.time,
//				"\nTopic = ", msg.(*SubscriptionMsg).Topic,
//				"\nValue = ", msg.(*SubscriptionMsg).Value,
//				"\nOperator = ", msg.(*SubscriptionMsg).Operator,
//				"\nIsAlpha ?= ", msg.(*SubscriptionMsg).IsAlpha,
//			)
//		} else{
//			fmt.Println(
//				"\nfrom = ", msg.(*PublishMsg).Message.from,
//				"\ntime = ", msg.(*PublishMsg).Message.time,
//				"\nTopic = ", msg.(*PublishMsg).Topic,
//				"\nValue = ", msg.(*PublishMsg).Value,
//			)
//		}
//	}
//}
//
////Test addSubScription(1) (dif all [Topic, Value, Operator])
//func Test_addSubscription_allDif(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	// To Init sub_mng
//	mos := Moscato{sub_mng: *newSubmng()}
//
//	// Make Data set(Subscription)
//	var msgList []MsgUnit
//	dataLen := 100
//	msgList = makeMsgList(dataLen, false)
//
//	//Watch Data set
//	//watchData(msgList, dataLen, true)
//
//	for i := 0; i < dataLen; i++ {
//		msg := msgList[i]
//		ip := msg.(*SubscriptionMsg).Message.from
//		Topic := msg.(*SubscriptionMsg).Topic
//		Value := msg.(*SubscriptionMsg).Value
//		Operator := msg.(*SubscriptionMsg).Operator
//		subnumber := mos.sub_mng.count_sub
//		isSingle := true
//
//		// 0. Check addSubscription
//		err := mos.sub_mng.addSubscription(msg)
//		assert.Equal(t, nil, err)
//
//		// 1. Check if ip mapping is correct
//		assert.Equal(t, subnumber, mos.sub_mng.ip2sub[ip][len(mos.sub_mng.ip2sub[ip])-1], "Ip mapping is failed")
//
//		// 2. Check topicNode
//		topicPtr := mos.sub_mng.list.head
//		for topicPtr != nil {
//			if Compare(Topic, topicPtr.Topic) == 0 {
//				break
//			}
//			topicPtr = topicPtr.next
//		}
//
//		assert.Equal(t, Topic, topicPtr.Topic, "topicNode Add is failed")
//
//		// Check isSingle ?
//		if len(Operator) == 3 && Operator[1] == "&&" {
//			isSingle = false
//		}
//
//		// 3. Check Value in ValueNode & Check Operator in ValueNode
//		if !isSingle || (len(Operator) == 3 && Operator[1] == "||") {
//			valptr1 := topicPtr.list.getValueNodePos([]int64{Value[0]})
//			valptr2 := topicPtr.list.getValueNodePos([]int64{Value[1]})
//
//			assert.Equal(t, []int64{Value[0]}, valptr1.val)
//			assert.Equal(t, []int64{Value[1]}, valptr2.val)
//
//			assert.NotEqual(t, -1, checkOperatorList(isSingle, subnumber, Operator[0], valptr1))
//			assert.NotEqual(t, -1, checkOperatorList(isSingle, subnumber, Operator[2], valptr2))
//
//		} else {
//			valptr := topicPtr.list.getValueNodePos(Value)
//			assert.Equal(t, Value, valptr.val)
//			assert.NotEqual(t, -1, checkOperatorList(isSingle, subnumber, Operator[0], valptr))
//		}
//	}
//}
//
//// Test addSubScription(2) (same [Topic, Value] dif [Operator])
//func Test_addSubscription_same_topicNvalue(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//
//	// To Init sub_mng
//	mos := Moscato{sub_mng: *newSubmng()}
//
//	// Fix Topic & Value
//	topicLen := rand.Int()%10 + 1
//	staticTopic := make([]int64, topicLen)
//	staticValue := []int64{rand.Int63()}
//
//
//	// Make Data
//	var msgList []MsgUnit
//	dataLen := 100
//	msgList = makeMsgList(dataLen, false)
//
//	for i := 0; i < topicLen; i++ {
//		staticTopic[i] = rand.Int63()
//	}
//
//	// Fix Same Topic & Value
//	for i := 0; i < dataLen; i++ {
//		Operator := msgList[i].(*SubscriptionMsg).Operator
//		msgList[i].(*SubscriptionMsg).Topic = staticTopic
//
//		if len(Operator) == 1{
//			msgList[i].(*SubscriptionMsg).Value = staticValue
//		}
//	}
//
//	// Watch Data set
//	//watchData(msgList, dataLen, true)
//
//	for i := 0; i < dataLen; i++ {
//		msg := msgList[i]
//		err := mos.sub_mng.addSubscription(msg)
//		assert.Equal(t, nil, err)
//	}
//
//	topicPtr := mos.sub_mng.list.head
//	for topicPtr != nil {
//		if Compare(topicPtr.Topic, staticTopic) == 0 {
//			break
//		}
//		topicPtr = topicPtr.next
//	}
//	watchValueNode(topicPtr.list.head)
//
//}
//
//func watchValueNode(ptr *valueNode) {
//	valPtr := ptr
//	for valPtr != nil{
//		fmt.Println("Value = ", valPtr.val)
//		if len(valPtr.single2sub_s) != 0 {
//			fmt.Println("Single2sub (<) List = ", valPtr.single2sub_s)
//		}
//		if len(valPtr.single2sub_es) != 0 {
//			fmt.Println("Single2sub (<=) List = ", valPtr.single2sub_es)
//		}
//		if len(valPtr.single2sub_b) != 0 {
//			fmt.Println("Single2sub (>) List = ", valPtr.single2sub_b)
//		}
//		if len(valPtr.single2sub_eb) != 0 {
//			fmt.Println("Single2sub (>=) List = ", valPtr.single2sub_eb)
//		}
//		if len(valPtr.single2sub_e) != 0 {
//			fmt.Println("Single2sub (==) List = ", valPtr.single2sub_e)
//		}
//
//		if len(valPtr.range2sub_s) != 0 {
//			fmt.Println("range2sub (<) List = ", valPtr.range2sub_s)
//		}
//		if len(valPtr.range2sub_es) != 0 {
//			fmt.Println("range2sub (<=) List = ", valPtr.range2sub_es)
//		}
//		if len(valPtr.range2sub_b) != 0 {
//			fmt.Println("range2sub (>) List = ", valPtr.range2sub_b)
//		}
//		if len(valPtr.range2sub_eb) != 0 {
//			fmt.Println("range2sub (>=) List = ", valPtr.range2sub_eb)
//		}
//		valPtr = valPtr.next
//	}
//}
