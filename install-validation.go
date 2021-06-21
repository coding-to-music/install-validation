//
// https://stackoverflow.com/questions/19439430/go-golang-traverse-through-struct/19439503
//
package main

import (
	"fmt"
	"os/exec"
)

type InstallItems struct {
	run            bool
	attemptInstall bool
	desc           string
	installcmd     string
	validatecmd    string
}

// OPTIONS="-y -q"

var InstallItemsCity = []InstallItems{
	{true, true, "always good to update the package manager apt",
		"sudo apt update",
		"sudo apt update"},
	{true, true, "echo $PATH", "echo $PATH", "echo $PATH"},
	{false, true, "whoami", "whoami", "whoami"},
	// {false, true, "ansible", "sudo apt install -y ansible", "ansible --version"},
	// {false, true, "docker as sudo", "sudo docker run hello-world", "sudo docker run hello-world"},
	// {false, true, "docker as regular-user", "docker run hello-world", "docker run hello-world"},
	// {false, true, "tree", "sudo apt install -y tree", "tree --version"},
	// {false, true, "ncdu", "sudo apt install -y ncdu", "ncdu -v"},
	// {false, true, "vscode", "code --version", "code --version"},
	// {false, true, "kind", "kind --version", "kind --version"},
	// {false, true, "doctl", "doctl version", "doctl version"},
	// {false, true, "python3", "sudo apt install -y python3", "python3 --version"},
	// {false, true, "see how we are configured with GitHub", "git config --list", "git config --list"},
	{true, true, "Check that GitHub can be reached via SSH", "ssh -vT git@github.com", "ssh -vT git@github.com"},
	// {true, true, "NodeJS", "sudo apt install -y node", "node --version"},
	// {true, true, "python", "sudo apt install -y python", "python --version"},

	// {true, true, "kubectl 0", "kubectl version", "kubectl version"},
	// {true, true, "kubectl 1: Update the apt package index and install packages needed to use the Kubernetes apt repository", "sudo apt-get install -y apt-transport-https ca-certificates curl", "kubectl version"},
	// {true, true, "kubectl 2: Download the Google Cloud public signing key", "sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg", "kubectl version"},
	// {true, true, "kubectl 3: Add the Kubernetes apt repository", "echo 'deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main' | sudo tee /etc/apt/sources.list.d/kubernetes.list", "kubectl version"},
	// {true, true, "kubectl 4: Update apt package index with the new repository and install kubectl", "sudo apt-get install -y kubectl", "kubectl version"},
	// {true, true, "kubectl 5", "kubectl version", "kubectl version"},

	// {true, true, "golang", "go version", "go version"},
	// {true, true, "whereis go", "whereis go", "whereis go"},
	// {true, true, "test the ansible connection", "ansible -m ping TestClient", "ansible -m ping TestClient"},
}

var (
	ValidateAvailableCount = 0
	ValidateRequestedCount = 0
	ValidateErrorCount     = 0
	ValidatePassedCount    = 0
	InstallAvailableCount  = 0
	InstallRequestedCount  = 0
	InstallErrorCount      = 0
	InstallPassedCount     = 0
)

var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red

	descColor             = Teal
	installcmdColor       = Yellow
	validatecmdColor      = Red
	successColor          = Green
	errorColor            = Red
	InstallRequestedColor = Magenta
	InstallErrorColor     = Purple
	InstallPassedColor    = White

	// fmt.Println(installcmdColor("ValidateAvailableCount = ", ValidateAvailableCount))
	// fmt.Println(validatecmdColor("ValidateRequestedCount = ", ValidateRequestedCount))
	// fmt.Println(successColor("ValidatePassedCount = ", ValidatePassedCount))
	// fmt.Println(errorColor("ValidateErrorCount = ", ValidateErrorCount))

)

// resultingOutputColor = is successColor or errorColor

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// type InstallItems struct {
// run            bool
// attemptInstall bool
// desc           string
// installcmd     string
// validatecmd    string

