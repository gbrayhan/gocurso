## Server en go  
Go nos permite levantar servidor con librerias que estan dentro del skd, la implementacion es bantante sencilla como se muestra a continuacion:  

```golang  
package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Servidor Go")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080",nil)
}
```  

## Consumiendo una api en **go**
  
Go nos permite usar diferentes librerias para leer una api que retorna valores json de una manera nativa.   

```golang
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type Post struct {
	UserId int `json : "userId"`
	Id int `json : "id"`
	Title string `json : "title"`
	Body string `json : "body"`
}

func main() {
	var cliente = &http.Client{Timeout: 10 * time.Second}
	var url = "https://jsonplaceholder.typicode.com/posts"

	resp, err := cliente.Get(url)
	if err != nil {
		panic(err.Error())
	}
	var usuarios []Post
	err = json.NewDecoder(resp.Body).Decode(&usuarios)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(usuarios[0].Body)
}
```  