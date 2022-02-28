package integration_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"most-used-word/src/config"
	"most-used-word/src/model"
	"net/http"
	"os"
	"testing"
)

const ENDPOINT_TOP_USE = "/top-used"
const CONTENT_TYPE = "application/json"

func TestMain(t *testing.M) {
	log.Println("Start integration test!")
	exitVal := t.Run()
	log.Println("End integration test!")

	os.Exit(exitVal)
}

func TestWithStatusOK(t *testing.T) {

	bodyContent := model.ContentReceive{
		Title:     "top 1 word",
		TopNumber: "1",
		Content:   "the lazy fox jumped over the brown dog",
	}

	bodyByte, err := json.Marshal(bodyContent)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	responseBody := bytes.NewBuffer(bodyByte)
	serverConfig := config.GetServerConfig()
	resp, err := http.Post(serverConfig.Protocol+"://"+serverConfig.ServerHost+":"+serverConfig.ServerPort+ENDPOINT_TOP_USE, CONTENT_TYPE, responseBody)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Errorf("Expect status code 200, got : %s", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	respContent := model.TopUsedReturn{}

	err = json.Unmarshal(body, &respContent)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if respContent.Result[0].Word != "the" || respContent.Result[0].NumberOccur != 2 {
		t.Errorf("Expected output {\"the\" : 2, got %+v", respContent.Result[0])
		return
	}

}

func TestWithStatusBadRequest(t *testing.T) {
	type InvalidStruct struct {
		InvalidValue string
	}

	bodyContent := InvalidStruct{
		InvalidValue: "string",
	}

	bodyByte, err := json.Marshal(bodyContent)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	responseBody := bytes.NewBuffer(bodyByte)

	serverConfig := config.GetServerConfig()
	resp, err := http.Post(serverConfig.Protocol+"://"+serverConfig.ServerHost+":"+serverConfig.ServerPort+ENDPOINT_TOP_USE, CONTENT_TYPE, responseBody)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	respContent := model.ErrorResponse{}

	err = json.Unmarshal(body, &respContent)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if respContent.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code return 400, got: %d", respContent.StatusCode)
		return
	}
}
