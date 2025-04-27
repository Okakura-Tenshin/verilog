// middlewares/log_full_request.go
package middlewares

import (
    "bytes"
    "io/ioutil"
    "log"
    "net/http/httputil"

    "github.com/gin-gonic/gin"
)

func LogFullRequest() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 读出原始 body
        rawBody, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Printf("[LogFullRequest] read body error: %v\n", err)
            // 继续让请求执行，但 body 可能为空
            c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
            c.Next()
            return
        }
        // 2. 恢复 body，保证后续 Bind/Handler 能正常读取
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))

        // 3. Dump 整个 HTTP 请求，包括请求行、headers 以及 body
        dump, err := httputil.DumpRequest(c.Request, true)
        if err != nil {
            log.Printf("[LogFullRequest] dump request error: %v\n", err)
        } else {
            log.Printf("[LogFullRequest]\n%s\n", string(dump))
        }

        // 4. 再次恢复 body，以防后续业务代码再读一次
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))

        // 5. 继续处理
        c.Next()
    }
}
