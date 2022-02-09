package main

import (
	"UserCrud/pb"
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

	if resp.StatusCode != 200 {
		t.Errorf("Error while creating user:" + bodyData)
	}

	var user pb.User

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
		t.Errorf("Error while creating user:" + bodyData)
	}

	var user pb.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println(err)
		t.Error()
	}
	id := user.UUID

	endPoint = BaseUrl + "/users/" + id

	data = url.Values{}
	data.Set("email", user.EmailAddress)
	data.Set("firstname", user.FirstName)
	data.Set("LastName", user.LastName)
	data.Set("DOB", user.DOB.String())
	data.Set("PhoneNumber", user.PhoneNumber)
	data.Set("UUID", user.UUID)

	payload = strings.NewReader(data.Encode())
	req, _ = http.NewRequest("PUT", endPoint, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	bodyData = string(body)

	if resp.StatusCode != 200 {
		t.Errorf("Error while Updating User: " + bodyData)
	}
}

func TestClientDeleteUser(t *testing.T) {
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
		t.Errorf("Error while creating user:" + bodyData)
	}

	var user pb.User

	// err = json.NewDecoder(resp.Body).Decode(&user)
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println(err)
		t.Error()
	}
	id := user.UUID

	endPoint = BaseUrl + "/users/" + id

	data = url.Values{}
	payload = strings.NewReader(data.Encode())
	req, _ = http.NewRequest("DELETE", endPoint, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	bodyData = string(body)

	if resp.StatusCode != 200 {
		t.Errorf("Error while Deleting User: " + bodyData)
	}
}

func TestClientListUsers(t *testing.T) {
	endPoint := BaseUrl + "/users/"

	resp, err := http.Get(endPoint)
	body, _ := ioutil.ReadAll(resp.Body)
	bodyData := string(body)
	if resp.StatusCode != 200 {
		t.Errorf("Get User Error: " + bodyData)
	}
	if err != nil {
		panic(err)
	}
}
