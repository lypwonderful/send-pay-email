package readfile

import (
	llog "github.com/lypwonderful/llog"
	"github.com/tealeg/xlsx"
)

type XlsxToMailStruct struct {
	Content  [][][]string
	FilePath string
	MMap map[string]ContentDetail
}
type ContentDetail struct {
	Name    string
	Pay     string
	Remarks string
	Email   string
}

const mLen = 4


func (xms *XlsxToMailStruct) ReadXlsxFile() (err error) {
	if xms.Content, err = xlsx.FileToSlice(xms.FilePath); err != nil {
		llog.Errorln("FileToSlice Error-", err)
		return
	}
	llog.V(3).Infoln("Read Succ From ",xms.FilePath)
	return err
}
func checkFileSliceIsOK(ckSlic [][]string, memberLen int, firstLineNum int) {
	firstFlag := len(ckSlic[0]) == firstLineNum
	secondFlag := ckSlic[0][0] == "Name"
	if firstFlag && secondFlag {
		if memberLen == 1 {
			llog.Fatal("Empty File!!!")
		}
		llog.Infoln("File Slice is ok,Start Change to mapEmail")
	}
}
func (xms *XlsxToMailStruct) GetMemberContentDetail() {
	xms.MMap = make(map[string]ContentDetail,0)
	for _, val := range xms.Content {
		checkFileSliceIsOK(val, len(val), 4)
		for _, v := range val {
			if len(v) == mLen && v[0] != "" && v[0] != "Name" {
				mStruct := &ContentDetail{
					v[0],
					v[1],
					v[2],
					v[3],
				}
				xms.MMap[v[mLen-1]] = *mStruct
			}

		}

	}
}
