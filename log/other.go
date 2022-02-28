package log

import (
	"github.com/Mengdch/goUtil/FileTools"
	"github.com/linxGnu/goseaweedfs"
	"net/http"
	"time"
)

const (
	mHost  = ""
	upHost = ""
)

func Download() {
}

func download(download, path string) error {
	if len(download) == 0 || len(path) == 0 {
		return nil
	}
	var req *http.Request
	req, err := http.NewRequest(http.MethodGet, download, nil)
	if err != nil {
		return nil
	}

	body := GetBody(req)
	err = fileFunc.WriteFileByte(path, body)
	Log("download:"+download, "")
	return err
}
func Upload(filePath string) string {
	sw, err := goseaweedfs.NewSeaweed(mHost, []string{upHost}, 160000, &http.Client{Timeout: 30 * time.Second})
	if err != nil {
		Log("upload.AssignFileId", err.Error())
		return ""
	}
	if len(sw.Filers()) > 0 {
		a, e := sw.AssignFileId("", "")
		if e != nil {
			Log("upload.AssignFileId", e.Error())
			return ""
		}
		r, e := sw.Filers()[0].UploadFile(filePath, a.FileID, "", "")
		if e != nil {
			Log("upload.Filer.UploadFile"+filePath+"("+a.FileID+"):", e.Error())
			return ""
		}
		if len(r.Name) == 0 {
			return ""
		}
		return upHost + "/" + a.FileID
	}
	_, fp, err := sw.UploadFile(filePath, "", "")
	if err != nil {
		Log("upload.UploadFile"+filePath+":", err.Error())
		return ""
	}
	return upHost + "/" + fp.FileID
}
