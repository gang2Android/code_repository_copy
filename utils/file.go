package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CopyDir(src string, dest string) {
	fmt.Println(src, dest)
	// 清除目标文件夹下的文件(.git除外)
	err1 := filepath.Walk(dest+string(os.PathSeparator), func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		if strings.Contains(path, ".git") {
			if !strings.Contains(path, ".gitignore") {
				return nil
			}
		}
		if strings.LastIndex(path, string(os.PathSeparator)) == (len(path) - 1) {
			fmt.Println("1")
		} else {
			err2 := os.RemoveAll(path)
			if err2 != nil {
				fmt.Println(err2)
			}
		}
		return nil
	})
	if err1 != nil {
		fmt.Println("删除目标文件夹的内容失败-" + err1.Error())
		return
	}

	srcOriginal := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		fmt.Println(src)
		fmt.Println(srcOriginal)
		fmt.Println(dest)
		if !f.IsDir() {
			if strings.Contains(src, ".git") {
				fmt.Println("跳过.git")
				return nil
			}
			destNew := strings.Replace(src, srcOriginal, dest, -1)
			fmt.Println(destNew)
			fmt.Println("CopyFile:" + src + " to " + destNew)
			_, err1 := CopyFile(src, destNew)
			if err1 != nil {
				fmt.Println("复制文件异常", err1)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	fmt.Println("dst:" + dst)
	dst_slices := strings.Split(dst, "\\")
	dst_slices_len := len(dst_slices)
	dest_dir := ""
	for i := 0; i < dst_slices_len-1; i++ {
		dest_dir = dest_dir + dst_slices[i] + "\\"
	}
	fmt.Println("dest_dir:" + dest_dir)
	b, err := PathExists(dest_dir)
	if b == false {
		err := os.MkdirAll(dest_dir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
