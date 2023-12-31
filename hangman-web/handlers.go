package main

import (
	"Hangman/modules"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"time"
	"unicode"
)

// Table User
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Sessions []Session
	Hangman  []HangmanData
}

// Table Session
type Session struct {
	gorm.Model
	SessionID  uuid.UUID `gorm:"primary_key"`
	UserID     int
	Data       string `gorm:"type:text"`
	CreatedAt  time.Time
	LastAccess time.Time
	ExpireAt   time.Time

	User User `gorm:"association_foreignkey:UserID"`
}

// Table HangmanData
type HangmanData struct {
	gorm.Model
	UserID         int    `gorm:"column:user_id"`
	Live           int    `gorm:"type:int"`
	GameDifficulty string `gorm:"type:varchar(255)"`
	Word           string `gorm:"type:varchar(255)"`
	Result         string `gorm:"type:varchar(255)"`
	UsedLW         string `gorm:"type:varchar(255)"`
	Index          string `gorm:"type:varchar(255)"`
	UserLetter     string `gorm:"type:varchar(255)"`
	Picture        string `gorm:"type:varchar(255)"`
	AlreadyUsed    string `gorm:"type:varchar(255		)"`
	BadInput       string `gorm:"type:varchar(255)"`
	OnlyLowerCase  string `gorm:"type:varchar(255)"`
	Reveal         string `gorm:"type:varchar(255)"`
	GCompleted     string `gorm:"type:varchar(255)"`
	StartTime      time.Time
	Elapsed        float64 `gorm:"type:int"`
	User           User    `gorm:"association_foreignkey:UserID"`
}

// Table Scoreboard
type ScoreBoard struct {
	gorm.Model
	UserID   int    `gorm:"column:user_id"`
	Username string `gorm:"type:varchar(255)"`
	Score    int    `gorm:"type:int"`
	User     User   `gorm:"association_foreignkey:UserID"`
}

type Data struct {
	User      User
	GameState HangmanData
}

// Function of the start page
func Start(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "start.html", nil)
}

