# chromedpdemo

[原创分享 Golang 使用 Chromedp 绕过反爬抓取微信公众号文章](https://gocn.vip/topics/10991?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)

- chromedp爬取微信公众号

- 安装Chromedp

    ```
    go get -u github.com/chromedp/chromedp
    ```
- 执行程序

    ```
    go run main.go
    ```

- 问题及解决方法

    - 问题

        ```
        2021/03/06 10:14:07 [scrapeNewArticle] chromedp Run fail_2,err: context deadline exceeded
        ```
    - 解决方法
        ```
        请求超时时间太短，加长时间
        // 设置超时时间
	    ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
        ```
