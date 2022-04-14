package modules

import (
	"fmt"
)

//func TestCompare(t *testing.T) {
//	var security = NewSecurity()
//	var sm SecurityManager
//	sm = security
//
//	ksm := KeyShareMsg{Message: Message{From: "1.1.1.1", Version: "1", Time: "2", Kind: 1}, key: "1234"}
//	sm.RegKey(ksm)
//	sm.GetNodeKey(ksm.Message.From)
//	fmt.Println(sm.GetNodeKey(ksm.Message.From))
//	var targetKey []int64
//	targetKey = []int64{1234, 1235, 1236}
//	fmt.Println(sm.ReEncrypt(sm.GetNodeKey(ksm.Message.From), 0, targetKey))
//	//fmt.Println(sm.CompareDigit(1236, 1234))
//
//}

func CreatePubMsg(msg Message, topic string, value string, content string) *PublishMsg {
	toPubMsg := new(PublishMsg)
	toPubMsg.Message = msg

	intArr := []rune(topic)
	//fmt.Print("Topic length ")
	//fmt.Println(len(intArr))
	//fmt.Println(len(toPubMsg.Topic))
	for index := 0; index < len(intArr); index++ {
		toPubMsg.Topic = append(toPubMsg.Topic, int64(intArr[index]))
	}
	//fmt.Println(len(toPubMsg.Topic))
	intArr = []rune(value)
	for index := 0; index < len(intArr); index++ {
		toPubMsg.Value = append(toPubMsg.Value, int64(intArr[index]))
	}
	intArr = []rune(content)
	for index := 0; index < len(intArr); index++ {
		toPubMsg.Content = append(toPubMsg.Content, int64(intArr[index]))
	}

	return toPubMsg
}

func EncryptionMsg(msg *PublishMsg, gyKey int64, privateKey int64) *PublishMsg {
	for index := range msg.Topic {
		msg.Topic[index] = msg.Topic[index] + gyKey
	}
	for index := range msg.Value {
		msg.Value[index] = msg.Value[index] + gyKey
	}
	for index := range msg.Content {
		msg.Content[index] = msg.Content[index] + gyKey + privateKey
	}

	return msg
}

func DecryptionMsg(msg *PublishMsg, gyKey int64, privateKey int64) {
	for index := range msg.Topic {
		msg.Topic[index] = msg.Topic[index] - gyKey - privateKey
	}
	for index := range msg.Value {
		msg.Value[index] = msg.Value[index] - gyKey - privateKey
	}
	for index := range msg.Content {
		msg.Content[index] = msg.Content[index] - gyKey - privateKey
	}

	var runeArr []rune
	for index := range msg.Topic {
		runeArr = append(runeArr, rune(int(msg.Topic[index])))
	}
	fmt.Println("Topic is: " + string(runeArr))
	runeArr = nil

	for index := range msg.Value {
		runeArr = append(runeArr, rune(int(msg.Value[index])))
	}
	fmt.Println("Value is: " + string(runeArr))
	runeArr = nil

	for index := range msg.Content {
		runeArr = append(runeArr, rune(int(msg.Content[index])))
	}
	fmt.Println("Content is: " + string(runeArr))
	runeArr = nil

}

func printMsg(msg *PublishMsg) {
	var runeArr []rune
	for index := range msg.Topic {
		runeArr = append(runeArr, rune(int(msg.Topic[index])))
	}
	fmt.Println("Topic is: " + string(runeArr))
	runeArr = nil

	for index := range msg.Value {
		runeArr = append(runeArr, rune(int(msg.Value[index])))
	}
	fmt.Println("Value is: " + string(runeArr))
	runeArr = nil

	for index := range msg.Content {
		runeArr = append(runeArr, rune(int(msg.Content[index])))
	}
	fmt.Println("Content is: " + string(runeArr) + "\n\n")
	runeArr = nil
}

// From "1.1.1.1" to "3.3.3.3" node
//func TestReEnc(t *testing.T) {
//	var security = NewSecurity()
//	var sm SecurityManager
//	sm = security
//	security.KeyMap["1.1.1.1"] = "56789"
//	security.KeyMap["3.3.3.3"] = "99999"
//
//	//fmt.Println(sm.GetNodeKey("3.3.3.3"))
//	msg := Message{From: "1.1.1.1", Version: "1", Time: "123", Kind: 3}
//	//fmt.Println(msg)
//	publishMsg := CreatePubMsg(msg, "soccer123한글", "playerList", "Son and 10 players")
//	fmt.Println(publishMsg)
//	fmt.Println("original publish message is...")
//	printMsg(publishMsg)
//	encPublishMsg := EncryptionMsg(publishMsg, 1234, 56789)
//	fmt.Println("encrypt publish message by publisher's private key")
//	printMsg(encPublishMsg)
//	fmt.Println(encPublishMsg)
//	reEncPublishMsg := sm.ReEncPubMsg(*encPublishMsg, "3.3.3.3")
//	fmt.Println("re-encrypt publish message by subscriber's private key")
//	printMsg(reEncPublishMsg)
//	//fmt.Println(reEncPublishMsg)
//	fmt.Println("decrypted publish message is...")
//	DecryptionMsg(reEncPublishMsg, 1234, 99999)
//}
