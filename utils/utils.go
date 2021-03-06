package utils

import (
	"encoding/json"
	"etneca-logbook/models"
	"etneca-logbook/repository"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sulat"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

type Alert struct {
	Result  bool
	Message message
}

type Danger struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type message struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

var SGC = sulat.SGC{}

func init() {
	gotenv.Load()
	SGC = sulat.SGC{
		SendGridAPIKey:   "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",
		SendGridEndPoint: "/v3/mail/send",
		SendGridHost:     "https://api.sendgrid.com",
	}
}

func SentNewMessage(response http.ResponseWriter, result bool, message string) {
	var danger Danger
	if message == "invalid syntax" {
		response.WriteHeader(http.StatusBadRequest)
	}
	danger.Message = message
	danger.Result = result

	json.NewEncoder(response).Encode(danger)
}

func SentMessage(response http.ResponseWriter, message models.Constants) {

	if message.Message.En == "invalid syntax" {
		response.WriteHeader(http.StatusBadRequest)
	} else if message.Message.En == "token expired" {
		response.WriteHeader(http.StatusUnauthorized)
	} else if message.Message.En == "token not found" {
		response.WriteHeader(http.StatusNotFound)
	} else if message.Message.En == "invalid token" {
		response.WriteHeader(http.StatusForbidden)
	}

	json.NewEncoder(response).Encode(message)
}

func Decrypt(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func Encrypt(text string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(bytes)
}

func GenerateToken(authen models.Authen, types string) (string, error) {
	var secret string
	var expires int64

	if types == "access" {
		secret = os.Getenv("ACCESS_TOKEN")
		expires = time.Now().Add(time.Minute * 2).Unix()
	} else if types == "refresh" {
		secret = os.Getenv("REFRESH_TOKEN")
		expires = time.Now().Add(time.Minute * 5).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authen.ID,
		"email":    authen.Email,
		"password": authen.Password,
		"exp":      expires,
	})

	tokenString, _ := token.SignedString([]byte(secret))

	if types == "refresh" {

		tk := time.Unix(expires, 0)
		sub := time.Now()
		stringID := authen.ID.Hex()
		err := repository.SetToken(stringID, tokenString, tk, sub)
		if err != nil {
			return "", err
		}
	}
	return tokenString, nil
}

func ValidAccessToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an Error")
		}
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})
	return token, err
}

func ValidRefreshToken(refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an Error")
		}
		return []byte(os.Getenv("REFRESH_TOKEN")), nil
	})
	return token, err
}

func ValidPath(refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an Error")
		}
		return []byte(os.Getenv("KEY_CRYPTO")), nil
	})
	return token, err
}

func ParseJson(Token string) (string, bool) {
	var ID string
	token, _ := jwt.Parse(Token, nil)
	claims, err := token.Claims.(jwt.MapClaims)
	for key, val := range claims {
		if key == "id" {
			ID = val.(string)
		}
	}
	return ID, err
}

func ParseEmail(Token string) (string, bool) {
	var email string
	token, _ := jwt.Parse(Token, nil)
	claims, err := token.Claims.(jwt.MapClaims)
	for key, val := range claims {
		if key == "email" {
			email = val.(string)
		}
	}
	return email, err
}

func GenerateEmpID(emp models.Employee) (string, error) {
	var id string

	com, _ := repository.FindCompany(emp.Company)
	year := emp.StartDateTime.Format("06")
	month := emp.StartDateTime.Format("01")
	e, _ := repository.LastEmployee()

	splitID := strings.SplitAfter(e.ID, "")
	lastID := splitID[6] + splitID[7] + splitID[8]
	num, _ := strconv.Atoi(lastID)

	if num < 9 {
		s := strconv.Itoa(num + 1)
		lastID = "00" + s
	} else if num < 99 {
		s := strconv.Itoa(num + 1)
		lastID = "0" + s
	} else if num < 999 {
		lastID = strconv.Itoa(num + 1)
	}
	id = com.Acronym + year + month + lastID

	return id, nil
}

