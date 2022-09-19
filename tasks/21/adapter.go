package main

import "fmt"

// motherBoard is a struct that uses for connecting various devices
type motherBoard struct{}

// Connect is a fucntion that is uses for connecting via interface
func (sd *motherBoard) Connect(device StorageDevice) {
	device.ConnectDevice()
	fmt.Println("Device has been connected")
}

// StorageDevice is an interface for connecting with devices
type StorageDevice interface {
	ConnectDevice()
}

// flashDrive is a struct that can be connected
type flashDrive struct{}

// ConnectViaSATA is a function that connects fd and motherboard
func (fd *flashDrive) ConnectViaSATA() {
	fmt.Println("Flash drive connected to mother board")
}

// hardDiskDrive is a struct that can be connected
type hardDiskDrive struct{}

// ConnectDevice is a function that connects hdd and motherboard
func (hdd *hardDiskDrive) ConnectDevice() {
	fmt.Println("HDD connected to mother board")
}

// solidStateDrive is a struct that can be connected
type solidStateDrive struct {
	fd *flashDrive
}

// ConnectDevice is a function that connects adapter and motherboard
func (adapter *solidStateDrive) ConnectDevice() {
	fmt.Println("SSD connected to mother board via Flash Drive adapter")
	adapter.fd.ConnectViaSATA()
}

func main() {
	hdd := hardDiskDrive{}
	fd := flashDrive{}
	board := motherBoard{}

	adapter := solidStateDrive{
		fd: &fd,
	}

	board.Connect(&hdd)
	board.Connect(&adapter)
}
