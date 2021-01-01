package dosktop

import (
	"github.com/supercom32/dosktop/constants"
	"github.com/supercom32/dosktop/internal/math"
	"github.com/supercom32/dosktop/internal/memory"
	"github.com/supercom32/dosktop/internal/stringformat"
	"fmt"
	"github.com/gdamore/tcell"
	"os"
	"os/signal"
	"syscall"
)

/*
defaultValueType is a structure that holds common information about the
current terminal session that needs to be shared.
 */
type defaultValueType struct {
	screen         tcell.Screen
	layerAlias     string // What happens when last layer is deleted? This needs to be updated.
	terminalWidth  int
	terminalHeight int
	screenLayer    memory.LayerEntryType
	isDebugEnabled bool
}

/*
commonResource is a variable used to hold shared data that is accessed
by this package.
*/
var commonResource defaultValueType

/*
InitializeTerminal allows you to initialize Dosktop for the first time.
This method must be called first before any operations take place. The
parameters 'width' and 'height' represent the display size of the
terminal instance you wish to create. In addition, the following
information should be noted:

- If you pass in a zero or negative value for ether width or height a panic
will be generated to fail as fast as possible.
*/
func InitializeTerminal(width int, height int) {
	if width <=0 || height <= 0 {
		panic(fmt.Sprintf("The specified terminal width and height of '%d, %d' is invalid!", width, height))
	}
	memory.InitializeScreenMemory()
	memory.InitializeButtonMemory()
	memory.InitializeImageMemory()
	memory.InitializeTextStyleMemory()
	memory.InitializeTimerMemory()
	commonResource.terminalWidth = width
	commonResource.terminalHeight = height
	if !commonResource.isDebugEnabled {
		screen, err := tcell.NewScreen()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if err := screen.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		commonResource.screen = screen
		commonResource.screen.EnableMouse()
		setupCloseHandler()
		go setupEventUpdater()
	}
}

/*
setupEventUpdater is a background method that monitors all events coming
into the terminal session. When an event is detected, it is recorded and
monitoring continues.
*/
func setupEventUpdater() {
	for {
		PrintDebugLog("EVENT IN - " + getTimeAndDate())
		updateEventQueues()
		PrintDebugLog("EVENT OUT - " + getTimeAndDate())
	}
}

/*
setupCloseHandler enables the trapping of all unexpected system calls and shuts
down the terminal gracefully. This means all terminal settings should be reset
back to normal if anything unexpected happens to the user or if the process is
killed.
*/
func setupCloseHandler() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		<-channel
		commonResource.screen.Fini()
		os.Exit(1)
	}()

}

/*
RestoreTerminalSettings allows the user to gracefully return the terminal
back to its normal settings. This should be called once your application
is finished using Dosktop so that the users terminal environment is not
left in a bad state.
*/
func RestoreTerminalSettings() {
	commonResource.screen.Fini()
}

/*
Inkey allows you to read keyboard input from the user's terminal. This
method returns the character pressed or a keyword representing the
special key pressed (For example: 'a', 'A', 'escape', 'f10', etc.).
In addition, the following information should be noted:

- If more than one keystroke is recorded, it is stored sequentially
in the input buffer and this method needs to be called repeatedly in
order to read them.
*/
func Inkey() string {
	return memory.KeyboardMemory.GetKeystrokeFromKeyboardBuffer()
}

/*
Layer allows you to specify a default layer alias that you wish to use when
interacting with methods which have a non-layer alias method signature.
Non-layer alias method signatures can be identified by finding methods which
have both a layer and non-layer version. This makes interacting with methods
faster, as the user does not need to provide the layer alias context in which
he is working on. For example:

	// On the layer with a layer alias of "MyForegreoundLayer" print a string.
	dosktop.PrintLayer("MyForegroundLayer", "Hello World")

	// Set the default layer alias to be "MyForegroundLayer".
	dosktop.Layer("MyForegroundLayer")

	// Since we set the default layer, we don't need to call the method
	// PrintLayer anymore. Instead, we can use the shorter method Print.
	dosktop.Print("Hello World")
*/
func Layer(layerAlias string) {
	commonResource.layerAlias = layerAlias
}

/*
AddLayer allows you to add a text layer to the current terminal display. You
can add as many layers as you wish to suite your applications needs. Text
layers are useful for setting up windows, modal dialogs, viewports, game
foregrounds and backgrounds, and even effects like parallax scrolling. In
addition, the following information should be noted:

- If you specify location for your layer that is outside the visible
terminal display, then only the visible portion of your text layer will be
rendered. Likewise, if your text layer is larger than the visible area of your
terminal display, then only the visible portion of it will be displayed.

- If you pass in a zero or negative value for ether width or height a panic
will be generated to fail as fast as possible.

- The z order priority controls which text layer should be drawn first and
which text layer should be drawn last. Layers that have a higher priority
will be drawn on top of layers that have a lower priority. In the event
that two layers have the same priority, they will be drawn in random order.
This is to ensure that programmers do not attempt to rely on any specific
behavior that might be a coincidental side effect.

- The parent alias specifies which text layer is the parent of the one being
created. Having a parent layer means that the child layer will only render
on the parent and not the main terminal. This allows you to have text layers
within text layers that can be moved or manipulated relative to the parent.
If you pass in a value of "" for the parent alias, then no parent is used
and the layer is rendered directly to the terminal display. This feature
is useful for creating 'Window' effects where content is contained within
something else.

- When adding a new text layer, it will become the default
working text layer automatically. If you wish to set another text layer
as your default, use 'Layer' to explicitly set it.
*/
func AddLayer(layerAlias string, xLocation int, yLocation int, width int, height int, zOrderPriority int, parentAlias string) {
	if width <= 0 || height <= 0 {
		panic(fmt.Sprintf("Could not create the text layer '%s' since the width and height of (%d, %d) is invalid.", layerAlias, width, height))
	}
	memory.AddLayer(layerAlias, xLocation, yLocation, width, height, zOrderPriority, parentAlias)
	commonResource.layerAlias = layerAlias
}

