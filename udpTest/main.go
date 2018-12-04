package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Data struct {
	Ip     string
	Header `json:"Header"`
	Data1  `json:"Data"`
}
type Header struct {
	MsgType string `json:"MsgType"`
	Device string `json:"Device"`
	FromMac string `json:"FromMac"`
	ToMac string `"json:"ToMac"`
}
type Data1 struct {

}

var x = false

func timer (){
	for{
		time.Sleep(time.Second * 5)
		if x==true{
			os.Exit(25656)
		}else{
			x = true
		}
	}
}

var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr

func main(){

	const url ="http://localhost:8080/test"
	OpenURL(url)
	receive := make(chan [] Data, 1000)
	send := make(chan Data)

	/*jsonFile, err := os.Open("package.json")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("json file opened")

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var a Data
	if err:= json.Unmarshal(byteValue, &a); err != nil{
		panic("bye bye")
	}*/

	a := Data{Ip: "",
		Header: Header{
			MsgType:"Ping",
			Device:"FSIP",
			FromMac:"b8:27:eb:ca:90:06",
			ToMac:"*",
		},
		}
	var x []Data
	
	go timer()
	go Broadcast(send)
	send <- a

	go Listen(receive)
	x = <-receive
	//fmt.Println("lolki", x)

	time.Sleep(time.Second*2)
	router := gin.New()
	/*t, err := LoadTemplate()
	if err != nil {
		panic(err)
	}*/
	//router.SetHTMLTemplate(t)
	router.Use(gin.Recovery())
	router.LoadHTMLFiles("test.tmpl")
	//router.LoadHTMLGlob("test.tmpl")
	router.GET("/test", func(c *gin.Context) {
		for i:= 0; i< len(x); i++{
			if x[i].Ip!="" {
				c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"title": "Info",
			"Device": x[i].Header.Device,
			"IP": x[i].Ip,
		})}
		}

	})
	router.GET("/ping", MyPingLogger())


	router.Run(":8080")

}

func MyPingLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		x = false
		c.JSON(200,gin.H{"data": "pong"})
		
	}

}

func OpenURL(url string) error {
	return openBrowser(url)
}
func openBrowser(url string) error {
	r := strings.NewReplacer("&", "^&")
	return runCmd("cmd", "/c", "start", r.Replace(url))
}
func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	setFlags(cmd)
	return cmd.Run()
}

func setFlags(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

