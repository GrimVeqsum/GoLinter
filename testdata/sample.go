package testdata

import (
	"log"
)

func Sample() {
    log.Println("Starting server")        // ❌ Uppercase
    log.Println("server started")         // ✅ Correct
    log.Println("пароль пользователя")    // ❌ Not English
    log.Println("user password: 1234")    // ❌ Sensitive
    log.Println("connection failed!")     // ❌ Special char
}