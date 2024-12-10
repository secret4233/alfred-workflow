package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow

	icon = &aw.Icon{
		Value: aw.IconClock.Value,
		Type:  aw.IconClock.Type,
	}
)

const (
	StrBackslash = "${backslash}"
)

func run() {

	var err error

	args := wf.Args()

	if len(args) == 0 {
		return
	}

	defer func() {
		if err == nil {
			wf.SendFeedback()
			return
		}
	}()

	// 处理 now
	input := strings.Join(args, " ")
	input = strings.Replace(input, "\\\\\\", StrBackslash, -1)
	input = strings.Replace(input, "\\", "", -1)
	input = strings.Replace(input, StrBackslash, "\\", -1)
	var ans bytes.Buffer
	err = json.Indent(&ans, []byte(input), "", "	")
	if err != nil {
		log.Println(err)
		fmt.Println(input)
	}

	jsonFormat := ans.String()
	wf.NewItem("格式化后的字符串").
		Subtitle("json format").
		Icon(icon).
		Arg(jsonFormat).
		Valid(true)
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
