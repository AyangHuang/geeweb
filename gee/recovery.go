package gee

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func trace(err string) string {

	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	var str strings.Builder
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.String(500, "服务器出错")
			}
		}()
		c.Next()
	}
}
