package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	APIRequests()
	new := ArtistStruct{Tab: ArtistTab}
	tmpl.Execute(w, new)

}


func Artiste(w http.ResponseWriter, r *http.Request){
	fmt.Println("eeeeeee")
	http.ServeFile(w, r, "tmpl/artist.html")
}