// Function of the start page
func Level(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	sessionID, err := uuid.FromString(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	// Get the session from the database
	var session Session
	if err := db.Where("session_id = ?", sessionID).First(&session).Error; err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	var user User
	var GameState HangmanData
	if err := db.Where("user_id = ? and g_completed = ? and live > ?", session.UserID, "false", 0).First(&GameState).Error; err == nil {
		http.Redirect(w, r, "/game", http.StatusFound)
		return
	}

	if err := db.Where("id = ?", session.UserID).First(&user).Error; err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	//Set value from button on html page
	easy := r.FormValue("Easy")
	medium := r.FormValue("Medium")
	hard := r.FormValue("Hard")
	if easy != "" {
		//Launch the game in easy mode
		Difficulty(easy, &GameState)
		GameState.Picture = "/Assets/zombie.png"
		GameState.UserID = session.UserID
		GameState.Elapsed = 0
		if err := db.Create(&GameState).Error; err != nil {
			http.Error(w, "Impossible to create DataBase", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/game", http.StatusFound)
	} else if medium != "" {
		//Launch the game in medium mode
		Difficulty(medium, &GameState)
		GameState.Picture = "/Assets/zombie.png"
		GameState.UserID = session.UserID
		GameState.Elapsed = 0
		if err := db.Create(&GameState).Error; err != nil {
			http.Error(w, "Impossible to create DataBase", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/game", http.StatusFound)
	} else if hard != "" {
		//Launch the game in hard mode
		Difficulty(hard, &GameState)
		GameState.Picture = "/Assets/zombie.png"
		GameState.UserID = session.UserID
		GameState.Elapsed = 0
		if err := db.Create(&GameState).Error; err != nil {
			http.Error(w, "Impossible to create DataBase", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/game", http.StatusFound)
	}
	data := Data{User: user}
	err = tpl.ExecuteTemplate(w, "level.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Function of the Game page
func Game(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	sessionID, err := uuid.FromString(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	// Get the session from the database
	var session Session
	if err := db.Where("session_id = ?", sessionID).First(&session).Error; err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	var user User

	var GameState HangmanData
	if err := db.Where("user_id = ? and g_completed = ? and live > ?", session.UserID, "false", 0).First(&GameState).Error; err != nil {
		http.Error(w, "No Data", http.StatusUnauthorized)
		return
	}
	if err := db.Where("id = ?", session.UserID).First(&user).Error; err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	t, err := template.ParseFiles("./Templates/game.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var tryWord bool
	var resultCorrectLetter bool

	//Evert db.Save(&GameState) are here to save data in database

	GameState.GCompleted = "false"
	GameState.UserLetter = r.FormValue("Enter")
	db.Save(&GameState)

	//test if user put letter
	if GameState.UserLetter != "" {
		if GameState.Live > 0 {
			GameState.OnlyLowerCase = ""
			GameState.BadInput = ""
			db.Save(&GameState)
			for _, letter := range GameState.UserLetter {
				if string(letter) < "a" || string(letter) > "z" {
					GameState.OnlyLowerCase = "Error : you can only chose lowercase characters for Hangman"
				}
			}
			db.Save(&GameState)
			//Test if it's onlylowercase
			if GameState.OnlyLowerCase == "" {
				if len(GameState.UserLetter) == 1 || len(GameState.UserLetter) == len(GameState.Word) {
					var try bool
					try, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result = modules.Repetition(false, GameState.UsedLW, GameState.UserLetter, GameState.AlreadyUsed, GameState.Result)
					db.Save(&GameState)
					if try == false {
						GameState.Index, tryWord = modules.TryLetter(GameState.UserLetter, GameState.Word)
						GameState.Result, resultCorrectLetter = modules.UpdateResult(GameState.UserLetter, GameState.Index, tryWord, GameState.Result)
						db.Save(&GameState)
						if resultCorrectLetter == false && tryWord == false {
							GameState.Live -= 1
							db.Save(&GameState)
							GameState.Picture = PrintJose(GameState.Live, GameState.Picture)
							//Set timer
							elapsed := time.Now()
							diff := elapsed.Sub(GameState.StartTime)
							GameState.Elapsed += diff.Seconds()
							db.Save(&GameState)
						} else if resultCorrectLetter == false && tryWord == true {
							GameState.Live -= 2
							db.Save(&GameState)
							//Set timer
							elapsed := time.Now()
							diff := elapsed.Sub(GameState.StartTime)
							GameState.Elapsed += diff.Seconds()
							db.Save(&GameState)
						}

					}
					//Test if the word is found
					if modules.TestFinish(GameState.Result) == true { // Winner
						GameState.GCompleted = "true"
						GameState.Picture = PrintJose(GameState.Live, GameState.Picture)
						//Set timer
						elapsed := time.Now()
						diff := elapsed.Sub(GameState.StartTime)
						GameState.Elapsed += diff.Seconds()
						db.Save(&GameState)
						var score ScoreBoard
						if err := db.Where("user_id = ?", session.UserID).First(&score).Error; err != nil {
							http.Error(w, "Invalid session", http.StatusUnauthorized)
							return
						}
						score.Score += ScoreCalcul(GameState.GameDifficulty, GameState.Live, GameState.Elapsed)
						db.Save(&score)
						http.Redirect(w, r, "/scoreboard?message=Popup", http.StatusFound)
					}
					//Test if he loses
					if GameState.Live == 0 || GameState.Live < 0 { // Loser
						GameState.GCompleted = "true"
						GameState.Picture = PrintJose(GameState.Live, GameState.Picture)
						//Set timer
						elapsed := time.Now()
						diff := elapsed.Sub(GameState.StartTime)
						GameState.Elapsed += diff.Seconds()
						db.Save(&GameState)
						var score ScoreBoard
						if err := db.Where("user_id = ?", session.UserID).First(&score).Error; err != nil {
							http.Error(w, "Invalid session", http.StatusUnauthorized)
							return
						}
						score.Score -= 15
						db.Save(&score)
						http.Redirect(w, r, "/scoreboard?message=Popup", http.StatusFound)
					}
				} else {
					GameState.BadInput = "Error : you can only put 1 letter or a word with the same number of letter than the word to search!!!!!!"
					db.Save(&GameState)
				}
			}
		}
		elapsed := time.Now()
		diff := elapsed.Sub(GameState.StartTime)
		GameState.Elapsed += diff.Seconds()
	}
	GameState.StartTime = time.Now()
	db.Save(&GameState)
	t.Execute(w, GameState)
}

// Function of the Scoreboard page
func Scoreboard(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	sessionID, err := uuid.FromString(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	// Get the session from the database
	var Sessions Session
	if err := db.Where("session_id = ?", sessionID).First(&Sessions).Error; err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("./Templates/scoreboard.html")
	var GameState HangmanData
	var Users User
	if err := db.Where("user_id = ?", Sessions.UserID).Last(&GameState).Error; err != nil {
		http.Error(w, "No Data", http.StatusUnauthorized)
		return
	}
	if err := db.Where("id = ?", Sessions.UserID).Last(&Users).Error; err != nil {
		http.Error(w, "No Data", http.StatusUnauthorized)
		return
	}
	var Scores []ScoreBoard
	db.Find(&Scores)
	sort.Slice(Scores, func(i, j int) bool {
		return Scores[i].Score > Scores[j].Score
	})
	var DataGames int
	var Symbol string
	if GameState.Live <= 0 {
		DataGames = -15
	} else {
		DataGames = ScoreCalcul(GameState.GameDifficulty, GameState.Live, GameState.Elapsed)
		Symbol = "+"
	}

	var TotalPoint ScoreBoard
	if err := db.Where("user_id = ?", Sessions.UserID).Last(&TotalPoint).Error; err != nil {
		http.Error(w, "No Data", http.StatusUnauthorized)
		return
	}
	var Picture string
	if TotalPoint.Score <= 50 {
		Picture = "/Assets/luciole.png"
	} else if TotalPoint.Score > 50 && TotalPoint.Score <= 100 {
		Picture = "/Assets/radiant.png"
	} else if TotalPoint.Score > 100 && TotalPoint.Score <= 300 {
		Picture = "/Assets/claqueur.png"
	} else {
		Picture = "/Assets/PUANTS.png"
	}

	t.Execute(w, map[string]interface{}{"Scores": Scores, "GameState": GameState, "DataGames": DataGames, "Users": Users, "TotalPoint": TotalPoint, "Symbol": Symbol, "Picture": Picture})
}

// Function of the Inscription page
func Inscription(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//Destroy cookies --> Deconnexion
	expiration := time.Now().AddDate(-1, 0, 0)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	username := r.FormValue("username")
	// check password criteria
	password := r.FormValue("password")
	if username != "" && password != "" {
		// variables that must pass for password creation criteria
		var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
		pswdNoSpaces = true
		for _, char := range password {
			switch {
			// func IsLower(r rune) bool
			case unicode.IsLower(char):
				pswdLowercase = true
			// func IsUpper(r rune) bool
			case unicode.IsUpper(char):
				pswdUppercase = true
			// func IsNumber(r rune) bool
			case unicode.IsNumber(char):
				pswdNumber = true
			// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				pswdSpecial = true
			// func IsSpace(r rune) bool, type rune = int32
			case unicode.IsSpace(int32(char)):
				pswdNoSpaces = false
			}
		}
		if 6 < len(password) && len(password) < 60 {
			pswdLength = true
		}
		// check if username already exists in database
		var users User
		err := db.Where("username = ?", username).First(&users).Error
		if err == nil {
			tpl.ExecuteTemplate(w, "inscription.html", struct{ ErrorMessage string }{ErrorMessage: "username already exists"})
			return
		}

		// check if password meets conditions
		var errorMessage string
		if !pswdLowercase {
			errorMessage += "Error: password must contain at least one lowercase letter. "
		}
		if !pswdUppercase {
			errorMessage += "Error: password must contain at least one uppercase letter. "
		}
		if !pswdNumber {
			errorMessage += "Error: password must contain at least one number. "
		}
		if !pswdSpecial {
			errorMessage += "Error: password must contain at least one special character. "
		}
		if !pswdLength {
			errorMessage += "Error: password must be between 7 and 59 characters. "
		}
		if !pswdNoSpaces {
			errorMessage += "Error: password cannot contain spaces. "
		}
		if errorMessage != "" {
			tpl.ExecuteTemplate(w, "inscription.html", struct{ ErrorMessage string }{ErrorMessage: errorMessage})
			return
		}
		// create bcrypt hash from password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// insert username and hashed password in database
		users = User{Username: username, Password: string(hashedPassword)}
		db.Create(&users)
		var score ScoreBoard
		score = ScoreBoard{UserID: int(users.ID), Username: users.Username, Score: 0}
		db.Create(&score)
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "inscription.html", nil)
}

// Function of the Connexion page
func Connexion(w http.ResponseWriter, r *http.Request) {
	// Get the form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	//Destroy cookies --> Deconnexion
	expiration := time.Now().AddDate(-1, 0, 0)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	// Check if username and password are not empty
	if username != "" || password != "" {

		// Get the user from the database
		var user User
		var errorMessage string
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			if err != nil {
				errorMessage = "Error: Bas Username.. "
			}
		}

		// Compare the bcrypt password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			if err != nil {
				errorMessage = "Error: Bas Password. "
			}
		}
		if errorMessage != "" {
			tpl.ExecuteTemplate(w, "connexion.html", struct{ ErrorMessage string }{ErrorMessage: errorMessage})
			return
		}
		// Create a new session
		session := &Session{
			SessionID:  uuid.NewV4(),
			UserID:     int(user.ID),
			Data:       strconv.Itoa(int(user.ID)),
			CreatedAt:  time.Now(),
			LastAccess: time.Now(),
			ExpireAt:   time.Now().Add(24 * time.Hour),
		}
		if err := db.Save(&session).Error; err != nil {
			http.Error(w, "Error creating session", http.StatusInternalServerError)
			return
		}

		// Create a new cookie and set the session ID as the value
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   session.SessionID.String(),
			Expires: session.ExpireAt,
		}
		http.SetCookie(w, cookie)

		// Redirect the user to the homepage
		http.Redirect(w, r, "/level", http.StatusFound)
	}
	tpl.ExecuteTemplate(w, "connexion.html", nil)
}

// Function of the EasterEggs page
func EasterEggs(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	sessionID, err := uuid.FromString(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}

	// Get the session from the database
	var session Session
	if err := db.Where("session_id = ?", sessionID).First(&session).Error; err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("./Templates/EasterEgss.html")
	var score ScoreBoard
	if err := db.Where("user_id = ?", session.UserID).First(&score).Error; err != nil {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	score.Score += 1500
	db.Save(&score)
	http.Redirect(w, r, "/level", http.StatusSeeOther)
	t.Execute(w, nil)

}
