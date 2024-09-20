package main

import (
	"embed"
	"go-fastdfs-web-go/src/models"
	"go-fastdfs-web-go/src/server"
	"go-fastdfs-web-go/src/setting"
	"io"
	"io/fs"
	"os"
)

//go:embed static/* template/* conf/*
var staticFS embed.FS

func init() {
	err := extractFilesIfNotExists(staticFS, "conf", "./conf")
	if err != nil {
		panic(err)
	}
	if err := extractFilesIfNotExists(staticFS, "template", "./template"); err != nil {
		panic(err)
	}
	if err := extractFilesIfNotExists(staticFS, "static", "./static"); err != nil {
		panic(err)
	}
	setting.LoadSetting()
	models.InitDataBase()
}

func extractFilesIfNotExists(fs2 embed.FS, sourceDir, targetDir string) error {
	// 检查目标目录是否存在

	return fs.WalkDir(fs2, sourceDir, func(path string, d fs.DirEntry, err error) error {

		// 如果是目录，则创建目录
		if d.IsDir() {
			return os.MkdirAll(path, 0755)
		}

		// 打开嵌入的文件
		ff, err := fs2.Open(path)
		if err != nil {
			return err
		}
		defer ff.Close()

		// 创建目标文件
		tf, err := os.Create(path)
		if err != nil {
			return err
		}
		defer tf.Close()

		// 复制内容
		_, err = io.Copy(tf, ff)
		return err
	})

}
func main() {

	server.Run()
}
