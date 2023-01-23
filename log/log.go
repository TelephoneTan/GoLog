package log

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	tz     *time.Location
	lock   sync.Mutex
	sb     = &strings.Builder{}
	red    = color.New(color.FgRed)
	yellow = color.New(color.FgYellow)
	green  = color.New(color.FgGreen)
	white  = color.New(color.FgWhite)
	e      = red
	w      = yellow
	s      = green
	i      = white
	stdout = color.Output
	stderr = color.Error
)

func init() {
	var err error
	tz, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("无法初始化时区：", err)
	}
}

func t(sb *strings.Builder) {
	sb.WriteString(time.Now().In(tz).Format("2006/01/02 15:04:05.000000000 "))
}

func sp(sb *strings.Builder, format string, a ...interface{}) {
	if format == "" {
		_, _ = fmt.Fprint(sb, a...)
	} else {
		_, _ = fmt.Fprintf(sb, format, a...)
	}
}

func p(w io.Writer, colour *color.Color, format string, a ...interface{}) {
	lock.Lock()
	defer lock.Unlock()
	t(sb)
	sp(sb, format, a...)
	msg := sb.String()
	if !strings.HasSuffix(msg, "\r") && !strings.HasSuffix(msg, "\n") {
		sb.WriteRune('\n')
		msg = sb.String()
	}
	_, _ = colour.Fprint(w, msg)
	sb.Reset()
}

func E(a ...interface{}) {
	p(stderr, e, "", a...)
}

func EF(format string, a ...interface{}) {
	p(stderr, e, format, a...)
}

func W(a ...interface{}) {
	p(stderr, w, "", a...)
}

func WF(format string, a ...interface{}) {
	p(stderr, w, format, a...)
}

func I(a ...interface{}) {
	p(stdout, i, "", a...)
}

func IF(format string, a ...interface{}) {
	p(stdout, i, format, a...)
}

func S(a ...interface{}) {
	p(stdout, s, "", a...)
}

func SF(format string, a ...interface{}) {
	p(stdout, s, format, a...)
}

func F(a ...interface{}) {
	E(a...)
	os.Exit(1)
}

func FF(format string, a ...interface{}) {
	EF(format, a...)
	os.Exit(1)
}