/*
DeleteLayer allows you to remove a text layer. If you wish to reuse a text
layer for a future purpose, you may also consider making the layer invisible
instead of deleting it. In addition, the following information should be noted:

- When a text layer is deleted, all child text layers are recursively deleted
as well.

- If any dynamically drawn TUI controls reference the deleted layer, they will
still be present. However, because the layer they were created for no longer
exists, they will never be rendered. Consider removing any TUI controls before
deleting the layer they reference. If you delete a layer that is referenced
by dynamic TUI controls, creating a new layer with the same layer alias will
allow them to be rendered again.

- If you attempt to delete a text layer which is currently set as your default
text layer, then a panic will be generated in order to fail as fast as
possible.

- If you attempt to delete a text layer that does not exist, then the operation
will be ignored.
*/
func DeleteLayer(layerAlias string) {
	if memory.IsLayerExists(layerAlias) {
		layerEntry := memory.GetLayer(layerAlias)
		deleteLayer(*layerEntry, false)
	}
}

/*
deleteLayer allows you to remove a text layer. If you wish to reuse a text
layer for a future purpose, you may also consider making the layer invisible
instead of deleting it. In addition, the following information should be noted:

- The parameter 'isChild' is used to determine if the current layer being
deleted is a child or not. If you are calling this method yourself, this value
should always be 'false' since it is only use when called recursively.

- When a text layer is deleted, all child text layers are recursively deleted
as well.

- If any dynamically drawn TUI controls reference the deleted layer, they will
still be present. However, because the layer they were created for no longer
exists, they will never be rendered. Consider removing any TUI controls before
deleting the layer they reference. If you delete a layer that is referenced
by dynamic TUI controls, creating a new layer with the same layer alias will
allow them to be rendered again.

- If you attempt to delete a text layer which is currently set as your default
text layer, then a panic will be generated in order to fail as fast as
possible.
*/
func deleteLayer(layerEntry memory.LayerEntryType, isChild bool) {
	if layerEntry.LayerAlias == commonResource.layerAlias {
		panic(fmt.Sprintf("The text layer '%s' could not be deleted since it is the default text layer!", layerEntry.LayerAlias))
	}
	memory.DeleteLayer(layerEntry.LayerAlias)
}


/*
AddTextStyle allows you to add a new printing style which can be used for
writing dialog text with. Dialog text printing differs from regular
printing, since it allows for "typewriter" drawing effects as well as
changing text attributes on the fly (color, isBold, etc). In addition,
the following information should be noted:

- This method expects you to pass in a 'textStyleEntry' type obtained by
calling 'dosktop.NewTextStyle()'. This entry type contains all the
options available for your text style and can be configured easily
by setting each attribute accordingly.
*/
func AddTextStyle(textStyleAlias string, textStyleEntry memory.TextStyleEntryType) {
	memory.AddTextStyle(textStyleAlias, textStyleEntry)
}

/*
DeleteTextStyle allows you to remove a text style that was added previously.
In addition, the following information should be noted:

- If you attempt to delete an entry that does not exist, then no operation
will be performed.
*/
func DeleteTextStyle(textStyleAlias string) {
	memory.DeleteTextStyle(textStyleAlias)
}

/*
NewTextStyle allows you to obtain a new text style entry which can be
used when printing dialog text. By configuring attributes for your text
style entry and adding your entry to Dosktop, simple markup commands can
be used to switch between dialog printing styles automatically.
For example:

	// Create a new text style entry to configure.
	myTextStyleEntry := dosktop.NewTextStyle()
	// Configure your text style so that the foreground color is red.
	myTextStyleEntry.ForegroundColor = dosktop.GetRGBColor(255, 0, 0)
	// Add a new text style called "RedColor".
	dosktop.AddTextStyle("RedColor", myTextStyleEntry)
*/
func NewTextStyle() memory.TextStyleEntryType {
	return memory.NewTextStyleEntry()
}

/*
NewTuiStyleEntry allows you to obtain a new style entry which can be used
for specifying how TUI controls and other TUI drawing operations should
occur. For example:

	// Create a new TUI style entry to configure.
	myTuiStyleEntry := dosktop.NewTuiStyleEntry()
	// Configure the style entry so that the upper left corner character
	// for drawing a window is the '╔' character. Here we use '\u' to
	// denote the specific byte value in our code since not all editors may
	// know how to display the actual character.
	myTuiStyleEntry.UpperLeftCorner = '\u2554'
	// Draw a window on the text layer with the alias of "ForegroundLayer",
	// using the TUI style entry "myTuiStyleEntry", at layer location (0, 0),
	// with a width and height of 10x10 characters.
	dosktop.DrawWindow("ForegroundLayer", myTuiStyleEntry, 0, 0, 10, 10)
*/
func NewTuiStyleEntry() memory.TuiStyleEntryType {
	return memory.NewTuiStyleEntry()
}

