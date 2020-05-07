package emailrep

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func rep(email string) string {
	/* local variable definition */

	var mail = email

	//	resp, err := http.Get("https://emailrep.io/" + mail)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1" + mail)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handling err
		fmt.Printf("Sorry we got hit by an error!\n")
	}
	defer resp.Body.Close()
	fmt.Println(string(body))

	return string(body)
}
