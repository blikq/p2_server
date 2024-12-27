package server

import (
	// "fmt"
	"log"
	"net"
	"encoding/json"
)



var Entities []*Entity

func Start() {
	id_ := 0

	// Start the server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("error while listening:", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}

		entity := NewEntity(id_, 0, 0)
		id_++
		Entities = append(Entities, entity)

		resp, err := json.Marshal(entity)
		if err != nil {
			log.Println("error while marshaling entity to JSON:", err)
			continue
		}

		_, err = conn.Write(resp)
		if err != nil {
			log.Printf("error while writing to conn: %v, exiting ...", err)
			return
		}

		go handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("connected to:", conn.RemoteAddr())

	for {
		// Read update from the connection
		buffer := make([]byte, 1024) // 1KB buffer to read data
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("error while reading from conn: %v, exiting ...", err)
			return
		}

		data := buffer[:n]
		log.Println("received:", string(data))
		var entity Entity
		err = json.Unmarshal(data, &entity)
		if err != nil {
			log.Println("error while unmarshaling JSON:", err)
			continue
		}

		// Update the entity
		Entities[entity.Id].Move(entity.X_pos, entity.Y_pos)

		// Send the updated entity back
		jsonResp, err := json.Marshal(Entities[entity.Id])
		if err != nil {
			log.Println("error while marshaling entity to JSON:", err)
			continue
		}

		_, err = conn.Write(jsonResp)
		if err != nil {
			log.Printf("error while writing to conn: %v, exiting ...", err)
			return
		}


	}
}


