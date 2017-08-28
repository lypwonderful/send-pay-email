package main

import (
	llog "github.com/lypwonderful/llog"
	"send-pay-emali/pkg/readfile"

	"send-pay-emali/pkg/sendmail"
)

var (
	logPath  = "F:/goproject/src/send-pay-emali/output"
	filePath = "F:/goproject/src/send-pay-emali/file/2017827.xlsx"
)

func logInit() {
	llog.SetLogPath(logPath)
	llog.Init(0, "INFO")

}
func main() {
	logInit()
	llog.Infoln("====== start pay-email =====")

	xlxsMail := readfile.XlsxToMailStruct{
		FilePath: filePath,
	}
	xlxsMail.ReadXlsxFile()
	xlxsMail.GetMemberContentDetail()
	sendmail.SendToEmail(xlxsMail)
	llog.Flush()
}
