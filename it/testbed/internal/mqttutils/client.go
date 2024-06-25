package mqttutils

import (
	"fmt"
	"time"

	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/types/pair"
	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/types/syncmap"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type QoS byte

const (
	MQTT_QOS_LOW    QoS = 0
	MQTT_QOS_MEDIUM QoS = 1
	MQTT_QOS_HIGH   QoS = 2
)

type Credentials pair.Pair[string, string]

type Client interface {
	Connect() error
	Unsubscribe(topic string) error
	Disconnect(waitfor ...time.Duration)
	Publish(topic string, payload any, retained ...bool) error
	Subscribe(topic string, callback mqtt.MessageHandler) error
	SubscribeOnce(topic string, callback mqtt.MessageHandler) error
}

func NewClient(mqtt mqtt.Client, qos QoS) Client {
	return &_Client{
		qos:           qos,
		delegate:      mqtt,
		subscriptions: syncmap.New[string, bool](),
	}
}

type _Client struct {
	qos           QoS
	delegate      mqtt.Client
	subscriptions syncmap.SyncMap[string, bool]
}

func (self *_Client) Connect() error {
	return self.await(self.delegate.Connect())
}

func (self *_Client) Disconnect(waitfor ...time.Duration) {
	waitfor = append(waitfor, time.Millisecond)

	self.delegate.Disconnect(uint(max(time.Millisecond, waitfor[0]).Milliseconds()))
}

func (self *_Client) Unsubscribe(topic string) (err error) {
	self.subscriptions.Update(topic, func(subscribed bool, found bool) (bool, bool) {
		err = self.await(self.delegate.Unsubscribe(topic))
		return false, false // returning 'false' as second result deletes the record
	})

	return
}

func (self *_Client) Subscribe(topic string, callback mqtt.MessageHandler) (err error) {
	self.subscriptions.Update(topic, func(subscribed bool, found bool) (bool, bool) {
		if subscribed {
			err = fmt.Errorf("already subscribed to topic %s", topic)
		} else {
			err = self.await(self.delegate.Subscribe(topic, byte(self.qos), callback))
		}

		return true, true // returning 'true' as second result sets the record to the specified value
	})

	return
}

func (self *_Client) SubscribeOnce(topic string, callback mqtt.MessageHandler) (err error) {
	return self.Subscribe(topic, func(client mqtt.Client, message mqtt.Message) {
		defer func() { _ = self.Unsubscribe(topic) }()
		callback(client, message)
	})
}

func (self *_Client) Publish(topic string, payload any, retained ...bool) error {
	return self.await(self.delegate.Publish(topic, byte(self.qos), len(retained) > 0 && retained[0], payload))
}

func (self *_Client) await(token mqtt.Token) error {
	token.Wait()
	return token.Error()
}
