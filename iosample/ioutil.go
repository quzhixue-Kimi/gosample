package iosample

import (
	"fmt"
	"io/ioutil"
	"log"
)

func CopyFileByIOUtil(srcFile, destFile string) (int, error) {
	if input, err := ioutil.ReadFile(srcFile); err != nil {
		return 0, err
	} else {
		if err = ioutil.WriteFile(destFile, input, 0644); err != nil {
			return 0, err
		}
		return len(input), nil
	}
}

func ListFiles(dirname string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			ListFiles(filename, level+1)
		}
	}
}
