package login

import (
	rands "crypto/rand"
	"io"
	"math/rand"
	"net"
	"strings"
	"time"
)

var Numbers = []rune("123456789")
var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func isDomainValid(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	_, err := net.LookupMX(domain) // Looks up Mail Exchange (MX) records
	return err == nil
}

func RandAlphanumeric(n int) string {
	a := make([]byte, n)
	if _, err := io.ReadFull(rands.Reader, a); err != nil {

		return ""
	}
	for i := range a {
		a[i] = letters[int(a[i])%len(letters)]
	}
	return string(a)
}

func GenerateRandomOtp(n int) string {
	if n == 0 {
		n = 6
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = Numbers[rand.Intn(len(Numbers))]
	}
	return string(b)
}

// returns today's date for calculations wrt UTC columns in DB
func todaysDate() (today string) {
	// IST date - current DATE
	//loc, hdh := time.LoadLocation("Asia/Kolkata")
	//fmt.Println(hdh)
	//fmt.Println(loc)
	location := time.FixedZone("IST", 5*60*60+30*60)
	istDate := time.Now().In(location).Format("2006-01-02")
	utcDate := time.Now().Format("2006-01-02")
	today = ""
	if istDate != utcDate {
		today = utcDate
	} else {
		today = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}
	return today + " 18:30:01"
}
