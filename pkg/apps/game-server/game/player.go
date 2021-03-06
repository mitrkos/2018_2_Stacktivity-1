package game

import (
	"2018_2_Stacktivity/models"
	"2018_2_Stacktivity/storage"
	"encoding/json"
	"log"

	"sync"

	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	mu     sync.Mutex
	user   *models.User
	enemy  *Player
	room   *Room
	conn   *websocket.Conn
	logic  models.PlayerLogic
	isOpen bool
}

type IncomingMessage struct {
	Player  *Player
	Message *models.Message
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 60 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func NewPlayer(user *models.User, conn *websocket.Conn) *Player {
	return &Player{mu: sync.Mutex{}, user: user, conn: conn, isOpen: true}
}

func (p *Player) CheckConn() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		p.conn.Close()
	}()

	for {
		<-ticker.C
		p.mu.Lock()
		p.conn.SetWriteDeadline(time.Now().Add(writeWait))
		err := p.conn.WriteMessage(websocket.PingMessage, nil)
		p.mu.Unlock()
		if err != nil {
			log.Printf("can't send message to player %s\n", p.user.Username)
			log.Println(err.Error())
			p.isOpen = false
			return
		}
	}
}

func (p *Player) Listen() {
	defer p.conn.Close()
	for {
		m := &models.Message{}
		err := p.conn.ReadJSON(m)
		if websocket.IsUnexpectedCloseError(err) {
			log.Printf("player %d was disconnected", p.user.ID)
			if p.room != nil {
				p.room.Unregister <- p
			}
			p.isOpen = false
			log.Println("player deleted")
			return
		}
		if p.room != nil {
			im := &IncomingMessage{
				Player:  p,
				Message: m,
			}
			log.Println("Read event", m.Event)
			p.room.Message <- im
		}
	}
}

func (p *Player) Send(s *models.Message) {
	log.Println("sending message to " + p.user.Username)
	log.Println("Event:", s.Event)
	p.mu.Lock()
	err := p.conn.WriteJSON(s)
	log.Println("sending message to " + p.user.Username + " done")
	p.mu.Unlock()
	if err != nil {
		log.Printf("can't send message to player %s\n", p.user.Username)
		log.Println(err.Error())
		p.room.Unregister <- p
	}
}

func (p *Player) StartMultiplayer() {
	players := make([]string, 2)
	players[0] = p.user.Username
	players[1] = p.enemy.user.Username

	log.Println("Get level ", p.room.levelNum)
	var level models.Level

	dbLevel, err := storage.GetUserStorage().GetLevelByNumber(p.room.levelNum)
	if err != nil {
		log.Println("PIZDA RULIU")
		log.Println(err.Error())
		return
	}
	if err := json.Unmarshal([]byte(dbLevel.Level), &level); err != nil {
		log.Println("HUITA KAKAYA-TO")
		log.Println(err.Error())
		return
	}

	m := &models.Message{
		Event:   models.StartGame,
		Players: &players,
		Level:   &level,
	}
	p.Send(m)
}
