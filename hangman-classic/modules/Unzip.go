package modules

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Unzip() {
	DownloadAnim()
	dst := "Animation"
	archive, err := zip.OpenReader("Animation.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("Invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
} // Unzip and create file
func DownloadAnim() {
	url := "https://drive.google.com/uc?export=download&id=1NLgmPituSyN--PAROmC-ieXlDswENjsQ"
	fileName := "Animation.zip"

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)

} // Download all files of animation
