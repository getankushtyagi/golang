package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("23_files/example.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("filename ",f.Name())
	fmt.Println(f.Stat())

	// there are mutliple ways to read the file, we can read it in chunks or we can read it all at once
	// we can read it in chunks using the Read method
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("read bytes ", n)
	fmt.Println(string(buf[:n]))

	// we can read it all at once using the ioutil package
	// data, err := ioutil.ReadFile("23_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// defer f.Close()


	// this is how we read the file 

	buf2 := make([]byte, 1024)
	fmt.Println(f.Read(buf2))

	// let see some other methods available in os package
	// we can use os.Stat to get the file information
	info, err := os.Stat("23_files/example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file size ", info.Size())
	fmt.Println("file mode ", info.Mode())
	fmt.Println("file mod time ", info.ModTime())

	// we can use os.Remove to delete the file
	// err = os.Remove("23_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// we can use os.Rename to rename the file
	// err = os.Rename("23_files/example.txt", "23_files/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
}
