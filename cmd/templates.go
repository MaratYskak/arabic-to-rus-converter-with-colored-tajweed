package main

import (
	"html/template"
	"path/filepath"
)

type templateData struct {
	Letter  string
	Tajweed string
	Space   bool
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	// Use the filepath.Glob function to get a slice of all filepaths with
	// the extension '.page.tmpl'. This essentially gives us a slice of all the
	// 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}
	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file path
		// and assign it to the name variable.
		name := filepath.Base(page)
		// Parse the page template file in to a template set.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Use the ParseGlob method to add any 'layout' templates to the
		// template set (in our case, it's just the 'base' layout at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}
		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}
		// Add the template set to the cache, using the name of the page
		// (like 'home.page.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}

func (app *application) initAlphabet() map[rune]string {
	return map[rune]string{
		'ّ': "shadda",
		'ٌ': "ун",
		'ٍ': "ин",
		'ً': "ан",
		'ة': "т",
		' ': " ",
		'َ': "а",
		'ِ': "и",
		'ُ': "у",
		'ا': "а",
		'ب': "б",
		'ت': "т",
		'ث': "с?",
		'ق': "қ",
		'ج': "j",
		'ح': "х",
		'خ': "х!",
		'ف': "ф",
		'ي': "й",
		'أ': "а",
		'إ': "и",
		'د': "д",
		'ذ': "з?",
		'ل': "л",
		'ر': "р",
		'ز': "з",
		'ض': "д!",
		'ص': "с!",
		'و': "w",
		'ه': "h",
		'م': "м",
		'ك': "к",
		'ن': "н",
		'ع': "'",
		'غ': "ғ",
		'س': "с",
		'ش': "ш",
		'ط': "т!",
		'ظ': "з!",
	}
}

func (app *application) initIhfa() map[rune]bool {
	return map[rune]bool{
		'ت': true,
		'ث': true,
		'ج': true,
		'د': true,
		'ذ': true,
		'س': true,
		'ش': true,
		'ص': true,
		'ض': true,
		'ط': true,
		'ظ': true,
		'ف': true,
		'ق': true,
		'ك': true,
		'ز': true,
	}
}

func (app *application) initHards() map[rune]bool {
	return map[rune]bool{
		'ض': true,
		'ص': true,
		'ق': true,
		'غ': true,
		'خ': true,
		'ظ': true,
		'ط': true,
	}
}

func (app *application) initQamariya() map[rune]bool {
	return map[rune]bool{
		'أ': true,
		'إ': true,
		'آ': true,
		'ق': true,
		'ف': true,
		'غ': true,
		'ع': true,
		'ه': true,
		'خ': true,
		'ح': true,
		'ج': true,
		'ي': true,
		'ب': true,
		'م': true,
		'ك': true,
		'و': true,
	}
}

func (app *application) initDataSlice() []*templateData {
	return []*templateData{}
}
