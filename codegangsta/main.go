package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/blog", blog)
    r.HandleFunc("/pricing", pricing)

    dash := r.PathPrefix("/dashboard").Subrouter()
    dash.HandleFunc("/", dashboardIndex)
    dash.HandleFunc("/things", dashboardThings)
    dash.HandleFunc("/things/crush", dashboardCrushAllTheThings)

    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/things", apiThings)
    api.HandleFunc("/things/crush", apiCrushAllTheThings)

    sirMuxalot := http.NewServeMux()
    sirMuxalot.Handle("/", r)
    sirMuxalot.Handle("/api/", negroni.New(
        negroni.HandlerFunc(APIMiddleware),
        negroni.Wrap(r),
    ))
    sirMuxalot.Handle("/dashboard/", negroni.New(
        negroni.HandlerFunc(DashboardMiddleware),
        negroni.Wrap(r),
    ))

    n := negroni.Classic()
    n.UseHandler(sirMuxalot)
    http.ListenAndServe(":3000", n)
}

func blog(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Bro, do you even code?</h1>")
}

func pricing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<p>Get access to the code crusher API for just $100/mo!!!</p><p>Don't have $100? You should get a job, but don't <a href=\"https://www.youtube.com/watch?v=l8ZJu-f-XOE\">take er jerbs<a>.</p>")
}

func dashboardIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<p>See your crusher usage and your API keys here.</p>")
}

func dashboardThings(w http.ResponseWriter, r *http.Request) {

    time.Sleep(10 * time.Second)
    fmt.Fprintf(w, "<h1>Here are your things:</h1> <ul><li><img src=\"http://devstickers.com/assets/img/pro/5j1q.png\"></li><li><img src=\"http://assets.bwbx.io/images/irJV4Lf1p7pI/v1/-1x-1.jpg\"></li><li><img src=\"http://assets.bwbx.io/images/izimLvFvzUIE/v1/-1x-1.jpg\"></li><li><img src=\"http://assets.bwbx.io/images/iMU90m8Aw6zA/v1/-1x-1.jpg\"></li></ul>")
}

func dashboardCrushAllTheThings(w http.ResponseWriter, r *http.Request) {
    // This is where we would normally crush all the things, but I'm not about to
    // share those secrets with you.
    fmt.Fprintf(w, "<h1>You crushed a thing!</h1>")
}

func apiThings(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("[{\"image\": \"http://devstickers.com/assets/img/pro/5j1q.png\"},{\"image\": \"http://assets.bwbx.io/images/irJV4Lf1p7pI/v1/-1x-1.jpg\"},{\"image\": \"http://assets.bwbx.io/images/izimLvFvzUIE/v1/-1x-1.jpg\"},{\"image\": \"http://assets.bwbx.io/images/iMU90m8Aw6zA/v1/-1x-1.jpg\"}]"))
}

func apiCrushAllTheThings(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("{\"message\": \"This feature is classified.\"}"))
}

func DashboardMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("The klout API is down. F it, let them in. YOLO SWAG 420 BLAZE IT!")
    w.Write([]byte("The klout API is down. F it, let them in. YOLO SWAG 420 BLAZE IT!"))
    next(w, r)
    w.Write([]byte("After processing"))
}

func APIMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("IP checks out.")
    next(w, r)
}
