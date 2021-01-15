/*
Dosktop is a package which allow users to easily create rich text-based terminal
applications and games. Dosktop differs from other terminal packages by
providing an extremely simple API that abstracts out all low-level terminal
operations, while avoiding complicated event-driven TUI designs.

The BASICs

If your familiar with programming on classical computers and OS environments
like the Commodore 64 or MS-DOS, you will find that Dosktop behaves similar
to these systems. Dosktop is a procedural package that gives users control
over the terminal with simple BASIC-like commands. It maintains this
simplicity by managing much of it's functionality and resources internally.
While users do not need to keep track of these resources themselves, there
are still a few concepts they need to be familiar with.

Aliases

Aliases are a way for a user to reference a resource in Dosktop without
actually needing to manage and store it themselves. Things like creating
a text layer, a button, or a timer, all use aliases to identify what the
user creates and how they can reference that resource at a later time.
For example:

	// Create a new text layer with the layer alias "Foreground", placed at
	// screen location (0, 0), with a width and height of 20x20 characters,
	// a zOrder drawing priority of 0, with no parent layer associated with it.
	dosktop.AddLayer("Foreground", 0, 0, 20, 20, 0, "")

Once created, the user no longer needs to manage the resource. If they
wish to manipulate it, they simply need to reference the Alias they want
to access. For example:

	// Move the text layer with the alias "Foreground" to screen X and Y
	// location (2, 2).
	dosktop.MoveLayerByAbsoluteValue("Foreground", 2, 2)

Furthermore, an alias is always unique and distinct for any given
resource. That means, while it is not recommended, it is possible to give
two different types of resources the same alias name. For example:

	// Create a timer with the timer alias "MyAliasName", with a duration of
	// 1000 milliseconds, with it's enabled status as being "true".
	dosktop.AddTimer("Foreground", 1000, true)

Now we have a layer that has an alias called "Foreground" and a timer
that has an alias called "Foreground".

Attribute Entries

Often when working with software, flexibility comes at the expense of
simplicity. Having methods with dozens of options and parameters makes
functionality more flexible, but also much more difficult to manage and
understand. That's why to keep things simple, Dosktop uses Attribute Entries
for configuring most customizable features. Attribute Entries are simply
structures that hold all kinds of information about how the user would like for
something to operate. By simply generating an entry, configuring it, and
passing it in as a parameter, functionality will automatically know how to
behave if one should want to do something outside the default. In addition,
if you want a feature to behave differently under specific cases, you can
configure multiple attribute entries and simply provide which one you need at
any given time. For example:

	// Create a style entry which lets you configure most TUI related features.
	styleEntry := dosktop.NewTuiStyleEntry()

	// Configure the foreground color to use when rendering TUI buttons.
	styleEntry.ButtonForegroundColor = dosktop.GetRGBColor(0, 255, 0)

	// Add a button to a layer with the alias called "Foreground", with a
	// button alias called "OkButton", with a button label of "OK", with
	// our created style entry, with a layer location of (2,2), with
	// a button width and height of 10x3 characters.
	dosktop.AddButton("Foreground", "OkButton", "OK", styleEntry, 2, 2, 10, 3)

You will note that not all style attributes need to be set. Any attributes
which are not modified will be left at their default settings. In addition,
if you wish to quickly create multiple attribute entries based off a single
one, you can simply do the following:

	// Create a new text style entry for printing text dialog messages.
	defaultTextStyle := dosktop.NewTextStyleEntry()

	// Configure some attributes for our new entry.
	defaultTextStyle.ForegroundColor = dosktop.GetRGBColor(255, 0, 0)

	// Create a new text style entry by cloning an existing style entry.
	boldTextStyle := dosktop.NewTuiStyleEntry(defaultTextStyle)

	// Set our new style to include the bold attribute
	defaultTextStyle.isBold = true

As you can see, cloning attribute entries is a speedy way to configure
multiple items with similar base settings.

Failing Fast

Almost all functionality in Dosktop is designed to perform consistently.
That is, it will perform most tasks predictably without external runtime
conditions changing its behaviour. As a result, when a problem occurs
Dosktop will usually prefer to throw a Panic rather than choose an alternative
solution or throw an error. This is because any problems encountered will
generally be a developer issue and it should be addressed rather than covered
up or hidden by default behaviour. For example:

	// Attempt to initialize Dosktop with a terminal size of (-1, 0)
	dosktop.InitializeTerminal(-1, 0)

In this case, the user is attempting to create a terminal session with
invalid dimensions. Rather than choose a sane default or return a recoverable
error, Dosktop will panic since clearly the developer made an error and should
correct it. For functionality that could vary depending on runtime conditions,
errors will be returned as expected.

Double buffering

When updating the screen with with rapid changes or changes that may take a
long time, it is desirable to wait until all drawing is completed before
performing a refresh. This allows you to eliminate flickering, screen
tearing, or other artifacts that may appear when display data is being
updated at the same time as a screen refresh. In Dosktop, all changes to
the terminal remain in memory until the user calls 'UpdateDisplay' to refresh
the screen. This allows the user to make as many changes as they want before
writing their changes to screen in a single pass.
*/
package dosktop

