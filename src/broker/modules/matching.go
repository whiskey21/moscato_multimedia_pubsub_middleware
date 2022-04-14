package modules

import (
	_ "errors"
)

type match_manager struct {
	match_count int // Number of successful matches
}

// Matching operation method
// 1. Traverse TopicList for match (msg.topic == topicNode.topic)
// 2. Traverse TopicNode.ValueList for match (msg.value <operator> valueNode.value)
// 3. Add ipaddr to ipList when successful matching
// 4. Add ipaddr to ipList when range matching
// 5. Delete Duplicated Ipaddrs
// @Return (list of IP addresses of matched subs, Pub Msg, error)
func (moscato *Moscato) Matching(msg MsgUnit) {

	//topic := msg.(PublishMsg).Topic
	//value := msg.(PublishMsg).Value
	//sub_mng := moscato.SubscriptionManager
	//
	//// iplist for return
	//ipList := make([]string, 0)
	//
	//// list for matched range subscriptions
	//big := make([]int, 0)
	//small := make([]int, 0)
	//
	//// Find (topicNode[Topic] == msg.Topic) Node
	//topicPtr := sub_mng.list
	//pos := topicPtr.getTopicNodePos(topic)
	//
	//// Don't Exist topicNode
	//if pos == nil {
	//	moscato.SendQueue <- myType{nil, msg, errors.New("Don't Exist Matching Topic")}
	//} else {
	//	// Traverse all valueNode -> and Match
	//	valPtr := pos.list.head
	//	for valPtr != nil {
	//		compare := Compare(valPtr.val, value)
	//		if compare > 0 { // sub.val > pub.val
	//			// single : { >, >= }
	//
	//			// (1) case : >
	//			for i := 0; i < len(valPtr.single2sub_b); i++ {
	//				sub := valPtr.single2sub_b[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// (2) case : >=
	//			for i := 0; i < len(valPtr.single2sub_eb); i++ {
	//				sub := valPtr.single2sub_eb[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//			// range : { >, >= }
	//
	//			// (1) case : >
	//			for i := 0; i < len(valPtr.range2sub_b); i++ {
	//				sub := valPtr.range2sub_b[i]
	//				big = append(big, sub)
	//			}
	//
	//			// (2) case : >=
	//			for i := 0; i < len(valPtr.range2sub_eb); i++ {
	//				sub := valPtr.range2sub_eb[i]
	//				big = append(big, sub)
	//			}
	//
	//		} else if compare < 0 { // sub.val < pub.val
	//
	//			// single : { <, <= }
	//
	//			// (1) case : <
	//			for i := 0; i < len(valPtr.single2sub_s); i++ {
	//				sub := valPtr.single2sub_s[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// (2) case : <=
	//			for i := 0; i < len(valPtr.single2sub_es); i++ {
	//				sub := valPtr.single2sub_es[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// range : { <, <= }
	//
	//			// (1) case : <
	//			for i := 0; i < len(valPtr.range2sub_s); i++ {
	//				sub := valPtr.range2sub_s[i]
	//				small = append(small, sub)
	//			}
	//			// (2) case : <=
	//			for i := 0; i < len(valPtr.range2sub_es); i++ {
	//				sub := valPtr.range2sub_es[i]
	//				small = append(small, sub)
	//			}
	//
	//		} else { // sub.val == pub.val
	//
	//			// single : { <=, >=, ==}
	//
	//			// (1) case : <=
	//			for i := 0; i < len(valPtr.single2sub_es); i++ {
	//				sub := valPtr.single2sub_es[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// (2) case : >=
	//			for i := 0; i < len(valPtr.single2sub_eb); i++ {
	//				sub := valPtr.single2sub_eb[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// (3) case : ==
	//			for i := 0; i < len(valPtr.single2sub_e); i++ {
	//				sub := valPtr.single2sub_e[i]
	//				ip := sub_mng.sub2ip[sub]
	//				ipList = append(ipList, ip)
	//			}
	//
	//			// range : { <=, >= }
	//
	//			// (1) case : <=
	//			for i := 0; i < len(valPtr.range2sub_es); i++ {
	//				sub := valPtr.range2sub_es[i]
	//				small = append(small, sub)
	//			}
	//
	//			// (2) case : >=
	//			for i := 0; i < len(valPtr.range2sub_eb); i++ {
	//				sub := valPtr.range2sub_eb[i]
	//				big = append(big, sub)
	//			}
	//
	//		}
	//		valPtr = valPtr.next
	//	}
	//
	//	// Add the intersection IP address of two sets (large and small) to the return list
	//	intersectSubnumbers := intersect.Hash(small, big)
	//	list := reflect.ValueOf(intersectSubnumbers)
	//	for i := 0; i < list.Len(); i++ {
	//		sub := list.Index(i).Interface().(int)
	//		ip := sub_mng.sub2ip[sub]
	//		ipList = append(ipList, ip)
	//	}
	//	moscato.MatchingManager.match_count++
	//
	//	// To delete Duplicated Ipaddr
	//	keys := make(map[string]bool)
	//	retIpList := []string{}
	//	for _, val := range ipList {
	//		if _, saveVal := keys[val]; !saveVal{
	//			keys[val] = true
	//			retIpList = append(retIpList, val)
	//		}
	//	}
	//
	//	// Return {IpaddressList, PubMsg, ErrorMsg} to SendQueue
	//	moscato.SendQueue <- myType{retIpList, msg, nil}
	//}
}
