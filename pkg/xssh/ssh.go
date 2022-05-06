package xssh

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

type xssh struct {
}

var host string
var port string
var user string
var command string
var keyPath string
var keyPass string

func readPubKey(file string) ssh.AuthMethod {
	var key ssh.Signer
	var err error
	var b []byte
	b, err = ioutil.ReadFile(file)
	mustExec(err, "failed to read public key")
	if !strings.Contains(string(b), "ENCRYPTED") {
		key, err = ssh.ParsePrivateKey(b)
		mustExec(err, "failed to parse private key")
	} else {
		key, err = ssh.ParsePrivateKeyWithPassphrase(b, []byte(keyPass))
		mustExec(err, "failed to parse password-protected private key")
	}
	return ssh.PublicKeys(key)
}

func main() {
	conf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			readPubKey(keyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // XXX: Security issue
	}
	client, err := ssh.Dial("tcp", strings.Join([]string{host, ":", port}, ""), conf)
	mustExec(err, "failed to dial SSH server")
	session, err := client.NewSession()
	mustExec(err, "failed to create SSH session")
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(command)
	mustExec(err, "failed to run command over SSH")
	log.Printf("%s: %s", command, b.String())
}

func mustExec(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:\n  %s", msg, err)
	}
}
