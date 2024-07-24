package internal

import at "github.com/Matej-Chmel/go-any-to-string"

// Embeds Options for any-to-string library.
// Provides methods for changing options by both
// Interview and Interview2 structs.
type EmbeddedOptions struct {
	options *at.Options
}

// Constructs new EmbeddedOptions
func NewEmbeddedOptions() EmbeddedOptions {
	return EmbeddedOptions{
		options: at.NewOptions(),
	}
}

// Returns a pointer to the underlying options
func (e *EmbeddedOptions) GetOptions() *at.Options {
	return e.options
}

// Sets the underlying options.
// If nil is passed, options are set to a default value.
func (e *EmbeddedOptions) SetOptions(val *at.Options) {
	if val == nil {
		e.options = at.NewOptions()
	} else {
		e.options = val
	}
}

// Changes options so that byte, uint8, rune and int32 are all
// printed as characters
func (e *EmbeddedOptions) ShowBytesAsString() {
	e.options.ByteAsString = true
	e.options.RuneAsString = true
}

// Changes options so that field names of structs
// in input and output are displayed
func (e *EmbeddedOptions) ShowFieldNames() {
	e.options.ShowFieldNames = true
}
