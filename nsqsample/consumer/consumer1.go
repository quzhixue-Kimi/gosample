package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type myMessageHandler struct {
}

func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	log.Printf("msg.Timestamp=%v,msg.nsqaddress=%s,msg.body=%s \n", time.Unix(0, m.Timestamp).Format("2006-01-02 03:04:05"), m.NSQDAddress, string(m.Body))

	return nil
}

func main() {
	config := nsq.NewConfig()
	if consumer, err := nsq.NewConsumer("whatisup", "showcase", config); err != nil {
		log.Fatalf("error occurr %v\n", err)
		return
	} else {
		consumer.AddHandler(&myMessageHandler{})
		if err = consumer.ConnectToNSQLookupd("10.68.180.137:4161"); err != nil {
			log.Fatalf("error occurr %v\n", err)
			return
		}
		//consumer.Stop()
	}
	for {
		time.Sleep(time.Second)
	}
}
