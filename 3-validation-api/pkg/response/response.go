package response

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Json(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ToFile(email string, verificationUUID string) {
	file, err := os.OpenFile("verifications.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	data := map[string]string{email: verificationUUID}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func FromFile(email string, verificationUUID string) bool {
	file, err := os.OpenFile("verifications.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var data map[string]string
		if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}
		if val, ok := data[email]; ok && val == verificationUUID {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	return false
}
