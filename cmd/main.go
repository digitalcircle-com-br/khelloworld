package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		buf := bytes.Buffer{}
		_ = request.Write(&buf)
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		bs := buf.Bytes()
		_, _ = writer.Write([]byte(fmt.Sprintf("Got request at %v\n", time.Now())))
		_, _ = writer.Write([]byte("=============\n"))
		_, _ = writer.Write(bs)
		_, _ = writer.Write([]byte("=============\n"))
		_, _ = writer.Write([]byte(fmt.Sprintf("Working Dir: %s\n", wd)))
		_, _ = writer.Write([]byte(fmt.Sprintf("User: %s\n", usr.Name)))
		_, _ = writer.Write([]byte(fmt.Sprintf("User Home Dir: %s\n", usr.HomeDir)))
		_, _ = writer.Write([]byte("Env\n"))
		env := os.Environ()
		for _, s := range env {
			_, _ = writer.Write([]byte("    * " + s + "\n"))
		}

		println(string(bs) + "\n")

	})
	log.Print("App starting at :8080")
	_ = http.ListenAndServe(":8080", nil)

}
