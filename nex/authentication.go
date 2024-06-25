package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/dr-luigi/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)
	globals.AuthenticationServer.ByteStreamSettings.UseStructureHeader = true
	
	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 5, 2))
	globals.AuthenticationServer.AccessKey = "2a724fa4"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("======= Dr. Luigi - Auth =======")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("================================")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
