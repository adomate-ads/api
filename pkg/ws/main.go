package ws

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	website_parse "github.com/adomate-ads/api/pkg/website-parse"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var Wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Job struct {
	conn *websocket.Conn
	msg  []byte
}

func Wshandler(w http.ResponseWriter, r *http.Request, jobs chan<- Job, wg *sync.WaitGroup) {
	conn, err := Wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		if messageType == websocket.TextMessage {
			wg.Add(1)
			jobs <- Job{conn: conn, msg: message}
		}
	}
}

func Worker(wg *sync.WaitGroup, jobs <-chan Job) {
	for job := range jobs {
		processMessage(job)
		wg.Done()
	}
}

type Message struct {
	Step      uint     `json:"step"`
	Domain    string   `json:"domain,omitempty"`
	Locations []string `json:"locations,omitempty"`
	Services  []string `json:"services,omitempty"`
	Budget    uint     `json:"budget,omitempty"`
}

type Response struct {
	Status  uint   `json:"status,omitempty"` //200(OK), 500(Error)
	Message string `json:"message,omitempty"`
}

func processMessage(job Job) {
	var msg Message
	err := json.Unmarshal(job.msg, &msg)
	if err != nil {
		resp := Response{
			Status:  500,
			Message: "Error unmarshalling message",
		}
		respString, err := json.Marshal(resp)
		if err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
	}

	if msg.Step == 1 {
		if msg.Domain != "" {
			//TODO - Get Domain Screenshot
			if err := website_parse.Screenshot(msg.Domain); err != nil {
				discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
			}
			//TODO - Get Domain Locations
			//sampleLocation := []string{"Houston, TX", "Dallas, TX", "Austin, TX"}
			//TODO - Get Domain Services
			//sampleService := []string{"Dental Services", "Medical Services", "Eye Care Services"}

			preregistration := models.PreRegistration{
				Domain: msg.Domain,
			}
			err := preregistration.CreatePreRegistration()
			if err != nil {
				resp := Response{
					Status:  500,
					Message: "Error creating pre-registration",
				}
				respString, err := json.Marshal(resp)
				if err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
				if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
				return
			} else {
				preregString, err := json.Marshal(preregistration)
				if err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
					return
				}
				resp := Response{
					Status:  200,
					Message: string(preregString),
				}
				respString, err := json.Marshal(resp)
				if err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
				if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
				return
			}
		} else {
			resp := Response{
				Status:  500,
				Message: "Domain is required",
			}
			respString, err := json.Marshal(resp)
			if err != nil {
				discord.SendMessage(discord.Warn, "Websocket Error: "+err.Error(), "")
			}
			if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
				discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
			}
			return
		}
	} else if msg.Step == 2 {
		prereg, err := models.GetPreRegistrationByDomain(msg.Domain)
		if err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		if msg.Locations != nil {
			for _, location := range msg.Locations {
				loc := models.PreRegLocation{
					PreRegistrationID: prereg.ID,
					Location:          location,
				}
				err := loc.CreatePreRegLocation()
				if err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
			}
		}
		if msg.Services != nil {
			for _, service := range msg.Services {
				serv := models.PreRegService{
					PreRegistrationID: prereg.ID,
					Service:           service,
				}
				err := serv.CreatePreRegService()
				if err != nil {
					discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
				}
			}
		}

		resp := Response{
			Status:  200,
			Message: "",
		}
		respString, err := json.Marshal(resp)
		if err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		return
	} else if msg.Step == 3 {
		//Response will be a desired budget
		prereg, err := models.GetPreRegistrationByDomain(msg.Domain)
		if err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		if msg.Budget != 0 {
			prereg.Budget = msg.Budget
			err := prereg.UpdatePreRegistration()
			if err != nil {
				discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
			}
		}
	} else {
		resp := Response{
			Status:  500,
			Message: "Step is required",
		}
		respString, err := json.Marshal(resp)
		if err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		if err := job.conn.WriteMessage(websocket.TextMessage, respString); err != nil {
			discord.SendMessage(discord.Error, "Websocket Error: "+err.Error(), "")
		}
		return
	}

}
