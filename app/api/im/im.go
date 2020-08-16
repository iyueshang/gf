package im

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
)

var (
	users = gmap.New(true)
	names = gset.NewStrSet(true)
	cache = gcache.New()
	conn  *ghttp.WebSocket
)

const (
	SendInterval = time.Second
)

type Msg struct {
	Type string      `json:"type" v:"type@required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	Form string      `json:"form" v:""`
}

func Im(r *ghttp.Request) {
	fmt.Println(r.Get("token"))
	msg := &Msg{}
	ws, err := r.WebSocket()
	conn = ws
	if err != nil {
		g.Log().Error(err)
		return
	}

	for {
		msgType, msgByte, err := ws.ReadMessage()

		switch msgType {
		case websocket.TextMessage:
			fmt.Println("text")
		case websocket.BinaryMessage:
		case websocket.CloseMessage:
		case websocket.PingMessage:
		case websocket.PongMessage:
		}

		fmt.Println(msgType)
		fmt.Println(msgByte)
		if err != nil {
			fmt.Println(err)
			g.Log().Error(err)
			return
		}
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			_ = write(Msg{
				Type: "error",
				Data: "消息格式不正确:" + err.Error(),
				Form: "",
			})
			continue
		}
		if e := gvalid.CheckStruct(msg, nil); e != nil {
			_ = write(Msg{
				Type: "error",
				Data: e.String(),
				Form: "",
			})
			continue
		}
		g.Log().Cat("chat").Println(msg)

		switch msg.Type {
		case "send":
			intervalKey := fmt.Sprintf("%p", ws)
			if !cache.SetIfNotExist(intervalKey, struct{}{}, SendInterval) {
				_ = write(Msg{Type: "error", Data: "您的消息发送得过于频繁，请休息下再重试", Form: ""})
				continue
			}
			// 有消息时，群发消息
			if msg.Data != nil {
				if err = writeGroup(Msg{"send", ghtml.SpecialChars(gconv.String(msg.Data)), ghtml.SpecialChars("test")}); err != nil {
					g.Log().Error(err)
				}
			}

		}
	}

}

func write(msg Msg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return conn.WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
}

func writeGroup(msg Msg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}

	users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
		}
	})
	return nil
}
