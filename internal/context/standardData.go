package context

import (
	"regexp"

	"github.com/Thanhlongnguyen1012/5GC_smf/internal/models"
)

var supiRegex = "^(imsi-[0-9]{5,15}|nai-.+|gci-.+|gli-.+|.+)$"
var peiRegex = "^(imei-[0-9]{15}|imeisv-[0-9]{16}|mac((-[0-9a-fA-F]{2}){6})(-untrusted)?|eui((-[0-9a-fA-F]{2}){8})|.+)$"

func VerifySupi(supi string) bool {
	re, _ := regexp.Compile(supiRegex)
	return re.MatchString(supi)
}
func VerifyPei(pei string) bool {
	re, _ := regexp.Compile(peiRegex)
	return re.MatchString(pei)
}
func VerifyAccessType(accessType models.AccessType) bool {
	return accessType == models.AccessType__3_GPP_ACCESS
}
