package ext

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
)

func getBitcoinCliPath() string {
	return os.Getenv("BITCOIN_CLI")
}

func IsValidBTCAddress(address string) bool {
	cmd := exec.Command(getBitcoinCliPath(), "getaddressinfo", address)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "exit status 1" {
			log.Println("bitcoin cli not running ")
		} else if err.Error() == "exit status 5" {
			log.Println("address not valid")
		} else {
			log.Println("bitcoin deamon not installed or other unexpected error")
		}
		return false
	}
	var dat map[string]interface{}
	err = json.Unmarshal(out, &dat)

	if err != nil {
		log.Fatal("error while unmarshaling json, btc address invalid?")
		return false
	}
	if dat["scriptPubKey"] == "" {
		log.Println("unexepected error whlie parsing address")
		return false
	}

	return true
}
