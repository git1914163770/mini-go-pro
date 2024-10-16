package web

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 获取触发panic的堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // 跳过3层caller
	var str strings.Builder
	str.WriteString(message + "\n\n")

	for i := 0; i < n; i++ {
		pc := pcs[i]
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n%s:%d %s", file, line, fn.Name()))
	}

	return str.String()

}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				http.Error(c.Writer, message, http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
