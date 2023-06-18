package model

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Network struct {
	players 	[]PlayerConnection
	host		net.Conn
}

type PlayerConnection struct {
	playerName string
	conn net.Conn
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

/*
Listener that throws incomming connections to the handler.

Returns an error if the listener can not be initiated.
*/
func (n *Network) Listener() error {
	listen, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go n.handleConnection(conn)
	}
}

/*
Adds the connection to the network struct along with a generated player name.
*/
func (n *Network) handleConnection(conn net.Conn) {
	player := PlayerConnection{
		playerName: "online player " + fmt.Sprint(n.CountOnlinePlayers()),
		conn: conn,
	}
	n.players = append(n.players, player)
}

/*
Establish a connection with the host.

Returns an error if there is a problem with the dial function.
*/
func (n *Network) DialHost() error {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		return err
	}
	n.host = conn
	return nil
}

func (n *Network) Respond(response string) error {
	conn := n.host
	_, respErr := conn.Write([]byte(response))
	if respErr != nil {
		return respErr
	}
	return nil
}

func (n *Network) Listen() net.Conn {
	return n.host
}

/*
Returns how many player connections have been established.
*/
func (n *Network) CountOnlinePlayers() int {
	return len(n.players)
}

/*
Sends a CLI prompt to a player and awaits an integer response.

Returns various potential errors.
	* If the players connection can not be found an error is returned.
	* If the connections returns an error it is passed along.
	* If the response parsing returns an error it is passed along.
*/
func (n *Network) Play(playerName string, prompt string) (int, error) {
	playerIndex, err := n.findPlayer(playerName)
	if err != nil {
		return 0, err
	}
	prompt = "Play\n" + prompt
	conn := n.players[playerIndex].conn
	_, sendErr := conn.Write([]byte(prompt))
	if sendErr != nil {
		return 0, sendErr
	}
	response := make([]byte, 4096)
	_, listErr := conn.Read(response)
	if listErr != nil {
		return 0, listErr
	}
	respInt, err :=  strconv.Atoi(string(response[0]))
	if err != nil {
		return 0, err
	}
	return int(respInt), nil
}

/*
Sends a display message to a player.

Returns an error if the player can not be found or the connection returns an error.
*/
func (n *Network) Display(playerName string, info string) error {
	playerIndex, err := n.findPlayer(playerName)
	if err != nil {
		return err
	}
	info = "Display\n" + info
	conn := n.players[playerIndex].conn
	_, sendErr := conn.Write([]byte(info))
	if sendErr != nil {
		return sendErr
	}
	return nil
}

/*
Send out standard info messages to all online players.
*/
func (n *Network) MassDisplay(info string) error {
	for i := 0; i < len(n.players); i++ {
		connErr := n.Display(n.players[i].playerName, info)
		if connErr != nil {
			return connErr
		}
	}
	return nil
}

func (n *Network) End(playerName string, info string) error {
	playerIndex, err := n.findPlayer(playerName)
	if err != nil {
		return err
	}
	info = "End\n" + info
	conn := n.players[playerIndex].conn
	_, sendErr := conn.Write([]byte(info))
	if sendErr != nil {
		return sendErr
	}
	return nil
}

func (n *Network) GameOver(winner string) {
	for i := 0; i < len(n.players); i++ {
		n.End(n.players[i].playerName, winner)
	}
}

func (n  *Network) findPlayer(name string) (int, error) {
	for i := 0; i < len(n.players); i++ {
		if n.players[i].playerName == name {
			return i, nil
		}
	}
	return -1, errors.New("did not find online player by name")
}

func (n *Network) findPlayerName(index int) (string, error) {
	if index < 0 || index >= n.CountOnlinePlayers() {
		return "", errors.New("index out of bounds")
	}
	return n.players[index].playerName, nil
}

/*
Returns a list of all online player's names
*/
func (n *Network) ListPlayers() []string {
	var playerList []string
	for i := 0; i < n.CountOnlinePlayers(); i++ {
		name, err := n.findPlayerName(i)
		if err != nil {
			continue
		}
		playerList = append(playerList, name)
	}
	return playerList
}

/*
Close the connections to all online players.
*/
func (n *Network) CloseConnections() {
	for i := 0; i < len(n.players); i++ {
		n.players[i].conn.Close()
	}
}