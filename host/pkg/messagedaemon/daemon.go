package messagedaemon

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-basic/uuid"
)

const TOPIC_CLIENT_ONLINE string = "/client/online"
const TOPIC_PING_CLIENT string = "/ping"
const TOPIC_PONG_SERVER string = "/pong"
const TOPIC_SERVICE_ADD string = "/service/add"
const TOPIC_SERVICE_GET string = "/service/get"

const TOPIC_PROVIDER_RSP_FMT string = "/provider/%s/%d"
const TOPIC_CALL_RSP_FMT string = "/caller/%s/%d"

const TOPIC_AVAILABLE_BUS_LISTENER string = "/bus/service/available"
const TOPIC_LOST_BUS_LISTENER string = "/bus/service/lost"

var (
	gMqClient  mqtt.Client
	gClientId  string
	clientMap  map[string]*Client
	serviceMap map[string]*Service
)

type Client struct {
	id string
	ch chan int
}

///////////////////JSON DATA/////////////////
type ClientId struct {
	ClientId string `json:"clientId"`
}

type ServiceAdd struct {
	ClientId    string   `json:"clientId"`
	MsgId       int      `json:"msgId"`
	ServiceName string   `json:"serviceName"`
	Methods     []string `json:"methods"`
}

type ServiceGet struct {
	ClientId    string `json:"clientId"`
	MsgId       int    `json:"msgId"`
	ServiceName string `json:"serviceName"`
}

type Service struct {
	Id      string   `json:"clientId"`
	Name    string   `json:"name"`
	Methods []string `json:"methods"`
}

type Status struct {
	Status  string   `json:"status"`
	Methods []string `json:"methods"`
	Data    string   `json:"data"`
}

type AvailableBusListener struct {
	ClientId string `json:"clientId"`
	Service  string `json:"serviceName"`
}

type LostBusListener struct {
	ClientId string   `json:"clientId"`
	Services []string `json:"services"`
}

func createNewClient(client *Client) {
	log.Println("daemon: create new client", client.id)
	for true {
		time.Sleep(10 * time.Second)
		msg := fmt.Sprintf("{\"clientId\": \"%s\"}", client.id)
		publish(TOPIC_PING_CLIENT, msg)
		select {
		case <-client.ch:
			// log.Println("daemon: client", client.id, "keepalived")
			break
		case <-time.After(8 * time.Second):
			// log.Println("daemon: client", client.id, "not response")
			services := make([]string, 0)
			for _, v := range serviceMap {
				if v.Id == client.id {
					services = append(services, v.Name)
					delete(serviceMap, v.Name)
				}
			}
			listener := LostBusListener{ClientId: client.id, Services: services}
			data, err := json.Marshal(listener)
			if err != nil {
				log.Println("daemon: failed encode listener")
				return
			}
			publish(TOPIC_LOST_BUS_LISTENER, string(data))
			delete(clientMap, client.id)
			return
		}
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := msg.Payload()
	// log.Printf("daemon: +recv: %s -> %s", topic, payload)
	if topic == TOPIC_CLIENT_ONLINE {
		//parse client id
		var clientId ClientId
		err := json.Unmarshal(payload, &clientId)
		if err != nil {
			log.Println("daemon: failed to parse message:", err.Error())
			return
		}
		//create thread to mantain the new client
		ch := make(chan int)
		clientMap[clientId.ClientId] = &Client{id: clientId.ClientId, ch: ch}
		go createNewClient(clientMap[clientId.ClientId])
		return
	}
	if topic == TOPIC_PONG_SERVER {
		//parse client id
		var clientId ClientId
		err := json.Unmarshal(payload, &clientId)
		if err != nil {
			log.Println("daemon: failed to parse message:", err.Error())
			return
		}
		if _, exist := clientMap[clientId.ClientId]; exist {
			client := clientMap[clientId.ClientId]
			client.ch <- 0
		}
		return
	}
	if topic == TOPIC_SERVICE_ADD {
		var add ServiceAdd
		err := json.Unmarshal(payload, &add)
		if err != nil {
			log.Println("daemon: failed to parse message:", err.Error())
			return
		}
		rspTopic := fmt.Sprintf(TOPIC_PROVIDER_RSP_FMT, add.ClientId, add.MsgId)
		if _, exist := serviceMap[add.ServiceName]; exist {
			log.Println("daemon: service name", add.ServiceName, "already exist")
			publish(rspTopic, fmt.Sprintf("{\"status\": \"fail\"}"))
			return
		}
		methods := make([]string, len(add.Methods))
		for k, v := range add.Methods {
			methods[k] = v
		}
		serviceMap[add.ServiceName] = &Service{Id: add.ClientId, Name: add.ServiceName, Methods: methods}

		go func() {
			publish(rspTopic, fmt.Sprintf("{\"status\": \"okay\"}"))
			listener := AvailableBusListener{ClientId: add.ClientId, Service: add.ServiceName}
			data, err := json.Marshal(listener)
			if err != nil {
				log.Println("daemon: failed encode listener")
				return
			}
			publish(TOPIC_AVAILABLE_BUS_LISTENER, string(data))
		}()
		return
	}
	if topic == TOPIC_SERVICE_GET {
		var get ServiceGet
		err := json.Unmarshal(payload, &get)
		if err != nil {
			log.Println("daemon: failed to parse message:", err.Error())
			return
		}
		rspTopic := fmt.Sprintf(TOPIC_CALL_RSP_FMT, get.ClientId, get.MsgId)
		if _, exist := serviceMap[get.ServiceName]; exist {
			var status Status
			status.Status = "okay"
			service := serviceMap[get.ServiceName]
			status.Methods = make([]string, len(service.Methods))
			for k, v := range service.Methods {
				status.Methods[k] = v
			}
			data, err := json.Marshal(status)
			if err != nil {
				log.Println("daemon: failed encode status")
				return
			}
			go publish(rspTopic, string(data))

			return
		}
		go publish(rspTopic, fmt.Sprintf("{\"status\": \"fail\"}"))
		return
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	initClient()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
}

func InitDaemon() {
	host := "127.0.0.1"
	port := 11883
	gClientId = "daemon"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(gClientId)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	gMqClient = mqtt.NewClient(opts)

	if token := gMqClient.Connect(); token.Wait() && token.Error() != nil {
		log.Println("daemon: ", token.Error())
	} else {
		log.Println("daemon: MQTT Initialize:", host, ":", port)
	}

	clientMap = make(map[string]*Client)
	serviceMap = make(map[string]*Service)
}

func publish(topic string, message string) {
	// log.Printf("daemon: +send: %s -> %s\n", topic, message)
	token := gMqClient.Publish(topic, 0, false, message)
	token.Wait()
}

func subscribe(topic string) {
	// log.Println("subscribe:", topic)
	token := gMqClient.Subscribe(topic, 1, nil)
	token.Wait()
}

func unsubscribe(topic string) {
	// log.Println("unsubscribe:", topic)
	token := gMqClient.Unsubscribe(topic)
	token.Wait()
}

func initClient() {
	log.Println("daemon: Connect to MQTT server")
	subscribe(TOPIC_CLIENT_ONLINE)
	subscribe(TOPIC_PONG_SERVER)
	subscribe(TOPIC_SERVICE_ADD)
	subscribe(TOPIC_SERVICE_GET)
}
