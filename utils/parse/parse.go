package parse

import "time"

func SetTime() *time.Time {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, time.Now().Format(layout))
	return &t
}
