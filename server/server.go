package server

import (
	"fmt"
	"log" // import the log package
	"net"
	"encoding/json"
)

type Ent struct {
	X_pos int `json:"x_pos"`
	Y_pos int `json:"y_pos"`
}

var Entities []*Entity

func Start() {
	id_dis := 0
	// will use redis later
	var entities []*Entity
	// Start the server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error while listening: ", err)
	}
	run(ln, id_dis, entities)



}

func handleConnection(conn net.Conn, id int, entities []*Entity) {
	// handle connection
	fmt.Println("connected to: ", conn.RemoteAddr())

	for {
		var buffer [1024]byte // 1KB buffer to read data
		_, err := conn.Read(buffer[:])
		if err != nil {
			log.Printf("err while reading from conn: %v, exiting ...", err)
			return
		}

		fmt.Println("", string(buffer[:]))
		// a := ` {"X_pos":10,"Y_pos":20}`
		a := string(buffer[:])
		var entity Ent
        fmt.Printf("%v", a)
		_ = json.Unmarshal([]byte(a), &entity)
		fmt.Println("entity: ", entity.X_pos)

		// entities[id].Move(entity.X_pos, entity.Y_pos)

		// for _, e := range entities {
		// 	fmt.Println("entity: ", e)
		// }
 

		// fmt.Println("received: ", string(buffer[:]))

		// resp := []byte{'P', 'O', 'N', 'G', '\n'}
		// _, err = conn.Write(resp)
		// if err != nil {
		// 	log.Printf("err while reading from conn: %v, exiting ...", err)
		// 	return
		// }
	}
}

func run(ln net.Listener, id_dis int, entities []*Entity) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			
			fmt.Println("error tcp: ", err)
			continue
		}
		id_dis++
		entities = append(entities, NewEntity(id_dis, 0, 0))
		go handleConnection(conn, id_dis, entities)
	}
}
