package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"github.com/Mengdch/browser"
	"github.com/Mengdch/browser/log"
	"github.com/Mengdch/goUtil/FileTools"
	"github.com/Mengdch/goUtil/TypeTools"
	"go.etcd.io/bbolt"
	"path/filepath"
	"time"
)

var userAgent string

func main() {
	defer log.CatchPanic("main")
	url := flag.String("url", "https://www.baidu.com", "链接")
	title := flag.String("title", "aa", "标题")
	ico := flag.String("icon", "", "图标")
	dev := flag.String("dev", "", "调试目录")
	ua := flag.String("ua", "", "UserAgent")
	max := flag.Bool("max", true, "初始最大化")
	width := flag.Int("width", 1600, "宽")
	height := flag.Int("height", 900, "高")
	flag.Parse()
	userAgent = *ua
	if len(*url) == 0 {
		return
	}
	jsFunc := map[int32]func(string) string{
		1: func(sha string) string {
			s := string(findOne(sha, keyTable))
			return s
		},
	}
	err := browser.StartFull(*url, *title, *ico, userAgent, *dev, *max, true, true, *width, *height, finish, save, jsFunc, nil, nil)
	if err != nil {
		log.Log(*title+":"+*url, err.Error())
	}
}

func finish(url string, success bool) {
	if !success {
		return
	}
	name := getUrl(url)
	if len(name) == 0 {
		return
	}
	val := fileFunc.ReadFileByte(name)
	sum256 := sha256.Sum256(val)
	sha := hex.EncodeToString(sum256[:])
	saveSha(sha, name, url, len(val))
}
func save(url, path string) {
	saveUrl(url, path)
}

var db *bbolt.DB

func init() {
	name := filepath.Join(fileFunc.NowPath(), "data")
	var err error
	db, err = bbolt.Open(name, 0666, nil)
	if err != nil {
		log.Log("initdb:"+name, err.Error())
		return
	}
}

const (
	urlTable = "url"
	keyTable = "key"
)

type shaData struct {
	Url    string    `json:"url"`
	File   string    `json:"file"`
	Len    int       `json:"len"`
	Finish time.Time `json:"finish"`
}

func saveUrl(url, file string) {
	addOne(url, file, urlTable)
}
func getUrl(url string) string {
	return string(findOne(url, urlTable))
}
func saveSha(sha, file, url string, count int) {
	addOne(sha, TypeTools.OutJson(shaData{url, file, count, time.Now()}), keyTable)
}
func getSha(sha string) shaData {
	var d shaData
	err := json.Unmarshal(findOne(sha, keyTable), &d)
	if err != nil {
		log.Log("getSha", err.Error())
		return d
	}
	return d
}
func findOne(key, table string) (value []byte) {
	if db == nil {
		return
	}
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return nil
		}
		value = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		log.Log("findOneData:"+table+":"+key, err.Error())
		return
	}
	return
}
func addOne(key, value, table string) (suc bool) {
	if db == nil {
		return
	}
	err := db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(value))
	})
	if err != nil {
		log.Log("addOne:"+table+":"+key+"-"+value+"-", err.Error())
		return false
	}
	return true
}
