package main
import(
	"fmt"
	"time"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request){
	ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
		err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}
	
func main(){
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
// this also is giving the output: hello --> which is same as the output given by http-server.go .. idk hy shld gpt.