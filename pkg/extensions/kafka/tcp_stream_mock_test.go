package kafka

import (
	"sync"

	"github.com/kubeshark/base/pkg/api"
)

type tcpStream struct {
	pcapId         string
	itemCount      int64
	identifyMode   bool
	emittable      bool
	isClosed       bool
	isTargetted    bool
	reqResMatchers []api.RequestResponseMatcher
	sync.Mutex
}

func NewTcpStream() api.TcpStream {
	return &tcpStream{}
}

func (t *tcpStream) SetProtocol(protocol *api.Protocol) {}

func (t *tcpStream) GetPcapId() string {
	return t.pcapId
}

func (t *tcpStream) GetIndex() int64 {
	return t.itemCount
}

func (t *tcpStream) GetIsIdentifyMode() bool {
	return t.identifyMode
}

func (t *tcpStream) IncrementItemCount() {
	t.itemCount++
}

func (t *tcpStream) SetAsEmittable() {
	t.emittable = true
}

func (t *tcpStream) GetReqResMatchers() []api.RequestResponseMatcher {
	return t.reqResMatchers
}

func (t *tcpStream) GetIsTargetted() bool {
	return t.isTargetted
}

func (t *tcpStream) GetIsClosed() bool {
	return t.isClosed
}
