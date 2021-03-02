package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	_ "github.com/quzhixue-Kimi/gosample/mysqlsample"
)

type People struct {
	Name string
	Age  int
}

func (p *People) String() string {
	return fmt.Sprintf("[People]name=%v,age=%v", p.Name, p.Age)
}

func (p *People) sayGoodBye() string {
	return fmt.Sprintf("%v says goodbye to u!", p.Name)
}

func (p People) sayHello(name string) {
	fmt.Println(name + " says hello to u!")
}

type Latencies struct {
	Request int `json:"request"`
	Kong    int `json:"kong"`
	Proxy   int `json:"proxy"`
}

func (l Latencies) String() string {
	return fmt.Sprintf("[lantencies]request=%v,kong=%v,proxy=%v", l.Request, l.Kong, l.Proxy)
}

type Service struct {
	Host           string `json:"host"`
	CreatedAt      int64  `json:"created_at"`
	ConnectTimeout int    `json:"connect_timeout"`
	Id             string `json:"id"`
	Protocol       string `json:"protocol"`
	Name           string `json:"name"`
	ReadTimeout    int    `json:"read_timeout"`
	Port           int    `json:"port"`
}

func (s Service) String() string {
	return fmt.Sprintf("[service]host=%v,id=%v,name=%v,port=%v,protocol=%v", s.Host, s.Id, s.Name, s.Port, s.Protocol)
}

type Headers struct {
	Host           string `json:"host"`
	AcceptEncoding string `json:"accept-encoding"`
	UserAgent      string `json:"user-agent"`
	Accept         string `json:"accept"`
	Connection     string `json:"connection"`
}

func (h Headers) String() string {
	return fmt.Sprintf("[headers]host=%v,accept=%v,connection=%v", h.Host, h.Accept, h.Connection)
}

type Request struct {
	Headers `json:"headers"`
	Method  string `json:"method"`
	Uri     string `json:"uri"`
	Url     string `json:"url"`
	Size    int    `json:"size"`
}

func (r Request) String() string {
	return fmt.Sprintf("[request]headers=%v,method=%v,uri=%v,url=%v,size=%v", r.Headers, r.Method, r.Uri, r.Url, r.Size)
}

type Try struct {
	BalanceLatency int    `json:"balance_latency"`
	Port           int    `json:"port"`
	BalanceStart   int    `json:"balance_start"`
	Ip             string `json:"ip"`
}

func (t Try) String() string {
	return fmt.Sprintf("[tries]balance_latency=%v,ip=%v", t.BalanceLatency, t.Ip)
}

type Result struct {
	Latencies `json:"latencies"`
	Service   `json:"service"`
	Request   `json:"request"`
	ClientIp  string `json:"client_ip"`
	Tries     []Try  `json:"tries"`
}

//func sendData(ch1 chan int) {
//	for i := 0; i < 10; i++ {
//		ch1 <- i
//		time.Sleep(500 * time.Millisecond)
//		fmt.Println("写入数据 ", i)
//	}
//	close(ch1)
//}
//
//func readData(ch1 chan int, exit chan bool) {
//	for {
//		v, ok := <-ch1
//		if !ok {
//			fmt.Println("已经读取了所有的数据，", ok)
//			break
//		}
//		time.Sleep(700 * time.Millisecond)
//		fmt.Println("取出数据：", v)
//	}
//	exit <- true
//	close(exit)
//}

func main() {

	p := &People{
		Name: "Kimi",
		Age:  37,
	}
	fmt.Println(p)
	p.sayHello(p.Name)
	fmt.Println(p.sayGoodBye())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/log", func(c *gin.Context) {
		var result Result
		if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
			fmt.Println(err)
			return
		} else {
			if err = json.Unmarshal(body, &result); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(result.Latencies)
				fmt.Println(result.Service)
				fmt.Println(result.Tries)
				fmt.Println(result.Request)
				fmt.Println(result.Headers)
			}
		}
	})
	r.POST("/log1", func(c *gin.Context) {
		if body, err := httputil.DumpRequest(c.Request, true); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(body))
		}
	})
	if err := r.Run(":3000"); err != nil {
		fmt.Println(err)
		return
	}

	//	mysqlsample.MySqlSample()
	//  mysqlsample.MySqlSample1()
	//jobs := make(chan int, 5)
	//results := make(chan int, 5)

	//for i := 1; i <= 2; i++ {
	//	go workpool.WorkerPool(i, jobs, results)
	//}

	//for i := 1; i <= 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)

	////for v := range results {
	////	fmt.Println(v)
	////}

	////for {
	////	if v, ok := <-results; !ok {
	////		break
	////	} else {
	////		fmt.Println(v)
	////	}
	////}

	//for i := 1; i <= 5; i++ {
	//	x := <-results
	//	fmt.Println(x)
	//}

	//	selectsample.SelectChannel()

	//msg := make(chan *io.LogInfo, 5)
	//exit := make(chan bool, 1)

	//l := io.NewLogInfo("hello", "Info")

	//for {
	//	go l.FileReadOperator(msg)
	//	go l.FileCreateOperator(msg, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}

	//ch1 := make(chan int, 5)
	//exit := make(chan bool, 1)
	//for {
	//	go sendData(ch1)
	//	go readData(ch1, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}
}
