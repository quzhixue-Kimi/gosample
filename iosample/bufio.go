package iosample

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func (l *LogInfo) FileReadOperator() {
	if file, err := os.Open("/tmp/whatisup.txt"); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		b2 := bufio.NewReader(file)
		//n := 0
		for {
			if data, err := b2.ReadString('\n'); err == io.EOF {
				break
			} else {
				//n++
				log := string(data)
				res := strings.Split(log, "|")
				fmt.Print(NewLogInfo(res[0], res[1]))
			}
		}
		//fmt.Println(n)
	}
}

func (l *LogInfo) FileCreateOperator() {
	if file, err := os.OpenFile("/tmp/whatisup.txt", os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		w1 := bufio.NewWriter(file)

		for i := 1; i <= 1000; i++ {
			_, err := w1.WriteString(l.String() + "\n")
			if err != nil {
				fmt.Printf("Error occurred: %v\n", err)
			}
		}
		w1.Flush()
	}
}
