package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "holy shit , this is home page!")
}

// 只能读取一次body
func readbody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "params %v", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is ", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "paras form error", err)
	}
	fmt.Fprintf(w, "before parse form %v", r.Form)

}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/readbody", readbody)
	http.HandleFunc("/isnil", getBodyIsNil)
	http.HandleFunc("/query", queryParams)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/getheader", header)
	http.HandleFunc("/getform", form)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
