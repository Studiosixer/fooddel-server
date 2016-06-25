package main

type Driver struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Available bool   `json:"completed"`
}

type Drivers []Driver
