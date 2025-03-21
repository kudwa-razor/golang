package main
import (
	"fmt"
	"net/http"
)

//Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. Here our simple response is "hello\n"
func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

//reading all the HTTP request headers and echoing them into the response body.
func main(){
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)


	http.ListenAndServe(":8090", nil)
}
