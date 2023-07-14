package storage

import (
	"fmt"
	"os"
)

/**
用文件保存数据需要解决的问题：
	1. 写入时应该保留之前的数据，就是说不能直接覆盖式的写入。
	2. 少量数据和大量数据的写入方式是不一样的。
	3. 并发读取时可能会读取到不完整的数据
	4. 数据在什么时候进行持久化存盘操作，就是说作为存储中介向外提供的保存接口和实际存盘应该是异步的。
	   调用方调用的保存接口传递数据是存放在内存的。
	5. 打开文件后写入到一半断点或崩溃怎么办。
		Ans: 如果是覆盖式写入，可以写入到一个临时文件，完成后重命名>删除源文件>移动文件到源文件目录

 https://build-your-own.org/database/01_files
*/

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	return err
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

func SaveData3(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	err = fp.Sync() // fsync
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

// TODO
func randomInt() int {
	return 666
}
