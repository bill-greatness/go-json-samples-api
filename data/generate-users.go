package data

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	models "github.com/bill-greatness/goxide/models/users"
)

func init() {
	// seeding the random number generator for "constantly-changing fields"
	rand.Seed(time.Now().UnixNano())

}

func randRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func getRandomChoice(col []string) string {
	// return a specific string by random.
	return col[rand.Intn(len(col))]
}

func getRandomChoiceN(col []float32) float32 {
	return col[rand.Intn(len(col))]
}

func generateUsers(total int) []*models.User {
	var (
		firstNames      []string  = []string{"James", "William", "Frank", "Nick", "Nicholas", "Jayden", "Christian", "John", "Pete", "Kai"}
		lastNames       []string  = []string{"Johnson", "Powell", "Martin", "Frick", "Incoom", "Lukas", "Havertz", "Clinton", "Smith", "Charlie"}
		longitudes      []float32 = []float32{3.442, -16.232, 88.223, 13.441, 4.443, 7.659, -2.221, 8.5564, 14.231, 12.231}
		latitudes       []float32 = []float32{-43.221, 0.3432, 54.221, 19.002, 18.334, 11.221, -9.3431, 12.11, 0.334, 8.92}
		mailExtenstions []string  = []string{"@gmail.com", "@ymail.com", "@xyz.com", "@mailer.com", "@apple.com", "@test.com", "@bby.com", "@somewhere.com"}
		occupations     []string  = []string{"Teacher", "Engineer", "Accountant", "Trader", "Businessman", "Technician", "Lawyer", "Politician", "Farmer", "Contractor"}
		genderGroup     []string  = []string{"Male", "Female", "Trans", "N/A"}
		maritalStatuses []string  = []string{"Single", "Married", "Divorced", "Separated", "Complicated", "Widowed"}
		streets         []string  = []string{"Pencil St.", "BX Trio St.", "FishEye St.", "Randomness Lane St.", "Perry Hikes St.", "Litos St.", "Fintone St.", "Alfafa St.", "Jane Newman St.", "Michael Jackson St.", "Cryton Avenue"}
		photoURLs       []string  = []string{"https://www.nicepng.com/png/detail/741-7413169_placeholder-female.png", "https://vitalehealthservices.com/wp-content/uploads/2019/07/male-placeholder-image.jpeg"}
		countries       []string  = []string{"Ghana", "Luxembourg", "Algeria", "Nigeria", "Kuwait", "UAE",
			"United Kingdom", "China", "Russia", "Mexico", "Egypt", "Croatia", "Germany", "Poland", "Argentina",
			"France", "Italy", "Spain", "Belgium"}
	)

	Info := []*models.User{}

	for item := 0; item < total; item++ {
		firstName := getRandomChoice(firstNames)
		lastName := getRandomChoice(lastNames)
		email := strings.ToLower(firstName) + strings.ToLower(lastName) + strconv.Itoa(rand.Intn(999)) + getRandomChoice(mailExtenstions)
		origin := getRandomChoice(countries)
		year := randRange(1920, 2002)
		dateOfBirth := strconv.Itoa(randRange(0, 29)) + "/" + strconv.Itoa(randRange(1, 12)) + "/" + strconv.Itoa(year)
		phoneLine := "(" + strconv.Itoa(randRange(0o1, 276)) + ")" + "-" + strconv.Itoa(randRange(100, 999)) + "-" + strconv.Itoa(randRange(1000, 9999))
		age := time.Now().Year() - year
		gender := getRandomChoice(genderGroup)
		photo := getRandomChoice(photoURLs)
		street := getRandomChoice(streets)
		job := getRandomChoice(occupations)
		maritalStatus := getRandomChoice(maritalStatuses)

		content := &models.User{
			ID:              item + 1,
			Name:            firstName + " " + lastName,
			Occupation:      job,
			Age:             age,
			CountryOfOrigin: origin,
			PhotoURL:        photo,
			Email:           email,
			Gender:          gender,
			MaritalStatus:   maritalStatus,
			DateOfBirth:     dateOfBirth,
			PhoneLine:       phoneLine,
			Address: &models.AddressInfo{
				Street: street,
				Cordinates: &models.CordinateInfo{
					Latitude:  getRandomChoiceN(latitudes),
					Longitude: getRandomChoiceN(longitudes),
				},
			},
		}
		Info = append(Info, content)

	}

	return Info
}
