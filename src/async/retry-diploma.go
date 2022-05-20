package async

import (
	"github.com/42School/blockchain-service/src/dao/diplomas"
	"github.com/42School/blockchain-service/src/metrics"
	"github.com/42School/blockchain-service/src/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func RetryDiploma() {
	url := tools.FtEndPoint + tools.RetryPath
	for {
		time.Sleep(1 * time.Minute)
		//time.Sleep(30 * time.Minute)
		copyList := tools.RetryQueue
		for e := copyList.Front(); e != nil; {
			if e != nil {
				diploma, _ := e.Value.(diplomas.DiplomaImpl)
				var dp diplomas.Diploma = diploma
				log.WithFields(dp.LogFields()).Debug("Try to retry a diploma")
				hash, bool := dp.EthWriting()
				if bool == true {
					data := "{'Status':true,'Message':'The writing in blockchain has been done, it will be confirmed in 10 min.','Data':{'Hash': " + hash + ",'Level':0,'Skills':[]}}"
					http.Post(url, "Content-Type: application/json", strings.NewReader(data))
					tools.RetryQueue.Remove(e)
					tools.Db.DeleteOneRetry(diploma)
					metrics.GaugeRetryQueue.Dec()
					e = copyList.Front()
				} else {
					diploma.Counter += 1
					log.Info(diploma.Counter)
					tools.RetryQueue.InsertBefore(diploma, e)
					tools.RetryQueue.Remove(e)
					e = e.Next()
				}
			}
		}
	}
}
