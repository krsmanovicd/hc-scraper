package main

//missing swaggerui
//missing dockerfile
//aborted because of authentication failure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type InstagramUser struct {
	Username    string `json: "username"`
	Description string `json: "description"`
	// Followers int    `json: "followers"`
	// Following int    `json: "following"`
	// Posts     int    `json: "posts"`
}

func InstagramScraper() {

	instagram_user := InstagramUser{}

	collector := colly.NewCollector()

	//authentication not working
	// err := collector.Post("https://instagram.com/accounts/login/", map[string]string{"username": "**user**", "password": "**pass**"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	collector.AllowedDomains = []string{"www.instagram.com", "instagram.com"}

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Response recieved", r.StatusCode)
	})

	//username - collects username from visited page
	//description - collects number of posts, followers and following from visited page
	//the idea was to, later, divide the description into 3 different variables (posts, followers and following)
	collector.OnHTML("section._aa_l", func(e *colly.HTMLElement) {
		temp := e.DOM
		instagram_user = InstagramUser{
			Username:    temp.Find("div._aa_m").Find("_aacl _aacs _aact _aacx _aada").Text(),
			Description: temp.Find("ul._aa_7").Find("li._aa_5").Find("a.qi72231t nu7423ey n3hqoq4p r86q59rh b3qcqh3k fq87ekyn bdao358l fsf7x5fv rse6dlih s5oniofx m8h3af8h l7ghb35v kjdc1dyq kmwttqpk srn514ro oxkhqvkx rl78xhln nch0832m cr00lzj9 rn8ck1ys s3jn8y49 icdlwmnq _a6hd").Find("div._aacl _aacp _aacu _aacx _aad6 _aade").Text(),
			// Followers: ,
			// Following: ,
			// Posts:,
		}
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//this method would later be changed so we can type username and then it would be added to the url
	//this was just a test version
	collector.Visit("https://www.instagram.com/utopiantravel/")
	writeJSON(instagram_user)

}

// temporarily writing to json
func writeJSON(data InstagramUser) {
	f, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}

	_ = ioutil.WriteFile("ig_user.json", f, 0644)
}

//test server
// func MainServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "okok")
// }

func main() {

	//http.HandleFunc("/", MainServer)
	//http.ListenAndServe(":8080", nil)
	InstagramScraper()

}
