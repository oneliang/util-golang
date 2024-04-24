package test

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker = "tcp://120.79.179.187:1884"
	// broker   = "8.134.59.213:1883"
	clientID = "stu_app_63_10"
	// topic    = "students/homework/cancelled"
	topic = "test/#"
)

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func TestMqttGetMsh(t *testing.T) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetCleanSession(false)
	opts.SetUsername("ponylearn-student-mobile")
	opts.SetPassword("123456")

	client := mqtt.NewClient(opts)
	//连接
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to broker: %v", token.Error())
	}
	//订阅
	if token := client.Subscribe(topic, 2, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("aaa")
	// 等待中断信号，以便优雅地关闭 MQTT 客户端
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("b")
	return
	// 断开 MQTT 连接
	client.Disconnect(250)
}

const (
	// MQTT 服务器地址
	mqttServer = "tcp://120.79.179.187:1883"

	mqttPwd = "123456"
	// 客户端ID，应该是唯一的
	mqttClientID = "client-client-111"
	// 服务端ID
	mqttServerID = "server-client"
	// 订阅的主题
	onlineTopic  = "online"
	offlineTopic = "offline"
)

func handleSignals(client mqtt.Client) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel
	client.Disconnect(250)
}

// @Description: 测试MQtt客户端连接发送遗嘱消息
// @Author: Chenlin 2024-03-29 17:03:41
func TestMQttClient(t *testing.T) {
	opts := mqtt.NewClientOptions().
		AddBroker(mqttServer).
		SetPassword(mqttPwd).
		SetClientID(mqttClientID).
		SetAutoReconnect(true). //设置自动重连接
		SetWill(offlineTopic, clientID, 1, false)
	//设置重连接回调
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("ReConnect to broker")
	})
	// 创建MQTT客户端
	client := mqtt.NewClient(opts)
	// 断开连接
	defer client.Disconnect(250)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}
	if token := client.Subscribe(onlineTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to online topic: %v", token.Error())
	}
	// 等待一段时间以接收消息
	time.Sleep(20 * time.Minute)
	// 发布上线消息
	// token := client.Publish(onlineTopic, 1, false, clientID)
	// token.Wait()
	// if token.Error() != nil {
	// 	log.Fatalf("Error publishing online message: %v", token.Error())
	// }
	// 模拟设备在线，发送消息
	// for {
	// token := client.Publish(onlineTopic, 1, false, clientID)
	// token.Wait()
	// if token.Error() != nil {
	// 	log.Fatalf("Error publishing online message: %v", token.Error())
	// }
	// time.Sleep(30 * time.Minute) // 每隔30秒发送一次在线消息
	// }
}

// @Description: 测试mqtt服务端接收链接信息和遗嘱消息
// @Author: Chenlin 2024-03-29 15:04:32
func TestMQttServer(t *testing.T) {
	opts := mqtt.NewClientOptions().
		AddBroker(mqttServer).
		SetPassword(mqttPwd).
		SetClientID(mqttServerID)
	// 创建MQTT客户端
	client := mqtt.NewClient(opts)
	// 断开连接
	defer client.Disconnect(250)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}
	// 定义消息处理函数
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
	}
	// 订阅 online 和 offline 主题
	if token := client.Subscribe(onlineTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to online topic: %v", token.Error())
	}
	if token := client.Subscribe(offlineTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to offline topic: %v", token.Error())
	}
	// 等待一段时间以接收消息
	time.Sleep(10 * time.Minute)
	fmt.Println("Disconnected from MQTT broker")
}
