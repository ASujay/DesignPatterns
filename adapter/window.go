package main


// this the unknown/incompatible service

import "fmt"

type Windows struct {
    
}

func (m *Windows) InsertIntoUSBPort(){
    fmt.Println("USB connector is plugged into windows machine.")
}
