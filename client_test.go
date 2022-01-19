package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var BaseUrl = "http://localhost:8080"
var client = &http.Client{}

func TestClientCreateUser(t *testing.T) {
	endPoint := BaseUrl + "/users"

	data := url.Values{}
	data.Set("email", "jonathan.patta@gmail.com")
	data.Set("firstname", "Jonathan")

	payload := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST", endPoint, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodyData := string(body)

	if resp.StatusCode != 200 {
		t.Errorf(bodyData)
	}
}

func TestClientGetUser(t *testing.T) {
	endPoint := BaseUrl + "/users"

	data := url.Values{}
	data.Set("email", "jonathan.patta@gmail.com")
	data.Set("firstname", "Jonathan")

	payload := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST", endPoint, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodyData := string(body)

	fmt.Println(bodyData)

	if resp.StatusCode != 200 {
		t.Errorf("Error while creating user:" + bodyData)
	}

	var user User

	// err = json.NewDecoder(resp.Body).Decode(&user)
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println(err)
		t.Error()
	}
	id := user.UUID

	endPoint = BaseUrl + "/users/" + id

	resp, err = http.Get(endPoint)
	body, _ = ioutil.ReadAll(resp.Body)
	bodyData = string(body)
	if resp.StatusCode != 200 {
		t.Errorf("Get User Error: " + bodyData)
	}
	if err != nil {
		panic(err)
	}
}

func TestClientUpdateUser(t *testing.T) {

}