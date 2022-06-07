package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

/* не получилось заставить программу падать при наличии любого кол-ва открытых файловых дескрипторов,
поэтому, я просто вызываю панику на каждом 64-м созданном файле, отлавливаю, и обрабатываю её
*/

func createMyEmptyFile(path, prefix string, count int) {
	var arrFD []*os.File
	var errFD []string
	for i := 1; i <= count; i++ {
		name := path + prefix + "_" + strconv.Itoa(i)

		func(string, int) {
			defer func() {
				if v := recover(); v != nil {
					fmt.Println("try to recover")
					errFD = append(errFD, name)
					for _, value := range arrFD {
						value.Close()
						fmt.Printf("%s, has been closed \n", value.Name())
					}
					arrFD = arrFD[:0]
					fmt.Println(arrFD)
				}
			}()
			if i%64 == 0 {
				panic("WARNING: There are only 64 file descriptors (hard limit) available, which limit the number of simultaneous connections.")
			} else {
				file, err := os.Create(name)
				if err != nil {
					fmt.Printf("file %s can't be created ", name)
				} else {
					fmt.Println(file.Name())
					arrFD = append(arrFD, file)
				}
			}
		}(name, i)

	}

	for _, value := range arrFD {
		value.Close()
		fmt.Printf("%s, has been closed \n", value.Name())
	}

	for _, value := range errFD {
		file, err := os.Create(value)
		if err != nil {
			fmt.Println("Not recovered")
		} else {
			fmt.Println(value)
			file.Close()
			fmt.Printf("%s, has been closed \n", value)
		}
	}
}

func main() {
	path := flag.String("path", "", "path to file")
	prefix := flag.String("prefix", "", "file prefix")
	fileN := flag.Int("count", 1, "number of created files")
	flag.Parse()
	createMyEmptyFile(*path, *prefix, *fileN)
}
