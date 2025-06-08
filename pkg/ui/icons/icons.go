package icons

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CircleStack() Node {
	return icon("CircleStack",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"),
		),
	)
}

func Eyes() Node {
	return icon("Eyes",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"),
		),
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"),
		),
	)
}

func UserCircle() Node {
	return icon("UserCircle",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M17.982 18.725A7.488 7.488 0 0 0 12 15.75a7.488 7.488 0 0 0-5.982 2.975m11.963 0a9 9 0 1 0-11.963 0m11.963 0A8.966 8.966 0 0 1 12 21a8.966 8.966 0 0 1-5.982-2.275M15 9.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"),
		),
	)
}

func Globe() Node {
	return icon("Globe",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9S9.515 3 12 3m0 0a8.997 8.997 0 0 1 7.843 4.582M12 3a8.997 8.997 0 0 0-7.843 4.582m15.686 0A11.953 11.953 0 0 1 12 10.5c-2.998 0-5.74-1.1-7.843-2.918m15.686 0A8.959 8.959 0 0 1 21 12c0 .778-.099 1.533-.284 2.253m0 0A17.919 17.919 0 0 1 12 16.5c-3.162 0-6.133-.815-8.716-2.247m0 0A9.015 9.015 0 0 1 3 12c0-1.605.42-3.113 1.157-4.418"),
		),
	)
}

func Home() Node {
	return icon("Home",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"),
		),
	)
}

func Info() Node {
	return icon("Info",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"),
		),
	)
}

func Mail() Node {
	return icon("Mail",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75"),
		),
	)
}

func Archive() Node {
	return icon("Archive",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m20.25 7.5-.625 10.632a2.25 2.25 0 0 1-2.247 2.118H6.622a2.25 2.25 0 0 1-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125Z"),
		),
	)
}

func PencilSquare() Node {
	return icon("PencilSquare",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10"),
		),
	)
}

func Document() Node {
	return icon("Document",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z"),
		),
	)
}

func Exit() Node {
	return icon("Exit",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15"),
		),
	)
}

func Enter() Node {
	return icon("Enter",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15M12 9l-3 3m0 0 3 3m-3-3h12.75"),
		),
	)
}

func UserPlus() Node {
	return icon("UserPlus",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M18 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 1 1-6.75 0 3.375 3.375 0 0 1 6.75 0ZM3 19.235v-.11a6.375 6.375 0 0 1 12.75 0v.109A12.318 12.318 0 0 1 9.374 21c-2.331 0-4.512-.645-6.374-1.766Z"),
		),
	)
}

func QuestionCircle() Node {
	return icon("QuestionCircle",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 5.25h.008v.008H12v-.008Z"),
		),
	)
}

func XCircle() Node {
	return icon("XCircle",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m9.75 9.75 4.5 4.5m0-4.5-4.5 4.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"),
		),
	)
}

func MagnifyingGlass() Node {
	return icon("MagnifyingGlass",
		El("path",
			Attr("stroke-linecap", "round"),
			Attr("stroke-linejoin", "round"),
			Attr("d", "m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"),
		),
	)
}

func icon(id string, els ...Node) Node {
	return cache.SetIfNotExists(fmt.Sprintf("icon.%s", id), func() Node {
		return SVG(
			Attr("xmlns", "http://www.w3.org/2000/svg"),
			Attr("fill", "none"),
			Attr("viewBox", "0 0 24 24"),
			Attr("stroke-width", "1.5"),
			Attr("stroke", "currentColor"),
			Class("w-5 h-5"),
			Group(els),
		)
	})
}
