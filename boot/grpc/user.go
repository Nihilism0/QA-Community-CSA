package boot

import "CSAwork/server"

func UserGrpcSetup() {
	go server.Loginserver()
	go server.Register()
}
