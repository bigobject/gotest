package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	_, fileInfos, err := getVoiceFiles("./test")
	if err != nil {
		fmt.Println("ClearOverFlowVoice failed, err:", err)
		return
	}

	fmt.Println("fileInfos has ", len(fileInfos), " files")

	sort.Slice(fileInfos, func(i, j int) bool { return fileInfos[i].ModTime().Before(fileInfos[j].ModTime()) })

	for _, node := range fileInfos {
		fmt.Println("FullName:", node.FullName, ", ModTime:", node.ModTime(), ", Size:", node.Size())
	}
}

type MyFileInfo struct {
	FullName string
	os.FileInfo
}

func getVoiceFiles(path string) (totalSize int64, files []MyFileInfo, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, []MyFileInfo{}, err
	}

	for _, node := range dir {
		if node.IsDir() {
			subPath := path + "/" + node.Name()
			subsize, subfiles, suberr := getVoiceFiles(subPath)
			if suberr != nil {
				fmt.Println("getFiles failed for path:", subPath, ", suberr:", suberr)
				continue
			}

			if 0 == subsize {
				files = append(files, MyFileInfo{path + "/" + node.Name(), node})
				continue
			}

			files = append(files, subfiles...)
			totalSize += subsize

		} else {
			if strings.Contains(node.Name(), ".mp3") || strings.Contains(node.Name(), ".wav") || strings.Contains(node.Name(), ".pcm") {
				files = append(files, MyFileInfo{path + "/" + node.Name(), node})
				totalSize += node.Size()
			}
		}
	}

	return totalSize, files, nil
}
