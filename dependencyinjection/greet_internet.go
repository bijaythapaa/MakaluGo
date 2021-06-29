package dependencyinjection

import (
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	OtherGreet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
