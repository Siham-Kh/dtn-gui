package helpers

import (
	"github.com/gotk3/gotk3/gtk"
)

func fill(progress *gtk.ProgressBar, fraction float64) bool {

	/*Fill in the bar with the new fraction*/
	progress.SetFraction(fraction)

	/*Ensures that the fraction stays below 1.0*/
	if fraction < 1.0 {
		return true
	} else {
		return false
	}
}

func AppLabel() *gtk.Label {
	label, _ := gtk.LabelNew("Data Transfer node")
	// label.ModifyFontEasy("DejaVu Serif 15")
	return label
}

func authors() []string {
	return []string{"Siham Khoussi <siham.khoussi@nist.gov>", "Abdella Battou <abdella.battou@nist.gov>"}
}
