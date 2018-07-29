package msg

type Response struct {
	seq int // 消息的标号，用于配对请求和应答，自增
	msg []byte
}

type Request struct {
	seq int
	msg []byte
}

type Session struct {
	seq      int
	Request  Request
	Response Response
	success  func()
	fail     func()
}

type Manager struct {
	seq      int
	Sessions map[int]Session

	MsgIn  chan interface{}
	MsgOut chan interface{}
}

func (m *Manager) SeqInc() {
	m.seq += 1
}

func (m *Manager) Send() {
	// 不需要回调
}

func (m *Manager) Query() {
	// 请求的时候，可以设置回调，等消息回来的时候，调用回调
}

func (m *Manager) OnMsg() {
	// 检查seq，如果存在，就调用相应的回调
	// 检查seq，如果不存在，就返回相同seq的request
}