func TimeFormat(timeString string) time.Time {
	split := strings.Split(timeString, " ")
	if split[1] == "Jan" {
		split[1] = "01"
	} else if split[1] == "Feb" {
		split[1] = "02"
	} else if split[1] == "Mar" {
		split[1] = "03"
	} else if split[1] == "Apr" {
		split[1] = "04"
	} else if split[1] == "May" {
		split[1] = "05"
	} else if split[1] == "Jun" {
		split[1] = "06"
	} else if split[1] == "Jul" {
		split[1] = "07"
	} else if split[1] == "Aug" {
		split[1] = "08"
	} else if split[1] == "Sep" {
		split[1] = "09"
	} else if split[1] == "Oct" {
		split[1] = "10"
	} else if split[1] == "Nov" {
		split[1] = "11"
	} else if split[1] == "Dec" {
		split[1] = "12"
	}

	timeString = split[3] + "-" + split[1] + "-" + split[2] + "T" + split[4] + ".000Z"
	timeTime, _ := time.Parse("2006-01-02T15:04:05.000Z", timeString)
	return timeTime
}

func generatePath(text string) string {
	key := os.Getenv("KEY_CRYPTO")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": text,
	})
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func SentMail(email string) string {
	path := generatePath(email)
	var link = "www.myurl.com/forget/"
	link = link + path
	var FullHTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html dir="ltr" xmlns="http://www.w3.org/1999/xhtml">
	<head>
	    <meta name="viewport" content="width=device-width" />
	    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	    <link rel="icon" type="image/ico" sizes="16x16" href="https://itrepablik.com/static/assets/images/favicon.ico">
	    <title>Email Notifications</title>
	</head>
	<body style="margin:0px; background: #f8f8f8; ">
	    <div width="100%" style="background: #f8f8f8; padding: 0px 0px; font-family:arial; line-height:28px; height:100%;  width: 100%; color: #514d6a;">
	        <div style="max-width: 700px; padding:50px 0;  margin: 0px auto; font-size: 14px">
	            <table border="0" cellpadding="0" cellspacing="0" style="width: 100%; margin-bottom: 20px">
	                <tbody>
	                    <tr>
	                        <td style="vertical-align: top; padding-bottom:30px;" align="center">
	                            <a href="https://itrepablik.com" target="_blank">
	                                <img src="https://itrepablik.com/static/assets/images/ITRepablik_top_logo.png" style="width:230px; height:auto;" alt="xtreme admin" style="border:none">
	                            </a>
	                        </td>
	                    </tr>
	                </tbody>
	            </table>

	            <div style="padding: 40px; background: #fff;">
	                <table border="0" cellpadding="0" cellspacing="0" style="width: 100%;">
	                    <tbody>
	                        <tr>
	                            <td style="border-bottom:1px solid #f6f6f6;">
	                                <h1 style="font-size:14px; font-family:arial; margin:0px; font-weight:bold;">Hi UserName,</h1>
	                            </td>
	                        </tr>
	                        <tr>
	                            <td style="padding:10px 0 30px 0;">
	                                <p>A request to reset your password has been made. If you did not make this request, simply ignore this email. If you did make this request, please reset your password:</p>
	                                <center>
	                                <a href="` + link + `" style="display: inline-block; padding: 11px 30px; margin: 20px 0px 30px; font-size: 15px; color: #fff; background: #4fc3f7; border-radius: 60px; text-decoration:none;">Reset Password</a>
	                                </center>
	                                <b>- Thanks (ITRepablik.com Team)</b>
	                            </td>
	                        </tr>
	                        <tr>
	                            <td style="border-top:1px solid #f6f6f6; padding-top:20px; color:#777">
	                                If the button above does not work, try copying and pasting the URL into your browser.<br/>
	                                <a href="#">` + link + `</a><br/>
	                                If you continue to have problems, please feel free to contact us at <a href="mailto:support@itrepablik.com">support@itrepablik.com</a>
	                            </td>
	                        </tr>
	                    </tbody>
	                </table>
	            </div>

	            <div style="text-align: center; font-size: 12px; color: #b2b2b5; margin-top: 20px">
	                <p> Powered by ITRepablik.com
	                    <br>
	                    <a href="javascript: void(0);" style="color: #b2b2b5; text-decoration: underline;">Unsubscribe</a>
	                </p>
	            </div>
	        </div>
	    </div>
	</body>
	</html>`
	mailOpt := &sulat.SendMail{
		Subject: "Reset Password!",
		From:    sulat.NewEmail("kritmet", "kengforsteam@gmail.com"),
		To:      sulat.NewEmail("Eiei Manz", "kritmet.w@gmail.com"),
	}
	htmlContent, err := sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML:       true,
		FullHTMLTemplate: FullHTML,
	})
	if err != nil {
		itrlog.Fatal(err)
	}
	_, err = sulat.SendEmailSG(mailOpt, htmlContent, &SGC)
	if err != nil {
		itrlog.Fatal(err)
	}
	return link
}
