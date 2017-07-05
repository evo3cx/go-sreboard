package bukalapak

import (
	"bytes"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

func handleErr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func HandleHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	homeText := `
		Simmple Service run By Golang,
		Service ini akan semua perintah pada terminal melalui param URL

		Contoh:
			127.0.0.1:8080/ls_-la

			output: akan menghasilkan list file/directory pada folder server

	`
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(homeText))
}

func HandleService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := ps.ByName("command")
	cmd, args := splitWithCommand(params)

	var data bytes.Buffer
	var err error

	w.Header().Set("Content-Type", "text/plain")
	//if user request exec top, create long polling
	if cmd == "top" {
		for {
			data, err = runCommand(cmd, args...)
			if err != nil {
				handleErr(w, err)
				return
			}
			io.WriteString(w, data.String())
			time.Sleep(time.Second * 2)
		}
	}

	data, err = runCommand(cmd, args...)
	if err != nil {
		handleErr(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data.Bytes())
}

func runCommand(command string, parrams ...string) (bytes.Buffer, error) {
	cmd := exec.Command(command, parrams...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return out, err
	}

	return out, nil
}

func splitWithCommand(params string) (command string, args []string) {
	st := strings.Split(params, "_")
	command = st[0]

	if command == "top" {
		args = []string{"-b", "-n", "1"}
		return
	}

	for i := 1; i < len(st); i++ {
		args = append(args, st[i])
	}
	return
}
