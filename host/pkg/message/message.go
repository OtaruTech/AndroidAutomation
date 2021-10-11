package message

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-basic/uuid"
)

const TOPIC_CLIENT_ONLINE string = "/client/online"
const TOPIC_PING_CLIENT string = "/ping"
const TOPIC_PONG_SERVER string = "/pong"
const TOPIC_SERVICE_ADD string = "/service/add"
const TOPIC_SERVICE_GET string = "/service/get"
const TOPIC_PROVIDER_RSP_FMT string = "/provider/%s/%d"
const TOPIC_METHOD_FMT string = "/method/%s/%s"
const TOPIC_CALLER_RSP_FMT string = "/caller/%s/%d"
const TOPIC_AVAILABLE_BUS_LISTENER string = "/bus/service/available"
const TOPIC_LOST_BUS_LISTENER string = "/bus/service/lost"

type MethodHandle func(Message) Message
type ServiceAvailableHandle func(clientId string, service string)
type ServiceLostHandle func(clientId string, services []string)

type ClientConnectedHandle func()
type ClientDisconnectedHandle func()

type Message struct {
	buf string
}

func (message *Message) Empty() bool {
	if message.buf == "" {
		return true
	} else {
		return false
	}
}

func (message *Message) SetBool(key string, val bool) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetInt(key string, val int) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetFloat(key string, val float64) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetString(key string, val string) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetMessage(key string, val *Message) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val.buf
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetIntArray(key string, val []int) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	obj[key] = val
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetStringArray(key string, val []string) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	strArr := make([]string, len(val))
	for _, v := range val {
		strArr = append(strArr, v)
	}
	obj[key] = strArr
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) SetMessageArray(key string, val []Message) {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	valString := make([]string, 0)
	for _, v := range val {
		valString = append(valString, v.buf)
	}
	obj[key] = valString
	out, _ := json.Marshal(obj)
	message.buf = string(out)
}

func (message *Message) GetBool(key string, defVal bool) bool {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		return obj[key].(bool)
	}
	return defVal
}

func (message *Message) GetInt(key string, defVal int) int {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		return (int)(obj[key].(float64))
	}
	return defVal
}

func (message *Message) GetFloat(key string, defVal float64) float64 {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		return obj[key].(float64)
	}
	return defVal
}

func (message *Message) GetString(key string, defVal string) string {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		return obj[key].(string)
	}
	return defVal
}

func (message *Message) GetMessage(key string, defVal *Message) *Message {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		msgString := obj[key].(string)
		return &Message{buf: msgString}
	}
	return defVal
}

func (message *Message) GetIntArray(key string) []int {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		intArr := make([]int, len(obj[key].([]interface{})))
		for k := range intArr {
			intArr[k] = (int)(obj[key].([]interface{})[k].(float64))
		}
		return intArr
	}
	return []int{}
}

func (message *Message) GetStringArray(key string) []string {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		strArr := make([]string, len(obj[key].([]interface{})))
		for k := range strArr {
			strArr[k] = obj[key].([]interface{})[k].(string)
		}
		return strArr
	}
	return []string{}
}

func (message *Message) GetMessageArray(key string) []Message {
	var obj map[string]interface{} = make(map[string]interface{})
	json.Unmarshal([]byte(message.buf), &obj)
	if _, exist := obj[key]; exist {
		msgArr := make([]Message, len(obj[key].([]interface{})))
		for k := range msgArr {
			msgArr[k].buf = obj[key].([]interface{})[k].(string)
		}
		return msgArr
	}
	return []Message{}
}

type Method struct {
	topic  string
	handle MethodHandle
}

type LocalService struct {
	Name      string
	MethodMap map[string]*Method
}

type RemoteService struct {
	Name    string
	Methods []string
}

type MessageBusListener struct {
	AvailableHandle  ServiceAvailableHandle
	LostHandle       ServiceLostHandle
	ConnectHandle    ClientConnectedHandle
	DisconnectHandle ClientDisconnectedHandle
}

func (service LocalService) RegisterMethod(name string, handle MethodHandle) {
	if _, exist := service.MethodMap[name]; exist {
		log.Printf("message: register method failed, name %s already exist\n", name)
		return
	}
	topic := "/method/" + service.Name + "/" + name
	method := &Method{topic: topic, handle: handle}
	service.MethodMap[name] = method
}

