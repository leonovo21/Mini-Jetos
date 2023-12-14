package fil

import (
	"fmt"
	"log"
	"net/http"
)

func Smain() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	fmt.Println("Starting server at port 6969\n")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatal(err)
	}
}
func helloHandler() {

}
