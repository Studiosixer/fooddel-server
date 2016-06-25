package main

import "fmt"

var currentId int

var drivers Drivers

// Give us some seed data
func init() {
	RepoCreateDriver(Driver{Name: "John the fast"})
	RepoCreateDriver(Driver{Name: "Joe the slow"})
}

func RepoFindDriver(id int) Driver {
	for _, t := range drivers {
		if t.Id == id {
			return t
		}
	}
	// return empty Driver if not found
	return Driver{}
}

func RepoCreateDriver(t Driver) Driver {
	currentId += 1
	t.Id = currentId
	drivers = append(drivers, t)
	return t
}

func RepoDestroyDriver(id int) error {
	for i, t := range drivers {
		if t.Id == id {
			drivers = append(drivers[:i], drivers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Driver with id of %d to delete", id)
}
