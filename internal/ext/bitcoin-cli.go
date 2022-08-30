package ext

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getBitcoinCliPath() string {
	return os.Getenv("BITCOIN_CLI")
}

func getWalletPassword() string {
	return os.Getenv("BITCOIN_CLI_WALLET_PASS")
}

func IsValidBTCAddress(address string) bool {
	cmd := exec.Command(getBitcoinCliPath(), "getaddressinfo", address)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "exit status 1" {
			log.Println("bitcoin cli not running ")
		} else if err.Error() == "exit status 1" {
			log.Println("bitcoin deamon not installed or other unexpected error")
		}
		return false
	}
	var dat map[string]interface{}
	err = json.Unmarshal(out, &dat)

	if err != nil {
		log.Println("error while unmarshaling json, btc address invalid?")
		return false
	}

	if dat["scriptPubKey"] == "" {
		log.Println("unexepected error whlie parsing address")
		return false
	}

	return true
}

func UnlockWallet() bool {
	cmd := exec.Command(getBitcoinCliPath(), "unlockwallet", getWalletPassword())
	out, err := cmd.CombinedOutput()

	if err != nil {
		if err.Error() == "exit status 1" {
			log.Println("bitcoin cli not running ")
		} else if err.Error() == "exit status 1" {
			log.Println("bitcoin deamon not installed or other unexpected error")
		}
		return false
	}
	log.Println(string(out))
	return true
}

func GetBtcBalance() float64 {
	cmd := exec.Command(getBitcoinCliPath(), "getbalance")
	out, err := cmd.Output()

	if err != nil {
		if err.Error() == "exit status 1" {
			log.Println("bitcoin cli not running ")
		} else if err.Error() == "exit status 1" {
			log.Println("bitcoin deamon not installed or other unexpected error")
		}
		return 0.0
	}

	balance, err := strconv.ParseFloat(strings.TrimSuffix(string(out), "\n"), 32)

	if err != nil {
		log.Println(err)
		log.Println("unable to parse bitcoin balance")
		return 0
	}
	return balance
}
