package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pl := fmt.Println
	err := os.Mkdir("d:/subdir", 0755)
	check(err)

	defer os.RemoveAll("d:/subdir")

	createEmptyFile := func(name string) {
		d_arr := []byte("")
		check(os.WriteFile(name, d_arr, 0644))
	}

	createEmptyFile("d:/subdir/file1")

	err = os.MkdirAll("d:/subdir/parent/child", 0755)
	check(err)

	createEmptyFile("d:/subdir/parent/file2")
	createEmptyFile("d:/subdir/parent/file3")
	createEmptyFile("d:/subdir/parent/child/file4")

	c, err := os.ReadDir("d:/subdir/parent")
	check(err)

	pl("listing d:/subdir/parent")

	for _, entry := range c {
		pl("", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("d:/subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	pl("Listing d:/subdir/parent/child")

	for _, entry := range c {
		pl(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	pl("Visiting subdir")
	err = filepath.Walk("d:/subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil
}
