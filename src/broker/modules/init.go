package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Moscato struct {
	queue               MsgQueue
	SendQueue           chan myType
	MicroServiceManager NodeManager `inject:""`
	MatchingManager     match_manager
	SubscriptionManager sub_manager
	SecureManager       SecurityManager `inject:""`
}

type myType struct {
	subList []string
	pubMsg  MsgUnit
	cosine  []float64
	err     error
}

type Reply struct { //RPC리턴값
	CompleteLog string //제대로 받았는지 확인하는 로그
}

type Receiver struct { //RPC 서버에 등록하기 위한 변수
	moscato *Moscato
}

type Args struct { // 매개변수
	JsonMsg []byte
	Kind    int
}

/*
MS→MM

-MM 실행되면 MM서버는 열려있음(MS은 자동으로 Client)(포트 8150)

-Send2MM 호출 → rpc call로 MmReceive호출해서 MM로 전달(json형식)

-MmReceive에서 msgType 검사 후 그것에 맞게 msgUnit으로 MM의 Receive로 보냄

-MM의 Receive에서 해당 Message를 처리
*/

func (receiver Receiver) MmReceive(args Args, reply *Reply) error { //
	// 메세지 별로 나눠서 언마샬
	switch args.Kind {

	case KSM:
		var msg KeyShareMsg
		err := json.Unmarshal(args.JsonMsg, &msg)
		if err != nil {
			return err
		}
		go func() {
			_, err := receiver.moscato.Receive(msg)
			if err != nil {

			}
		}()
		reply.CompleteLog = "KSM received"
	case PM:
		var msg PublishedImage
		//fmt.Println(args.JsonMsg)
		err := json.Unmarshal(args.JsonMsg, &msg)
		//fmt.Println(msg)
		if err != nil {
			return err
		}

		go func() {
			_, err := receiver.moscato.Receive(msg)
			if err != nil {

			}
		}()
		reply.CompleteLog = "PM received"
	case SM:
		var msg SubscriptionImage
		err := json.Unmarshal(args.JsonMsg, &msg)
		if err != nil {
			return err
		}
		go func() {
			_, err := receiver.moscato.Receive(msg)
			if err != nil {

			}
		}()
		reply.CompleteLog = "SM received"
	case RM:
		var msg RegisterMsg
		err := json.Unmarshal(args.JsonMsg, &msg)
		if err != nil {
			return err
		}
		go func() {
			_, err := receiver.moscato.Receive(msg)
			if err != nil {
				fmt.Println(err)
			}
		}()
		reply.CompleteLog = "RM received"
	case WM:
		var msg WithdrawMsg
		err := json.Unmarshal(args.JsonMsg, &msg)
		if err != nil {
			return err
		}
		go func() {
			_, err := receiver.moscato.Receive(msg)
			if err != nil {

			}
		}()
		reply.CompleteLog = "WM received"
	default:
		return errors.New("message type Error: Not registered message type")
	}
	//reply.CompleteLog = "received completely"
	return nil
}

