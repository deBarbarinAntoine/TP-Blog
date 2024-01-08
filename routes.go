package TPBlog

import "net/http"

// routes initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/album/jul", albumJulHandler)
	http.HandleFunc("/track/sdm", trackSdmHandler)
}
