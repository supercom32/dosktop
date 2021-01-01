package dosktop

import (
	"bytes"
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/memory"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"strings"
	"fmt"
)


/*
UnloadImage allows you to remove an image from memory. Since images can
take up a large amount of space, it is recommended to unload images
any time you are done working with them. However, you may consider
retaining images if they are frequently used and repeatedly loading
them would be less effective. In addition, the following information should
be noted:

- If you pass in an image alias that does not exist, then the delete
operation will be ignored.
*/
func UnloadImage(imageAlias string) {
	memory.DeleteImage(imageAlias)
}

/*
LoadImage allows you to load an image into memory without performing
any ansi conversions ahead of time. This takes up more memory for larger images
but allows you to render those images at arbitrary resolutions. For example,
loading a large image to retain detail and dynamically rendering that
image later depending on the available terminal resolution detected.
*/
func LoadImage(imageFile string, imageAlias string) error {
	imageEntry, err := getImageEntryFromFileSystem(imageFile)
	if err != nil {
		return err
	}
	memory.AddImage(imageAlias, imageEntry)
	return err
}

/*
LoadPreRenderedImage allows you to pre-render an image before loading it into
memory. This enables you to save memory by rendering larger images ahead of
time instead of storing the image data for later use. For example, you can
take a large image and pre-render it at a much lower resolution suitable for
the terminal. In addition, the following information should be noted:

- If you load a pre-rendered image, you are not able to draw them dynamically
at various resolutions. The image can only be drawn with the settings specified
at load time.

- If you specify a value of 0 for ether the width or height, then that
dimension will be automatically calculated to a value that best maintain
the images aspect ratio.

- If you specify a value less than or equal to 0 for both the width and
height, a panic will be generated to fail as fast as possible.

- When pre-rendering an image, it should be noted that each text cell assigned
contains a top and bottom pixel. This is done to provide as much resolution as
possible for your image. That means for a pre-rendered image with a size of
10x10 characters, the actual image being rendered would be 10x20 pixels tall.
If the user wishes to maintain proper aspect ratios, they must manually select
a height that appropriately compensates for this effect, or leave the height
value as 0 to have it done automatically.

- The blur sigma controls how much blurring occurs after your image has been
resized. This allows you to soften your image before it is rendered in ansi
so that hard edges are removed. A value of 0.0 means no blurring will occur,
with higher values increasing the blur factor.
*/
func LoadPreRenderedImage(imageFile string, imageAlias string, widthInCharacters int, heightInCharacters int, blurSigma float64) error {
	imageEntry, err := getImageEntryFromFileSystem(imageFile)
	if err != nil {
		return err
	}
	imageEntry.LayerEntry = getImageLayer(imageEntry.ImageData, widthInCharacters, heightInCharacters, blurSigma)
	imageEntry.ImageData = nil
	memory.AddImage(imageAlias, imageEntry)
	return err
}

/*
LoadBase64Image allows you to load a base64 encoded image into memory without
performing any ansi conversions ahead of time. This takes up more memory for
larger images but allows you to render those images at arbitrary resolutions.
For example, loading a large image to retain detail and dynamically rendering
that image later depending on the available terminal resolution detected.
Since base64 encoded images can be stored in strings, they are ideal for
directly embedding them into applications.
*/
func LoadBase64Image(imageDataAsBase64 string, imageAlias string) error {
	imageEntry := memory.NewImageEntry()
	imageData, err := getImageFromBase64String(imageDataAsBase64)
	if err != nil {
		return err
	}
	imageEntry.ImageData = imageData
	memory.AddImage(imageAlias, imageEntry)
	return err
}

