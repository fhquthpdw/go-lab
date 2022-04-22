/*
Copyright © 2020 daochun.zhao <daochun.zhao@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/janeczku/go-spinner"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// scannerCmd represents the scanner command
var scannerCmd = &cobra.Command{
	Use:   "scanner",
	Short: "scanner all files",
	Long: `scanner all files, and list duplicate files, and output to files, and delete duplicate files`,
	RunE: ScannerRun,
}

var outputFormat *string	// 输出格式
var output *string		// 输出位置 console, stderr, file ...
var root string			// 扫描起始位置
var routineNumber int64	// goroutine数量
var fileType string		// 过滤文件类型
var fileTypeArr []string		// 过滤文件类型
var fileSize string		// 过滤文件大小 20m, 20M, 20k, 20K, 20G, 20g
var fileSizeByte int64		// 过滤文件大小 20m, 20M, 20k, 20K, 20G, 20g

func init() {
	rootCmd.AddCommand(scannerCmd)

	scannerCmd.PersistentFlags().StringVar(&root, "root", "", "scan root path")
	scannerCmd.PersistentFlags().Int64Var(&routineNumber, "routineNumber", 0, "goroutine number")
	scannerCmd.PersistentFlags().StringVar(&fileType, "fileType", "", "filter file type")
	scannerCmd.PersistentFlags().StringVar(&fileSize, "fileSize", "", "filter file size")

	outputFormat = scannerCmd.Flags().StringP("outputFormat", "f", "text", "output format text,json,xml,yaml")
	output = scannerCmd.Flags().StringP("output", "o", "", "where to output the result, default is console or file")
}

var (
	mutex sync.Mutex
	goroutinesPool = make(chan struct{}, 10000)
	filesize = make(map[string]string)
	wg sync.WaitGroup
)
/////
func ScannerRun(cmd *cobra.Command, args []string) (err error) {
	if root == "" {
		if root, err = os.Getwd(); err != nil {
			return errors.New("can not get root path")
		}
	}

	// 解晰参数
	if fileType != "" {
		fileTypeArr = strings.Split(fileType, ",")
	}
	fileSizeByte = parseFileSize(fileSize)

	// 旋转 loading...
	s := spinner.StartNew("scanning...")
	//s.SetCharset([]string{">  ", " >>", ">>>"})
	s.SetSpeed(time.Millisecond * 100)

	// 递归扫描
	wg.Add(1)
	goroutinesPool <- struct{}{}
	go scan(root, &wg)
	wg.Wait()
	//time.Sleep(2 * time.Second) // for spinner
	s.Stop()

	// 输出
	// TODO: 格式化输出
	for file, size := range filesize {
		fmt.Println(file, size)
	}

	return nil
}
func scan(rootPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	files, _ := ioutil.ReadDir(rootPath)
	for _, file := range files {
		rootPath = strings.TrimRight(rootPath, "/")
		fullPath := rootPath + "/" + file.Name()
		if file.IsDir() {
			wg.Add(1)
			goroutinesPool <- struct{}{}
			scan(fullPath, wg)
		} else {
			if checkType(fullPath) &&
				checkSize(file.Size()){
				mutex.Lock()
				filesize[fullPath] = formatOutputSize(file.Size())
				mutex.Unlock()
			}
		}
	}
	<-goroutinesPool
}
func checkType(filePath string) bool {
	if len(fileTypeArr) == 0 {
		return true
	}
	ext := strings.ReplaceAll(filepath.Ext(filePath), ".", "")
	return checkSliceContains(fileTypeArr, ext)
}
func checkSize(fileSize int64) bool {
	if fileSizeByte == 0 {
		return true
	} else {
		if fileSize >= fileSizeByte {
			return true
		}
	}
	return false
}
func checkSliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func parseFileSize(fileSize string) int64 {
	if fileSize == "" {
		return 0
	}

	fileSize = strings.ToLower(fileSize)
	lastC := fileSize[len(fileSize)-1:]
	n, err := strconv.ParseInt(fileSize[:len(fileSize)-1], 10, 64)
	if err != nil {
		return 0
	}

	switch lastC {
	case "k":
		return 1024 * n
	case "m":
		return 1024 * 1024 * n
	case "g":
		return 1024 * 1024 * 1024 * n
	}

	return 0
}

// human size
const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
)
func formatOutputSize(sizeByte int64) string {
	sizeByteFloat := float64(sizeByte)
	if sizeByte < KB {
		return fmt.Sprintf("%dB", sizeByte)
	} else {
		if sizeByte >= KB && sizeByte < MB {
			return fmt.Sprintf("%.2fK", sizeByteFloat / 1024)
		} else if sizeByte >= MB && sizeByte < GB {
			return fmt.Sprintf("%.2fM", sizeByteFloat / 1024 / 1024)
		} else {
			return fmt.Sprintf("%.2fG", sizeByteFloat / 1024 / 1024 / 1024)
		}
	}
	return ""
}