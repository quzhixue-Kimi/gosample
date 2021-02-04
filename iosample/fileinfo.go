package iosample

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func FileInfoSample() {
	if fileInfo, err := os.Stat("/tmp/goflyway.log"); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		fmt.Printf("%T\n", fileInfo)
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.Size())
		fmt.Println(fileInfo.IsDir())
		fmt.Println(fileInfo.ModTime())
		fmt.Println(fileInfo.Mode())
	}
}

func FilePathSample() {
	fileName1 := "/tmp/goflyway.log"
	fileName2 := "samplecode.go"
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))
	fmt.Println(path.Join(fileName1, ".."))
}

func ReadFileByIO() {
	if file, err := os.Open("/tmp/goflyway.log"); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		bs := make([]byte, 512)
		//n := -1
		for {
			n, err := file.Read(bs)
			if n == 0 || err == io.EOF {
				fmt.Println("读取到了文件的末尾，结束读取操作。。")
				break
			}
			fmt.Println(string(bs[:n]))
		}
	}
}

func WriteFileByIO() {
	newFile := "/tmp/test.log"
	if file, err := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	} else {
		defer file.Close()
		if _, err := file.WriteString("hello world\n"); err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		}
	}
}

func CopyFileByIO(srcFile, destFile string) (int64, error) {
	if file1, err := os.Open(srcFile); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return 0, err
	} else {
		defer file1.Close()
		if file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			return 0, err
		} else {
			defer file2.Close()
			return io.Copy(file2, file1)
		}
	}
}