/*
NewSelectionEntry allows you to obtain an entry used for specifying what
options you want to make available for a given menu prompt. For example:

	// Create a new TUI style entry with default settings.
	tuiStyleEntry := dosktop.NewTuiStyleEntry()
	// Create a new selection entry to populate our menu entries with.
	selectionEntry := dosktop.NewSelectionEntry()
	// Add a selection with the alias "Opt1" and a display value of "OK".
	selectionEntry.AddSelection("Opt1", "OK")
	// Add a selection with the alias "Opt2" with the display value of "CANCEL".
	selectionEntry.AddSelection("Opt2", "CANCEL")
	// Prompt the user with a vertical selection menu, on the text layer
	// with the alias "ForegroundLayer", using a default TUI style entry,
	// a selection entry with two options, at the layer location (0, 0),
	// with a menu width and height of 15x15 characters.
	selectionMade := dosktop.GetSelectionFromVerticalMenu ("ForegroundLayer", tuiStyleEntry, selectionEntry, 0, 0, 15, 15)
*/
func NewSelectionEntry() memory.SelectionEntryType {
	return memory.NewSelectionEntry()
}

/*
SetAlpha allows you to set the alpha value for a given text layer. This lets
you perform pseudo transparencies by making the layer foreground and background
colors blend with the layers underneath it to the degree specified. In
addition, the following information should be noted:

- An alpha value of 1.0 is equal to 100% visible, while an alpha value of
0.0 is 0% visible. Specifying a value outside this range indicates that
you want to over amplify or under amplify the color transparency effect.

- If the percent change specified is outside of the RGB color range (for
example, if you specified 200%), then the color will simply bottom or max
out at RGB(0, 0, 0) or RGB(255, 255, 255) respectively.
*/
func SetAlpha(layerAlias string, alphaValue float32) {
	layerEntry := memory.GetLayer(layerAlias)
	layerEntry.DefaultAttribute.ForegroundTransformValue = alphaValue
	layerEntry.DefaultAttribute.BackgroundTransformValue = alphaValue
}

/*
GetColor allows you to obtain an RGB color from a predefined color palette
list. This list corresponds to the 16 color ANSI standard, where color
0 is Black and 15 is Bright White.  In addition, the following information
should be noted:

- If you specify a color index less than 0 or greater than 15 a panic
will be generated to fail as fast as possible.
*/
func GetColor(colorIndex int) int32 {
	if colorIndex < 0 || colorIndex > len(constants.AnsiColorByIndex) {
		panic(fmt.Sprintf("The specified color index '%d' is invalid!", colorIndex))
	}
	return constants.AnsiColorByIndex[colorIndex]
}

/*
GetRGBColor allows you to obtain a specific RGB color based on the red, green, and
blue index values provided. In addition, the following information should be noted:

- If you specify a color channel index less than 0 or greater than 255 a panic
will be generated to fail as fast as possible.
*/
func GetRGBColor(redColorIndex int32, greenColorIndex int32, blueColorIndex int32) int32 {
	if redColorIndex < 0 || redColorIndex > 255 || greenColorIndex < 0 || greenColorIndex > 255 ||
		blueColorIndex < 0 || blueColorIndex > 255 {
		panic(fmt.Sprintf("The specified RGB color index '%d, %d, %d' is invalid!", redColorIndex, greenColorIndex, blueColorIndex))
	}
	return int32(tcell.NewRGBColor(redColorIndex, greenColorIndex, blueColorIndex))
}

/*
Color allows you to set default colors on your text layer for printing with.
The color index specified corresponds to the 16 color ANSI standard, where
color 0 is Black and 15 is Bright White. If you wish to change colors settings
for a text layer that is not currently set as your default, use 'ColorLayer'
instead.
*/
func Color(foregroundColorIndex int, backgroundColorIndex int) {
	ColorLayer(commonResource.layerAlias, foregroundColorIndex, backgroundColorIndex)
}

/*
ColorLayer allows you to set default colors on your specified text layer for
printing with. The color index specified corresponds to the 16 color ANSI
standard, where color 0 is Black and 15 is Bright White. If you do not wish
to specify a text layer, you can use the method 'Color' which will simply
change the color for the default text layer previously set.
 */
func ColorLayer(layerAlias string, foregroundColorIndex int, backgroundColorIndex int) {
	if foregroundColorIndex < 0 || foregroundColorIndex > len(constants.AnsiColorByIndex) {
		panic(fmt.Sprintf("The specified foreground color index '%d' for layer '%s' is invalid!", foregroundColorIndex, layerAlias))
	}
	if backgroundColorIndex < 0 || backgroundColorIndex > len(constants.AnsiColorByIndex) {
		panic(fmt.Sprintf("The specified background color index '%d' for layer '%s' is invalid!", backgroundColorIndex, layerAlias))
	}
	layerEntry := memory.GetLayer(layerAlias)
	layerEntry.DefaultAttribute.ForegroundColor = constants.AnsiColorByIndex[foregroundColorIndex]
	layerEntry.DefaultAttribute.BackgroundColor = constants.AnsiColorByIndex[backgroundColorIndex]
}

