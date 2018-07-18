package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"fmt"
	//"path"
	"io/ioutil"
	"html/template"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {

	tmpl += ".html"

	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", e)

				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))

			}
		}()

		fn(w, r)

	}
}

var templates map[string]*template.Template

func init() {

	templates = make(map[string]*template.Template)

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading templates:", templatePath)

		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		//		s := `
		//	<form method="POST" action="upload" enctype="multipart/form-data">
		//Choose an image to upload: <input name="image" type="file"/>
		//<input type="submit" value="Upload"/>
		//</form>`
		//		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//		io.WriteString(w, s)
		//		//fmt.Fprintf(w, s)
		//		//w.
		//		//w.Write([]byte(s))
		//		//template.New('webpage')

		//t,err:=template.ParseFiles("upload.html")
		//if err!=nil{
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//
		//t.Execute(w, nil)
		//return

		renderHtml(w, "upload", nil)

	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)

		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "view?id="+filename, http.StatusFound)
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//log.Println(r.Form)
		//log.Println(r.URL.Query())
		//id := r.URL.Query().Get("id")
		////log.Println(typeof(ids))
		////if len(ids) == 0 {
		////	http.Error(w, "error id", http.StatusInternalServerError)
		////	return
		////}
		////log.Println(ids[0], typeof(ids[0]))
		////filename := string(ids[0])
		////log.Println(filename)
		//s := `<img src="uploads/` + id + `" />`
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		////io.WriteString(w, s)
		//fmt.Fprintf(w, s)

		imageId := r.FormValue("id")
		imagePath := UPLOAD_DIR + "/" + imageId

		if exits := isExists(imagePath); !exits {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "image")
		http.ServeFile(w, r, imagePath)

	}

}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	//var listHtml string
	//
	//for _, fileInfo := range fileInfoArr {
	//	imgId := fileInfo.Name()
	//	listHtml += "<li><a href=\"/view?id=" + imgId + "\">" + imgId + "</a></li>"
	//}
	//
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//io.WriteString(w, "<ol>"+listHtml+"</ol>")

	locals := make(map[string]interface{})
	images := []string{} // 创建一个空的切片
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	fmt.Println(images)
	locals["images"] = images
	//t, err:= template.ParseFiles("list.html")
	//if err!=nil{
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//t.Execute(w, locals)

	renderHtml(w, "list", locals)

}

//func staticHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("in static handler...")
//	fileServer.ServeHTTP(w, r)
//}

//var fileServer http.Handler

func main() {

	//fileServer = http.FileServer(http.Dir(path.Dir(os.Args[0])))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/", safeHandler(listHandler))

	//http.HandleFunc("/", staticHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
