package main

import (
	"errors"
	"fmt"
	aw "github.com/deanishe/awgo"
	"regexp"
	"strconv"
	"strings"
	"time"
)
var (
	workflow *aw.Workflow

	icon = &aw.Icon{
		Value: aw.IconClock.Value,
		Type:  aw.IconClock.Type,
	}

	layouts = []string{
		"2006-01-02 15:04:05.999 MST",
		"2006-01-02 15:04:05.999 -0700",
		time.RFC3339,
		time.RFC3339Nano,
		time.UnixDate,
		time.RubyDate,
		time.RFC1123Z,
	}

	moreLayouts = []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999",
	}

	regexpTimestamp = regexp.MustCompile(`^[1-9]{1}\d+$`)
)

func run() {

	var err error

	args := workflow.Args()

	if len(args) == 0 {
		return
	}

	defer func() {
		if err == nil {
			workflow.SendFeedback()
			return
		}
	}()

	// 处理 now
	input := strings.Join(args, " ")
	if input == "now" {
		processNow()
		return
	}

	// 处理时间戳
	if regexpTimestamp.MatchString(input) {
		v, e := strconv.ParseInt(args[0], 10, 32)
		if e == nil {
			processTimestamp(time.Unix(v, 0))
			return
		}
		err = e
		return
	}

	// 处理时间字符串
	err = processTimeStr(input)
}


func processNow() {

	now := time.Now()

	// prepend unix timestamp
	secs := fmt.Sprintf("%d", now.Unix())
	workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(icon).
		Arg(secs).
		Valid(true)

	// process all time layouts
	processTimestamp(now)
}

// process all time layouts
func processTimestamp(timestamp time.Time) {
	for _, layout := range layouts {
		v := timestamp.Format(layout)
		workflow.NewItem(v).
			Subtitle(layout).
			Icon(icon).
			Arg(v).
			Valid(true)
	}
}

func processTimeStr(timestr string) error {

	timestamp := time.Time{}
	layoutMatch := ""

	layoutMatch, timestamp, ok := matchedLayout(layouts, timestr)
	if !ok {
		layoutMatch, timestamp, ok = matchedLayout(moreLayouts, timestr)
		if !ok {
			return errors.New("no matched time layout found")
		}
	}

	// prepend unix timestamp
	secs := fmt.Sprintf("%d", timestamp.Unix())
	workflow.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(icon).
		Arg(secs).
		Valid(true)

	// other time layouts
	for _, layout := range layouts {
		if layout == layoutMatch {
			continue
		}
		v := timestamp.Format(layout)
		workflow.NewItem(v).
			Subtitle(layout).
			Icon(icon).
			Arg(v).
			Valid(true)
	}

	return nil
}

func matchedLayout(layouts []string, timestr string) (matched string, timestamp time.Time, ok bool) {

	for _, layout := range layouts {
		v, err := time.Parse(layout, timestr)
		if err == nil {
			return layout, v, true
		}
	}
	return
}

func main() {
	workflow = aw.New()
	workflow.Run(run)
}
