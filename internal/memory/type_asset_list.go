package memory

type ImageListEntryType struct {
	FileName string
	FileAlias string
}

type PreloadedImageListEntryType struct {
	FileName string
	FileAlias string
	WidthInCharacters int
	HeightInCharacters int
	BlurSigma float64
}

type AssetListType struct {
	PreloadedImageList []PreloadedImageListEntryType
	ImageList []ImageListEntryType
}

func NewAssetList() AssetListType {
	var assetList AssetListType
	return assetList
}

func (shared *AssetListType) AddImage(fileName string, fileAlias string) {
	var newImageListEntryType ImageListEntryType
	newImageListEntryType.FileName = fileName
	newImageListEntryType.FileAlias = fileAlias
	shared.ImageList = append(shared.ImageList, newImageListEntryType)
}

func (shared *AssetListType) AddPreloadedImage(fileName string, fileAlias string, widthInCharacters int, heightInCharacters int, blurSigma float64) {
	var preloadedImageListEntryType PreloadedImageListEntryType
	preloadedImageListEntryType.FileName = fileName
	preloadedImageListEntryType.FileAlias = fileAlias
	preloadedImageListEntryType.WidthInCharacters = widthInCharacters
	preloadedImageListEntryType.HeightInCharacters = heightInCharacters
	preloadedImageListEntryType.BlurSigma = blurSigma
	shared.PreloadedImageList = append(shared.PreloadedImageList, preloadedImageListEntryType)
}

func (shared *AssetListType) Clear() {
	shared.ImageList = nil
	shared.PreloadedImageList = nil
}