func (service RemoteService) CallMethod(name string, message Message) (output Message) {
	var call MethodCall
	call.ServiceName = service.Name
	call.MethodName = name
	call.ClientId = gClientId
	call.MsgId = getUuid()
	call.InputMsg = message.buf

	topic := fmt.Sprintf(TOPIC_METHOD_FMT, service.Name, name)
	rspTopic := fmt.Sprintf(TOPIC_CALLER_RSP_FMT, gClientId, call.MsgId)
	subscribe(rspTopic)
	data, err := json.Marshal(&call)
	if err != nil {
		log.Println("message: failed to encode method call:", err.Error())
		unsubscribe(rspTopic)
		return Message{}
	}
	publish(topic, string(data))
	waitTopicQueue = append(waitTopicQueue, rspTopic)
	var outputString string
	requestChan := make(chan Status)
	methodRequestMap[call.MsgId] = requestChan
	select {
	case rsp := <-requestChan:
		if rsp.Status != "okay" {
			unsubscribe(rspTopic)
			delete(methodRequestMap, call.MsgId)
			return Message{}
		}
		outputString = rsp.Data
		break
	case <-time.After(2 * time.Second):
		delete(methodRequestMap, call.MsgId)
		for k, v := range waitTopicQueue {
			if v == rspTopic {
				waitTopicQueue = append(waitTopicQueue[:k], waitTopicQueue[k+1:]...)
			}
		}
		unsubscribe(rspTopic)
		return Message{}
	}

	unsubscribe(rspTopic)
	delete(methodRequestMap, call.MsgId)
	return Message{buf: outputString}
}

//////////////////////JSON DATA///////////////////
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

