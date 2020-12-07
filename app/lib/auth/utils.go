package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func OpenOauthUrl(url string) error {

	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "openbsd":
		fallthrough
	case "linux":
		cmd = exec.Command("xdg-open", url)
	}

	if cmd != nil {
		cmd.Stdout = os.Stdout
		err := cmd.Start()

		if err != nil {
			log.Printf("Failed to open Browser %s", err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Printf("Failed to open browser %s", err)
		}
		return nil
	}
	return nil
}

func checkBrowserPort(port int) bool {
	l, err := net.Listen("tcp", ":"+string(rune(port)))

	if l != nil {
		defer l.Close()
		return true
	}

	if err != nil {
		return true
	}
	return false
}

func GetRedirectPort() int {
	isPortOpen := false
	port := 4000
	for !isPortOpen {
		if checkBrowserPort(port) {
			isPortOpen = true
		} else {
			port = port + 1
		}
	}
	return port
}

func GetOauthUrl(port int) string {

	oauthUrl := fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?"+
		"redirect_uri=%s&"+
		"prompt=select_account&"+
		"response_type=code&"+
		"client_id=%s&"+
		"scope=https://www.googleapis.com/auth/spreadsheets+https://www.googleapis.com/auth/userinfo.email&"+
		"access_type=offline", "http://localhost:"+strconv.Itoa(port)+"/oauth", "CLIENT_ID")

	log.Printf("opening in browser %s", oauthUrl)
	return oauthUrl
}

func SaveCredentialsToFile(response []byte) error {
	userHomeDirectory, err := os.UserHomeDir()

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(userHomeDirectory+"/.fuse-wfh", response, 0644)

	if err != nil {
		return err
	}

	fmt.Println("Saved credentials")
	return nil
}

func ParseCredentials() Credentials {
	var credentials Credentials
	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {

		err = errors.New("please login first")
	}
	contents, err := ioutil.ReadFile(homeDir + "/.fuse-wfh")

	if err != nil {
		err = errors.New("please login first")
	}

	err = json.Unmarshal(contents, &credentials)
	if err != nil {
		err = errors.New("error parsing user credentials")
	}

	return credentials
}