/*
ColorRGB allows you to set default colors on your text layer for printing with.
This method allows you to specify colors using RGB color index values within
the range of 0 to 255. If you wish to change colors settings for a text layer
that is not currently set as your default, use 'ColorLayerRGB' instead.
 */
func ColorRGB(foregroundRedIndex int32, foregroundGreenIndex int32, foregroundBlueIndex int32, backgroundRedIndex int32, backgroundGreenIndex int32, backgroundBlueIndex int32) {
	ColorLayerRGB(commonResource.layerAlias, foregroundRedIndex, foregroundGreenIndex, foregroundBlueIndex, backgroundRedIndex, backgroundGreenIndex, backgroundBlueIndex)
}

/*
ColorLayerRGB allows you to set default colors on your specified text layer
for printing with. This method allows you to specify colors using RGB color
index values within the range of 0 to 255. If you do not wish to specify a
text layer, you can use the method 'ColorRGB' which will simply change the
color for the default text layer previously set.
*/
func ColorLayerRGB(layerAlias string, foregroundRed int32, foregroundGreen int32, foregroundBlue int32, backgroundRed int32, backgroundGreen int32, backgroundBlue int32) {
	foregroundColor := GetRGBColor(foregroundRed, foregroundGreen, foregroundBlue)
	backgroundColor := GetRGBColor(backgroundRed, backgroundGreen, backgroundBlue)
	colorLayer24Bit(layerAlias, foregroundColor, backgroundColor)
}

/*
colorLayer24Bit allows you to color a layer using a 24-bit color expressed as
an int32. This is useful for internal methods that already have a 24-bit color
and do not require to compute it again.
*/
func colorLayer24Bit(layerAlias string, foregroundColor int32, backgroundColor int32) {
	layerEntry := memory.GetLayer(layerAlias)
	layerEntry.DefaultAttribute.ForegroundColor = foregroundColor
	layerEntry.DefaultAttribute.BackgroundColor = backgroundColor
}

/*
MoveLayerByAbsoluteValue allows you to move a text layer by an absolute value.
This is useful if you know exactly what position you wish to move your text
layer to. In addition, the following information should be noted:

- If you move your layer outside the visible terminal display, only the visible
display area will be rendered. Likewise, if your text layer is a child of
a parent layer, then only the visible display area will be rendered on the
parent.
*/
func MoveLayerByAbsoluteValue(layerAlias string, xLocation int, yLocation int) {
	layerEntry := memory.GetLayer(layerAlias)
	layerEntry.ScreenXLocation = xLocation
	layerEntry.ScreenYLocation = yLocation
}

/*
MoveLayerByRelativeValue allows you to move a text layer by a relative value.
This is useful for windows, foregrounds, backgrounds, or any kind of
animations or movement you may wish to do in increments. For example:

	// Move the text layer with the alias "ForegroundLayer" one character to
	// the left and two characters down from its current location.
	dosktop.MoveLayerByRelativeValue("ForegroundLayer", -1, 2)

In addition, the following information should be noted:

- If you move your layer outside the visible terminal display, only the visible
display area will be rendered. Likewise, if your text layer is a child of
a parent layer, then only the visible display area will be rendered on the
parent.
*/
func MoveLayerByRelativeValue(layerAlias string, xLocation int, yLocation int) {
	layerEntry := memory.GetLayer(layerAlias)
	layerEntry.ScreenXLocation += xLocation
	layerEntry.ScreenYLocation += yLocation
}

/*
Locate allows you to set the default cursor location on your specified text
layer for printing with. This is useful for when you wish to print text
at different locations of your text layer at any given time. If you wish to
change the cursor location for a text layer that is not currently set as your
default, use 'LocateLayer' instead. In addition, the following information
should be noted:

- If you pass in a location value that falls outside the dimensions of the
default text layer, a panic will be generated to fail as fast as possible.

- Valid text layer locations start at position (0,0) for the upper left corner.
Since location values do not start at (1,1), valid end positions for the bottom
right corner will be one less than the text layer width and height. For
example:

	// Create a new text layer with the alias "ForegroundLayer", at location
	// (0,0), with a width and height of 15x15, a z order priority of 1,
	// and no parent layer associated with it.
	dosktop.AddLayer("ForegroundLayer", 0, 0, 15, 15, 1, "")
	// Set the text layer with the alias "ForegroundLayer" as our default.
	dosktop.Layer("ForegroundLayer")
	// Move our cursor location to the bottom right corner of our text layer.
	dosktop.Locate(14, 14)
*/
func Locate(xLocation int, yLocation int) {
	LocateLayer(commonResource.layerAlias, xLocation, yLocation)
}

