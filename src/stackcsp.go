package lab4

type stackAction int

const (
	length stackAction = iota
	push
	pop
)

type stackCommand struct {
	action stackAction
	val    interface{}
	result chan<- interface{}
}

type CspStack struct {
	top     *Element
	size    int
	cmdChan chan stackCommand
}

func NewCspStack() *CspStack {
	s := &CspStack{cmdChan: make(chan stackCommand)}
	go s.run()
	return s
}

func (cs *CspStack) Len() int {
	reply := make(chan interface{})
	cs.cmdChan <- stackCommand{action: length, result: reply}
	return (<-reply).(int)
}

func (cs *CspStack) Push(value interface{}) {
	cs.cmdChan <- stackCommand{action: push, val: value}
}

func (cs *CspStack) Pop() (value interface{}) {
	reply := make(chan interface{})
	cs.cmdChan <- stackCommand{action: pop, result: reply}
	return <-reply
}

func (cs *CspStack) run() {
	for cmd := range cs.cmdChan {
		switch cmd.action {
		case length:
			cmd.result <- cs.size
		case push:
			cs.top = &Element{cmd.val, cs.top}
			cs.size++
		case pop:
			if cs.size > 0 {
				value := cs.top.value
				cs.top = cs.top.next
				cs.size--
				cmd.result <- value
			} else {
				cmd.result <- nil
			}
		}
	}
}
