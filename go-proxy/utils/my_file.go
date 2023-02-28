package utils

import (
	"fmt"
	"os"
)

// 写入文件,文件不存在则创建,如在则追加内容
func WriteFile(path string, str string) {
	_, b := IsFile(path)
	var f *os.File
	var err error
	if b {
		//打开文件，
		// f, _ = os.OpenFile(path, os.O_APPEND, 0666)
		// f, _ = os.OpenFile(path, os.O_RDWR, 0666)
		f, _ = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)

	} else {
		//新建文件
		f, err = os.Create(path)
	}

	//使用完毕，需要关闭文件
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("err  =  ", err)
		}
	}()

	if err != nil {
		fmt.Println("err  =  ", err)
		return
	}
	_, err = f.WriteString(str)
	if err != nil {
		fmt.Println("err  =  ", err)
	}
}

// 判断路径是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func IsDir(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && f.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}

// func  main()  {
// 	path  :=  "./demo.txt"
// 	str  :=  "abcd\r\nefg我爱学习"
// 	WriteFile(path,  str)
// }

// func WriteToFile(
// 	fileName string,
// ) {
// 	file, err := os.Open("xxx")
// 	defer func() {
// 		file.Close()
// 	}()
// 	if err != nil && os.IsNotExist(err) {
// 		err := os.Mkdir(fileName, os.ModePerm)
// 		file = os.Create("xx")
// 	}

// 	filePath := "D:/goProject/src/go_demo/files/test.txt"
// 	// O_APPEDN 追加方式
// 	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		fmt.Printf("open file error=%v\n", err)
// 		return
// 	}
// 	defer file.Close()
// 	str := "hello golang\n"
// 	writer := bufio.NewWriter(file)
// 	for i := 0; i < 5; i++ {
// 		writer.WriteString(str)
// 	}
// 	writer.Flush()
// }
