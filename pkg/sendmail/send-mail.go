package sendmail

import (
	"net/smtp"
	llog "github.com/lypwonderful/llog"
	"send-pay-emali/pkg/readfile"
)
func body(name ,pay ,rmark string) string{
	mailbody:=`<table border="1">
		<tr>
			<th>Name</th>
			<th>Pay</th>
			<th>Remark</th>
		</tr>
	   <tr>
		<td>`+name+`</td>
		<td>`+pay+`</td>
		<td>`+rmark+`</td>
	   </tr>
   </table>`
	return mailbody
}
func getMsg(strMsg string) []byte {
	contentType := "Content-Type: text/" + "html" + "; charset=UTF-8"
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		contentType+
		"\r\n\r\n" +
		strMsg+"\r\n")
	return msg
}
func SendToEmail(mailStruct readfile.XlsxToMailStruct){
	to :=make([]string,1)
	auth := smtp.PlainAuth("", "lypwonderful@163.com", "994828337lyp", "smtp.163.com")
	for keyMail,valStruct:=range mailStruct.MMap{
		to[0]=keyMail
		msg :=getMsg(body(valStruct.Name,valStruct.Pay,valStruct.Remarks))
		if err := smtp.SendMail("smtp.163.com:25", auth, "lypwonderful@163.com", to, msg);err!=nil{
			llog.Infof("SendMail Msg To %s Error:%s",to[0],err)
		}else{
			llog.Infof("== SendMail Msg To %s Succ ==",to[0])
		}
	}
}
