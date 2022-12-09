package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "home.page.html", nil)
}

func (app *application) howto(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/howto" {
		app.notFound(w)
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "howto.page.html", nil)
}

func (app *application) result(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/result" {
		app.notFound(w)
		return
	}

	text := r.FormValue("input")

	DataSlice := tdSlice{}

	ArabicText := []rune(text)
	// colors := map[string]string{

	// 	"black":   "\033[1;30m",
	// 	"red":     "\033[1;31m",
	// 	"green":   "\033[1;32m",
	// 	"yellow":  "\033[1;33m",
	// 	"blue":    "\033[1;94m",
	// 	"purple":  "\033[38;5;56m",
	// 	"magenta": "\033[1;35m",
	// 	"teal":    "\033[1;36m",
	// 	"white":   "\033[1;37m",
	// 	"orange":  "\033[38;5;208m",
	// 	"clear":   "\033[0m",
	// }
	mapa := map[rune]string{
		'ّ': "shadda",
		'ٌ': "ун",
		'ٍ': "ин",
		'ً': "ан",
		' ': " ",
		'َ': "а",
		'ِ': "и",
		'ُ': "у",
		'ا': "а",
		'ب': "б",
		'ت': "т",
		'ث': "с?",
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
		'و': "у",
		'ه': "h",
		'م': "м",
		'ن': "н",
		'ع': "'",
		'غ': "ғ",
		'س': "с",
		'ش': "ш",
		'ط': "т!",
		'ظ': "з!",
	}

	ihfa := map[rune]bool{
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

	qamariya := map[rune]bool{
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

	// result := ""
	skip := 0
	for i, v := range ArabicText {
		if skip > 0 {
			skip--
			continue
		}
		//конец слова
		if i == len(ArabicText)-1 {
			break
		}
		//я
		if v == 'ي' && ArabicText[i+1] == 'َ' {
			skip = 1
			DataSlice = append(DataSlice, &templateData{"я", "", false})
			// result += "я"
			continue
		}
		//ю
		if v == 'ي' && ArabicText[i+1] == 'ُ' {
			skip = 1
			DataSlice = append(DataSlice, &templateData{"ю", "", false})
			// result += "я"
			continue
		}
		//ихфа. нун в конце слова и без огласовки
		if v == 'ن' && ArabicText[i+1] == ' ' && ihfa[ArabicText[i+2]] {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"н-", "ihfa", false})
			continue
		}
		//ихфа. нун в конце слова с сукуном
		if v == 'ن' && ArabicText[i+1] == 'ْ' && ArabicText[i+2] == ' ' && ihfa[ArabicText[i+3]] {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"н-", "ihfa", false})
			continue
		}
		//ихфа. фатха танвин
		if v == 'ً' && ArabicText[i+1] == ' ' && ihfa[ArabicText[i+2]] {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"ан-", "ihfa", false})
			continue
		}
		//ихфа. кясра танвин
		if v == 'ٍ' && ArabicText[i+1] == ' ' && ihfa[ArabicText[i+2]] {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"ин-", "ihfa", false})
			continue
		}
		//ихфа. дамма танвин
		if v == 'ٌ' && ArabicText[i+1] == ' ' && ihfa[ArabicText[i+2]] {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"ун-", "ihfa", false})
			continue
		}
		//шадда
		if (v == 'َ' || v == 'ِ' || v == 'ُ') && ArabicText[i+1] == 'ّ' {
			DataSlice = append(DataSlice, &templateData{mapa[ArabicText[i-1]], "", false})
			DataSlice = append(DataSlice, &templateData{mapa[v], "", false})
			skip = 1
			continue
		}
		//гунна нун с шаддой
		if ArabicText[i+1] == 'ّ' && v == 'ن' {
			DataSlice = append(DataSlice, &templateData{"н", "gunna", false})
			skip = 1
			continue
		}
		//гунна мим мим
		if v == 'م' && ArabicText[i+1] == ' ' && ArabicText[i+2] == 'م' {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"м-м", "gunna", false})
			// result += gchalk.Green("м-м")
			// fmt.Print("\033[0m")
			continue
		}
		//гунна нун ба
		if v == 'ن' && ArabicText[i+1] == ' ' && ArabicText[i+2] == 'ب' {
			skip = 2
			DataSlice = append(DataSlice, &templateData{"м-м", "gunna", false})
			DataSlice = append(DataSlice, &templateData{"б", "", false})
			// result += gchalk.Green("м-мб")
			// fmt.Print("\033[0m")
			continue
		}
		if i < len(ArabicText)-5 {
			if v == 'إ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"и", "", false})
				// result += "и"
				continue
			}
			if v == 'أ' && ArabicText[i+1] == 'َ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"а", "", false})
				// result += "а"
				continue
			}
			if v == 'أ' && ArabicText[i+1] == 'ُ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"у", "", false})
				// result += "у"
				continue
			}
			//ды
			if v == 'ض' && ArabicText[i+1] == 'ِ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"д!ы", "", false})
				// result += "д!ы"
				continue
			}
			//до
			if v == 'ض' && ArabicText[i+1] == 'َ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"д!о", "", false})
				// result += "д!о"
				continue
			}
			//сы
			if v == 'ص' && ArabicText[i+1] == 'ِ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"с!ы", "", false})
				// result += "с!ы"
				continue
			}
			//со
			if v == 'ص' && ArabicText[i+1] == 'َ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"с!о", "", false})
				// result += "с!о"
				continue
			}
			//ғо
			if v == 'غ' && ArabicText[i+1] == 'َ' {
				skip = 1
				DataSlice = append(DataSlice, &templateData{"ғо", "", false})
				// result += "ғо"
				continue
			}
			//филь
			if ArabicText[i+1] == ' ' && v == 'ي' && (ArabicText[i+2] == 'ا' && ArabicText[i+3] == 'ل') {
				skip = 3
				DataSlice = append(DataSlice, &templateData{" ", "", false})
				// result += " "
				continue
			}
			//фии
			if ArabicText[i+1] == ' ' && v == 'ي' && (ArabicText[i+2] != 'ا' && ArabicText[i+3] != 'ل') {
				DataSlice = append(DataSlice, &templateData{"и", "", false})
				// result += "и"
				continue
			}
			if ArabicText[i+1] == ' ' && ArabicText[i+2] == 'ا' && ArabicText[i+3] == 'ل' {
				//
				//
				//<<<<< фиксим баг "мин-ариль"
				if mapa[v] != " " && mapa[v] != "shadda" {
					DataSlice = append(DataSlice, &templateData{mapa[v], "", false})
				}
				//
				//
				//>>>>>

				//лунные буквы, когда у "аль" лям хоть с сукуном, хоть без
				if qamariya[ArabicText[i+4]] || (ArabicText[i+4] == 'ْ' && qamariya[ArabicText[i+5]]) {
					DataSlice = append(DataSlice, &templateData{"ль", "", false})
					// result += "ль"
				} else {
					if ArabicText[i+5] == 'ض' {
						DataSlice = append(DataSlice, &templateData{"д", "", false})
						// result += "д"
					}
					if ArabicText[i+5] == 'ص' {
						DataSlice = append(DataSlice, &templateData{"с!", "", false})
						// result += "с!"
					}
					if ArabicText[i+5] == 'ث' {
						DataSlice = append(DataSlice, &templateData{"с?", "", false})
						// result += "с"
					}
					if ArabicText[i+5] == 'ش' {
						DataSlice = append(DataSlice, &templateData{"ш", "", false})
						// result += "ш"
					}
					if ArabicText[i+5] == 'س' {
						DataSlice = append(DataSlice, &templateData{"с", "", false})
						// result += "с"
					}
					if ArabicText[i+5] == 'ت' {
						DataSlice = append(DataSlice, &templateData{"т", "", false})
						// result += "т"
					}
					if ArabicText[i+5] == 'ن' {
						DataSlice = append(DataSlice, &templateData{"н", "", false})
						// result += "н"
					}
					if ArabicText[i+5] == 'ظ' {
						DataSlice = append(DataSlice, &templateData{"з!", "", false})
						// result += "з"
					}
					if ArabicText[i+5] == 'ط' {
						DataSlice = append(DataSlice, &templateData{"т!", "", false})
						// result += "т"
					}
					if ArabicText[i+5] == 'ذ' {
						DataSlice = append(DataSlice, &templateData{"з?", "", false})
						// result += "з"
					}
					if ArabicText[i+5] == 'د' {
						DataSlice = append(DataSlice, &templateData{"д", "", false})
						// result += "д"
					}
					if ArabicText[i+5] == 'ز' {
						DataSlice = append(DataSlice, &templateData{"з", "", false})
						// result += "з"
					}
					if ArabicText[i+5] == 'ر' {
						DataSlice = append(DataSlice, &templateData{"р", "", false})
						// result += "р"
					}
				}
				skip = 3
				DataSlice = append(DataSlice, &templateData{"", "", true})
				// result += " "
				continue
			}
		}
		if mapa[v] == " " {
			DataSlice = append(DataSlice, &templateData{"", "", true})
		}
		if mapa[v] != " " && mapa[v] != "shadda" {
			DataSlice = append(DataSlice, &templateData{mapa[v], "", false})
		}

		// fmt.Print(mapa[v])
	}

	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "convert.page.html", DataSlice)
}
