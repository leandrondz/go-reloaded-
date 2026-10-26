package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Transform(texte string) string {
	mots := strings.Fields(texte)

	for i := 1; i < len(mots); i++ {
		mot := mots[i]

		if mot == "(up)" {
			mots[i-1] = strings.ToUpper(mots[i-1])
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		}

		if mot == "(low)" {
			mots[i-1] = strings.ToLower(mots[i-1])
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		}

		if mot == "(cap)" {
			runes := []rune(strings.ToLower(mots[i-1]))
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				mots[i-1] = string(runes)
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		}

		if mot == "(hex)" {
			nombre, err := strconv.ParseInt(mots[i-1], 16, 64)
			if err == nil {
				mots[i-1] = strconv.FormatInt(nombre, 10)
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		}

		if mot == "(bin)" {
			nombre, err := strconv.ParseInt(mots[i-1], 2, 64)
			if err == nil {
				mots[i-1] = strconv.FormatInt(nombre, 10)
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		}

		if (mot == "(up," || mot == "(low," || mot == "(cap,") &&
			i+1 < len(mots) {

			action := strings.TrimPrefix(mot, "(")
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			nombre, err := strconv.Atoi(nombreTexte)

			if err == nil && nombre > 0 && nombre <= i {
				for j := i - nombre; j < i; j++ {
					if action == "up," {
						mots[j] = strings.ToUpper(mots[j])
					}

					if action == "low," {
						mots[j] = strings.ToLower(mots[j])
					}

					if action == "cap," {
						runes := []rune(strings.ToLower(mots[j]))
						if len(runes) > 0 {
							runes[0] = unicode.ToUpper(runes[0])
							mots[j] = string(runes)
						}
					}
				}

				mots = append(mots[:i], mots[i+2:]...)
				i--
			}
		}
	}

	texte = strings.Join(mots, " ")

	texte = strings.ReplaceAll(texte, "' ", "'")
	texte = strings.ReplaceAll(texte, " '", "'")

	var resultat strings.Builder

	for i := 0; i < len(texte); i++ {
		c := texte[i]

		if c == '.' || c == ',' || c == '!' ||
			c == '?' || c == ':' || c == ';' {

			for resultat.Len() > 0 {
				s := resultat.String()
				if s[len(s)-1] != ' ' {
					break
				}
				resultat.Reset()
				resultat.WriteString(s[:len(s)-1])
			}

			resultat.WriteByte(c)

			if i+1 < len(texte) && texte[i+1] != ' ' &&
				texte[i+1] != '.' && texte[i+1] != ',' &&
				texte[i+1] != '!' && texte[i+1] != '?' &&
				texte[i+1] != ':' && texte[i+1] != ';' {
				resultat.WriteByte(' ')
			}
		} else {
			resultat.WriteByte(c)
		}
	}

	texte = strings.Join(strings.Fields(resultat.String()), " ")
	mots = strings.Fields(texte)

	for i := 0; i < len(mots)-1; i++ {
		if mots[i] == "a" || mots[i] == "A" {
			suivant := strings.TrimLeft(mots[i+1], "\"'([{")

			if len(suivant) > 0 {
				premiere := unicode.ToLower([]rune(suivant)[0])

				if strings.ContainsRune("aeiou", premiere) {
					if mots[i] == "A" {
						mots[i] = "An"
					} else {
						mots[i] = "an"
					}
				}
			}
		}
	}

	return strings.Join(mots, " ")
}
