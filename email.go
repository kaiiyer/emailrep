package emailrep

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Rep function will return the reputation of the email address passed as the argument
func Rep(email string) string {
	/* local variable definition */

	var mail = email

	resp, err := http.Get("https://emailrep.io/" + mail)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handling err
		fmt.Printf("Sorry we got hit by an error!\n")
	}
	defer resp.Body.Close()

	return string(body)
}
