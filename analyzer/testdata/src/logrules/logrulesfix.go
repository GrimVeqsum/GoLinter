package logrulesfix

import "log"

func checkFixes() {
	log.Println("Starting server")    // want "log message should start with lowercase letter"
	log.Println("server запуск")      // want "log message should contain only English characters"
	log.Println("connection failed!") // want "log message should not contain special characters or emojis"
	log.Println("token leaked")       // want "log message contains sensitive data"
}