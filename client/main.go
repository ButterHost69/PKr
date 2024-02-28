package main

import (
	"ButterHost69/PKr-client/myserver"
	"ButterHost69/PKr-client/utils"
	"ButterHost69/PKr-client/models"
	"fmt"
	"sync"
	// "time"
)

const (
	INIT_MY_PORT   = ":3000"
	INIT_MY_DOMAIN = "localhost"
	INIT_CON_TYPE  = "tcp"
)

func main() {
	var wg sync.WaitGroup
	// var recvPort string
	utils.ClearScreen()
	fmt.Println("`` Client Started ``")
	models.CreateUserIfNotExists()

	for {
		var opt int
		menu := Title.Render("~ Select An Option ~\n")
		menu += "\n"
		menu += Option.Render("1. Initialize Connection\n2. Manage Connections\n3. GET ALL Files\n4. PUSH Files\n7. Quit\n")

		fmt.Println(MenuBorder.Render(menu))
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			utils.ClearScreen()
			menu := Title.Render("~ Select Initailaization Mode ~\n")
			menu += "\n"
			menu += Option.Render("1. Listen\n2. Dial\n")
			fmt.Println(MenuBorder.Render(menu))

			fmt.Scanln(&opt)
			switch opt {
			case 1:
				fmt.Print(InputFieldLabels.Render(" Listening Connection..."))
				fmt.Println("\n")

				var port string
				var domain string
				var conType string

				fmt.Print(InputFieldLabels.Render(" Enter Domain [eg: 'localhost']: "))
				fmt.Scanln(&domain)
				fmt.Print(InputFieldLabels.Render(" Enter Port [eg: '3000']: "))
				fmt.Scanln(&port)
				fmt.Print(InputFieldLabels.Render(" Enter Connection Type [eg: 'tcp']: "))
				fmt.Scanln(&conType)
				wg.Add(1)
				server := myserver.InitListener(domain, ":"+port, conType, &wg)
				fmt.Print("\n")
				conslug := server.StartGRPCInitConnection()
				myserver.SetWorkSpaceFolders(conslug)

			case 2:
				fmt.Print(InputFieldLabels.Render(" Dialing Connection..."))
				fmt.Println("\n")

				var port string
				var domain string
				var conType string
	
				fmt.Print(InputFieldLabels.Render(" Enter Domain [eg: 'localhost']: "))
				fmt.Scanln(&domain)
				fmt.Print(InputFieldLabels.Render(" Enter Port [eg: '3000']: "))
				fmt.Scanln(&port)
				fmt.Print(InputFieldLabels.Render(" Enter Connection Type [eg: 'tcp']: "))
				fmt.Scanln(&conType)
				
				wg.Add(1)
				server := myserver.InitSender(domain, ":"+port, conType, &wg)
				fmt.Print("\n")
				conslug, err := server.DialGRPCInitConnection()
				if err != nil {
					fmt.Println("Error Occured in Dialing....")
				} else {
					fmt.Println("\n")
					myserver.SetWorkSpaceFolders(conslug)
				}
			}
		}
	}
}
