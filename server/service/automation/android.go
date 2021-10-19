package automation

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/automation"
)

type AndroidService struct {
}

func (androidService *AndroidService) GetLatestOTA(req automation.LatestOtaRequest) (err error, rsp automation.LatestOtaResponse) {
	cmd := exec.Command("curl", "-u", "readonly:readonly", req.OtaPath)
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to download", req.OtaPath, err)
		return
	}
	ret := string(data)
	lines := strings.Split(ret, "\n")
	var maxVersion int
	maxVersion = 0
	for _, line := range lines {
		if strings.HasPrefix(line, "<a href=\"") {
			verStr := GetBetweenStr(line, "<a href=\"", "/\">")[9:]
			var version int
			_, err := fmt.Sscanf(verStr, req.OtaFormat, &version)
			if err == nil {
				if version > 0 && version > maxVersion {
					maxVersion = version
				}
			} else {
				continue
			}
		}
	}
	otaUrl := fmt.Sprintf("%s"+req.OtaFormat+"/otas/user/"+req.FileFormat, req.OtaPath, maxVersion, maxVersion)
	rsp.LatestVersion = maxVersion
	rsp.LatestOtaUrl = otaUrl
	return
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
