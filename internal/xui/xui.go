package xui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "net/http/cookiejar"
	// "net/url"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type postRequest struct {
	Client 		ClientData	`json:"client"`
	InboundIds 	[]int32		`json:"inboundIds"`
}

type ClientData struct {
	ID   		string	`json:"id"`
	SubID		string	`json:"subId"`
	Email 		string 	`json:"email"`
	Enable 		bool	`json:"enable"`
	ExpiryTime	int64	`json:"expiryTime"`
	TotalGb		int64	`json:"totalGB"`
	LimitIP		int64	`json:"limitIp"`
}

type ClientResult struct {
	ID   	string	
	SubID	string	
	Link	string
}


func Conn(tgUserID int64, totalGB int, limitDevice int) (*ClientResult ,error) {
	uuidUser := uuid.New().String()
	subId := strings.ReplaceAll(uuidUser,"-","")[:16]

	fmt.Print(uuidUser)
	const  clientAdd string = "/panel/api/clients/add"
	err := godotenv.Load()
	xuiHost := os.Getenv("XUI_HOST")
	xuiToken := os.Getenv("XUI_TOKEN")
	tgUserIDStr := strconv.FormatInt(tgUserID,10)
	// xuiUser := os.Getenv("XUI_USER")
	

	
	
	req, err := json.Marshal(&postRequest{
		Client: ClientData{
			ID: uuidUser,
			SubID: subId,
			Email: tgUserIDStr,
			Enable: true,
			ExpiryTime: 0,
			TotalGb: int64(totalGB) * 1024 * 1024 * 1024,
			LimitIP: int64(limitDevice),
		},
		InboundIds: []int32{1},
	})
	if err != nil {
		log.Print("Ошибка marshal")
	}
	urlAdd := xuiHost + clientAdd
	httpReq, err := http.NewRequest("POST", urlAdd, bytes.NewBuffer(req))
	httpReq.Header.Set("Content-Type","application/json")
	httpReq.Header.Set("Authorization","Bearer " + xuiToken)

	client := http.Client{}
	resp, err := client.Do(httpReq)
	
	if err != nil {
		log.Print("ошибка создания httpReq")
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	
	subBaseURL := os.Getenv("XUI_SUB_URL")
    link := fmt.Sprintf("%s/sub/%s", subBaseURL, subId)
	return &ClientResult{
		ID: uuidUser,
		SubID: subId,
		Link: link,
	}, nil
	
}

