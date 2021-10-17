/***

# rename files in dirs

-current_dir
|- dir1
|- 001.jpg
|- 002.jpg
|- dir2
|- 001.jpg
|- 002.jpg

Convert to

-current_dir
|- dir1-001.jpg
|- dir1-002.jpg
|- dir2-001.jpg
|- dir2-002.jpg

*/

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        fmt.Println(f.Name())
        info, err := os.Stat(f.Name());
        if  err == nil && info.IsDir() {
            fmt.Println(info.IsDir())
            dirs, _ := ioutil.ReadDir( f.Name() + "/L-size")
            for _, d := range dirs {
                fmt.Println(d.Name())
                e := os.Rename(f.Name() + "/L-size/" + d.Name(), f.Name()+"-"+d.Name())
                if e != nil {
                    log.Fatal(e)
                }
            }
        }
    }
}