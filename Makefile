
all:
	go build webfe.go
	sudo setcap 'cap_net_bind_service=+ep' webfe

