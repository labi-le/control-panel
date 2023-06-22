package pm

import (
	"github.com/ShinyTrinkets/overseer"
	"github.com/labi-le/control-panel/pkg/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Command struct {
	Name string   `json:"name"`
	Exec string   `json:"exec"`
	Args []string `json:"args"`
}

func PMCommand(c echo.Context, m *Methods, cmd Command) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		overOpt := overseer.Options{
			Buffered:  false,
			Streaming: true,
		}

		over := overseer.NewOverseer()
		log.Infof("Client connected %s", ws.Request().RemoteAddr)

		cmd := over.Add(cmd.Name, cmd.Exec, cmd.Args, overOpt)

		if err := utils.ManageProc(cmd, over, ws); err != nil {
			m.badResponseWS(ws, err)
			return
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (a *API) UpdatePackage(c echo.Context) error {
	return PMCommand(c, m, Command{
		Name: "update",
		Exec: "apt",
		Args: []string{
			"update",
			"-y",
		},
	})
}

func (a *API) InstallPackage(c echo.Context) error {
	//return PMCommand(c, m, Command{
	//	Name: "install",
	//	Exec: "apt",
	//	Args: []string{
	//		"install",
	//		"-y",
	//		c.Param("package"),
	//	},
	//})
	return PMCommand(c, m, Command{
		Name: "install",
		Exec: "pacman",
		Args: []string{
			"--noconfirm",
			"-Syyuu",
			c.Param("package"),
		},
	})
}

func (m *API) DeletePackage(c echo.Context) error {
	return PMCommand(c, m, Command{
		Name: "remove",
		Exec: "apt",
		Args: []string{
			"remove",
			"-y",
			c.Param("package"),
		},
	})

	//return PMCommand(c, m, Command{
	//	Name: "remove",
	//	Exec: "pacman",
	//	Args: []string{
	//		"-Rs",
	//		"--noconfirm",
	//		c.Param("package"),
	//	},
	//})
}
