package ambrosio

import (
	"fmt"
	"net/http"
    "regexp"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

type Listener interface {
	Listen(port int)
    NewAction()
}

type Ambrosio struct {
	name    string
    router  *mux.Router
	negroni *negroni.Negroni
    Behaviours []Behaviour
}

type Behaviour struct {
    Pattern string
    Handler func([]string) (string, bool)
}

type Handler func(http.ResponseWriter, *http.Request)

func NewAmbrosio(name string) *Ambrosio {
    /*helpBehaviour := Behaviour {
        "/help",
        func(matches []string, w http.ResponseWriter, req *http.Request) {
            fmt.Println("help")
        },
    }*/
	ambrosio := Ambrosio{
		name,
        mux.NewRouter(),
		negroni.Classic(),
        []Behaviour{},//helpBehaviour inside brackets
	}
	return &ambrosio
}

func (a Ambrosio) Listen(port int) {

	a.router.HandleFunc("/ask", func(w http.ResponseWriter, req *http.Request) {
        var handledOnce = false
		var actionResult string
		var actionError bool
        action := req.FormValue("action")

        for _, b := range a.Behaviours {

            matched, _ := regexp.MatchString(b.Pattern, action)

            if (matched == true) {
				matches := regexp.MustCompile(b.Pattern).FindStringSubmatch(action)
                handledOnce = true
				actionResult, actionError = b.Handler(matches)
            }

        }
        if(!handledOnce) {
            fmt.Fprintf(w, "Unkown command")
        } else if(actionError == false){
			fmt.Fprintf(w, actionResult)
		} else {
			fmt.Fprintf(w, actionResult)
		}

	})

	// start the server
	n := negroni.Classic()
	n.UseHandler(a.router)
	n.Run(fmt.Sprintf(":%d", port))
}

func (a Ambrosio) NewBehaviour(b Behaviour) []Behaviour{

    a.Behaviours = append(a.Behaviours, b)
    return a.Behaviours
}