//Recieve - MM가 MS로부터 메세지 전달받음
func (moscato *Moscato) Receive(msg MsgUnit) (MsgUnit, error) {
	logger := NewMyLogger()
	defer logger.Sync()

	//rpc call
	var msg_type = msg.CheckType()
	//메세지 타입에 따라 다르게 처리
	switch msg_type {

	case KSM: //Key share msg

	case PM: //Publish msg
		//log.Println("PM received")
		fromNodeName, _ := moscato.MicroServiceManager.GetIpaddr(msg.(PublishedImage).From)
		logger.Info("PM received from:[" + fromNodeName + "]")
		moscato.queue.push(moscato.preProcessMsg(msg))

	case SM: //Subscription msg
		//log.Println("SM received")
		fromNodeName, _ := moscato.MicroServiceManager.GetIpaddr(msg.(SubscriptionImage).From)
		logger.Info("SM received from:[" + fromNodeName + "]")
		err := moscato.SubscriptionManager.addSubscription(moscato.preProcessMsg(msg))
		if err != nil {
			logger.Warn(err.Error())
			//return nil, err
		}

	case RM: //Register msg
		var newMsg RegisterMsg
		newMsg = msg.(RegisterMsg)
		logger.Info("RM received from:[" + newMsg.From + "]")

		newNode := MSNode{newMsg.From, newMsg.From}
		resultAddNode := moscato.MicroServiceManager.AddMicroservice(newNode)
		if resultAddNode {
			logger.Info("Node added successful")
		} else {
			logger.Error("Node is already added, ignore RM")
			//log.Println("Node is already added, ignore RM")
			return msg, nil
		}

		addr, _ := moscato.MicroServiceManager.GetIpaddr(newMsg.From)

		moscato.SecureManager.RegKey(newMsg)
		logger.Debug("Registered microservice: address " + addr +
			" / key " + strconv.FormatUint(uint64(moscato.SecureManager.GetNodeKey(newMsg.From)), 10))

		// ackRM 메세지 전송
		go moscato.Send2MS(addr, newMsg)

	case WM: //Withdraw msg
		fromNodeName, _ := moscato.MicroServiceManager.GetIpaddr(msg.(WithdrawMsg).From)
		logger.Info("WM received from:[" + fromNodeName + "]")
		//ip := msg.(WithdrawMsg).From
		//sublist := moscato.SubscriptionManager.ip2sub[ip]
		//fmt.Println("prev list = ", sublist)
		moscato.MicroServiceManager.RemoveMicroservice(msg.(WithdrawMsg).From)
		moscato.SecureManager.RemoveSecureKey(msg.(WithdrawMsg).From)
		//moscato.SubscriptionManager.delete(ip)
		//sublist2 := moscato.SubscriptionManager.ip2sub[ip]
		//fmt.Println("after list =", sublist2)

	default:
		return nil, errors.New("Message type Error: Not registered message type")
	}

	return msg, nil
}

//MS로 보낼때 쓸 함수

/*
MM→MS

-MS 실행되면 MS서버는 열려있음(MM은 자동으로 Client)(포트 8160)

-Send2MS 호출 → rpc call로 MsReceive호출해서 MS로 전달(json형식)

-MsReceive에서 msgType 검사 후 그것에 맞게 msgUnit으로 MS의 Receive로 보냄

-MS의 Receive에서 해당 Message를 처리
*/

