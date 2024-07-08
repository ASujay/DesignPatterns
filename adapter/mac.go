package main

// this the known/compatible service

import "fmt"

type Mac struct {
    
}

func (m *Mac) InsertIntoLightningPort(){
    fmt.Println("Lightning connector is plugged into mac machine.")
}
