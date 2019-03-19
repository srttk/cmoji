package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Cmd struct {
	token     string
	channelID string
}

func NewCmd(token, channelID string) Cmd {
	return Cmd{
		token:     token,
		channelID: channelID,
	}
}

func (c Cmd) ListEmoji() (map[string]string, error) {
	fmt.Println(c)
	return listEmoji(c.token)
}

func (c Cmd) StampEmoji(emoji string, emojiMap map[string]string) error {
	fmt.Println(c)
	text := fmt.Sprintf("stamp `%s`", emoji)
	imgURL := emojiMap[strings.Trim(emoji, ":")]
	a := newAttachment(text, imgURL, "#FFAACC")
	arg := newArgument(c.token, c.channelID, "")
	arg.setAttachments(a)

	return postMessage(c.token, arg)
}

func (c Cmd) SendEmojiMap(emojiMap map[string]string) error {
	fmt.Println(c)
	var keys []string
	for k := range emojiMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	buf := new(bytes.Buffer)
	for _, v := range keys {
		fmt.Fprintf(buf, "%s - :%s: | ", v, v)
	}
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	arg := newArgument(c.token, c.channelID, string(b))

	return postMessage(c.token, arg)
}
