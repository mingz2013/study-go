package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"fmt"
	"path"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		s := `
	<form method="POST" action="upload" enctype="multipart/form-data">
Choose an image to upload: <input name="image" type="file"/>
<input type="submit" value="Upload"/>
</form>`
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//io.WriteString(w, s)
		fmt.Fprintf(w, s)
		//w.
		//w.Write([]byte(s))
		//template.New('webpage')
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			log.Println("create error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "view?id="+filename, http.StatusFound)
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println(r.Form)
		log.Println(r.URL.Query())
		id := r.URL.Query().Get("id")
		//log.Println(typeof(ids))
		//if len(ids) == 0 {
		//	http.Error(w, "error id", http.StatusInternalServerError)
		//	return
		//}
		//log.Println(ids[0], typeof(ids[0]))
		//filename := string(ids[0])
		//log.Println(filename)
		s := `<img src="uploads/` + id + `" />`
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//io.WriteString(w, s)
		fmt.Fprintf(w, s)
	}

}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("in static handler...")
	fileServer.ServeHTTP(w, r)
}

var fileServer http.Handler

func main() {
	fileServer = http.FileServer(http.Dir(path.Dir(os.Args[0])))
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/", staticHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
