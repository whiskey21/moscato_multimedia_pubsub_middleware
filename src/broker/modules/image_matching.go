package modules

import (
	"errors"
	"math"
)

func Cosine(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= length_b {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("Vectors should not be null (all zeros)")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func (moscato *Moscato) ImageMatching(msg MsgUnit) {

	topic := msg.(PublishedImage).Topic
	NewMyLogger().Debug(FloatSlice2String(topic))
	sub_mng := moscato.SubscriptionManager

	// iplist for return
	//ipList := make([]string, 0)

	topicPtr := sub_mng.list
	ptr, cosine := topicPtr.getTopicBySimilarity(topic)

	// No topicNode
	if len(ptr) == 0 {
		//fmt.Println("error")
		var empt = make([]float64, 0)
		moscato.SendQueue <- myType{nil, msg, empt, errors.New("No Exist Matching Topic")}
	} else {

		//find exact topic match
		iplist := make([]string, 0)
		var cosineList []float64
		// 2중 for loop 개선필요
		for i := 0; i < len(ptr); i++ {
			for j := 0; j < topicPtr.size; j++ {
				pnt := sub_mng.subnum2conv[j]
				compare := ExactCompare(ptr[i].topic, pnt)
				if compare == 1 {
					ipAddr := sub_mng.sub2ip[j]
					cosineValue := cosine[i]
					iplist = append(iplist, ipAddr)
					cosineList = append(cosineList, cosineValue)
				}
			}
		}
		moscato.MatchingManager.match_count++
		moscato.SendQueue <- myType{iplist, msg, cosineList, nil}
	}
}
