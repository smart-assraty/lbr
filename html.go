package main

import (
	"html/template"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/pkg/sftp"
)

type HTML struct {
	domName      string
	templateName string
	location     string
}

func upload(docName string, location string) (err error) {
	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password("1q2w3e$R"))
	config := ssh.ClientConfig{User: "root", Auth: auths, HostKeyCallback: ssh.InsecureIgnoreHostKey()}

	conn, connErr := ssh.Dial("tcp", "172.19.0.3:22", &config)
	checkError("connErr", connErr)
	client, errClient := sftp.NewClient(conn)
	checkError("errClient", errClient)

	srcFile, srcErr := os.Open("html/" + docName)
	checkError("srcFile", srcErr)
	defer srcFile.Close()
	checkError("Mkdir", client.MkdirAll("/var/www/html/"+location))
	dstFile, dstErr := client.Create("/var/www/html/" + location + docName)
	checkError("file", dstErr)
	defer dstFile.Close()

	dstFile.ReadFrom(srcFile)
	return nil
}

func buildFromTemplate[T any](users T, html HTML) {
	userTemplate, errTemplate := template.ParseFiles("./templates/" + html.templateName + ".gohtml")
	checkError("errTemplate", errTemplate)

	fileName := strings.ToLower(html.domName)
	file, errFile := os.Create("html/" + fileName)
	checkError("errFile:", errFile)

	execute1 := userTemplate.Execute(file, users)
	checkError("Execute1", execute1)

	upload(fileName, html.location)
	checkError("Remove", os.Remove("html/"+fileName))

	defer file.Close()
}