/*
Locate allows you to set the default cursor location on your specified text
layer for printing with. This is useful for when you wish to print text
at different locations of your text layer at any given time. If you do not
wish to specify a text layer, you can use the method 'Locate' which will
simply change the cursor location for the default text layer previously set.
In addition, the following information should be noted:

- If you pass in a location value that falls outside the dimensions of the
specified text layer, a panic will be generated to fail as fast as possible.

- Valid text layer locations start at position (0,0) for the upper left corner.
Since location values do not start at (1,1), valid end positions for the bottom
right corner will be one less than the text layer width and height. For
example:

	// Create a new text layer with the alias "ForegroundLayer", at location
	// (0,0), with a width and height of 15x15, a z order priority of 1,
	// and no parent layer associated with it.
	dosktop.AddLayer("ForegroundLayer", 0, 0, 15, 15, 1, "")
	// Move our cursor location to the bottom right corner of our text layer.
	dosktop.LocateLayer(14, 14)
*/
func LocateLayer(layerAlias string, xLocation int, yLocation int) {
	layerEntry := memory.GetLayer(layerAlias)
	if xLocation < 0 || yLocation < 0 ||
		xLocation >= layerEntry.Width || yLocation >= layerEntry.Height {
		panic(fmt.Sprintf("The specified location (%d, %d) is out of bounds for layer '%s' with a size of (%d, %d).", xLocation, yLocation, layerAlias, layerEntry.Width, layerEntry.Height))
	}
	layerEntry.CursorXLocation = xLocation
	layerEntry.CursorYLocation = yLocation
}

/*
Print allows you to write text to the default text layer. If you wish to
print to a text layer that is not currently set as the default, use
'PrintLayer' instead. In addition, the following information should be noted:

- When text is written to the text layer, the cursor position is also updated
to reflect its new location. Like a typewriter, the cursor position moves to
the start of the next line after each print statement.

- If the string to print ends up being too long to fit at its current location,
then only the visible portion of your string will be printed.

- If printing has not yet finished and there are no available lines left, then
all remaining characters will be discarded and printing will stop.
*/
func Print(textToPrint string) {
	PrintLayer(commonResource.layerAlias, textToPrint)
}

/*
PrintLayer allows you to write text to a specified text layer. If you do not
wish to specify a text layer, you can use the method 'Print' which will
simply print to the default text layer previously set. In addition, the
following information should be noted:

- When text is written to the text layer, the cursor position is also updated
to reflect its new location. Like a typewriter, the cursor position moves to
the start of the next line after each print statement.

- If the string to print ends up being too long to fit at its current location,
then only the visible portion of your string will be printed.

- If printing has not yet finished and there are no available lines left, then
all remaining characters will be discarded and printing will stop.
*/
func PrintLayer(layerAlias string, textToPrint string) {
	layerEntry := memory.GetLayer(layerAlias)
	if layerEntry.CursorYLocation >= layerEntry.Height {
		layerEntry.CursorYLocation = layerEntry.Height - 1
		layerEntry.CharacterMemory = scrollCharacterMemory(layerEntry)
	}
	arrayOfRunes := stringformat.GetRunesFromString(textToPrint)
	printLayer(layerEntry, layerEntry.DefaultAttribute, layerEntry.CursorXLocation, layerEntry.CursorYLocation, arrayOfRunes)
	layerEntry.CursorXLocation = 0
	layerEntry.CursorYLocation = layerEntry.CursorYLocation + 1
}

/*
printLayer allows you to write text to a text layer. This is useful
for internal methods that want to write text to a text layer directly, without
effecting user settings (such as current cursor location, etc). In addition,
the following information should be noted:

- If the location to print falls outside of the range of the text layer,
then only the visible portion of your text will be printed.
*/
func printLayer(layerEntry *memory.LayerEntryType, attributeEntry memory.AttributeEntryType, xLocation int, yLocation int, textToPrint []rune) {
	layerWidth := layerEntry.Width
	layerHeight := layerEntry.Height
	cursorXLocation := xLocation
	cursorYLocation := yLocation
	characterMemory := layerEntry.CharacterMemory
	for _, currentCharacter := range textToPrint {
		if cursorXLocation >= 0 && cursorXLocation < layerWidth && cursorYLocation >= 0 && cursorYLocation < layerHeight {
			characterMemory[cursorYLocation][cursorXLocation].AttributeEntry = memory.NewAttributeEntry(&attributeEntry)
			characterMemory[cursorYLocation][cursorXLocation].Character = currentCharacter
			characterMemory[cursorYLocation][cursorXLocation].LayerAlias = layerEntry.LayerAlias
		}
		if stringformat.IsRuneCharacterWide(currentCharacter) {
			cursorXLocation += 2
		} else {
			cursorXLocation++
		}
		if cursorXLocation >= layerWidth {
			return
		}
	}
}

/*
Clear allows you to empty the default text layer of all its contents. If you
wish to clear a text layer that is not currently set as the default, use
'ClearLayer' instead.
*/
func Clear() {
	ClearLayer(commonResource.layerAlias)
}

/*
Clear allows you to empty the specified text layer of all its contents. If you
do not wish to specify a text layer, you can use the method 'Clear' which will
simply clear the default text layer previously set.
*/
func ClearLayer(layerAlias string) {
	layerEntry := memory.GetLayer(layerAlias)
	clearLayer(layerEntry)
}

