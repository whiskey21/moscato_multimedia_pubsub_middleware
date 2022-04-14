package modules

import (
	"fmt"
	"strconv"
)

// 키관리 부분, 노드 입력받고 키 반환하는 부분 구현
type Security struct {
	KeyMap map[string]string
}

//Security 생성자
func NewSecurity() *Security {
	logger := NewMyLogger()
	defer logger.Sync()
	security := &Security{map[string]string{}}
	defer logger.Debug("security setting complete.")
	return security
}

type SecurityManager interface {
	RegKey(rm RegisterMsg) // 원래 RegisterMsg
	GetNodeKey(nodeName string) int64
	GetImgNodeKey(nodeName string) float64
	ReEncrypt(fromKey int64, toKey int64, target []int64) []int64
	ReEncImgPubMsg(fromPubMsg PublishedImage, nodeName string) PublishedImage
	ReEncPubMsg(fromPubMsg PublishMsg, nodeName string) PublishMsg
	RemoveSecureKey(nodeName string) bool
	//CompareTopic(topic1 []int64, topic2 []int64) int
	//CompareDigit(topic1 int64, topic2 int64) int
	//CompareAlpha(topic1 []int64, topic2 []int64) int
}

/*
keyShareMsg 에서 각 노드의 private 키를 받아 keyMap 에 저장
*/
func (sc Security) RegKey(rm RegisterMsg) {
	sc.KeyMap[rm.Message.From] = strconv.FormatInt(rm.PrivateKey, 10)
}

/**
각 노드의 키를 주소를 이용하여 맵에서 가져옴
*/

func (sc Security) GetImgNodeKey(nodeName string) float64 {

	messageStringKey := sc.KeyMap[nodeName]
	mKey, err := strconv.ParseFloat(messageStringKey, 10)
	if err != nil {
		fmt.Println("GetNodeKey Error: key string to int64 parsing error.")
	}
	return mKey
}

func (sc Security) GetNodeKey(nodeName string) int64 {

	messageStringKey := sc.KeyMap[nodeName]
	mKey, err := strconv.ParseInt(messageStringKey, 10, 64)
	if err != nil {
		fmt.Println("GetNodeKey Error: key string to int64 parsing error.")
	}
	return mKey
}

/*
reEncrypt 해서 슬라이스 반환
*/
func (sc Security) ReEncrypt(fromKey int64, toKey int64, target []int64) []int64 {
	var tmpTarget []int64
	for index := range target {
		tmpTarget = append(tmpTarget, target[index]-fromKey+toKey)
	}

	return tmpTarget
}

func (sc Security) ImgReEncryptWithoutPrivateKey(toKey float64, target []float64) []float64 {
	var tmpTarget []float64
	for index := range target {
		tmpTarget = append(tmpTarget, target[index]+toKey)
	}

	return tmpTarget
}

func (sc Security) ReEncryptWithoutPrivateKey(toKey int64, target []int64) []int64 {
	var tmpTarget []int64
	for index := range target {
		tmpTarget = append(tmpTarget, target[index]+toKey)
	}

	return tmpTarget
}

// topic과 value는 m+k로만 존재하므로 ReEnc과정에서 subscriber의 개인키만 더해주면 된다.

func (sc Security) ReEncImgPubMsg(fromPubMsg PublishedImage, nodeName string) PublishedImage {
	toKey := sc.GetImgNodeKey(nodeName)
	//fromKey := sc.GetNodeKey(fromPubMsg.Message.From)

	toPubMsg := PublishedImage{}
	toPubMsg.Message = fromPubMsg.Message
	toPubMsg.Topic = sc.ImgReEncryptWithoutPrivateKey(toKey, fromPubMsg.Topic)

	fmt.Println("ReEncrypted Message")
	fmt.Println(FloatSlice2String(toPubMsg.Topic))
	return toPubMsg
}

func (sc Security) ReEncPubMsg(fromPubMsg PublishMsg, nodeName string) PublishMsg {
	toKey := sc.GetNodeKey(nodeName)
	fromKey := sc.GetNodeKey(fromPubMsg.Message.From)

	toPubMsg := PublishMsg{}
	toPubMsg.Message = fromPubMsg.Message
	toPubMsg.Topic = sc.ReEncryptWithoutPrivateKey(toKey, fromPubMsg.Topic)
	toPubMsg.Value = sc.ReEncryptWithoutPrivateKey(toKey, fromPubMsg.Value)
	toPubMsg.Content = sc.ReEncrypt(fromKey, toKey, fromPubMsg.Content)

	return toPubMsg
}

//Key제거 함수
func (sc *Security) RemoveSecureKey(nodeName string) bool {
	logger := NewMyLogger()
	defer logger.Sync()
	//삭제 전 존재여부 확인
	_, exists := sc.KeyMap[nodeName]

	if exists {
		delete(sc.KeyMap, nodeName)
		logger.Debug("[" + nodeName + "] : delete Key successful")
		return true
	} else {
		return false
	}
}
