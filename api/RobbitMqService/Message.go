package RobbitMqService

import (
	"SecKill/model"
	"encoding/json"
	"github.com/prometheus/common/log"
)

// 商品下单用于消息队列传输的消息
type Message struct {
	Username string
	Coupon   model.Coupon
}

func NewMessage(username string, coupon model.Coupon) *Message {
	return &Message{Username: username, Coupon: coupon}
}

func (m *Message) JsonToStr() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Error("json 解析出错：" + err.Error())

	}

	return string(bytes)
}

func (m *Message) StrToJson(dataStr []byte) *Message {
	if err := json.Unmarshal(dataStr, m); err != nil {
		log.Error("json 转换出错：" + err.Error())
	}
	return m
}
