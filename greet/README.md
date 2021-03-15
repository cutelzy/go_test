### go-zero框架通过goctl实现api
- 安装
    ```
    GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero
    GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl
    ```
- 开始
    
    - 快速生成api服务
        ```
        goctl api new greet
        cd greet
        go run greet.go -f etc/greet-api.yaml
        ```
    - 发送请求
        ```
        curl -i http://localhost:8888/from/you
        ```
    - 返回结果
        ```
        HTTP/1.1 200 OK
        Date: Sun, 30 Aug 2020 15:32:35 GMT
        Content-Length: 0
        ```
- 可以根据api文件生成前端需要的Java, TypeScript, Dart, JavaScript代码
    ```
    goctl api java -api greet.api -dir greet
    goctl api dart -api greet.api -dir greet
    ...
    ```