/*
clearLayer allows you to empty the specified text layer of all its contents.
This is useful for internal methods that want to clear a text layer directly.
*/
func clearLayer(layerEntry *memory.LayerEntryType) {
	*layerEntry = memory.NewLayerEntry(layerEntry.Width, layerEntry.Height)
}

/*
scrollCharacterMemory allows you to advance the specified text layer up by one
row. This means that the first row is discarded and all subsequent rows are
moved up by one position. The new row created at the bottom of the text layer
will be filled with spaces (" ") colored with the layers default attributes.
*/
func scrollCharacterMemory(layerEntry *memory.LayerEntryType) [][]memory.CharacterEntryType {
	layerWidth := layerEntry.Width
	characterMemory := layerEntry.CharacterMemory
	characterMemory = characterMemory[1:]
	characterObjectArray := make([]memory.CharacterEntryType, layerWidth)
	for currentCharacterCell := 0; currentCharacterCell < layerWidth; currentCharacterCell++ {
		characterEntry := memory.NewCharacterEntry()
		characterEntry.AttributeEntry = memory.NewAttributeEntry(&layerEntry.DefaultAttribute)
		characterEntry.Character = ' '
		characterObjectArray[currentCharacterCell] = characterEntry
	}
	characterMemory = append(characterMemory, characterObjectArray)
	layerEntry.CharacterMemory = characterMemory
	return characterMemory
}

/*
getRuneOnLayer allows you to obtain a specific rune at the location specified
on the given text layer. In addition, the following information should be
noted:

- If the location specified is outside the valid range of the text layer, then
a panic will be thrown to fail as fast as possible.
*/
func getRuneOnLayer(layerEntry *memory.LayerEntryType, xLocation int, yLocation int) rune {
	if xLocation < 0 || xLocation >= layerEntry.Width || yLocation < 0 || yLocation >= layerEntry.Height {
		panic(fmt.Sprintf("The specified location (%d, %d) is out of bounds for the layer entry with a size of (%d, %d).", xLocation, yLocation, layerEntry.Width, layerEntry.Height))
	}
	characterMemory := layerEntry.CharacterMemory
	return characterMemory[yLocation][xLocation].Character
}

/*
GetCellIdUnderMouseLocation allows you to obtain the cell ID for the text
directly under your mouse cursor. This is useful for tracking
elements on a screen, creating "hot spots", or interactive zones which you
want the user to interact with. In addition, the following information should be
noted:

- If multiple text layers are being displayed, the cell ID returned will be
from the top-most visible text cell.

- The cell ID returned will only reflect what is currently being displayed
on the terminal display. If you wish for any new changes to take effect,
call 'UpdateDisplay' to refresh the visible display area first.
 */
func GetCellIdUnderMouseLocation(layerAlias string) int {
	mouseXLocation, mouseYLocation, _, _ := memory.MouseMemory.GetMouseStatus()
	return getCellIdByLayerEntry(&commonResource.screenLayer, mouseXLocation, mouseYLocation)
}

/*
getCellIdByLayerAlias allows you to obtain a cell ID from a given text layer
by layer alias. This is simply a wrapper method that converts the text
layer alias into a layer entry and calls 'getCellIdByLayerEntry'.
*/
func getCellIdByLayerAlias(layerAlias string, mouseXLocation int, mouseYLocation int) int {
	layerEntry := memory.GetLayer(layerAlias)
	return getCellIdByLayerEntry(layerEntry, mouseXLocation, mouseYLocation)
}

/*
getCellIdByLayerEntry allows you to obtain a cell ID from a given text layer
by layer entry. In addition, the following information should be noted:

- If the location specified is outside the valid range of the text layer, then
a value of '-1' is returned instead.
*/
func getCellIdByLayerEntry(layerEntry *memory.LayerEntryType, xLocation int, yLocation int) int {
	returnValue := -1
	if xLocation < 0 || xLocation >= layerEntry.Width || yLocation < 0 || yLocation >= layerEntry.Height {
		return returnValue
	}
	if yLocation-layerEntry.ScreenYLocation >= 0 && xLocation-layerEntry.ScreenXLocation >= 0 &&
		yLocation-layerEntry.ScreenYLocation < len(layerEntry.CharacterMemory) && xLocation-layerEntry.ScreenXLocation < len(layerEntry.CharacterMemory[0]) {
		characterEntry := layerEntry.CharacterMemory[yLocation-layerEntry.ScreenYLocation][xLocation-layerEntry.ScreenXLocation]
		returnValue = characterEntry.AttributeEntry.CellId
	}
	return returnValue
}

/*
UpdateDisplay allows you to synchronize the terminals visible display area with
your current changes. In addition, the following information should be noted:

- All text layers are sorted from lowest to highest z order priority level.

- Layers with the same z order priority will appear in random display order.
This is to ensure that programmers do not attempt to rely on any specific
behavior that might be a coincidental side effect.
*/
func UpdateDisplay() {
	sortedLayerAliasSlice := memory.GetSortedLayerMemoryAliasSlice()
	baseLayerEntry := memory.NewLayerEntry(commonResource.terminalWidth, commonResource.terminalHeight)
	baseLayerEntry = renderLayers(&baseLayerEntry, sortedLayerAliasSlice)
	DrawLayerToScreen(&baseLayerEntry, false)
	commonResource.screenLayer = baseLayerEntry
}

