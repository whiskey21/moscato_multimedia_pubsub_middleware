package modules

import (
	"errors"
)

type MsgQueue struct {
	queue     chan MsgUnit
	QueueFunc QueueOperate
}

type QueueOperate interface {
	queue_init() error              //Queue 초기화 멤버함수
	push(msg Message) bool          //Queue push 멤버함수
	pop(wait bool) (Message, error) //Queue pop 멤버함수 (wait -> busy waiting 여부 결정)
}

//Message Queue를 초기화 해주는 함수
func (mq *MsgQueue) queue_init() error {
	logger := NewMyLogger()
	logger.Sync()

	if mq.queue != nil && len(mq.queue) != 0 {
		return errors.New("Queue Hadlerer Error: Already initialized.")
	} else if mq.queue == nil {
		mq.queue = make(chan MsgUnit, 1000)
		//log.Println("queue is initialized.")
		logger.Debug("queue is initialized")
		return nil
	} else {
		close(mq.queue)
		mq.queue = make(chan MsgUnit)
		return nil
	}
}

//push 함수
func (mq *MsgQueue) push(msg MsgUnit) bool {
	mq.queue <- msg
	return true
}

//pop 함수
func (mq *MsgQueue) pop(block bool) MsgUnit {
	if block == false {
		if len(mq.queue) == 0 {
			return nil
		} else {
			return <-mq.queue
		}
	} else {
		//queue에 데이터가 들어올 때까지 block
		return <-mq.queue
	}
}
