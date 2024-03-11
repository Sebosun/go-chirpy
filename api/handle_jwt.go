package api

import (
	"fmt"
	"net/http"
)

func (api *ApiConfig) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("twojstary")
}

/* func (api *ApiConfig) HandleRevoke(w http.ResponseWriter, r *http.Request) { */
/* } */
