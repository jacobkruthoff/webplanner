package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `First Server`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body, got %v wanted %v", actual, expected)
	}

}

func TestRouter(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status should be ok, go %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "First Server"

	if respString != expected {
		t.Errorf("response should be %s, got %s", expected, respString)
	}

}
