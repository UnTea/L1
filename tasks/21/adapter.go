package main

import "fmt"

type motherBoard struct{}

func (sd *motherBoard) Connect(device StorageDevice) {
	device.ConnectDevice()
	fmt.Println("Device has been connected")
}

type StorageDevice interface {
	ConnectDevice()
}

type flashDrive struct{}

func (fd *flashDrive) ConnectViaSATA() {
	fmt.Println("Flash drive connected to mother board")
}

type hardDiskDrive struct{}

func (hdd *hardDiskDrive) ConnectDevice() {
	fmt.Println("HDD connected to mother board")
}

type solidStateDrive struct {
	fd *flashDrive
}

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
