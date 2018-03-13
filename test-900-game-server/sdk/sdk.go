package sdk

import (
	"net/http"
	"fmt"
)

type SDK struct{}

func (s SDK) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello!")
}
