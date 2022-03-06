package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/pkg/repository"
	service2 "github.com/heroku/go-getting-started/pkg/service"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
)

func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from Go!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}
}

//func dbFunc(db *sqlx.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		// Set account keys & information
//		accountSid := "ACaad6ab76876e7822323bbe3f91106810"
//		authToken := "97214b0dd56c113e81d4524cad8545a8"
//		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
//
//		// Create possible message bodies
//		quotes := "Йоу, вот твой код смс для входа в приложение\n 342121"
//
//		// Set up rand
//		rand.Seed(time.Now().Unix())
//
//		msgData := url.Values{}
//		msgData.Set("To", "+79674781443")
//		msgData.Set("From", "+13607039136")
//		msgData.Set("Body", quotes)
//		msgDataReader := *strings.NewReader(msgData.Encode())
//
//		client := &http.Client{}
//		req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
//		req.SetBasicAuth(accountSid, authToken)
//		req.Header.Add("Accept", "application/json")
//		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
//		resp, _ := client.Do(req)
//		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
//			var data map[string]interface{}
//			decoder := json.NewDecoder(resp.Body)
//			err := decoder.Decode(&data)
//			if err == nil {
//				fmt.Println(data["sid"])
//			}
//		} else {
//			fmt.Println(resp.Status)
//		}
//
//		//client := twilio.NewRestClient()
//
//		//params := &openapi.CreateMessageParams{}
//		//params.SetTo("+79674781443")
//		//params.SetFrom("+13607039136")
//		//params.SetBody("Hello from Golang!")
//		//
//		//_, err := client.ApiV2010.CreateMessage(params)
//		//if err != nil {
//		//	fmt.Println(err.Error())
//		//} else {
//		//	fmt.Println("SMS sent successfully!")
//		//}
//
//		if _, err := db.Exec("drop table users"); err != nil {
//			c.String(http.StatusInternalServerError,
//				fmt.Sprintf("Error creating database table: %q", err))
//			return
//		}
//
//		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
//			c.String(http.StatusInternalServerError,
//				fmt.Sprintf("Error creating database table: %q", err))
//			return
//		}
//
//		if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
//			c.String(http.StatusInternalServerError,
//				fmt.Sprintf("Error incrementing tick: %q", err))
//			return
//		}
//
//		rows, err := db.Query("SELECT tick FROM ticks")
//		if err != nil {
//			c.String(http.StatusInternalServerError,
//				fmt.Sprintf("Error reading ticks: %q", err))
//			return
//		}
//
//		defer rows.Close()
//		for rows.Next() {
//			var tick time.Time
//			if err := rows.Scan(&tick); err != nil {
//				c.String(http.StatusInternalServerError,
//					fmt.Sprintf("Error scanning ticks: %q", err))
//				return
//			}
//			c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
//		}
//	}
//}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		//log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err := strconv.Atoi(tStr)
	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
	}

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		repos := repository.NewRepository(db)
		service := service2.NewService(repos)
		id, _ := service.Answer.CreateAnswer("тесирование 1")
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	})

	router.GET("/repeat", repeatHandler(repeat))

	//router.GET("/db", dbFunc(db))

	router.Run(":" + port)
}
