package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func serveGithubWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	logError(err)
	secret := []byte(os.Getenv("SECRET"))
	message := []byte(body)
	hash := hmac.New(sha1.New, secret)
	hash.Write(message)

	hashResult := hex.EncodeToString(hash.Sum(nil))

	remoteHash := r.Header.Get("X-Hub-Signature")

	if ("sha1=" + hashResult) == remoteHash {
		go exeCmd(os.Getenv("COMMAND"), os.Getenv("ARGUMENTS"), os.Getenv("COMMANDPATH"))
		fmt.Println("Hashes matched, ran command!")
	} else {
		fmt.Println("Hashes didn't match, didn't run the command either!")
	}

	fmt.Fprintf(w, hashResult)
}

func serveWebhook(w http.ResponseWriter, r *http.Request) {
	go exeCmd(os.Getenv("COMMAND"), os.Getenv("ARGUMENTS"), os.Getenv("COMMANDPATH"))
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func exeCmd(cmd string, args string, dir string) {
	command := exec.Command(cmd, args)
	command.Dir = dir
	out, _ := command.Output()
	fmt.Println(string(out))
}

func main() {
	// Normal webhook needs a security check
	//http.HandleFunc("/", serveWebhook)
	//http.HandleFunc("/github", serveGithubWebhook)
	
	http.HandleFunc("/", serveGithubWebhook)
	http.HandleFunc("/github", serveGithubWebhook)
	log.Fatal(http.ListenAndServe(":80", nil))
}
