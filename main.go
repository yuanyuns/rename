package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var name string


func init()  {
	flag.StringVar(&name, "name", "", "重命名前缀")
}

func main() {

	flag.Parse()
	files, _ := ioutil.ReadDir("./")

	for index, f := range files {
		rename(f.Name(),index)
	}

}

func rename(oldPath string, index int) {

	filenameWithSuffix :=path.Ext(oldPath)
	// 去除空格
	if filenameWithSuffix !="" && filenameWithSuffix != ".exe" && filenameWithSuffix!=".mod" && filenameWithSuffix!=".go" {
		newPath := fmt.Sprintf("%s%d%s",name,index,filenameWithSuffix)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("将 %s ==> %s  重命名失败 \n", oldPath,newPath)
			panic(err)
		}else {
			fmt.Printf("将 %s ==> %s  重命名成功 \n", oldPath,newPath)
		}
	}
}
