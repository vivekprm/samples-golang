package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseEnvironment() {
	//useGlobalEnv := true
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Printf("Environment Variable file (.env) is not present.  Relying on Global Environment Variables")
		//useGlobalEnv = false
	}

	setEnvVariable("CLIENT_ID", os.Getenv("CLIENT_ID"))
	setEnvVariable("SPA_CLIENT_ID", os.Getenv("SPA_CLIENT_ID"))
	setEnvVariable("ISSUER", os.Getenv("ISSUER"))

	if os.Getenv("CLIENT_ID") == "" {
		log.Printf("Could not resolve a CLIENT_ID environment variable.")
		os.Exit(1)
	}

	if os.Getenv("SPA_CLIENT_ID") == "" {
		log.Printf("Could not resolve a SPA_CLIENT_ID environment variable.")
		os.Exit(1)
	}

	if os.Getenv("ISSUER") == "" {
		log.Printf("Could not resolve a ISSUER environment variable.")
		os.Exit(1)
	}

}
func setEnvVariable(env string, current string) {
	if current != "" {
		return
	}

	file, _ := os.Open(".env")
	defer file.Close()

	lookInFile := bufio.NewScanner(file)
	lookInFile.Split(bufio.ScanLines)

	for lookInFile.Scan() {
		parts := strings.Split(lookInFile.Text(), "=")
		key, value := parts[0], parts[1]
		if key == env {
			os.Setenv(key, value)
		}
	}
}