type MethodCall struct {
	ServiceName string `json:"serviceName"`
	MethodName  string `json:"methodName"`
	ClientId    string `json:"clientId"`
	MsgId       int    `json:"msgId"`
	InputMsg    string `json:"input"`
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

var (
	msgId            int
	gMqClient        mqtt.Client
	gClientId        string
	localServiceMap  map[string]*LocalService
	remoteServiceMap map[string]*RemoteService
	waitTopicQueue   []string
	rspChan          chan Status
	methodRequestMap map[int]chan Status
	busListener      *MessageBusListener
	keepaliveTimer   *time.Timer
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := msg.Payload()
	// log.Printf("message: +recv: %s -> %s", topic, payload)
	keepaliveTimer.Reset(time.Minute)
	if topic == TOPIC_PING_CLIENT {
		var clientId ClientId
		err := json.Unmarshal(payload, &clientId)
		if err != nil {
			log.Println("message: failed to parse message:", payload)
			return
		}
		if clientId.ClientId == gClientId {
			//response pong to daemon
			go publish(TOPIC_PONG_SERVER, string(payload))
		}
		return
	}

	if topic == TOPIC_AVAILABLE_BUS_LISTENER {
		var availableBusListener AvailableBusListener
		err := json.Unmarshal(payload, &availableBusListener)
		if err != nil {
			log.Println("message: failed to parse message:", payload)
			return
		}
		if busListener != nil {
			go func() {
				if busListener.AvailableHandle != nil {
					busListener.AvailableHandle(availableBusListener.ClientId, availableBusListener.Service)
				}
			}()
		}
		return
	}

	if topic == TOPIC_LOST_BUS_LISTENER {
		var lostBusListener LostBusListener
		err := json.Unmarshal(payload, &lostBusListener)
		if err != nil {
			log.Println("message: failed to parse message:", payload)
			return
		}
		if busListener != nil {
			go func() {
				if busListener.LostHandle != nil {
					busListener.LostHandle(lostBusListener.ClientId, lostBusListener.Services)
				}
			}()
		}
		return
	}

	//service add/method call response
	for len(waitTopicQueue) > 0 {
		for k, v := range waitTopicQueue {
			if v == topic {
				var status Status
				err := json.Unmarshal(payload, &status)
				if err != nil {
					log.Println("message: failed to parse message:", err.Error())
					return
				}
				strs := strings.Split(topic, "/")
				if len(methodRequestMap) > 0 && len(strs) == 4 && strs[1] == "caller" {
					index, _ := strconv.Atoi(strs[3])
					request := methodRequestMap[index]
					request <- status
				} else {
					rspChan <- status
				}
				waitTopicQueue = append(waitTopicQueue[:k], waitTopicQueue[k+1:]...)
			}
		}
	}
	//method call
	for _, v := range localServiceMap {
		for _, vv := range v.MethodMap {
			if topic == vv.topic {
				var call MethodCall
				err := json.Unmarshal(payload, &call)
				if err != nil {
					log.Println("failed to parse message:", payload)
					break
				}
				rspTopic := fmt.Sprintf(TOPIC_CALLER_RSP_FMT, call.ClientId, call.MsgId)
				go func(method *Method, topic string) {
					output := method.handle(Message{buf: call.InputMsg})
					var rsp Status
					rsp.Status = "okay"
					rsp.Data = output.buf
					data, err := json.Marshal(&rsp)
					if err != nil {
						log.Println("failed to encode rsp data")
						return
					}
					go publish(topic, string(data))
				}(vv, rspTopic)
			}
		}
	}

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	initClient()
	if busListener.ConnectHandle != nil {
		busListener.ConnectHandle()
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("message: connection lost")
	if busListener.DisconnectHandle != nil {
		busListener.DisconnectHandle()
	}
}

func InitMessage(host string, port int) {
	uuid := uuid.New()
	gClientId = "client-" + uuid
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(gClientId)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	gMqClient = mqtt.NewClient(opts)

	if token := gMqClient.Connect(); token.Wait() && token.Error() != nil {
		log.Println("message: ", token.Error())
	} else {
		log.Println("message: MQTT Initialize:", host, ":", port)
	}
	localServiceMap = make(map[string]*LocalService)
	remoteServiceMap = make(map[string]*RemoteService)
	rspChan = make(chan Status)
	methodRequestMap = make(map[int]chan Status)

	// var msg Message
	// var innerMsg Message
	// innerMsg.SetString("str", "hello, world!")
	// innerMsg.SetInt("int", 100)
	// msg.SetMessage("msg", &innerMsg)
	// log.Println("innerMsg", innerMsg)
	// log.Println("msg", msg)
	// log.Println("getMsg", msg.GetMessage("msg", nil))
	// msgArr := make([]Message, 0)
	// for i := 0; i < 10; i++ {
	// 	var msg Message
	// 	msg.SetInt("index", i)
	// 	msg.SetString("str", fmt.Sprintf("hello%02d", i))
	// 	msgArr = append(msgArr, msg)
	// }
	// var message Message
	// message.SetMessageArray("msgArr", msgArr)
	// log.Println(message)
	// log.Println("msgArr", message.GetMessageArray("msgArr"))
}

func publish(topic string, message string) {
	// log.Printf("message: +send: %s -> %s\n", topic, message)
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

func getUuid() int {
	msgId += 1
	return msgId
}

func initClient() {
	var clientId ClientId
	log.Println("message: Connect to MQTT server")
	clientId.ClientId = gClientId
	bdata, err := json.Marshal(&clientId)
	if err != nil {
		log.Println("message: failed to encode client id")
		return
	}
	publish(TOPIC_CLIENT_ONLINE, string(bdata))
	subscribe(TOPIC_PING_CLIENT)
	keepaliveTimer = time.NewTimer(time.Minute)
	go func() {
		<-keepaliveTimer.C
		log.Println("client offline!!")
		if busListener.DisconnectHandle != nil {
			busListener.DisconnectHandle()
		}
	}()
}

//////////////////////public API//////////////////

func RegisterBusListener(listener MessageBusListener) {
	busListener = &listener
	subscribe(TOPIC_AVAILABLE_BUS_LISTENER)
	subscribe(TOPIC_LOST_BUS_LISTENER)
}

func CreateLocalService(name string) *LocalService {
	if _, exist := localServiceMap[name]; exist {
		log.Printf("message: create local service, name %s already exist\n", name)
		return nil
	}
	return &LocalService{Name: name, MethodMap: make(map[string]*Method)}
}

func AddService(service *LocalService) {
	var add ServiceAdd
	add.ClientId = gClientId
	add.MsgId = getUuid()
	add.ServiceName = service.Name
	add.Methods = make([]string, len(service.MethodMap))
	var id = 0
	for k, _ := range service.MethodMap {
		add.Methods[id] = k
		id++
	}

	data, err := json.Marshal(&add)
	if err != nil {
		log.Println("message: failed to encode service add data",
			err.Error())
		return
	}
	rspTopic := fmt.Sprintf(TOPIC_PROVIDER_RSP_FMT, gClientId, add.MsgId)
	subscribe(rspTopic)
	publish(TOPIC_SERVICE_ADD, string(data))
	waitTopicQueue = append(waitTopicQueue, rspTopic)
	//wait for response
	select {
	case rsp := <-rspChan:
		if rsp.Status != "okay" {
			//response failed
			log.Println("message: rsp status:", rsp.Status)
			unsubscribe(rspTopic)
			return
		}
		break
	case <-time.After(5 * time.Second):
		unsubscribe(rspTopic)
		return
	}
	unsubscribe(rspTopic)
	for _, v := range service.MethodMap {
		subscribe(v.topic)
	}
	localServiceMap[service.Name] = service
}

func GetService(name string) *RemoteService {
	var get ServiceGet
	get.ServiceName = name
	get.ClientId = gClientId
	get.MsgId = getUuid()
	rspTopic := fmt.Sprintf(TOPIC_CALLER_RSP_FMT, gClientId, get.MsgId)
	subscribe(rspTopic)
	data, err := json.Marshal(&get)
	if err != nil {
		log.Println("message: failed to encode service get data:", err.Error())
		unsubscribe(rspTopic)
		return nil
	}
	publish(TOPIC_SERVICE_GET, string(data))
	waitTopicQueue = append(waitTopicQueue, rspTopic)
	var service *RemoteService
	select {
	case rsp := <-rspChan:
		if rsp.Status != "okay" {
			// log.Println("message: rsp status:", rsp.Status)
			unsubscribe(rspTopic)
			return nil
		}
		service = &RemoteService{Name: name}
		service.Methods = make([]string, len(rsp.Methods))
		for _, v := range rsp.Methods {
			service.Methods = append(service.Methods, v)
		}
		break
	case <-time.After(5 * time.Second):
		log.Printf("message: get service timeout")
		unsubscribe(rspTopic)
		return nil
	}
	unsubscribe(rspTopic)
	remoteServiceMap[name] = service

	return service
}
