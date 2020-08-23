// 每一个可执行的go语言程序必定具备一个main包，并在该main包下具备可执行函数main的go文件
package main
// HelloGo.go
// 基于图灵API一个简单的聊天机器人
// 引入相关依赖
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)
// 请求体结构体
type requestBody struct {
	Key string `json:"key"`
	Info string	`json:"info"`
	UserId string	`json:"userid"` 
}

// 结果体结构体
type responseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string	`json:"list"`
	Url string	`json:"url"`
}
// 请求机器人
func process(inputChan <-chan string, userid string) {
	for {
		// 从甬道中接收输入
		input := <- inputChan
		if input == "EOF"{
			break
		}
		// 构造请求体
		reqData := &requestBody{
			Key: "792bcf45156d488c92e9d11da494b085",
			Info: input,
			UserId: userid,
		}
		// 转义为json
		byteData, _ := json.Marshal(&reqData)
		// 请求聊天机器人接口
		req, err := http.NewRequest("POST",
			"http://www.tuling123.com/openapi/api",
			bytes.NewReader(byteData))
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Network Error!")
		}else {
			// 将结果从json中解析并输出到命令行
			body, _ := ioutil.ReadAll(resp.Body)
			var respData responseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI: " + respData.Text)

		}
		resp.Body.Close()
	}
}

func main() {
	var input string
	fmt.Println("Enter 'EOF' to shut down: ")
	// 创建通道
	channel := make(chan string)
	// main 结束时关闭通道
	defer close(channel)
	// 启动 goroutine 运行机器人回答线程
	go process(channel, string(rand.Int63()))
	for {
		// 从命令行中读取输入
		fmt.Scanf("%s", &input)
		// 将输出放到通道中
		channel <- input
		// 结束程序
		if input == "EOF"{
			fmt.Println("Bye!")
			break
		}
	}
}