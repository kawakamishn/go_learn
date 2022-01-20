//解凍したときに現れる、__MACOSX/manual/._精算WEBマニュアル_一般ユーザー㈬国内出張202101.xls　などはキャッシュファイル

package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if strings.HasSuffix(os.Args[1], ".zip") {
		r, err := zip.OpenReader(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		for _, f := range r.File {
			fmt.Println(f)
		}
	} else if strings.HasSuffix(os.Args[1], ".tar") {
		r, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		tr := tar.NewReader(r)
		// Iterate through the files in the archive.
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				// end of tar archive
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(hdr.Name)
		}
	} else {
		fmt.Println("not supported file format")
	}

}
