package dosktop

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nwaples/rardecode"
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/yeka/zip"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strings"
)
var virtualFileSystemArchive string
var virtualFileSystemPassword string
var virtualFileSystemArchiveType int
var virtualFileSystemEncryptionKey string

/*
GetScrambledPassword allows you to scramble a password with a simple
XOR algorithm. This allows a user to provide a password for a virtual file
system without having to store it in their own program as plaintext. To use
this feature, simply pass in your desired password and scramble key to
obtain your encoded password. This password can then be used to mount a virtual
file system provided you use the same scramble key to decode it. In addition,
the following information should be noted:

- This method is not designed for cryptographic security. It is simply
a method of transposing a password so that it is not viewable in
plaintext. While this can prevent casual users from extracting a password
hardcoded into a binary executable, it will not prevent a determined user
from obtaining it.

- The length and randomness of your 'password' will directly influence the
usefulness of your chosen 'scrambleKey'. Consider using a long and
random password if you want maximum effectiveness.
*/
func GetScrambledPassword(password string, scrambleKey string) string {
	scrambledPassword := xorString(password, scrambleKey)
	scrambledPassword = base64.StdEncoding.EncodeToString([]byte(scrambledPassword))
	return scrambledPassword
}

/*
getUnscrambledPassword allows you to obtain an unscrambled password that was
created using the 'GetScrambledPassword' method. This is used by the virtual
file system to decode a password that was previously scrambled by the user in
order to avoid storing passwords in plaintext. In addition, the following
information should be noted:

- Password scrambling is not designed for cryptographic security. It is simply
a method of transposing a password so that it is not viewable in
plaintext. While this can prevent casual users from extracting a password
hardcoded into a binary executable, it will not prevent a determined user
from obtaining it.
*/
func getUnscrambledPassword(password string, scrambleKey string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(password)
	unscrambledPassword := xorString(string(decodedString), scrambleKey)
	return unscrambledPassword
}

/*
xorString allows you to perform an xor over a given string using a given
'scrambleKey'. This method is useful for when you don't want to store
something in plaintext. In addition, the following information should be
noted:

- String scrambling is not designed for cryptographic security. It is simply
a method of transposing data so that it is not viewable in plaintext. While
this can prevent casual users from extracting data hardcoded into a binary
executable, it will not prevent a determined user from obtaining it.

- While this method will properly XOR your string, it will not guarantee
that the obtained result is screen printable. Consider converting the
result to base64 if you require a value which can be correctly printed
and copied.

- The 'scrambleKey' will be MD5 hashed before being used. This ensures that
if any part of the scramble key has been modified, a totally different
XOR pattern will be generated. This precaution prevents partial decoding of
string data when some of the scramble key happens to be correct.
*/
func xorString(stringToXor string, scrambleKey string) string {
	var xoredString string
	hashedScrambleKey := getMD5Hash(scrambleKey)
	for i := 0; i < len(stringToXor); i++ {
		xoredString += string(stringToXor[i] ^ hashedScrambleKey[i % len(hashedScrambleKey)])
	}
	return xoredString
}

/*
getMD5Hash allows you to obtain an MD5 hash from a provided text string.
*/
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}


