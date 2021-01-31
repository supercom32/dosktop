package filesystem

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func DownloadFile(url string, filepath string) error {
	// Obtain the actual file data.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file for writing into.
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the response body data to the file.
	_, err = io.Copy(out, resp.Body)
	return err
}

func CopyFile(sourceFile string, targetFile string) error {
	sourceFileStat, err := os.Stat(sourceFile)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", sourceFile)
	}
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return err
}

func WriteBytesToFile(fileName string, bytesToWrite []byte, permissions int) error {
	perm := os.FileMode(uint32(permissions))
	err := ioutil.WriteFile(fileName, bytesToWrite, perm)
	return err
}

func AppendLinesToFile(fileName string, lineToWrite string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	file.WriteString(lineToWrite)
	return err
}

func GetListOfFiles(directoryPath string, regexMatcher string) ([]string, error) {
	return GetListOfDirectoryContents(directoryPath, regexMatcher,true, false)
}

func GetListOfDirectories(directoryPath string, regexMatcher string) ([]string, error) {
	return GetListOfDirectoryContents(directoryPath, regexMatcher, false, true)
}

func GetNormalizedDirectoryPath(directoryPath string) string {
	var normalizedDirectoryPath string = directoryPath
	if !strings.HasSuffix(directoryPath, "/")  && !strings.HasSuffix(directoryPath, "\\") {
		normalizedDirectoryPath = normalizedDirectoryPath + "/"
	}
	return normalizedDirectoryPath
}

func GetListOfDirectoryContents(directoryPath string, regexMatcher string, isFilesListed bool, isDirectoriesListed bool) ([]string, error) {
	var fileList []string
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return fileList, err
	}
	for _, file := range files {
		regex := regexp.MustCompile(regexMatcher)
		match := regex.FindStringSubmatch(file.Name())
		if len(match) > 0 {
			if file.IsDir() && isDirectoriesListed {
				fileList = append(fileList, file.Name())
			}
			if !file.IsDir() && isFilesListed {
				fileList = append(fileList, file.Name())
			}
		}
	}
	return fileList, err
}

func IsDirectoryEmpty(directoryName string) (bool, error) {
	f, err := os.Open(directoryName)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}


func GetFileSize(fileName string) (int64, error){
	var fileSize int64
	fi, err := os.Stat(fileName)
	if err != nil {
		return fileSize, err
	}
	fileSize = fi.Size()
	return fileSize, err
}

func GetWorkingDirectory() (string, error) {
	var parent string
	workingDirectory, err := os.Getwd()
	if err != nil {
		return parent, err
	}
	parent = filepath.Dir(workingDirectory)
	return parent, err
}

func GetParentDirectory(directoryPath string) string {
	normalizedDirectory := directoryPath
	if strings.HasSuffix(directoryPath, "/") || strings.HasSuffix(directoryPath, "\\") {
		normalizedDirectory = normalizedDirectory[:len(normalizedDirectory)-1]
	}
	parentDirectory := path.Dir(normalizedDirectory)
	return parentDirectory
}

func RenameFile(sourceFile string, targetFile string) error {
	err := os.Rename(sourceFile, targetFile)
	return err
}

func IsDirectoryExists(directoryPath string) bool {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFile(fileName string) error {
	err := os.Remove(fileName)
	return err
}

func MoveFile(sourceFile string, targetFile string) error {
	err := os.Rename(sourceFile, targetFile)
	return err
}

func GetFileContentsAsBytes(fileName string) ([]byte, error) {
	var fileContents []byte
	var err error
	fileContents, err = ioutil.ReadFile(fileName)
	if err != nil {
		return fileContents, err
	}
	return fileContents, err
}

func CreateDirectory(directoryPath string, permissions int) error {
	perm := os.FileMode(uint32(permissions))
	err := os.MkdirAll(directoryPath, perm)
	return err
}
func GetFileExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	return extension
}

func GetBaseFileName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetFileNameFromPath(fullyQualifiedFileName string) string {
	return filepath.Base(fullyQualifiedFileName)
}
func GetBaseDirectory(filePath string) string {
	directory, _ := path.Split(filePath)
	return directory
}

func IsFile(path string) (bool, error){
	var isFile bool
	fi, err := os.Stat(path)
	if err != nil {
		return isFile, err
	}
	mode := fi.Mode()
	if mode.IsRegular() {
		isFile = true
	}
	return isFile, err
}