func validate(data_arr []InstallItems) int {

	errors := 0
	for _, elem := range data_arr {

		ValidateAvailableCount++

		if elem.run {

			ValidateRequestedCount++

			fmt.Println(descColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			fmt.Println(descColor(elem.desc))

			fmt.Println(descColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

			fmt.Println(Yellow(elem.validatecmd))

			out, err := exec.Command("bash", "-cl", elem.validatecmd).Output()

			// cmd.Stderr = cmd.Stdout

			// cmd := exec.Command(elem.validatecmd)
			// cmd := exec.Command(app, elem.installcmd)
			// stdout, err := cmd.Output()

			if err != nil {
				ValidateErrorCount++
				fmt.Println(Red("Error"))
				fmt.Println(Red("err.Error"))
				fmt.Println(Red(err.Error()))
				fmt.Println(Red("string(out)"))
				fmt.Println(Red(string(out)))
				errors++

				// for {
				// 	tmp := make([]byte, 1024)
				// 	_, err := out.Read(tmp)
				// 	fmt.Print(string(tmp))
				// 	if err != nil {
				// 		break
				// 	}
				// }

				if elem.run {

					InstallAvailableCount++
					InstallRequestedCount++

					attemptout, attempterr := exec.Command("bash", "-cl", elem.installcmd).Output()

					if attempterr != nil {
						InstallErrorCount++
						fmt.Println(Red("Error"))
						fmt.Println(Red("err.Error"))
						fmt.Println(Red(attempterr.Error()))
						fmt.Println(Red("string(attemptout)"))
						fmt.Println(Red(string(attemptout)))
					} else {
						InstallPassedCount++
						// Print the output
						fmt.Println(Green("Success"))
						fmt.Println(Green(string(attemptout)))
					}
				}
			} else {
				ValidatePassedCount++

				// Print the output
				fmt.Println(Green("Success"))
				fmt.Println(Green(string(out)))
			}
		}
	}
	return errors
}

//

// 	func print-hello-world(data_arr []InstallItems) int {
// 	fmt.Println(Info("Hello World at the command line from server.go"))
// 	fmt.Println(Info("Invoke with go run main.go"))
// 	fmt.Println("\n")

// 	fmt.Println(Info("hello, Info world!"))
// 	fmt.Println(Warn("hello, Warn world!"))
// 	fmt.Println(Fatal("hello, Fatal world!"))
// 	fmt.Println("\n")

// 	fmt.Println(Black("hello, Black world!"))
// 	fmt.Println(Red("hello, Red world!"))
// 	fmt.Println(Green("hello, Green world!"))
// 	fmt.Println(Yellow("hello, Yellow world!"))
// 	fmt.Println(Purple("hello, Purple world!"))
// 	fmt.Println(Magenta("hello, Magenta world!"))
// 	fmt.Println(White("hello, White world!"))
// 	fmt.Println(Teal("hello, Teal world!"))
// 	fmt.Println(Green("hello, Green need Blue world!"))
// 	fmt.Println("\n")

// 	fmt.Println(descColor("hello, descColor world!"))
// 	fmt.Println(installcmdColor("hello, installcmdColor world!"))
// 	fmt.Println(validatecmdColor("hello, validatecmdColor world!"))
// 	fmt.Println(successColor("hello, successColor world!"))
// 	fmt.Println(errorColor("hello, errorColor world!"))
// 	fmt.Println(InstallRequestedColor("hello, InstallRequestedColor world!"))
// 	fmt.Println(InstallErrorColor("hello, InstallErrorColor world!"))
// 	fmt.Println(InstallPassedColor("hello, InstallPassedColor world!"))

// //	func print-hello-world() {

// 	// = Teal
// 	// installcmdColor  = Yellow
// 	// validatecmdColor = Red
// 	// successColor     = Green
// 	// errorColor       = Red

// 	fmt.Println("\n")

// 	fmt.Println("Hello, playground")
// 	fmt.Println("\n")

// 	content, err := ioutil.ReadFile("/home/tmc/testdata/hello")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("File contents: %s", content)
// 	fmt.Println("\n")

// }
func main() {

	// print - hello - world(InstallItemsCity)
	validate(InstallItemsCity)

	fmt.Println(descColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
	fmt.Println(descColor("Available Items"))

	fmt.Println(descColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

	fmt.Println("\n")
	fmt.Println(installcmdColor("ValidateAvailableCount = ", ValidateAvailableCount))
	fmt.Println(validatecmdColor("ValidateRequestedCount = ", ValidateRequestedCount))
	fmt.Println(successColor("ValidatePassedCount = ", ValidatePassedCount))
	fmt.Println(errorColor("ValidateErrorCount = ", ValidateErrorCount))
	fmt.Println("\n")

	fmt.Println(successColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
	fmt.Println(successColor("Run Results"))

	fmt.Println(InstallRequestedColor("InstallAvailableCount = ", InstallAvailableCount))
	fmt.Println(InstallRequestedColor("InstallRequestedCount = ", InstallRequestedCount))
	fmt.Println(InstallPassedColor("InstallPassedCount = ", InstallPassedCount))
	fmt.Println(InstallErrorColor("InstallErrorCount = ", InstallErrorCount))
	fmt.Println(successColor("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

}
