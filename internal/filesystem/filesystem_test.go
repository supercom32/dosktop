package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadFile(test *testing.T) {
	err := DownloadFile("https://bad_url", "/tmp/download.txt")
	assert.Errorf(test, err, "An error was expected to be generated when a bad URL is used!")
	err = DownloadFile("https://www.google.ca", "/tmp/index.html")
	assert.NoErrorf(test, err, "An error was not expected to be generated when downloading!")
}

func TestCopyFile(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	targetFile := "/tmp/target.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = CopyFile(sourceFile, targetFile)
	assert.NoErrorf(test, err, "An error was not expected when copying a sample file!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a source sample file!")
	err = DeleteFile(targetFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a target sample file!")
}

func TestAppendLinesToFile(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = AppendLinesToFile(sourceFile, "A sample line")
	assert.NoErrorf(test, err, "An error was not expected when trying to append to a file!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a source sample file!")
}

func TestGetListOfFiles(test *testing.T) {
	sourceDirectory := "/tmp/"
	_, err := GetListOfFiles(sourceDirectory, ".*")
	assert.NoErrorf(test, err, "An error was not expected when trying to list all files within a directory!")
}

func TestGetListOfDirectories(test *testing.T) {
	sourceDirectory := "/tmp/"
	_, err := GetListOfDirectories(sourceDirectory, ".*")
	assert.NoErrorf(test, err, "An error was not expected when trying to list all directories within a directory!")
}

func TestGetNormalizedDirectoryPath(test *testing.T) {
	sourceDirectory := "/tmp/someDirectory"
	obtainedValue := GetNormalizedDirectoryPath(sourceDirectory)
	expectedValue := sourceDirectory + "/"
	assert.Equalf(test, expectedValue, obtainedValue, "The directory path was not normalized as expected!")
}

func TestGetListOfDirectoryContents(test *testing.T) {
	sourceDirectory := "/tmp/"
	_, err := GetListOfDirectoryContents(sourceDirectory, ".*", true, true)
	assert.NoErrorf(test, err, "An error was not expected when trying to list the contents of a directory!")
}

func TestIsDirectoryEmpty(test *testing.T) {
	sourceDirectory := "/tmp/"
	sourceFile := "/tmp/source.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	obtainedValue, err := IsDirectoryEmpty(sourceDirectory)
	expectedValue := false
	assert.NoErrorf(test, err, "An error was not expected when trying check if a directory was empty or not!")
	assert.Equalf(test, expectedValue, obtainedValue, "The directory specified was expected to be not empty!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a source sample file!")
}

func TestGetWorkingDirectory(test *testing.T) {
	_, err := GetWorkingDirectory()
	assert.NoErrorf(test, err, "An error was not expected when trying to get the current working directory!")
}

func TestGetParentDirectory(test *testing.T) {
	obtainedValue := GetParentDirectory("/tmp/some_directory/sample_file.txt")
	expectedValue := "/tmp/some_directory"
	assert.Equalf(test, expectedValue, obtainedValue, "The parent directory did not match what was expected!")
}

func TestRenameFile(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	targetFile := "/tmp/target.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = RenameFile(sourceFile, targetFile)
	assert.NoErrorf(test, err, "No error was expected when trying to rename a file!")
	err = DeleteFile(targetFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a source sample file!")
}

func TestIsDirectoryExists(test *testing.T) {
	obtainedValue := IsDirectoryExists("/tmp")
	expectedValue := true
	assert.Equalf(test, expectedValue, obtainedValue, "The specified directory was expected to exist!")
}

func TestDeleteFile(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a source sample file!")
}

func TestMoveFile(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	targetFile := "/tmp/target.txt"
	err := WriteBytesToFile(sourceFile, []byte("sample_string"), 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = MoveFile(sourceFile, targetFile)
	assert.NoErrorf(test, err, "An error was not expected when copying a sample file!")
	obtainedValue := IsDirectoryExists(sourceFile)
	expectedValue := false
	assert.Equalf(test, expectedValue, obtainedValue, "The source directory that was moved was not expected to exist!")
	err = DeleteFile(targetFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a target sample file!")
}

func TestGetFileContentsAsBytes(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	fileContents := []byte("This is a sample string!")
	err := WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a target sample file!")
}

func TestGetFileSize(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	fileContents := []byte("This is a sample string!")
	err := WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	obtainedValue, err := GetFileSize(sourceFile)
	expectedValue := int64(24)
	assert.Equalf(test, expectedValue, obtainedValue, "The obtained file size did not match what was expected.")
	assert.NoErrorf(test, err, "An error was not expected when trying to get the size of the file!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "An error was not expected when trying to delete a target sample file!")
}


func TestCreateDirectory(test *testing.T) {
	sourceDirectory := "/tmp/newDirectory"
	err := CreateDirectory(sourceDirectory, 0666)
	assert.NoErrorf(test, err, "No error was expected to occur when trying to create a directory!")
	err = DeleteFile(sourceDirectory)
	assert.NoErrorf(test, err, "No error was expected to occur when trying to delete a directory!")
}

func TestGetFileExtension(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	fileContents := []byte("This is a sample string!")
	err := WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	obtainedValue := GetFileExtension(sourceFile)
	expectedValue := ".txt"
	assert.Equalf(test, expectedValue, obtainedValue, "The file extension obtained was not what was expected!")
	err = DeleteFile(sourceFile)
	assert.NoErrorf(test, err, "No error was expected to occur when trying to delete a file!")
}

func TestGetBaseFileName(test *testing.T) {
	obtainedValue := GetBaseFileName("my_file.eng.txt")
	expectedValue := "my_file.eng"
	assert.Equalf(test, expectedValue, obtainedValue, "The base filename obtained was not what was expected!")
}

func TestGetFileNameFromPath(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	fileContents := []byte("This is a sample string!")
	err := WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	obtainedValue := GetFileNameFromPath(sourceFile)
	expectedValue := "source.txt"
	assert.Equalf(test, expectedValue, obtainedValue, "The file name obtained from the provided path was not what was expected!")
	err = WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
}

func TestGetBaseDirectory(test *testing.T) {
	sourceFile := "/tmp/source.txt"
	fileContents := []byte("This is a sample string!")
	err := WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
	obtainedValue := GetBaseDirectory(sourceFile)
	expectedValue := "/tmp/"
	assert.Equalf(test, expectedValue, obtainedValue, "The base directory obtained was not what was expected!")
	err = WriteBytesToFile(sourceFile, fileContents, 0666)
	assert.NoErrorf(test, err, "An error was not expected when creating a sample file!")
}

func TestIsFile(test *testing.T) {
	sourceDirectory := "/tmp/newDirectory"
	err := CreateDirectory(sourceDirectory, 0666)
	assert.NoErrorf(test, err, "No error was expected to occur when trying to create a directory!")
	obtainedValue, err := IsFile(sourceDirectory)
	expectedValue := false
	assert.NoErrorf(test, err, "An error was not expected when checking if a path is a directory or not!")
	assert.Equalf(test, expectedValue, obtainedValue, "The value returned for checking if a path was a file or not was not as expected!")
	err = DeleteFile(sourceDirectory)
	assert.NoErrorf(test, err, "No error was expected to occur when trying to delete a directory!")
}