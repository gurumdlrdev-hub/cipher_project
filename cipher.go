package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Converts digits ('0'-'9') into special characters using cipher_code table
func convertDigitsToSpecialChars(digits string) string {
	var result strings.Builder

	for _, char := range digits {
		digitStr := string(char)
		var symbol string

		query := `SELECT code FROM cipher_code WHERE decodenumber = $1`
		err := DBs.CipherDB.QueryRow(query, digitStr).Scan(&symbol)

		if err != nil {
			switch digitStr {
			case "1":
				symbol = "@"
			case "2":
				symbol = "#"
			case "3":
				symbol = "$"
			case "4":
				symbol = "%"
			case "5":
				symbol = "*"
			case "6":
				symbol = "^"
			case "7":
				symbol = "&"
			case "8":
				symbol = "\\"
			case "9":
				symbol = "?"
			case "0":
				symbol = "!"
			default:
				symbol = digitStr
			}
		}

		result.WriteString(symbol)
	}

	return result.String()
}

// Converts special characters back into digits ('0'-'9')
func convertSpecialCharsToDigits(symbols string) string {
	var digits strings.Builder

	for _, char := range symbols {
		symbolStr := string(char)
		var digit string

		query := `SELECT decodenumber FROM cipher_code WHERE code = $1`
		err := DBs.CipherDB.QueryRow(query, symbolStr).Scan(&digit)

		if err != nil {
			switch symbolStr {
			case "@":
				digit = "1"
			case "#":
				digit = "2"
			case "$":
				digit = "3"
			case "%":
				digit = "4"
			case "*":
				digit = "5"
			case "^":
				digit = "6"
			case "&":
				digit = "7"
			case "\\":
				digit = "8"
			case "?":
				digit = "9"
			case "!":
				digit = "0"
			default:
				digit = symbolStr
			}
		}

		digits.WriteString(digit)
	}

	return digits.String()
}

// API Handler to process encoding requests
func apiEncode(c *gin.Context) {
	inputText := c.PostForm("encode_input")

	if strings.TrimSpace(inputText) == "" {
		c.HTML(http.StatusBadRequest, "encode.html", gin.H{"error": "Input text cannot be empty"})
		return
	}

	processedText := strings.ReplaceAll(inputText, " ", "~")

	var encodeResult []string
	runes := []rune(processedText)
	length := len(runes)

	for i := 0; i < length; {
		var cipherNum string
		found := false

		// 1. Try matching 2-character pairs in cipher_project first
		if i+1 < length {
			chunk2 := string(runes[i : i+2])
			query := `SELECT decodenumber FROM cipher_project WHERE code = $1`
			err := DBs.CipherDB.QueryRow(query, chunk2).Scan(&cipherNum)

			if err == nil {
				symbols := convertDigitsToSpecialChars(cipherNum)
				encodeResult = append(encodeResult, symbols)
				i += 2
				found = true
			}
		}

		// 2. Fallback to 1-character lookup
		if !found {
			chunk1 := string(runes[i])

			if chunk1 == "~" {
				symbols := convertDigitsToSpecialChars("0000")
				encodeResult = append(encodeResult, symbols)
				i += 1
				continue
			}

			query := `SELECT decodenumber FROM cipher_project WHERE code = $1`
			err := DBs.CipherDB.QueryRow(query, chunk1).Scan(&cipherNum)

			if err != nil {
				if err == sql.ErrNoRows {
					var symbol string
					altQuery := `SELECT code FROM cipher_code WHERE decodenumber = $1`
					altErr := DBs.CipherDB.QueryRow(altQuery, chunk1).Scan(&symbol)

					if altErr == nil {
						encodeResult = append(encodeResult, symbol)
					} else {
						encodeResult = append(encodeResult, "[?]")
					}
				} else {
					log.Println("Database query error:", err)
					c.HTML(http.StatusInternalServerError, "encode.html", gin.H{"error": "Database query error"})
					return
				}
			} else {
				symbols := convertDigitsToSpecialChars(cipherNum)
				encodeResult = append(encodeResult, symbols)
			}

			i += 1
		}
	}

	finalCipherOutput := strings.Join(encodeResult, " ")

	c.HTML(http.StatusOK, "encode.html", gin.H{
		"original":     inputText,
		"encodeOutput": finalCipherOutput,
	})
}

// API Handler to process decoding requests
func apiDecode(c *gin.Context) {
	cipherText := c.PostForm("decode_text")

	if strings.TrimSpace(cipherText) == "" {
		c.HTML(http.StatusBadRequest, "decode.html", gin.H{"error": "Cipher text cannot be empty"})
		return
	}

	// Split space-separated symbol groups
	symbolGroups := strings.Fields(cipherText)
	var decodedChars []string

	for _, group := range symbolGroups {
		// Step 1: Convert special character group back to digits (e.g., "!!!!" -> "0000")
		digitCode := convertSpecialCharsToDigits(group)

		var codeLetter string

		// ⚡ FIX: Explicitly handle space codes ("0000" or "9999")
		if digitCode == "0000" || digitCode == "9999" {
			decodedChars = append(decodedChars, " ")
			continue
		}

		// Step 2: Query cipher_project for the original letter/pair using 4-digit code
		query := `SELECT code FROM cipher_project WHERE decodenumber = $1`
		err := DBs.CipherDB.QueryRow(query, digitCode).Scan(&codeLetter)

		if err != nil {
			if err == sql.ErrNoRows {
				// Fallback lookup in cipher_code table
				altQuery := `SELECT code FROM cipher_code WHERE decodenumber = $1`
				altErr := DBs.CipherDB.QueryRow(altQuery, digitCode).Scan(&codeLetter)

				if altErr != nil {
					codeLetter = "[?]"
				}
			} else {
				log.Println("Database query error:", err)
				c.HTML(http.StatusInternalServerError, "encode.html", gin.H{"error": "Database query error"})
				return
			}
		}

		decodedChars = append(decodedChars, codeLetter)
	}

	// Step 3: Reconstruct text and replace any remaining ~ with spaces
	joinedText := strings.Join(decodedChars, "")
	originalText := strings.ReplaceAll(joinedText, "~", " ")

	c.HTML(http.StatusOK, "decode.html", gin.H{
		"decodedOutput": originalText,
	})
}
