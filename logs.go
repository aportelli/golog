/*
Copyright © 2022 Antonin Portelli <antonin.portelli@me.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package log

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/fatih/color"
)

var Dbg = Logger{
	Level:     2,
	Color:     color.New(color.FgYellow),
	StdLogger: log.New(os.Stdout, "DEBUG   : ", 0),
}

var Warn = Logger{
	Level:     0,
	Color:     color.New(color.FgHiMagenta, color.Bold),
	StdLogger: log.New(os.Stderr, "WARNING : ", 0),
}

var Err = Logger{
	Level:     0,
	Color:     color.New(color.FgRed, color.Bold),
	StdLogger: log.New(os.Stderr, "ERROR   : ", 0),
}

var Inf = Logger{
	Level:     1,
	Color:     color.New(color.FgHiBlack),
	StdLogger: log.New(os.Stdout, "INFO    : ", 0),
}

var Msg = Logger{
	Level:     0,
	Color:     color.New(),
	StdLogger: log.New(os.Stdout, "", 0),
}

var Emph = Logger{
	Level:     0,
	Color:     color.New(color.FgGreen, color.Bold),
	StdLogger: log.New(os.Stdout, "", 0),
}

func ErrorCheck(err error, message string) {
	if err != nil {
		Err.Println(err.Error())
		if message != "" {
			Err.Println(message)
		}
		Dbg.Println(string(debug.Stack()))
		os.Exit(1)
	}
}
