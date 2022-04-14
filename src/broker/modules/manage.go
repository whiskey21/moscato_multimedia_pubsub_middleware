package modules

import "sync"

//각 Microservice에 대한 정보 저장 노드
type MSNode struct {
	nodeName string //Nodename- 현재 데모에서는 IpAddress와 같음
	ipAddr   string
}

//Nodename 반환
func (node *MSNode) GetName() string {
	return node.nodeName
}

//IpAddress 반환
func (node *MSNode) GetIpaddr() string {
	return node.ipAddr
}

type NodeManager interface {
	GetIpaddr(nodeName string) (string, bool) //IpAddress반환
	AddMicroservice(node MSNode) bool         //MS추가
	RemoveMicroservice(nodeName string) bool  //MS삭제
}

//모든 Microservice정보 저장
type MStable struct {
	NodeTable map[string]MSNode
	mu        sync.RWMutex
}

//MStable 생성자
func NewMStable() *MStable {
	logger := NewMyLogger()
	defer logger.Sync()

	defer logger.Debug("node manager setting complete.")
	return &MStable{NodeTable: make(map[string]MSNode)}
}

//IpAddress반환
func (manager *MStable) GetIpaddr(nodeName string) (string, bool) {
	//해당 이름의 노드이름이 존재하는지 확인
	node, exists := manager.NodeTable[nodeName]

	//존재하지 않는 경우 nil리턴
	if !exists {
		return "", false
	} else {
		return node.ipAddr, true
	}
}

//MS추가
func (manager *MStable) AddMicroservice(node MSNode) bool {
	//삽입 전 존재여부 확인
	manager.mu.RLock()
	_, exists := manager.NodeTable[node.GetName()] ////해당 Node의 이름이 있는지 검색
	//manager.mu.Unlock()
	manager.mu.RUnlock()

	if exists {
		return false
	} else { //존재안한다면 추가 (존재할경우 이미 있는것이기 때문에 추가할 필요 없음)
		//manager.mu.Lock()
		manager.NodeTable[node.GetName()] = node
		//manager.mu.Unlock()
		return true
	}
}

//MS삭제
func (manager *MStable) RemoveMicroservice(nodeName string) bool {
	logger := NewMyLogger()
	defer logger.Sync()
	//삭제 전 존재여부 확인
	_, exists := manager.NodeTable[nodeName] //해당 이름을 가진 Node가 있는지 검색

	if exists {
		delete(manager.NodeTable, nodeName) //존재한다면 삭제
		logger.Info("[" + nodeName + "] : service quit")
		return true
	} else {
		return false
	}
}
