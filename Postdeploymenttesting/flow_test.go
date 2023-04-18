package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func rlcm_TemplateTest(getpost string, urlparams string, rawbody string) (string, int) {
	//token := generate_token_video()
	url := "http://localhost:8080/" + urlparams

	if getpost == "POST" {
		if rawbody != "" {
			body := []byte(rawbody)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			// //cookie := http.Cookie{Name: "SessionId", Value: sessionId}
			// //req.AddCookie(&cookie)
			// tlsConfig := &tls.Config{
			// 	InsecureSkipVerify: true,
			// }

			// tlsConfig.BuildNameToCertificate()
			// transport := &http.Transport{TLSClientConfig: tlsConfig}
			// client := &http.Client{Transport: transport}
			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				log.Print(err)
				return "", 500
			}
			defer resp.Body.Close()

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Print(err)
			}
			return (string(data)), resp.StatusCode
		} else {
			req, err := http.NewRequest("POST", url, nil)
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				log.Print(err)
				return "", 500
			}
			defer resp.Body.Close()

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Print(err)
			}
			return (string(data)), resp.StatusCode
		}
	} else {
		req, err := http.NewRequest(getpost, url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Print(err)
			return "", 500
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		return (string(data)), resp.StatusCode
	}

}
func Test_Postiveflow(t *testing.T) {
	url := "users/gopal/add"
	data, code := rlcm_TemplateTest("POST", url, `{"amount":500}`)
	if code != 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
	url = "users/gopal/balance"
	data, code = rlcm_TemplateTest("GET", url, `{"amount":500}`)
	if code != 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
	url = "users/gopal/history"
	data, code = rlcm_TemplateTest("GET", url, `{"amount":500}`)
	if code != 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
}

func Test_Negativeflow1(t *testing.T) {
	url := "users/gopal/add"
	data, code := rlcm_TemplateTest("GET", url, `{"amount":500}`)
	if code == 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
	url = "users/gopal/balance"
	data, code = rlcm_TemplateTest("PUT", url, `{"amount":500}`)
	if code == 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
	url = "users/gopal/history"
	data, code = rlcm_TemplateTest("PATCH", url, `{"amount":500}`)
	if code == 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
}
func Test_Negativeflow2(t *testing.T) {
	url := "users/gopal/add"
	data, code := rlcm_TemplateTest("POST", url, `{"amount":-500}`)
	if code == 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
}
func Test_Negativeflow3(t *testing.T) {
	url := "users/unknown/balance"
	data, code := rlcm_TemplateTest("GET", url, `{"amount":-500}`)
	if code != 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
}
func Test_Negativeflow4(t *testing.T) {
	url := "users/unknown/history"
	data, code := rlcm_TemplateTest("GET", url, `{"amount":-500}`)
	if code != 200 {
		t.Error("Error:", data, " expected:", 200, "actual: ", code)
	} else {
		fmt.Println("Test passed: ", data, " Status: ", code)
	}
}
