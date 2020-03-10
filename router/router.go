package router

import (
	"github.com/ruandao/dwg255-fish-hall/controllers"
	"net/http"
)

func init() {
	http.HandleFunc("/get_serverinfo", controllers.GetServerInfo)
	http.HandleFunc("/guest", controllers.Guest)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/get_user_status", controllers.GetUserStatus)
	http.HandleFunc("/get_message", controllers.GetMessage)
	http.HandleFunc("/enter_public_room", controllers.EnterPublicRoom)
	http.HandleFunc("/register_game_server", controllers.RegisterGameServer)

	http.HandleFunc("/qq/login", controllers.QQLogin)
	http.HandleFunc("/qq/message", controllers.QQCallback)
}
