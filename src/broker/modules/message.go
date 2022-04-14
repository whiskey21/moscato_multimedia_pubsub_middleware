package modules

import "encoding/json"

//*****메세지 타입 상수화
const (
	KGM = 1 + iota //KeyGenMessage
	KSM            //KeyShareMessage
	PM             //PublishMessage
	SM             //SubscriptionMessage
	RM             //RegisterMessage
	WM             //WithdrawMessage
)

//*****메세지 틀*****
type Message struct {
	From    string //메세지 만든 주체의 Ip주소
	Version string //메세지 버전
	Time    string //메세지 만든 시간
	Kind    int    //메세지 종류
}

type MsgUnit interface {
	ConvertToJson() ([]byte, error) //메세지를 JSON형식으로 바꿔주는 멤버함수
	CheckType() int                 //메세지의 종류를 반환하는 함수
}

//*****각 메세지 형식 및 정의**********

type sendPublishedImage struct {
	PublishedImage
	Cosine []float64
}

//KeyGen 명령 메세지
type PublishedImage struct {
	Message
	Topic []float64
}

type KeyGenMsg struct {
	Message
	iptable []string
}

//Key공유 메세지
type KeyShareMsg struct {
	Message
	key string
}

//전달할 내용을 담은 메세지
type PublishMsg struct {
	Message
	Topic   []int64 //대주제	ex)soccer
	Value   []int64 //topic의 세부적인 내용 ex)ManCity or 40
	Content []int64 //내용 ex)오늘 케빈 데 브라위너가 골을 넣었다
}

type SubscriptionImage struct {
	Message
	Topic []float64
}

//구독 정보를 담은 메세지
type SubscriptionMsg struct {
	Message
	Topic    []int64  //대주제 ex)soccer
	Value    []int64  //피연산자 ex)Mancity or 20 40
	Operator []string //연산자	ex) ==, < && >
	IsAlpha  bool     //value가 숫자인지 문자열인지, 문자열이면 단순비교, 숫자이면 범위연산
}

type RegisterImgMsg struct {
	Message
	PrivateKey float64
}

//Microservice 등록 메세지
type RegisterMsg struct {
	Message
	PrivateKey int64
}

//Microservice 탈퇴 메세지(없앰)
type WithdrawMsg struct {
	Message
}

func (msg SubscriptionImage) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

//******ConverToJson을 메세지 종류별로 실행가능하게 구현******
func (msg KeyGenMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg KeyShareMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg sendPublishedImage) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg PublishedImage) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg PublishMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg SubscriptionMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg RegisterImgMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg RegisterMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

func (msg WithdrawMsg) ConvertToJson() ([]byte, error) {
	js := msg
	jsonBytes, err := json.Marshal(js)
	return jsonBytes, err
}

//CheckType함수 구현
func (msg Message) CheckType() int {
	return msg.Kind //메세지 멤버변수 Kind 리턴
}

//KeyGenMsg 생성자
func NewKeyGenMsg(table *MStable) *KeyGenMsg {
	m := &KeyGenMsg{}
	for _, value := range table.NodeTable { // MicroService테이블에서 ip주소를 다 가져와서 iptable에 넣음
		m.iptable = append(m.iptable, value.GetIpaddr())
	}
	return m
}
