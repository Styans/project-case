package service

import (
	"ascii-art/internal/drawtext"
	"ascii-art/internal/handlers"
	"log"
	"net/http"
	"strconv"
)

func AsciiService() {
	fonts := drawtext.GetFonts("./internal/drawtext/fonts/")
	app := &handlers.Aplication{
		Fonts: fonts,
	}
	addr := 8080

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(addr),
		Handler: app.Route(),
	}
	err := srv.ListenAndServe()

	for err != nil {
		log.Println("Ошибка при запуске сервера:", err)
		addr += 1
		srv.Addr = ":" + strconv.Itoa(addr)
		err = srv.ListenAndServe()
	}
}
