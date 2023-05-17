package async

import (
	"fmt"
	"time"
)

func job(w int32) string {
	d := time.Duration(w) * time.Second
	time.Sleep(d)
	return fmt.Sprintf("%s: done", time.Now())
}
