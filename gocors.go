package gocors

// Webserver using gorilla mux and negroni.Classic with CORS middleware
import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "fmt"
)

type Ws struct {
    Address string
    Port string
    *mux.Router  // Embedded Gorilla Mux Router
}

func NewCORSRouter(address, port string) *Ws {
    ws := new(Ws)
    ws.Router = mux.NewRouter()
    ws.Address = address
    ws.Port = port
    return ws
}

func (s *Ws) Start(){
    // Create an instance of negroni Classic
    n := negroni.Classic()

    // Employ the CORS middleware at the top level
    n.Use(negroni.HandlerFunc(corsMiddleware))

    // ... then add the router
    n.UseHandler(s.Router)
    n.Run(fmt.Sprintf("%s:%s", s.Address,s.Port))
}

func corsMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    if origin := r.Header.Get("Origin"); origin != "" {
        // Set the access control to allow for the origin of the request
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }

    // For preflight OPTIONS requests return just the headers
    if r.Method == "OPTIONS" {
        return
    }

    // ...otherwise route to the next handler in the stack
    next(w, r)
}