/*
renderLayers allows you to render a list of text layers to the specified root
text layer. In addition, the following information should be noted:

- The root layer entry is considered the parent entry. Only text layers under
the specified parent will be rendered on it.

- If a text layer being rendered is a parent, then all child text layers will
be rendered on the parent before the parent is drawn. This is done by making
a recursive call to 'renderLayers' with the new parent layer.

- The list of text layers alias provided should be sorted so that layers with
a lower z order priority are rendered first.

- Any text layer which is marked as not visible will be ignored.

- All rendering occurs on a temporary text layer until it is ready to be
overlaid on the final (terminal ready) text layer. Buttons and other special
TUI controls are also dynamically rendered at this time so that the original
text layer data underneath them is preserved.
*/
func renderLayers(rootLayerEntry *memory.LayerEntryType, sortedLayerAliasSlice memory.LayerAliasZOrderPairList) memory.LayerEntryType {
	baseLayerEntry := memory.NewLayerEntry(0,0, rootLayerEntry)
	for currentListIndex := 0; currentListIndex < len(sortedLayerAliasSlice); currentListIndex++ {
		currentLayerEntry := memory.NewLayerEntry(0, 0, memory.GetLayer(sortedLayerAliasSlice[currentListIndex].Key))
		if currentLayerEntry.IsVisible {
			drawButtonsOnLayer(currentLayerEntry)
			if currentLayerEntry.IsParent && (currentLayerEntry.LayerAlias != baseLayerEntry.LayerAlias && currentLayerEntry.ParentAlias == baseLayerEntry.LayerAlias){
				renderedLayer := renderLayers(&currentLayerEntry, sortedLayerAliasSlice)
				overlayLayers(&renderedLayer, &baseLayerEntry)
			} else {
				if currentLayerEntry.ParentAlias == baseLayerEntry.LayerAlias {
					overlayLayers(&currentLayerEntry, &baseLayerEntry)
				}
			}
		}
	}
	return baseLayerEntry
}

/*
overlayLayersByLayerAlias allows you to overlay a text layer by its layer
alias. This is useful when you do not have actual layer data and only
know the alias of the layer you wish to overlay.
*/
func overlayLayersByLayerAlias(sourceLayerAlias string, targetLayerEntry *memory.LayerEntryType) {
	layerEntry := memory.GetLayer(sourceLayerAlias)
	overlayLayers(layerEntry, targetLayerEntry)
}

