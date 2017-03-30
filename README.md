# GoCORS
A simple webserver that allows for Cross Origin Requests in a development environment. Basically CORS middleware handled by codegangsta's negroni.Classic() and a gorilla/mux Router. 

### Dependencies

Ensure you have the following packages installed using 'go get':
- github.com/urfave/negroni
- github.com/gorilla/mux

### Usage

 -  NewCORSRouter(address, port string) *CORSRouter

    Returns a new CORSRouter Instance, with gorilla Router functionality embedded
    See Gorilla Mux [http://www.gorillatoolkit.org/pkg/mux] for details

 -  CORSRouter

    A struct with three exported fields: Address, Port and an embedded mux.Router

 - func (*CORSRouter) Start()

   Starts the Router using the user-defined Address and Port

### Example

    import (
        "net/http"
        cors "github.com/UcheEke/GoCORS"
    )

    func rootHandler(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Welcome to the CORS Server!"))
    }

    func main(){
        cr := cors.NewCORSRouter("","8080")
        cr.HandleFunc("/", rootHandler)
        cr.Start()
    }


