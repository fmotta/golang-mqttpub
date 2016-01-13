package main

import (
    "fmt"
    "os"
    "os/signal"
    "flag"
    "strings"

    "github.com/yosssi/gmq/mqtt"
    "github.com/yosssi/gmq/mqtt/client"
)

var host = flag.String("host", "localhost", "hostname of broker")
//var user = flag.String("user", "", "username")
//var pass = flag.String("pass", "", "password")
var retain = flag.Bool("retain", false, "retain message?")
var wait = flag.Bool("wait", false, "stay connected after publishing?")
var message = flag.String("message", "", "message")
var topic = flag.String("topic", "", "topic")
var port = flag.String("port", "1883", "port")
//var qos = flag.String("qos", string(mqtt.QoS0)[:1], "qos")
//var qos = flag.String("qos", "0", "qos") // something quirky with default of "0" results in no default
var qos = flag.String("qos", "00", "qos")

var Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
}

func main() {
    flag.Parse()

    if (flag.NFlag() < 2) {
    	fmt.Println("Too few arguments")
	Usage()
	os.Exit(0)
    }
    if ((*message == "") || (*topic == "")) {
    	fmt.Println("Need a topic and message to publish")
	Usage()
	os.Exit(0)
    }

    fmt.Println("topic: ", *topic, "\tmessage: ", *message)

    // Set up channel on which to send signal notifications.
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, os.Kill)

    // Create an MQTT Client.
    cli := client.New(&client.Options{
        // Define the processing of the error handler.
        ErrorHandler: func(err error) {
            fmt.Println(err)
        },
    })

    // Terminate the Client.
    defer cli.Terminate()

    s := []string{*host, *port}

    address := strings.Join(s, ":")
    fmt.Println("host info: ", address, "QoS option: ", *qos)


    var mQoS = mqtt.QoS0
     switch *qos {
	case "0":
		mQoS = mqtt.QoS0
	case "1":
		mQoS = mqtt.QoS1
	case "2":
		mQoS = mqtt.QoS2
	default:
		mQoS = mqtt.QoS0
    }

    // Connect to the MQTT Server.
    err := cli.Connect(&client.ConnectOptions{
        Network:  "tcp",
        Address:  address,
        ClientID: []byte("example-client"),
	CONNACKTimeout:  30,
	KeepAlive:       50,
    })
    if err != nil {
        panic(err)
    }

    err = cli.Publish(&client.PublishOptions{
        QoS:       mQoS,
        TopicName: []byte(*topic),
        Message:   []byte(*message),
    })
    if err != nil {
        panic(err)
    }


    // Wait for receiving a signal.
    // I cannot figure out why publish does not work without this command
    <-sigc

    // Disconnect the Network Connection.
    if err := cli.Disconnect(); err != nil {
        panic(err)
    }
}

