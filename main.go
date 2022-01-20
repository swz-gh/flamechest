package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/TheCreeper/go-notify"
)

type ChestFileInfo struct {
	FileID   string `json:"fileID"`
	FileName string `json:"fileName"`
	DelKey   string `json:"delKey"`
	Url      string `json:"url"`
}

type ChestResponse struct {
	Status   int           `json:"status"`
	Response string        `json:"response"`
	Content  ChestFileInfo `json:"content"`
}

func main() {
	println("Thanks for using flamechest")
	img, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}
	if string(img) == "screenshot aborted\n" {
		fmt.Println("Screenshot aborted")
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", os.Getenv("FLAMECHEST_ENDPOINT"), bytes.NewReader(img))
	req.Header.Add("authorization", os.Getenv("FLAMECHEST_AUTH_TOKEN"))
	req.Header.Add("content-type", "image/png")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var chestResponse ChestResponse
	json.Unmarshal(responseData, &chestResponse)

	fmt.Print(chestResponse.Content.Url)
	println("🔥", chestResponse.Content.Url)

	ntf := notify.NewNotification("🔥 Flamechest upload", chestResponse.Content.Url)
	ntf.Show()
}
