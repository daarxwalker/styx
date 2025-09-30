package styx

const (
	After                     = "::after"
	Backdrop                  = "::backdrop"
	Before                    = "::before"
	Column                    = "::column"
	Checkmark                 = "::checkmark"
	DetailsContent            = "::details-content"
	FileSelectorButton        = "::file-selector-button"
	FirstLetter               = "::first-letter"
	FirstLine                 = "::first-line"
	GrammarError              = "::grammar-error"
	MarkerPseudo              = "::marker"
	Picker                    = "::picker()"
	PickerIcon                = "::picker-icon"
	Placeholder               = "::placeholder"
	ScrollButton              = "::scroll-button()"
	ScrollMarker              = "::scroll-marker"
	ScrollMarkerGroupPseudo   = "::scroll-marker-group"
	Selection                 = "::selection"
	SpellingError             = "::spelling-error"
	TargetText                = "::target-text"
	ViewTransition            = "::view-transition"
	ViewTransitionGroup       = "::view-transition-group()"
	ViewTransitionNew         = "::view-transition-new()"
	ViewTransitionOld         = "::view-transition-old()"
	MozRangeTrack             = "::-moz-range-track"
	MozRangeThumb             = "::-moz-range-thumb"
	WebKitSliderRunnableTrack = "::-webkit-slider-runnable-track"
	WebKitSliderThumb         = "::-webkit-slider-thumb"
)

func Cue(selector string) string {
	if selector == "" {
		return "::cue"
	}
	return "::cue(" + selector + ")"
}

func Highlight(name string) string {
	return "::highlight(" + name + ")"
}

func Part(name string) string {
	return "::part(" + name + ")"
}

func Slotted(selector string) string {
	return "::slotted(" + selector + ")"
}

func ViewTransitionImagePair(name string) string {
	return "::view-transition-image-pair(" + name + ")"
}
