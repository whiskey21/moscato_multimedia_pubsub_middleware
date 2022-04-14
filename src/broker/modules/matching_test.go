package modules

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"math/rand"
//	"strconv"
//	"testing"
//	"time"
//)
//
//
//func makePubData(IsAlpha bool) MsgUnit{
//	rand.Seed(time.Now().UnixNano())
//	// Set Ipaddr
//	msg := Message{"", "1.0", "", PM}
//	for i := 0; i < 4; i++{
//		itoa := strconv.Itoa(rand.Int() % 256)
//		msg.From += itoa
//		if i != 3{
//			msg.From += "."
//		}
//	}
//
//	// Set Time
//	msg.Time += strconv.Itoa(rand.Int()%24) + ":"
//	msg.Time += strconv.Itoa(rand.Int()%60)
//
//	Topic := []int64{}
//	Value := []int64{}
//	content := []int64{}
//
//	// Set Topic
//	topicLen := rand.Int() % 10 + 1
//	for i := 0 ; i < topicLen; i++{
//		Topic = append(Topic, rand.Int63())
//	}
//
//	// Set Value
//	if IsAlpha {
//		valueLen := rand.Int() % 10 + 1
//		for i := 0 ; i < valueLen; i++{
//			Value = append(Value, rand.Int63())
//		}
//	} else {
//		Topic = append(Topic, rand.Int63())
//		Value = append(Value, rand.Int63())
//	}
//
//	// Set content
//	contentLen := rand.Int() % 10
//	for i := 0; i < contentLen; i++{
//		content = append(content, rand.Int63())
//	}
//
//	return &PublishMsg{msg, Topic, Value, content}
//}
//
//func makeSubData(IsAlpha bool, Topic []int64, Value []int64) MsgUnit{
//	rand.Seed(time.Now().UnixNano())
//	// Set Ipaddr
//	msg := Message{"", "1.0", "", SM}
//	for i := 0; i < 4; i++{
//		itoa := strconv.Itoa(rand.Int() % 256)
//		msg.From += itoa
//		if i != 3{
//			msg.From += "."
//		}
//	}
//
//	// Set Time
//	msg.Time += strconv.Itoa(rand.Int()%24) + ":"
//	msg.Time += strconv.Itoa(rand.Int()%60)
//
//	// Set Topic, Value, Operator
//	Operator := []string{}
//	candOp := []string{">", ">=", "<=" ,"<", "=="}
//	logicalOp := []string{"&&", "||"}
//
//	if IsAlpha {
//		Operator = append(Operator, "==")
//	} else {
//		randSeed := rand.Int() % 2
//		if randSeed == 0 {
//			lop := rand.Int() % 2
//			Operator = append(Operator, candOp[rand.Int()%2])
//			if lop == 0 {
//				Operator = append(Operator, logicalOp[0])
//				for {
//					x := rand.Int63()
//					if x > Value[0] {
//						Value = append(Value, x)
//						break
//					}
//				}
//			} else {
//				Operator = append(Operator, logicalOp[1])
//				Value = append(Value, rand.Int63())
//			}
//			Operator = append(Operator, candOp[rand.Int()%2+2])
//		} else {
//			op := candOp[rand.Int()%5]
//			Operator = append(Operator, op)
//		}
//	}
//
//	return &SubscriptionMsg{msg, Topic, Value, Operator, IsAlpha}
//}
//
//func Test_matching(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//	assert.Equal(t, 1, 1)
//	mos := Moscato{ sub_mng: *newSubmng(),}
//	mos.queue.queue_init()
//
//	// 1. Make Publish Data
//	pubLen := rand.Int()%100 + 1 // Set dataLength
//	var pubDataList []MsgUnit
//
//	for i := 0; i < pubLen; i++{
//		if i % 2 == 1 {
//			msg := makePubData(true)
//			mos.queue.push(msg)
//			pubDataList = append(pubDataList, msg)
//		} else{
//			msg := makePubData(false)
//			mos.queue.push(msg)
//			pubDataList = append(pubDataList, msg)
//		}
//	}
//
//	// 2. Creates subscription data with (the same <Topic, Value> or (the same <Topic, difValue>
//	//    And Add Subscription
//	subLen := pubLen
//	var subDataList []MsgUnit
//	for i := 0; i < subLen; i++{
//		Topic := pubDataList[i].(*PublishMsg).Topic
//		Value := pubDataList[i].(*PublishMsg).Value
//		if i % 2 == 1{
//			msg := makeSubData(true, Topic, Value)
//			mos.sub_mng.addSubscription(msg)
//			subDataList = append(subDataList, msg)
//
//		} else{
//			msg := makeSubData(false, Topic, Value)
//			mos.sub_mng.addSubscription(msg)
//			subDataList = append(subDataList, msg)
//		}
//	}
//
//	// 3. Watch Data
//
//  	// fmt.Println("PubData")
//	// watchData(pubDataList, pubLen, false)
//	// fmt.Println("\nSubData")
//	// watchData(subDataList, subLen, true)
//
//	// 4. Matching
//	for i := 0; i < pubLen; i++ {
//		matching, pubmsg, err := mos.MatchingManager.Matching(&mos.queue, &mos.sub_mng)
//
//		// Check if matching is correct
//		assert.Equal(t, nil, err)
//
//		fmt.Println("matching list = ", matching)
//		fmt.Println("pub msg = ", pubmsg)
//		fmt.Println("err ?= ", err)
//	}
//}
