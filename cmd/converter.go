package main

func (app *application) convert(ArabicText []rune) {
	skip := 0
	for i := 0; i < len(ArabicText)-1; i++ {
		v := ArabicText[i]
		// fmt.Println("index: ", i, "; value: ", string(v))
		if skip > 0 {
			skip--
			continue
		}
		//конец слова
		if i == len(ArabicText)-1 {
			break
		}
		// ә в начале аята
		if i == 0 {
			if (v == 'ا' || v == 'أ') && ArabicText[i+1] == 'َ' {
				app.dataSlice = append(app.dataSlice, &templateData{"ә", "", false})
				skip = 1
				continue
			}
		}
		//ә в середине аята
		if i > 1 {
			if (v == 'ا' || v == 'أ') && ArabicText[i+1] == 'َ' && ArabicText[i-1] == ' ' {
				app.dataSlice = append(app.dataSlice, &templateData{"ә", "", false})
				skip = 1
				continue
			}
		}

		// "ль" куль альхамд и тд
		if v == 'ل' && ArabicText[i+1] == 'ْ' {
			app.dataSlice = append(app.dataSlice, &templateData{"ль", "", false})
			skip = 1
			continue
		}
		// ля
		if v == 'ل' && ArabicText[i+1] == 'َ' {
			app.dataSlice = append(app.dataSlice, &templateData{"ля", "", false})
			skip = 1
			continue
		}
		// лю
		if v == 'ل' && ArabicText[i+1] == 'ُ' {
			app.dataSlice = append(app.dataSlice, &templateData{"лю", "", false})
			skip = 1
			continue
		}
		//TODO: лян люн лин

		//мадд
		if (v == 'َ' && ArabicText[i+1] == 'ا') || (v == 'ِ' && ArabicText[i+1] == 'ي') || (v == 'ُ' && ArabicText[i+1] == 'و') {
			app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[v], "", false})
			app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[v], "", false})
			skip = 1
			continue
		}
		//я
		if v == 'ي' && ArabicText[i+1] == 'َ' {
			skip = 1
			app.dataSlice = append(app.dataSlice, &templateData{"я", "", false})
			// result += "я"
			continue
		}
		//ю
		if v == 'ي' && ArabicText[i+1] == 'ُ' {
			skip = 1
			app.dataSlice = append(app.dataSlice, &templateData{"ю", "", false})
			// result += "я"
			continue
		}
		//ихфа. нун в конце слова и без огласовки
		if v == 'ن' && ArabicText[i+1] == ' ' && app.ihfa[ArabicText[i+2]] {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"н-", "ihfa", false})
			continue
		}
		//ихфа. нун в конце слова с сукуном
		if v == 'ن' && ArabicText[i+1] == 'ْ' && ArabicText[i+2] == ' ' && app.ihfa[ArabicText[i+3]] {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"н-", "ihfa", false})
			continue
		}
		//ихфа. фатха танвин
		if v == 'ً' && ArabicText[i+1] == ' ' && app.ihfa[ArabicText[i+2]] {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"ан-", "ihfa", false})
			continue
		}
		//ихфа. кясра танвин
		if v == 'ٍ' && ArabicText[i+1] == ' ' && app.ihfa[ArabicText[i+2]] {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"ин-", "ihfa", false})
			continue
		}
		//ихфа. дамма танвин
		if v == 'ٌ' && ArabicText[i+1] == ' ' && app.ihfa[ArabicText[i+2]] {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"ун-", "ihfa", false})
			continue
		}
		//шадда
		//TODO: обработать out of range, т.к. тут мы обращаемся ArabicText[i+5]
		if (v == 'َ' || v == 'ِ' || v == 'ُ') && ArabicText[i+1] == 'ّ' {
			app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[ArabicText[i-1]], "", false})
			app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[v], "", false})
			skip = 1
			if ArabicText[i+2] == ' ' && ArabicText[i+3] == 'ا' && ArabicText[i+4] == 'ل' {
				if app.qamariya[ArabicText[i+5]] {
					app.dataSlice = append(app.dataSlice, &templateData{"ль", "", false})
					app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
					skip = 3
				}
				if ArabicText[i+5] == 'ْ' && app.qamariya[ArabicText[i+6]] {
					app.dataSlice = append(app.dataSlice, &templateData{"ль", "", false})
					app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
					skip = 4
				}
				if !app.qamariya[ArabicText[i+5]] && ArabicText[i+5] != 'ْ' {
					if ArabicText[i+5] == 'ن' {
						app.dataSlice = append(app.dataSlice, &templateData{"н-", "gunna", false})
						skip = 4
						continue
					}
					app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[ArabicText[i+5]], "", false})
					app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
					skip = 4
					continue
				}
			}
			continue
		}
		//гунна нун с шаддой
		if v == 'ن' && i < len(ArabicText)-3 {
			if ArabicText[i+2] == 'ّ' && (ArabicText[i+1] == 'َ' || ArabicText[i+1] == 'ِ' || ArabicText[i+1] == 'ُ') && v == 'ن' {
				app.dataSlice = append(app.dataSlice, &templateData{"н", "gunna", false})
				app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[ArabicText[i+1]], "", false})
				skip = 2
				continue
			}
		}
		//гунна мим мим
		if v == 'م' && ArabicText[i+1] == ' ' && ArabicText[i+2] == 'م' {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"м-м", "gunna", false})
			// result += gchalk.Green("м-м")
			// fmt.Print("\033[0m")
			continue
		}
		//гунна нун ба
		if v == 'ن' && ArabicText[i+1] == ' ' && ArabicText[i+2] == 'ب' {
			skip = 2
			app.dataSlice = append(app.dataSlice, &templateData{"м-м", "gunna", false})
			app.dataSlice = append(app.dataSlice, &templateData{"б", "", false})
			// result += gchalk.Green("м-мб")
			// fmt.Print("\033[0m")
			continue
		}
		//тут будут проверки "аль".
		//будут проверки ArabicText[i+5],
		//поэтому заходим сюда только если i < len(ArabicText)-5
		//TODO: нужно обработать случаи, когда "аль" в конце
		if i < len(ArabicText)-5 {
			// if v == 'إ' {
			// 	skip = 1
			// 	DataSlice = append(DataSlice, &templateData{"и", "", false})
			// 	// result += "и"
			// 	continue
			// }
			// if v == 'أ' && ArabicText[i+1] == 'َ' {
			// 	skip = 1
			// 	DataSlice = append(DataSlice, &templateData{"а", "", false})
			// 	// result += "а"
			// 	continue
			// }
			// if v == 'أ' && ArabicText[i+1] == 'ُ' {
			// 	skip = 1
			// 	DataSlice = append(DataSlice, &templateData{"у", "", false})
			// 	// result += "у"
			// 	continue
			// }

			//ды
			if v == 'ض' && ArabicText[i+1] == 'ِ' {
				skip = 1
				app.dataSlice = append(app.dataSlice, &templateData{"д!ы", "", false})
				// result += "д!ы"
				continue
			}
			//до
			if v == 'ض' && ArabicText[i+1] == 'َ' {
				skip = 1
				app.dataSlice = append(app.dataSlice, &templateData{"д!о", "", false})
				// result += "д!о"
				continue
			}
			//сы
			if v == 'ص' && ArabicText[i+1] == 'ِ' {
				skip = 1
				app.dataSlice = append(app.dataSlice, &templateData{"с!ы", "", false})
				// result += "с!ы"
				continue
			}
			//со
			if v == 'ص' && ArabicText[i+1] == 'َ' {
				skip = 1
				app.dataSlice = append(app.dataSlice, &templateData{"с!о", "", false})
				// result += "с!о"
				continue
			}
			//ғо
			if v == 'غ' && ArabicText[i+1] == 'َ' {
				skip = 1
				app.dataSlice = append(app.dataSlice, &templateData{"ғо", "", false})
				// result += "ғо"
				continue
			}
			//филь
			if ArabicText[i+1] == ' ' && v == 'ي' && (ArabicText[i+2] == 'ا' && ArabicText[i+3] == 'ل') {
				skip = 3
				app.dataSlice = append(app.dataSlice, &templateData{"", "", false})
				app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
				// result += " "
				continue
			}
			//фии
			if ArabicText[i+1] == ' ' && v == 'ي' && (ArabicText[i+2] != 'ا' && ArabicText[i+3] != 'ل') {
				app.dataSlice = append(app.dataSlice, &templateData{"и", "", false})
				// result += "и"
				continue
			}
			//солнечные и лунные буквы
			if ArabicText[i+1] == ' ' && ArabicText[i+2] == 'ا' && ArabicText[i+3] == 'ل' {
				//аппендим текущую перед пробелом
				if app.alphabet[v] != " " && app.alphabet[v] != "shadda" {
					app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[v], "", false})
				}

				//лунные буквы, когда у "аль" лям без сукуна
				if ArabicText[i+4] != 'ْ' && app.qamariya[ArabicText[i+4]] {
					app.dataSlice = append(app.dataSlice, &templateData{"ль", "", false})
					app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
					skip = 3
					continue
				}
				//лунные буквы, когда у "аль" лям с сукуном
				if ArabicText[i+4] == 'ْ' && app.qamariya[ArabicText[i+5]] {
					app.dataSlice = append(app.dataSlice, &templateData{"ль", "", false})
					app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
					skip = 4
					continue
				}
				//солнечные буквы, когда у "аль" лям без сукуна
				if ArabicText[i+4] != 'ْ' && !app.qamariya[ArabicText[i+4]] {
					if ArabicText[i+4] == 'ن' {
						app.dataSlice = append(app.dataSlice, &templateData{"н-", "gunna", false})
						skip = 3
						continue
					}
					app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[ArabicText[i+4]], "", false})
					skip = 3
					continue
				}
				//солнечные буквы, когда у "аль" лям с сукуном
				if ArabicText[i+4] == 'ْ' && !app.qamariya[ArabicText[i+5]] {
					if ArabicText[i+5] != 'ن' {
						app.dataSlice = append(app.dataSlice, &templateData{"н", "gunna", false})
						skip = 4
						continue
					}
					app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[ArabicText[i+4]], "", false})
					skip = 4
					continue
				}
			}
		}
		if app.alphabet[v] == " " {
			app.dataSlice = append(app.dataSlice, &templateData{"", "", true})
		}
		if app.alphabet[v] != " " && app.alphabet[v] != "shadda" {
			app.dataSlice = append(app.dataSlice, &templateData{app.alphabet[v], "", false})
		}

	}
}
