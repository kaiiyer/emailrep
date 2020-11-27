package emailrep_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	//
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"Status\": \"Looks Okay\"}")
	}

	req := httptest.NewRequest("GET", "https://emailrep.io/", nil)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.Status == "200" {
		t.Fatal("Something went wrong!")
	}
}
