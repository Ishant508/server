package posting

import (
	"fmt"
	"net/http"
)

//Post 's the message
func Post(w http.ResponseWriter, r *http.Request) {
	t := post(string([]byte(r.URL.Path)[6:]))
	fmt.Printf(t.ReturnMessage + "\n")
	fmt.Fprintf(w, t.ReturnMessage)
}
