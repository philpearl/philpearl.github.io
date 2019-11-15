package badgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"strings"
	"testing"
)

func TestJSONFunny(t *testing.T) {
	var payload struct {
		Hello string `json:"hello"`
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// {"hello":""}
		payload.Hello = strings.Repeat("h", 499)
		data, _ := json.Marshal(&payload)
		fmt.Println(len(data))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(data)
		// w.Write([]byte("\n"))
	}))

	c := s.Client()

	ct := &httptrace.ClientTrace{
		PutIdleConn: putIdleConn,
	}
	cxt := httptrace.WithClientTrace(context.Background(), ct)

	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = req.WithContext(cxt)

	rsp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	fmt.Println(rsp.Header)

	if err := json.NewDecoder(rsp.Body).Decode(&payload); err != nil {
		t.Fatal(err)
	}

	fmt.Println(payload)

	// n, err := io.Copy(ioutil.Discard, rsp.Body)
	// if n != 0 {
	// 	t.Fatalf("expected n to be zero, is %d", n)
	// }
	// if err != nil {
	// 	t.Fatal(err)
	// }
	fmt.Println("done")
}

func putIdleConn(err error) {
	fmt.Println("PutIdleConn", err)
}
