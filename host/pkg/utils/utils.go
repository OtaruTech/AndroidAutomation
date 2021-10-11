package utils

import (
	"log"
	"os/exec"
)

func DownloadOTA(url string, localPath string) error {
	cmd := exec.Command("wget", "--user", "readonly", "--password", "readonly",
		url,
		"-O", localPath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to download", url, err)
	}
	return err
}

func Scp(local string, remote string, port string) {
	cmd := exec.Command("scp", "-P", port, local, remote)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to scp", err)
	}
}
