package dosktop

import (
	"bytes"
	"github.com/supercom32/dosktop/internal/memory"
	"errors"
	"fmt"
	"github.com/yeka/zip"
	"image"
	"io"
	"os"
)
var virtualFileSystemArchive string
var virtualFileSystemPassword string

/*
SetVirtualFileSystem allows you to specify a virtual file system to mount.
A virtual file system is a ZIP archive that contains all files in which
you wish to access. This is useful since instead of distributing multiple
files and folders with your application, you can simply package everything
inside a virtual file system and just include that instead. Since files on
your local file system are accessed the same way as on a virtual file
system, you can use external resources for testing with and package them up
into your virtual file system when your ready for distribution. In addition,
the following information should be noted:

- If the ZIP archive you wish to use as your virtual file system is password
protected, you must provide it here at mount time. If the archive is not
password protected, then you can simply set this parameter as "" since it
will be ignored.

- If password protection is used, your password should not be considered
secure. Password protection is designed for basic resource integrity
purposes only. It will prevent casual modifications of your resources,
but nothing more.

- If for some reason the virtual file system was unable to be mounted, an
error will be returned so that your application can handle this case
appropriately.
*/
func SetVirtualFileSystem(zipArchivePath string, password string) error {
	var error error
	virtualFileSystemArchive = zipArchivePath
	virtualFileSystemPassword = password
	openReader, returnedError := zip.OpenReader(zipArchivePath)
	if returnedError != nil {
		error = errors.New("SetVirtualFileSystem: Could not open '" + zipArchivePath)
		return error
	}
	defer openReader.Close()
	return error
}

/*
getImageEntryFromFileSystem allows you to obtain an image entry from the
default file system. If you have a virtual file system mounted, then the
image file will be retrieved from it instead of your local file system.
In addition, the following information should be noted:

- If for some reason the requested image could not be obtained, an
error will be returned so that your application can handle this case
appropriately.
*/
func getImageEntryFromFileSystem(imageFile string) (memory.ImageEntryType, error){
	imageData, err := getImageFromFileSystem(imageFile)
	if err != nil {
		return memory.NewImageEntry(), err
	}
	imageEntry := memory.NewImageEntry()
	imageEntry.ImageData = imageData
	return imageEntry, err
}

/*
getImageFromFileSystem allows you to obtain image data from a file from
the default file system. In addition, the following information should
be noted:

- If for some reason the requested image could not be obtained, an
error will be returned so that your application can handle this case
appropriately.
*/
func getImageFromFileSystem(imageFile string) (image.Image, error) {
	fileReadCloser, archiveReadCloser, err := getReadCloserFromFileSystem(imageFile)
	defer fileReadCloser.Close()
	defer archiveReadCloser.Close()
	imageData, _, err := image.Decode(fileReadCloser)
	if err != nil {
		panic (err)
		err = errors.New(fmt.Sprintf("Could not decode the image '%s': %s", imageFile, err.Error()))
		return nil, err
	}
	return imageData, err
}

/*
getImageFromFileSystem allows you to obtain text data from a file from
the default file system. In addition, the following information should
be noted:

- If for some reason the requested text data could not be obtained, an
error will be returned so that your application can handle this case
appropriately.
*/
func getTextFromFileSystem(textFile string) (string, error) {
	fileReadCloser, archiveReadCloser, err := getReadCloserFromFileSystem(textFile)
	defer fileReadCloser.Close()
	defer archiveReadCloser.Close()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(fileReadCloser)
	dataAsString := buffer.String()
	return dataAsString, err
}

/*
getReadCloserFromFileSystem allows you to get a read closer from a file from
the default file system. If you have a virtual file system mounted, then the
file will be retrieved from it instead of your local file system. In
addition, the following information should be noted:

- If a file is being accessed from a password protected virtual file
system, then the password provided at mount time will be used to decrypt
the file automatically.

- In addition to your 'io.ReadCloser', a '*zip.ReadCloser' is also returned.
This is for the virtual file system and is required to read any data from it.

- The returned 'io.ReadCloser' and '*zip.ReadCloser' should both be closed
once you are done working with it.
*/
func getReadCloserFromFileSystem(fileName string) (io.ReadCloser, *zip.ReadCloser, error) {
	var err error
	var fileReadCloser io.ReadCloser
	var archiveReadCloser zip.ReadCloser
	if virtualFileSystemArchive != "" {
		archiveReadCloser, err := zip.OpenReader(virtualFileSystemArchive)
		if err != nil {
			err = errors.New(fmt.Sprintf("Could not open '%s': %s", virtualFileSystemArchive, err.Error()))
			return fileReadCloser, archiveReadCloser, err
		}
		for _, currentFile := range archiveReadCloser.File {
			if currentFile.Name == fileName {
				if currentFile.IsEncrypted() {
					currentFile.SetPassword(virtualFileSystemPassword)
				}
				fileReadCloser, err = currentFile.Open()
				if err != nil {
					err = errors.New(fmt.Sprintf("Could not open the file '%s' from the virtual file system: %s", fileName, err.Error()))
					return fileReadCloser, archiveReadCloser, err
				}
			}
		}
		if fileReadCloser == nil {
			err = errors.New(fmt.Sprintf("Could not find the file '%s' from the virtual file system: %s", fileName, err.Error()))
			return fileReadCloser, archiveReadCloser, err
		}
	} else {
		fileReadCloser, err = os.Open(fileName)
		if err != nil {
			err = errors.New(fmt.Sprintf("Could not open the file '%s' from local disk: %s", fileName, err.Error()))
			return fileReadCloser, &archiveReadCloser, err
		}
	}
	return fileReadCloser, &archiveReadCloser, err
}