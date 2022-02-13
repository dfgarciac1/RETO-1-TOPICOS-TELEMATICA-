package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var connections []*websocket.Conn

type (
	Html struct {
		a   string
		SVG string
		img string
		p   string
	}
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error(), "  ")
		os.Exit(1)
	}
}

func ReadHTML(msg []byte, connections []*websocket.Conn) {
	tagHTML := new(Html)
	tagHTML.a = "a"
	tagHTML.SVG = "SVG"
	tagHTML.img = "img"
	url := string(msg)

	con, err := net.Dial("tcp", url)
	checkError(err)

	req := "HEAD / HTTP/1.0\r\n\r\n"

	_, err = con.Write([]byte(req))
	checkError(err)

	res, err := ioutil.ReadAll(con)
	checkError(err)

	fmt.Println(string(res))

	resp, err := http.Get("http://" + url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Print("HTLM", string(html))

	SVG := regexp.MustCompile(`<svg.*?>(.*)</svg>`)
	a := regexp.MustCompile(`<a.*?>(.*)</a>`)
	img := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	p := regexp.MustCompile(`<p\s*.*>\s*.*<\/p>`)
	submatchallSVG := SVG.FindAllStringSubmatch(string(html), -1)
	submatchallimg := img.FindAllStringSubmatch(string(html), -1)
	submatchalla := a.FindAllStringSubmatch(string(html), -1)
	submatchallp := p.FindAllStringSubmatch(string(html), -1)

	SaveDataHTML(submatchallSVG, tagHTML.SVG)
	SaveDataHTML(submatchalla, tagHTML.a)
	SaveDataHTML(submatchallimg, tagHTML.img)
	SaveDataHTML(submatchallp, tagHTML.p)

}

func SaveDataHTML(data [][]string, nameTag string) {
	for _, element := range data {
		name := CreateRandomName(len(data))
		f, err := os.Create(nameTag + "_" + name + ".txt")
		if err != nil {
			log.Print(err)
		}
		_, err2 := f.WriteString(element[1])

		if err2 != nil {
			log.Fatal(err2)
		}

		log.Print("Se proceso el mensaje y se creo el txt")
	}
}

func CreateRandomName(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

func UploadFile(msg []byte, connections []*websocket.Conn) {
	dec, err := base64.StdEncoding.DecodeString(string(msg))
	if err != nil {
		panic(err)
	}
	name := CreateRandomName(1)
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}

func main() {
	incomingHTML := make(chan []byte)
	incomingFile := make(chan []byte)
	incomingConnections := make(chan *websocket.Conn)

	Port := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[1])
		os.Exit(1)
	}
	log.Print("***********************************")
	log.Print("Server is running...")
	log.Print("PORT :", Port)

	go func() {
		for {
			select {
			case msg := <-incomingHTML:
				ReadHTML(msg, connections)
			case msg := <-incomingFile:
				UploadFile(msg, connections)
			case conn := <-incomingConnections:
				connections = append(connections, conn)
			}

		}
	}()

	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		conn, _ := upgrader.Upgrade(writer, request, nil)
		defer conn.Close()
		incomingConnections <- conn
		for {
			// Receive message
			_, message, _ := conn.ReadMessage()
			incomingHTML <- message
		}
	})

	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