/*
LoadPreRenderedBase64Image allows you to pre-render an image before loading it
into memory. This enables you to save memory by rendering larger images ahead
of time instead of storing the image data for later use. For example, you can
take a large image and pre-render it at a much lower resolution suitable for
the terminal. Since base64 encoded images can be stored in strings, they are
ideal for directly embedding them into applications. In addition, the following
information should be noted:

- If you load a pre-rendered image, you are not able to draw them dynamically
at various resolutions. The image can only be drawn with the settings specified
at load time.

- If you specify a value of 0 for ether the width or height, then that
dimension will be automatically calculated to a value that best maintain
the images aspect ratio. This is useful since it removes the need to
calculate this manually.

- If you specify a value less than or equal to 0 for both the width and
height, a panic will be generated to fail as fast as possible.

- When pre-rendering an image, it should be noted that each text cell assigned
contains a top and bottom pixel. This is done to provide as much resolution as
possible for your image. That means for a pre-rendered image with a size of
10x10 characters, the actual image being rendered would be 10x20 pixels tall.
If the user wishes to maintain proper aspect ratios, they must manually select
a height that appropriately compensates for this effect, or leave the height
value as 0 to have it done automatically.

- The blur sigma controls how much blurring occurs after your image has been
resized. This allows you to soften your image before it is rendered in ansi
so that hard edges are removed. A value of 0.0 means no blurring will occur,
with higher values increasing the blur factor.
*/
func LoadPreRenderedBase64Image(imageDataAsBase64 string, imageAlias string, widthInCharacters int, heightInCharacters int, blurSigma float64) error {
	imageEntry := memory.NewImageEntry()
	imageData, err := getImageFromBase64String(imageDataAsBase64)
	if err != nil {
		return err
	}
	imageEntry.LayerEntry = getImageLayer(imageData, widthInCharacters, heightInCharacters, blurSigma)
	memory.AddImage(imageAlias, imageEntry)
	return err
}

/*
getBase64PngFromImage allows you to covert raw image data into a base64 encoded
string. This is useful for embedding images directly in applications.
*/
func getBase64PngFromImage(imageToConvert image.Image) (string, error) {
	var imageAsBase64 string
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, imageToConvert)
	if err != nil {
		return imageAsBase64, err
	}
	imageAsBase64 = base64.StdEncoding.EncodeToString(buffer.Bytes())
	return imageAsBase64, err
}

/*
getImageFromBase64String allows you to obtain raw image data from a base64
encoded string. This is useful for when images are embedded  directly into
applications.
*/
func getImageFromBase64String(imageAsBase64 string) (image.Image, error){
	fileReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageAsBase64))
	imageData, _, err := image.Decode(fileReader)
	return imageData, err
}

/*
resizeImage allows you to resize an image.
*/
func resizeImage(imageData image.Image, width uint, height uint) image.Image {
	return resize.Resize(width, height, imageData, resize.Lanczos3)
	// return resize.Resize(width, height, imageData, resize.NearestNeighbor)
}

/*
getImageLayer allows you to specify an image and convert it into a text layer
suitable for drawing with. In addition, the following information should be
noted:

- If you specify a value of 0 for ether the width or height, then that
dimension will be automatically calculated to a value that best maintain
the images aspect ratio. This is useful since it removes the need to
calculate this manually.

- If you specify a value less than or equal to 0 for both the width and
height, a panic will be generated to fail as fast as possible.

- When pre-rendering an image, it should be noted that each text cell assigned
contains a top and bottom pixel. This is done to provide as much resolution as
possible for your image. That means for a pre-rendered image with a size of
10x10 characters, the actual image being rendered would be 10x20 pixels tall.
If the user wishes to maintain proper aspect ratios, they must manually select
a height that appropriately compensates for this effect, or leave the height
value as 0 to have it done automatically.

- The blur sigma controls how much blurring occurs after your image has been
resized. This allows you to soften your image before it is rendered in ansi
so that hard edges are removed. A value of 0.0 means no blurring will occur,
with higher values increasing the blur factor.
*/
func getImageLayer(sourceImageData image.Image, widthInCharacters int, heightInCharacters int, blurSigma float64) memory.LayerEntryType {
	if widthInCharacters <= 0 && heightInCharacters <= 0 {
		panic(fmt.Sprintf("The specified width and height of %dx%d for your image is not valid.", widthInCharacters, heightInCharacters))
	}
	calculatedPixelWidth := widthInCharacters
	calculatedPixelHeight := heightInCharacters * 2
	if widthInCharacters == 0 {
		calculatedPixelWidth = (heightInCharacters * 2 * sourceImageData.Bounds().Max.X) / sourceImageData.Bounds().Max.Y
	}
	if heightInCharacters == 0 {
		calculatedPixelHeight = (widthInCharacters * sourceImageData.Bounds().Max.Y) / sourceImageData.Bounds().Max.X
	}
	processedImageData := resizeImage(sourceImageData, uint(calculatedPixelWidth), uint(calculatedPixelHeight))
	if blurSigma > 0 {
		processedImageData =  imaging.Blur(processedImageData, blurSigma)
	}
	calculatedCharacterWidth := calculatedPixelWidth
	calculatedCharacterHeight := calculatedPixelHeight / 2
	layerEntry := memory.NewLayerEntry(calculatedCharacterWidth, calculatedCharacterHeight)
	currentImageYLocation := 0
	for currentYLocation := 0; currentYLocation < calculatedCharacterHeight; currentYLocation++ {
		for currentXLocation := 0; currentXLocation < calculatedCharacterWidth; currentXLocation++ {
			currentCharacter := layerEntry.CharacterMemory[currentYLocation][currentXLocation]
			currentCharacter.Character = constants.CharBlockUpperHalf
			upperPixel := processedImageData.At(currentXLocation, currentImageYLocation)
			redColorIndex, greenColorIndex, blueColorIndex, firstAlphaIndex := get8BitColorComponents(upperPixel)
			currentCharacter.AttributeEntry.ForegroundColor = GetRGBColor(int32(redColorIndex), int32(greenColorIndex), int32(blueColorIndex))
			if currentImageYLocation < calculatedCharacterHeight*2 {
				lowerPixel := processedImageData.At(currentXLocation, currentImageYLocation+1)
				redColorIndex, greenColorIndex, blueColorIndex, secondAlphaIndex := get8BitColorComponents(lowerPixel)
				currentCharacter.AttributeEntry.BackgroundColor = GetRGBColor(int32(redColorIndex), int32(greenColorIndex), int32(blueColorIndex))
				if firstAlphaIndex <= 150 || secondAlphaIndex <= 150 {
					currentCharacter.Character = constants.NullRune
				}
			}
			layerEntry.CharacterMemory[currentYLocation][currentXLocation] = currentCharacter
		}
		currentImageYLocation +=2
	}
	return layerEntry
}

