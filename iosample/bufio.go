package iosample

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type LogInfo struct {
	msg   string
	level string
}

func NewLogInfo(msg, level string) *LogInfo {
	return &LogInfo{
		msg:   msg,
		level: level,
	}
}

func (l *LogInfo) String() string {
	return fmt.Sprintf(l.msg + "|" + l.level)
}

func (l *LogInfo) FileReadOperator(msg chan string, exit chan bool) {
	if file, err := os.Open("/tmp/whatisup.txt"); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		//b2 := bufio.NewReader(file)
		//n := 0
		//for {
		//	if data, err := b2.ReadString('\n'); err == io.EOF {
		//		break
		//	} else {
		//		//n++
		//		time.Sleep(time.Millisecond * 500)
		//		log := string(data)
		//		res := strings.Split(log, "|")
		//		fmt.Print(NewLogInfo(res[0], res[1]))
		//	}
		//}
		for {
			if v, ok := <-msg; !ok {
				break
			} else {
				time.Sleep(time.Millisecond * 300)
				fmt.Println(v)
			}
		}
		exit <- true
		close(exit)
		//fmt.Println(n)
	}
}

func (l *LogInfo) FileCreateOperator(msg chan string) {
	if file, err := os.OpenFile("/tmp/whatisup.txt", os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		w1 := bufio.NewWriter(file)

		for i := 1; i <= 50; i++ {
			time.Sleep(time.Millisecond * 200)
			message := fmt.Sprintf(l.String() + "\n")
			msg <- message
			_, err := w1.WriteString(message)
			if err != nil {
				fmt.Printf("Error occurred: %v\n", err)
			}
		}
		w1.Flush()
		close(msg)
	}
}
