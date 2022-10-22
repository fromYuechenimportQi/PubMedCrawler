package src

import (
	"fmt"
	"github.com/xuthus5/BaiduTranslate"
	"strings"
)

// 申请的信息
func (this *PaperInfo) BaiduTranslate(appID, secretKey string) {
	fmt.Printf("Translating:%v .....\n", this.Title)
	bi := BaiduTranslate.BaiduInfo{AppID: appID, Salt: BaiduTranslate.Salt(5), SecretKey: secretKey, From: "en", To: "zh"}
	str := strings.Replace(this.Content, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	bi.Text = str
	this.Translate = bi.Translate()
}

func CheckValidation(appID, secretKey string) string {
	bi := BaiduTranslate.BaiduInfo{AppID: appID, Salt: BaiduTranslate.Salt(5), SecretKey: secretKey, From: "en", To: "zh"}
	return bi.Translate()
}
