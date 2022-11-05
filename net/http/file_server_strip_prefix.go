package main

import "net/http"

func main() {
	// To serve a directory one disk (/tmp) under an alternate URL
	// path (/tmpfile/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
}