/*
overlayLayers allows you to overlay one text layer on top of another text
layer. In addition, the following information should be noted:

- If the source text layer is set to be drawn outside the target layer,
then only the visible portion of the source text layer will be rendered.

- If the source text layer is found to be completely outside the range
of the target layer, then no rendering will occur.

- If the source rune to be drawn is null, then it will be considered
transparent.

- If a transparent rune has a foreground or background alpha value set,
then it will be drawn as a shadow with the color and intensity matching
the rune underneath it.
*/
func overlayLayers(sourceLayerEntry *memory.LayerEntryType, targetLayerEntry *memory.LayerEntryType) {
	sourceCharacterMemory := sourceLayerEntry.CharacterMemory
	targetCharacterMemory := targetLayerEntry.CharacterMemory
	sourceWidthToCopy := sourceLayerEntry.Width
	sourceHeightToCopy := sourceLayerEntry.Height
	// Calculate how much of the source Width to copy.
	sourceWidthToCopy = sourceLayerEntry.Width - int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenXLocation))
	if sourceLayerEntry.ScreenXLocation < 0 {
		if sourceWidthToCopy > targetLayerEntry.Width {
			sourceWidthToCopy = targetLayerEntry.Width
		}
	} else {
		if sourceWidthToCopy < targetLayerEntry.Width {
			sourceWidthToCopy = sourceLayerEntry.Width
		}
	}
	// Calculate how much of the source Height to copy.
	sourceHeightToCopy = sourceLayerEntry.Height - int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenYLocation))
	if sourceLayerEntry.ScreenYLocation < 0 {
		if sourceHeightToCopy > targetLayerEntry.Height {
			sourceHeightToCopy = targetLayerEntry.Height
		}
	} else {
		if sourceHeightToCopy < targetLayerEntry.Height {
			sourceHeightToCopy = sourceLayerEntry.Height
		}
	}
	// Adjust where rendering on the layer should start.
	startingSourceXLocation := 0
	startingSourceYLocation := 0
	startingTargetXLocation := 0
	startingTargetYLocation := 0
	if sourceLayerEntry.ScreenXLocation < 0 {
		startingSourceXLocation = int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenXLocation))
	} else {
		startingTargetXLocation = int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenXLocation))
	}
	if sourceLayerEntry.ScreenYLocation < 0 {
		startingSourceYLocation = int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenYLocation))
	} else {
		startingTargetYLocation = int(math.GetAbsoluteValueAsFloat64(sourceLayerEntry.ScreenYLocation))
	}
	if sourceWidthToCopy+startingTargetXLocation > targetLayerEntry.Width {
		sourceWidthToCopy = targetLayerEntry.Width - startingTargetXLocation
	}
	if sourceHeightToCopy+startingTargetYLocation > targetLayerEntry.Height {
		sourceHeightToCopy = targetLayerEntry.Height - startingTargetYLocation
	}
	// If the layer is totally off screen, don't bother to render it.
	if startingSourceXLocation+sourceWidthToCopy < 0 || sourceLayerEntry.ScreenXLocation+sourceWidthToCopy > targetLayerEntry.Width ||
		startingSourceYLocation+sourceHeightToCopy < 0 || sourceLayerEntry.ScreenYLocation+sourceHeightToCopy > targetLayerEntry.Height {
		return
	}
	// Perform the actual copy using the starting offsets previously calculated.
	for currentRow := 0; currentRow < sourceHeightToCopy; currentRow++ {
		for currentColumn := 0; currentColumn < sourceWidthToCopy; currentColumn++ {
			sourceCharacterEntry := &sourceCharacterMemory[currentRow+startingSourceYLocation][currentColumn+startingSourceXLocation]
			targetCharacterEntry := &targetCharacterMemory[currentRow+startingTargetYLocation][currentColumn+startingTargetXLocation]
			sourceAttributeEntry := sourceCharacterEntry.AttributeEntry
			targetAttributeEntry := targetCharacterEntry.AttributeEntry
			// Handle transformations
			if sourceCharacterEntry.Character == constants.NullRune {
				if sourceAttributeEntry.ForegroundTransformValue < 1 {
					targetAttributeEntry.ForegroundColor = GetTransitionedColor(targetAttributeEntry.ForegroundColor, GetRGBColor(0, 0, 0), sourceAttributeEntry.ForegroundTransformValue)
				}
				if sourceAttributeEntry.BackgroundTransformValue < 1 {
					targetAttributeEntry.BackgroundColor = GetTransitionedColor(targetAttributeEntry.BackgroundColor, GetRGBColor(0, 0, 0), sourceAttributeEntry.BackgroundTransformValue)
				}
				targetCharacterEntry.AttributeEntry = targetAttributeEntry
				targetCharacterMemory[currentRow+startingTargetYLocation][currentColumn+startingTargetXLocation] = *targetCharacterEntry
			} else {
				targetCharacterEntry.AttributeEntry = memory.NewAttributeEntry(&sourceAttributeEntry)
				targetCharacterEntry.Character = sourceCharacterEntry.Character
				targetCharacterEntry.LayerAlias = sourceCharacterEntry.LayerAlias
				// If there is no local color transforming being done on cells
				if sourceAttributeEntry.ForegroundTransformValue != 1 || sourceAttributeEntry.BackgroundTransformValue != 1 {
					if sourceAttributeEntry.ForegroundTransformValue < 1 {
						targetCharacterEntry.AttributeEntry.ForegroundColor = GetTransitionedColor(targetAttributeEntry.ForegroundColor, sourceAttributeEntry.ForegroundColor, sourceAttributeEntry.ForegroundTransformValue)
					}
					if sourceAttributeEntry.BackgroundTransformValue < 1 {
						targetCharacterEntry.AttributeEntry.BackgroundColor = GetTransitionedColor(targetAttributeEntry.BackgroundColor, sourceAttributeEntry.BackgroundColor, sourceAttributeEntry.BackgroundTransformValue)
					}
				} else {
					if sourceLayerEntry.DefaultAttribute.ForegroundTransformValue < 1 {
						targetCharacterEntry.AttributeEntry.ForegroundColor = GetTransitionedColor(targetAttributeEntry.ForegroundColor, sourceAttributeEntry.ForegroundColor, sourceLayerEntry.DefaultAttribute.ForegroundTransformValue)
					}
					if sourceLayerEntry.DefaultAttribute.BackgroundTransformValue < 1 {
						targetCharacterEntry.AttributeEntry.BackgroundColor = GetTransitionedColor(targetAttributeEntry.BackgroundColor, sourceAttributeEntry.BackgroundColor, sourceLayerEntry.DefaultAttribute.BackgroundTransformValue)
					}
				}
			}
		}
	}
}

/*
DrawLayerToScreen allows you to render a text layer to the visible terminal
screen. If debug is enabled, this method does nothing since the terminal
is virtual.
*/
func DrawLayerToScreen(layerEntry *memory.LayerEntryType, isForcedRefreshRequired bool) {
	if !commonResource.isDebugEnabled {
		width := layerEntry.Width
		height := layerEntry.Height
		for currentRow := 0; currentRow < height; currentRow++ {
			for currentCharacter := 0; currentCharacter < width; currentCharacter++ {
				style := tcell.StyleDefault
				attributeEntry := layerEntry.CharacterMemory[currentRow][currentCharacter].AttributeEntry
				style = style.Foreground(tcell.Color(attributeEntry.ForegroundColor))
				style = style.Background(tcell.Color(attributeEntry.BackgroundColor))
				style = style.Blink(attributeEntry.IsBlinking)
				style = style.Bold(attributeEntry.IsBold)
				style = style.Reverse(attributeEntry.IsReversed)
				style = style.Underline(attributeEntry.IsUnderlined)
				var character = layerEntry.CharacterMemory[currentRow][currentCharacter].Character
				r2 := []rune("")
				commonResource.screen.SetContent(currentCharacter, currentRow, character, r2, style)
			}
		}
		commonResource.screen.Show()
	}
}