func (moscato *Moscato) Send2MS(ipaddress string, msg MsgUnit) {
	logger := NewMyLogger()
	defer logger.Sync()

	client, err := rpc.Dial("tcp", ipaddress+":8150")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	reply := new(Reply)
	jmsg, _ := msg.ConvertToJson()
	args := Args{
		JsonMsg: jmsg,
		Kind:    msg.CheckType(),
	}
	err = client.Call("Receiver.Receive", args, reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	//log.Println(reply.CompleteLog) //잘 받았는지 확인 해줌
	// 마이크로 서비스에게 받은 메시지는 노란색으로 출력
	//log.Println(reply.CompleteLog)
	logger.Debug("HERE" + reply.CompleteLog)
}

//Matching을 용이하게 하기위한 메세지 가공 과정
func (moscato *Moscato) preProcessMsg(originalMsg MsgUnit) MsgUnit {
	if originalMsg.CheckType() == PM {
		pubMsg := originalMsg.(PublishedImage)
		for index := 0; index < len(pubMsg.Topic); index++ {
			pubMsg.Topic[index] = pubMsg.Topic[index] - moscato.SecureManager.GetImgNodeKey(pubMsg.From)
		}
		//for index := 0; index < len(pubMsg.Value); index++ {
		//	pubMsg.Value[index] = pubMsg.Value[index] - moscato.SecureManager.GetNodeKey(pubMsg.From)
		//}
		return pubMsg
	} else if originalMsg.CheckType() == SM {
		subMsg := originalMsg.(SubscriptionImage)
		for index := 0; index < len(subMsg.Topic); index++ {
			subMsg.Topic[index] = subMsg.Topic[index] - moscato.SecureManager.GetImgNodeKey(subMsg.From)
		}
		//for index := 0; index < len(subMsg.Value); index++ {
		//	subMsg.Value[index] = subMsg.Value[index] - moscato.SecureManager.GetNodeKey(subMsg.From)
		//}
		return subMsg
	}
	return nil
}

//암호화 해서 보내기
func (moscato *Moscato) SendWithEncrypt() MsgUnit {
	for {
		mt := <-moscato.SendQueue
		if mt.err == nil {
			for index := 0; index < len(mt.subList); index++ { //sublist들을 돌면서 매세지를 encrypt하여 메세지 보냄
				tmpNode := mt.subList[index]
				tmpNodeIpAddr, _ := moscato.MicroServiceManager.GetIpaddr(tmpNode)
				tmp := mt.cosine[index]
				var cosineSim []float64
				cosineSim = append(cosineSim, tmp)
				//moscato.SecureManager.ReEncPubMsg(mt.pubMsg.(PublishMsg), tmpNode)
				//fmt.Println("publish: ", mt.pubMsg)
				//moscato.Send2MS(tmpNodeIpAddr, mt.pubMsg)

				msg := sendPublishedImage{moscato.SecureManager.ReEncImgPubMsg(mt.pubMsg.(PublishedImage), tmpNode), cosineSim}
				moscato.Send2MS(tmpNodeIpAddr, msg)
				//moscato.Send2MS(tmpNodeIpAddr, moscato.SecureManager.ReEncImgPubMsg(mt.pubMsg.(PublishedImage), tmpNode), cosineSim)

				//moscato.Send2MS(tmpNodeIpAddr, mt.pubMsg.(PublishedImage))
			}

		}
		return nil
	}
}

func (moscato *Moscato) Run() {
	logger := NewMyLogger()
	defer logger.Sync()

	config := AppConfig{moscato}
	config.config()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		//withDraw(message, client)
		//fmt.Println(sig)
		_ = sig
		done <- true
		logger.Info("terminate Moscato Message Middleware")
		os.Exit(0)
	}()

	//모스카토 구조체 변수 초기화
	receiver := Receiver{moscato: moscato}
	err := moscato.queue.queue_init()
	moscato.SendQueue = make(chan myType)
	moscato.SubscriptionManager.Initialize()

	if err != nil {
		fmt.Println(err)
		return
	}

	send2MSFile, _ := os.Create("./send2MsTime.log")
	defer send2MSFile.Close()

	//go routine -> matching 동작
	go func() {
		for {
			msg := moscato.queue.pop(true)
			fmt.Println(FloatSlice2String(msg.(PublishedImage).Topic))
			encStartTime := time.Now()
			go moscato.ImageMatching(msg)
			go moscato.SendWithEncrypt()
			encElapsedTime := time.Since(encStartTime)
			fmt.Printf("MM enc시간: %d\n\n", encElapsedTime.Nanoseconds())
			fmt.Fprintln(send2MSFile, encElapsedTime.Nanoseconds())
		}
	}()

	//go moscato.CheckQueue()

	//rpc 등록 -> Receive 함수
	err = rpc.Register(receiver)
	if err != nil {
		println(err)
		return
	}

	go Listen()
	logger.Info("initializing complete")

	<-done
}

func Listen() {
	logger := NewMyLogger()
	defer logger.Sync()
	/*
		MS→MM일때 ⇒ port : 8160으로 열기

		(MM이 Server, MS가 Client)
	*/

	l, err1 := net.Listen("tcp", ":8160")

	if err1 != nil {
		logger.Fatal("Unable to listen on given port: " + err1.Error())
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
