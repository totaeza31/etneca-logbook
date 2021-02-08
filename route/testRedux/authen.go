package route

import (
	"etneca-logbook/controllers"

	"github.com/gorilla/mux"
)

func AuthenRoute(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/profile", controllers.VerifyAccess(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/token", controllers.VarifyRefresh(controllers.GetNewToken)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	router.HandleFunc("/forgot", controllers.GetNewPassword).Methods("POST")
	router.HandleFunc("/reset/{email}", controllers.ResetPassword).Methods("GET")

}

// func login(response http.ResponseWriter, request *http.Request) {
// 	url := "http://35.240.214.110/login"
// 	fmt.Println("URL:>", url)

// 	var jsonStr = []byte(`{"email":"sean_bean@gameofthron.es","password":"$2b$12$UREFwsRUoyF0CRqGNK0LzO0HM/jLhgUCNNIJ9RJAqMUQ74crlJ1Vu"}`)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("response Status:", resp.Status)
// 	fmt.Println("response Headers:", resp.Header)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println("response Body:", string(body))
// }