/*
SetVirtualFileSystem allows you to specify a virtual file system to mount.
A virtual file system is a ZIP or RAR archive that contains all files in which
you wish to access. This is useful since instead of distributing multiple
files and folders with your application, you can simply package everything
inside a virtual file system and just include that instead. Since files on
your local file system are accessed the same way as on a virtual file
system, you can use external resources for testing with and package them up
into your virtual file system when your ready for distribution. In addition,
the following information should be noted:

- If the archive you wish to use as your virtual file system is password
protected, you must provide it here at mount time. If the archive is not
password protected, then you can simply set this parameter as "" since it
will be ignored.

- If password protection is used, your password should not be considered
secure. Password protection is designed for basic resource integrity
purposes only. It will prevent casual modifications of your resources,
but nothing more.

- If you do not wish to store virtual file system passwords in plaintext,
you can provide a scrambled password instead. Simply pass in a
'scrambleKey' and your password will be decoded before use. If your
password does not require unscrambling, you can simply leave this parameter
as "". For more information on how to obtain a scrambled password, please
see the method 'GetScrambledPassword' for details.

- If for some reason the virtual file system was unable to be mounted, an
error will be returned so that your application can handle this case
appropriately.
*/
func SetVirtualFileSystem(archivePath string, password string, scrambleKey string) error {
	err := isArchiveFormatZip(archivePath)
	if err == nil {
		virtualFileSystemArchiveType = constants.VirtualFileSystemZip
		virtualFileSystemArchive = archivePath
		virtualFileSystemPassword = password
		virtualFileSystemEncryptionKey = scrambleKey
		return err
	}
	err = isArchiveFormatRar(archivePath, password)
	if err == nil {
		virtualFileSystemArchiveType = constants.VirtualFileSystemRar
		virtualFileSystemArchive = archivePath
		virtualFileSystemPassword = password
		virtualFileSystemEncryptionKey = scrambleKey
		return err
	}
	err = errors.New(fmt.Sprintf("Failed to open or decode '%s'.", archivePath))
	return err
}

/*
isArchiveFormatZip allows you to detect if the provided archive is in ZIP
format or not.
*/
func isArchiveFormatZip(archivePath string) error {
	readCloser, err := zip.OpenReader(archivePath)
	if err == nil {
		readCloser.Close()
	}
	return err
}

/*
isArchiveFormatRar allows you to detect if the provided archive is in RAR
format or not.
*/
func isArchiveFormatRar(archivePath string, password string) error {
	readCloser, err := rardecode.OpenReader(archivePath, password)
	if err == nil {
		readCloser.Close()
	}
	return err
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
	var imageData image.Image
	fileData, err := getFileDataFromFileSystem(imageFile)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not get image data from '%s': %s", imageFile, err.Error()))
		return nil, err
	}
	if strings.HasSuffix(strings.ToLower(imageFile), ".jpg") || strings.HasSuffix(strings.ToLower(imageFile), ".jpeg"){
		imageData, err = jpeg.Decode(bytes.NewReader(fileData))
	}
	if strings.HasSuffix(strings.ToLower(imageFile), ".png") {
		imageData, err = png.Decode(bytes.NewReader(fileData))
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not decode the image '%s': %s", imageFile, err.Error()))
		return nil, err
	}
	return imageData, err
}

/*
getTextFromFileSystem allows you to obtain text data from a file from
the default file system. In addition, the following information should
be noted:

- If for some reason the requested text data could not be obtained, an
error will be returned so that your application can handle this case
appropriately.
*/
func getTextFromFileSystem(textFile string) (string, error) {
	fileData, err := getFileDataFromFileSystem(textFile)
	dataAsString := string(fileData)
	return dataAsString, err
}

/*
getFileDataFromFileSystem allows you to get the contents of a file
from the default file system. If you have a virtual file system mounted, then
the file will be retrieved from it instead of your local file system. In
addition, the following information should be noted:

- If a file is being accessed from a password protected virtual file
system, then the password provided at mount time will be used to decrypt
the file automatically.
*/
func getFileDataFromFileSystem(fileName string) ([]byte, error) {
	var fileData []byte
	var err error
	if virtualFileSystemArchiveType == constants.VirtualFileSystemZip {
		fileData, err = getFileDataFromZipArchive(fileName)
	} else if virtualFileSystemArchiveType == constants.VirtualFileSystemRar {
		fileData, err = getFileDataFromRarArchive(fileName)
	}
	if err != nil {
		archiveError := err
		fileData, err = getFileDataFromLocalFileSystem(fileName)
		if err != nil {
			err = errors.New(fmt.Sprintf("Could not open the file '%s': %s, %s", fileName, archiveError.Error(), err.Error()))
		}
	}
	return fileData, err
}

