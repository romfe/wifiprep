package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var inter = flag.String("i", "noInt", "Interface to execute the tests")

func verifyError(out []byte, err error) {

	if err != nil {
		fmt.Println("[-] Error, exiting program D:")
		log.Fatal(err)
	} else {
		fmt.Println(string(out))
	}

}

func execute() {
	fmt.Println("[+] Bringing the interface down...")
	output, erro := exec.Command("bash", "-c", "sudo ifconfig "+*inter+" down").Output()
	verifyError(output, erro)
	fmt.Println("[+] Changing MAC address...")
	output, erro = exec.Command("bash", "-c", "sudo macchanger -A "+*inter+" down").Output()
	verifyError(output, erro)
	fmt.Println("[+] Entering monitor mode...")
	output, erro = exec.Command("bash", "-c", "sudo iwconfig "+*inter+" mode monitor").Output()
	verifyError(output, erro)
	fmt.Println("[+] Bringing the interface back up...")
	output, erro = exec.Command("bash", "-c", "sudo ifconfig "+*inter+" up").Output()
	verifyError(output, erro)
	fmt.Println("[+] All done, good luck ;)")
}

func init() {

	flag.Parse()

	if runtime.GOOS != "linux" {
		fmt.Println("[-] Dude/dudette, you gotta run this in a debian compatible Linux distro")
		fmt.Println("")
		os.Exit(1)
	}

	if os.Geteuid() != 0 {
		fmt.Println("[-] Dude/dudette, you gotta run this as root. Try sudo...")
		fmt.Println("")
		os.Exit(2)
	}

	if *inter == "noInt" {
		fmt.Println("[-] Please assign a wifi interface using -i ")
		os.Exit(3)
	}

	fmt.Println("[+] Starting...")
	fmt.Println("")
}

func main() {

	execute()

}