/*
get8BitColorComponents allows you to get red, green, and blue color components
from a specific color.
*/
func get8BitColorComponents(colorEntry color.Color) (int32, int32, int32, uint32) {
	redIndex, greenIndex, blueIndex, alphaIndex := colorEntry.RGBA()
	return int32(redIndex) / 257, int32(greenIndex) / 257, int32(blueIndex) / 257, alphaIndex / 257
}

/*
DrawImageToLayer allows you to draw a loaded image to the specified layer.
In addition, the following information should be noted:

- If the location specified falls outside the range for the layer, then
only the visible portion of the image will be drawn.

- If you are drawing an image which has already been pre-rendered, then
your width, height, and blur factor will be ignored.

- If you specify a value of 0 for ether the width or height, then that
dimension will be automatically calculated to a value that best maintain
the images aspect ratio. This is useful since it removes the need to
calculate this manually.

- If you specify a value less than or equal to 0 for both the width and
height, a panic will be generated to fail as fast as possible.

- When pre-rendering an image, it should be noted that each text cell assigned
contains a top and bottom pixel. This is done to provide as much resolution as
possible for your image. That means for a pre-rendered image with a size of
10x10 characters, the actual image being rendered would be 10x20 pixels tall.
If the user wishes to maintain proper aspect ratios, they must manually select
a height that appropriately compensates for this effect, or leave the height
value as 0 to have it done automatically.

- The blur sigma controls how much blurring occurs after your image has been
resized. This allows you to soften your image before it is rendered in ansi
so that hard edges are removed. A value of 0.0 means no blurring will occur,
with higher values increasing the blur factor.
*/
func DrawImageToLayer(layerAlias string, imageAlias string, xLocation int, yLocation int, widthInCharacters int, heightInCharacters int, blurSigma float64) {
	imageLayer := memory.ImageMemory[imageAlias].LayerEntry
	if memory.ImageMemory[imageAlias].ImageData != nil {
		imageData := memory.ImageMemory[imageAlias].ImageData
		imageLayer = getImageLayer(imageData, widthInCharacters, heightInCharacters, blurSigma)
	}
	drawImageToLayer(layerAlias, imageLayer, xLocation, yLocation)
}

/*
drawImageToLayer allows you to draw a loaded image to the specified layer.
 */
func drawImageToLayer(layerAlias string, imageLayer memory.LayerEntryType, xLocation int, yLocation int) {
	layerEntry := memory.GetLayer(layerAlias)
	imageLayer.ScreenXLocation = xLocation
	imageLayer.ScreenYLocation = yLocation
	overlayLayers(&imageLayer, layerEntry)
}