/*
getFileDataFromFileSystem allows you to get the contents of a file
from the local file system. If the contents of the file cannot be
retrieved, then an error is returned instead.
*/
func getFileDataFromLocalFileSystem(fileName string) ([]byte, error) {
	var fileData []byte
	fileReadCloser, err := os.Open(fileName)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not open the file '%s': %s", fileName, err.Error()))
		return fileData, err
	}
	defer fileReadCloser.Close()
	fileData, err = ioutil.ReadAll(fileReadCloser)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not read data from the file '%s': %s", fileName, err.Error()))
		return fileData, err
	}
	return fileData, err
}

/*
getFileDataFromFileSystem allows you to get the contents of a file from a ZIP
archive. If the contents of the file cannot be retrieved, then an error is
returned instead. In addition, the following information should be noted:

- If a file is being accessed from a password protected virtual file
system, then the password provided at mount time will be used to decrypt
the file automatically.
*/
func getFileDataFromZipArchive(fileName string) ([]byte, error) {
	var err error
	var fileReadCloser io.ReadCloser
	var fileData []byte
	archivePassword := virtualFileSystemPassword
	if virtualFileSystemEncryptionKey != "" {
		archivePassword = getUnscrambledPassword(archivePassword, virtualFileSystemEncryptionKey)
	}
	archiveReadCloser, err := zip.OpenReader(virtualFileSystemArchive)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not open '%s': %s", virtualFileSystemArchive, err.Error()))
		return fileData, err
	}
	defer archiveReadCloser.Close()
	for _, currentFile := range archiveReadCloser.File {
		if currentFile.Name == fileName {
			if currentFile.IsEncrypted() {
				currentFile.SetPassword(archivePassword)
			}
			fileReadCloser, err = currentFile.Open()
			if err != nil {
				err = errors.New(fmt.Sprintf("Could not open the file '%s' from the virtual file system: %s", fileName, err.Error()))
				return fileData, err
			}
			fileData, err = ioutil.ReadAll(fileReadCloser)
			if err != nil {
				err = errors.New(fmt.Sprintf("Could not read data from the file '%s': %s", fileName, err.Error()))
				return fileData, err
			}
			fileReadCloser.Close()
		}
	}
	if fileData == nil {
		err = errors.New(fmt.Sprintf("Could not find the file '%s' from the virtual file system.", fileName))
		return fileData, err
	}
	return fileData, err
}

/*
getFileDataFromFileSystem allows you to get the contents of a file from a RAR
archive. If the contents of the file cannot be retrieved, then an error is
returned instead. In addition, the following information should be noted:

- If a file is being accessed from a password protected virtual file
system, then the password provided at mount time will be used to decrypt
the file automatically. If a scramble key was provided at mount time, it will
be used to unscramble the password before being used.
*/
func getFileDataFromRarArchive(fileName string) ([]byte, error) {
	var fileData []byte
	archivePassword := virtualFileSystemPassword
	if virtualFileSystemEncryptionKey != "" {
		archivePassword = getUnscrambledPassword(archivePassword, virtualFileSystemEncryptionKey)
	}
	archiveReadCloser, err := rardecode.OpenReader(virtualFileSystemArchive, archivePassword)
	if err != nil {
		err = errors.New(fmt.Sprintf("Could not open '%s': %s", virtualFileSystemArchive, err.Error()))
		return fileData, err
	}
	defer archiveReadCloser.Close()
	for {
		fileHeader, err := archiveReadCloser.Next()
		if err == io.EOF {
			// If EOF then we are done reading the archive.
			if fileData == nil {
				err = errors.New(fmt.Sprintf("Could not find file '%s' in archive '%s': %s", fileName, virtualFileSystemArchive, err.Error()))
			}
			return fileData, err
		}
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed while scanning archive '%s': %s", virtualFileSystemArchive, err.Error()))
			return fileData, err
		}
		if fileHeader.Name == fileName {
			fileData, err = ioutil.ReadAll(archiveReadCloser)
			if err != nil {
				err = errors.New(fmt.Sprintf("Could not read data from the file '%s': %s", fileName, err.Error()))
				return fileData, err
			}
			return fileData, err
		}
	}
	return fileData, err
}