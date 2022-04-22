package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const SteamPort = "12345"
const JsonPort = "12340"

// 读取二进制文件字节流，再用这个字节流写入一个二进制文件
// 通过 http 发送二进制文件的字节流
func main() {
	// ***** LOCAL ***** //
	bContent := readFromBinary("a.zip")
	write2Binary(bContent, "local.zip")
	// ***** LOCAL ***** //

	// ***** application/octet-stream ***** //
	// server side
	go func() {
		getSteamAttachmentHandler := http.HandlerFunc(getSteamAttachment)
		http.Handle("/get-attachment-steam", getSteamAttachmentHandler)
		http.ListenAndServe(":"+SteamPort, nil)
	}()

	time.Sleep(time.Second)

	// client side
	if err := sendSteamAttachment("http://localhost:"+SteamPort+"/get-attachment-steam", "POST"); err != nil {
		fmt.Print(err.Error())
	}
	// ***** application/octet-stream ***** //

	// ***** JSON ***** //
	go func() {
		getJsonAttachmentHandler := http.HandlerFunc(getJsonAttachment)
		http.Handle("/get-attachment-json", getJsonAttachmentHandler)
		http.ListenAndServe(":"+JsonPort, nil)
	}()

	time.Sleep(time.Second)

	// client side
	if err := sendJsonAttachment("http://localhost:"+JsonPort+"/get-attachment-json", "POST"); err != nil {
		fmt.Print(err.Error())
	}
	// ***** JSON ***** //
}

type JsonData struct {
	Name string
	Data []byte
}

// Json
func sendJsonAttachment(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	b, err := ioutil.ReadFile("a.zip")
	if err != nil {
		return err
	}

	jsonData := JsonData{
		Name: "json.zip",
		Data: b,
	}
	jsonDataBytes, _ := json.Marshal(jsonData)

	req, err := http.NewRequest(method, url, bytes.NewReader(jsonDataBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("error: %d", rsp.StatusCode)
	}

	return nil
}

func getJsonAttachment(w http.ResponseWriter, request *http.Request) {
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonData := JsonData{}
	if err := json.Unmarshal(d, &jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpFile, err := os.Create(jsonData.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer tmpFile.Close()

	// TODO: check error
	tmpFile.Write(jsonData.Data)

	w.WriteHeader(http.StatusOK)
	return
}

// Json

// Steam
func sendSteamAttachment(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	b, err := ioutil.ReadFile("a.zip")
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("error: %d", rsp.StatusCode)
	}

	return nil
}

func getSteamAttachment(w http.ResponseWriter, request *http.Request) {
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpFile, err := os.Create("steam.zip")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer tmpFile.Close()

	// TODO: check error
	tmpFile.Write(d)

	w.WriteHeader(http.StatusOK)
	return
}

// Steam

//
func write2Binary(d []byte, fn string) {
	df, err := os.Create(fn)
	if err != nil {
		log.Fatal(err.Error())
	}

	df.Write(d)
}

func readFromBinary(fn string) []byte {
	b, _ := ioutil.ReadFile(fn)
	r := bytes.NewReader(b)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
