package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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

func GetLatestOTA(otaPath string, otaFormat string, fileFormat string) {
	cmd := exec.Command("curl", "-u", "readonly:readonly", otaPath)
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to download", otaPath, err)
	}
	ret := string(data)
	// log.Println("GetLatestOTA", ret)
	lines := strings.Split(ret, "\n")
	var maxVersion int
	maxVersion = 0
	for _, line := range lines {
		if strings.HasPrefix(line, "<a href=\"") {
			log.Println("->", GetBetweenStr(line, "<a href=\"", "/\">")[9:])
			verStr := GetBetweenStr(line, "<a href=\"", "/\">")[9:]
			var version int
			fmt.Sscanf(verStr, otaFormat, &version)
			log.Println("->", version)
			if version > 0 && version > maxVersion {
				maxVersion = version
			}
		}
	}
	otaUrl := fmt.Sprintf("%s"+otaFormat+"/otas/user/"+fileFormat, otaPath, maxVersion, maxVersion)
	log.Println(otaUrl)
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
