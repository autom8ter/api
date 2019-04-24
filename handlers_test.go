package api_test

import (
	"fmt"
	"github.com/autom8ter/api"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Setup(t *testing.T) *httptest.Server {
	err := api.InitSessions("")
	Fail(err, t)
	a := &api.Auth{
		Domain:       os.Getenv("AUTH0_DOMAIN"),
		ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
	}
	m := a.Mux("/dashboard")
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Logged Out")
	})
	m.HandleFunc("/dashboard", a.RequireLogin(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Logged In")
	}))

	return httptest.NewServer(m)
}

func TestHome(t *testing.T) {
	server := Setup(t)
	defer server.Close()
	fmt.Println("Server URL: ", server.URL)
	resp, err := http.Get(server.URL)
	Fail(err, t)
	bits, err := ioutil.ReadAll(resp.Body)
	Fail(err, t)
	fmt.Println(string(bits))
}

func TestLogin(t *testing.T) {
	server := Setup(t)
	defer server.Close()
	fmt.Println("Server URL: ", server.URL)
	resp, err := http.Get(server.URL+"/login")
	Fail(err, t)
	bits, err := ioutil.ReadAll(resp.Body)
	Fail(err, t)
	fmt.Println(string(bits))
}

func TestDashboard(t *testing.T) {
	server := Setup(t)
	defer server.Close()
	fmt.Println("Server URL: ", server.URL)
	resp, err := http.Get(server.URL+"/dashboard")
	Fail(err, t)
	bits, err := ioutil.ReadAll(resp.Body)
	Fail(err, t)
	fmt.Println(string(bits))
}

func Fail(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err.Error())
	}
}