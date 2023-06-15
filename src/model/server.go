package model

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Network struct {
	players []PlayerConnection
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
	conn := n.players[playerIndex].conn
	_, sendErr := conn.Write([]byte(prompt))
	if sendErr != nil {
		return 0, sendErr
	}
	response := make([]byte, 1024)
	_, listErr := conn.Read(response)
	if listErr != nil {
		return 0, listErr
	}
	respInt, convErr := strconv.ParseInt(string(response), 10, 64)
	if convErr != nil {
		return 0, convErr
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
	conn := n.players[playerIndex].conn
	_, sendErr := conn.Write([]byte(info))
	if sendErr != nil {
		return sendErr
	}
	return nil
}

func (n  *Network) findPlayer(name string) (int, error) {
	for i := 0; i < len(n.players); i++ {
		if n.players[i].playerName == name {
			return i, nil
		}
	}
	return -1, errors.New("did not find online player by name")
}