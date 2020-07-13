package queue

import (
	list "container/list"
	"log"
	"os"
)

const nameFileRetryQueue string = "RetryQueue.save.dev"

type Queue struct {
	List	*list.List
}

func New() *Queue {
	log.Println("Enter in New Queue")
	q := Queue{List: list.New()}
	_, err := os.OpenFile(nameFileRetryQueue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)
	if err != nil {
		log.Println(err)
	}
	return &q
}

func (q *Queue) CustomPushBack(_dp int) {
	q.List.PushBack(_dp)
	f, _ := os.OpenFile(nameFileRetryQueue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)
	f.Write([]byte("PushBack\n"))
}

func (q *Queue) CustomDelete(_dp *list.Element) {
	q.List.Remove(_dp)
	f, _ := os.OpenFile(nameFileRetryQueue, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)
	f.Write([]byte("Deleted\n"))
}