package context

import (
	"regexp"

	"honnef.co/go/tools/pattern"
)
var supiRegex = `^(imsi-[0-9]{5,15}|nai-.+|gci-.+|gli-.+|.+)$`  
var peiRegex = `^(imei-[0-9]{15}|imeisv-[0-9]{16}|mac((-[0-9a-fA-F]{2}){6})(-untrusted)?|eui((-[0-9a-fA-F]{2}){8})|.+)$`

var verifySupi(pattern, supi string) bool {

}