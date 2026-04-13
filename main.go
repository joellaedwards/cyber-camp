package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running on port: 8080")

	router := http.NewServeMux()
	router.Handle("/styles.css", http.FileServer(http.Dir(".")))
	router.HandleFunc("/", welcomeHandler)
	router.HandleFunc("/level1", page1Handler)
	router.HandleFunc("/checklevel1", checkLevel1Handler)
	router.HandleFunc("/level2-yup", page2Handler)
	router.HandleFunc("/coolOffTrailStuff", page3Handler)
	router.HandleFunc("/donezo", successHandler)

	http.ListenAndServe(":8080", router)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in welcome function getting html")
	html := welcomeHtml()
	fmt.Fprintln(w, html)
}

func page1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in page1 function getting html")
	html := page1Html("")
	fmt.Fprintln(w, html)
}

func page2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in page2 function getting html")
	html := page2Html()
	fmt.Fprintln(w, html)
}

func page3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in page3 function getting html")
	accessCookie, _ := r.Cookie("access")
	if accessCookie != nil && accessCookie.Value == "granted" {
		http.Redirect(w, r, "/donezo", http.StatusSeeOther)
	}
	html := page3Html()
	fmt.Fprintln(w, html)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in success function getting html")
	html := success()
	fmt.Fprintln(w, html)
}

func checkLevel1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in checkLevel1 function")
	r.ParseForm()
	_, option1 := r.Form["phishing_1"]
	_, option2 := r.Form["phishing_2"]
	_, option3 := r.Form["phishing_3"]
	_, option4 := r.Form["phishing_4"]

	var correct bool
	// 1 and 3 are the phishing ones
	if option1 && option3 {
		correct = true
	} else {
		correct = false
		html := page1Html("<p style=\"color: white;\">Try again. Some of your selections were incorrect.</p>")
		fmt.Fprintln(w, html)
		return
	}

	if option2 || option4 {
		correct = false
		html := page1Html("<p style=\"color: white;\">Try again. Some of your selections were incorrect.</p>")
		fmt.Fprintln(w, html)
		return
	}
	if correct {
		http.Redirect(w, r, "/level2-yup", http.StatusSeeOther)
	}
}