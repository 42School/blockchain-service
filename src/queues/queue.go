package queues

import (
	"container/list"
	"github.com/42School/blockchain-service/src/global"
)

func PushRetryQueue(_e list.Element) {
	copyList := global.RetryQueue
	for e := copyList.Front(); e != nil; e = copyList.Front() {
		if e != nil {
			if e.Value == _e.Value {
				return
			}
		}
	}
	global.RetryQueue.PushBack(_e)
}
