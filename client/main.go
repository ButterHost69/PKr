package main

import (
	"ButterHost69/PKr-client/models"
	"ButterHost69/PKr-client/myserver"
	"ButterHost69/PKr-client/utils"
	"strconv"

	// "bufio"
	"fmt"
	"os"
	"sync"
	// "time"
)

const (
	INIT_MY_PORT   = ":3000"
	INIT_MY_DOMAIN = "0.0.0.0"
	INIT_CON_TYPE  = "tcp"

	CMD_LISTN_PORT   = ":3069"
	CMD_LISTN_DOMAIN = "0.0.0.0"
	CMD_LISTN_TYPE  = "tcp"
)

func main() {
	var wg sync.WaitGroup
	// var recvPort string
	utils.ClearScreen()
	fmt.Println("`` Client Started ``")
	models.CreateUserIfNotExists()

	// Add Wait Group and then go func this son of a bitch
	// Creare an InitWGListener() ~
	// wg.Add(1)
	cmdListn := myserver.InitListener(CMD_LISTN_DOMAIN, CMD_LISTN_PORT, CMD_LISTN_TYPE)
	go cmdListn.StartGRPCCmdConnection()

	for {
		var opt int
		menu := Title.Render("~ Select An Option ~\n")
		menu += "\n"
		menu += Option.Render("1. Initialize Connection\n2. Manage Connections\n3. GET ALL Files\n4. PUSH Files\n7. Quit\n")

		fmt.Println(MenuBorder.Render(menu))
		fmt.Scan(&opt)
		utils.ClearScreen()
		switch opt {
		case 1:
			utils.ClearScreen()
			menu := Title.Render("~ Select Initailaization Mode ~\n")
			menu += "\n"
			menu += Option.Render("1. Listen\n2. Dial\n")
			fmt.Println(MenuBorder.Render(menu))

			fmt.Scan(&opt)
			switch opt {
			case 1:
				fmt.Print(InputFieldLabels.Render(" Listening Connection..."))
				fmt.Print("\n\n")

				var port string
				var domain string
				var conType string

				fmt.Print(InputFieldLabels.Render(" Enter Domain [eg: 'localhost']: "))
				fmt.Scan(&domain)
				fmt.Print(InputFieldLabels.Render(" Enter Port [eg: '3000']: "))
				fmt.Scan(&port)
				fmt.Print(InputFieldLabels.Render(" Enter Connection Type [eg: 'tcp']: "))
				fmt.Scan(&conType)
				// wg.Add(1)
				server := myserver.InitListener(domain, ":"+port, conType)
				fmt.Print("\n")
				conslug := server.StartGRPCInitConnection()
				fmt.Print("\n\n")
				myserver.SetWorkSpaceFolders(conslug)

			case 2:
				fmt.Print(InputFieldLabels.Render(" Dialing Connection..."))
				fmt.Print("\n\n")

				var port string
				var domain string
				var conType string

				fmt.Print(InputFieldLabels.Render(" Enter Domain [eg: 'localhost']: "))
				fmt.Scan(&domain)
				fmt.Print(InputFieldLabels.Render(" Enter Port [eg: '3000']: "))
				fmt.Scan(&port)
				fmt.Print(InputFieldLabels.Render(" Enter Connection Type [eg: 'tcp']: "))
				fmt.Scan(&conType)

				wg.Add(1)
				server := myserver.InitSender(domain, ":"+port, conType, &wg)
				fmt.Print("\n")
				conslug, err := server.DialGRPCInitConnection()
				if err != nil {
					fmt.Println("Error Occured in Dialing....")
				} else {
					fmt.Print("\n\n")
					myserver.SetWorkSpaceFolders(conslug)
				}

			}
		case 2:
			menu_loop:
			for{
				utils.ClearScreen()
				menu := Title.Render("~ Select an Option ~\n")
				menu += "\n"
				menu += Option.Render("1. List All Connections\n2. Send Request To a Connection\n3. Configure a Connection\n\n7. <- Back\n")
				fmt.Println(MenuBorder.Render(menu))
				fmt.Scan(&opt)
				switch opt {
				case 1:
					list_of_connections := models.GetAllConnections()
					fmt.Print("\nList of all Connections :\n\n")
					for _, con := range list_of_connections{
						fmt.Println(con.CurrentIP + ":" + con.CurrentPort  + "  |  " + con.ConnectionSlug)
					}
					fmt.Print("\nEnter c To Continue: ")
					var neko string
					fmt.Scan(&neko)
					
				case 2:
					// utils.ClearScreen()
					count := 0
					list_of_connections := models.GetAllConnections()
					fmt.Print("\nList of all Connections :\n\n")
					for _, con := range list_of_connections{
						count += 1
						fmt.Println(strconv.Itoa(count) + ". " + con.CurrentIP + ":" + con.CurrentPort + "  |  " + con.ConnectionSlug)
					}
					fmt.Print("\nChoose a Connection 1-...: ")
					var optConn int
					fmt.Scan(&optConn)
					curCon := list_of_connections[optConn - 1]
					// USE THIS AFTER SELECTION OF CONNECTION :~
					
					menu := Title.Render("~ Select an Option ~\n")
					menu += "\n"
					menu += Option.Render("1. GET Files\n2. IDK\n\n7. <- Back\n")
					fmt.Println(MenuBorder.Render(menu))
					fmt.Scan(&opt)
					switch opt {
					case 1:
						domain := curCon.CurrentIP // domain -> domain:port
						// fmt.Println(curCon.CurrentPort)
						server := myserver.InitSender(domain, curCon.CurrentPort, "tcp", &wg) //Domain already contains the port 
						server.DialCommandConnection(curCon)
						fmt.Print("\nEnter c To Continue: ")
						var neko string
						fmt.Scan(&neko)
					}
				case 3:
					utils.ClearScreen()
					menu := Title.Render("~ Select an Option ~\n")
					menu += "\n"
					menu += Option.Render("1. Add Another Workspace\n2. Remove Workspace\n\n7. <- Back\n")
					fmt.Println(MenuBorder.Render(menu))
					fmt.Scanln(&opt)
					switch opt {
					case 1:
						fmt.Println("add workspace")
					case 2:
						fmt.Println("remove workspace")
					}
				case 7:
					utils.ClearScreen()
					break menu_loop
				}
			}
		case 7:
			os.Exit(7)
		}
	}
}
