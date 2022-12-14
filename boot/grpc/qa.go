package boot

import "CSAwork/server"

func QAgrpcSetup() {
	go server.Qcreate()
	go server.Acreate()
}
