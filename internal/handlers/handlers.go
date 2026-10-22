package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("/index.html")
	if err != nil {
		http.Error(res, "не удалось открыть index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "не удалось обработать index.html", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Write(data)
}
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	//Тут функция это парсить форму, но пока непонятно что с нее ты получаешь потом и как их читать и дальше передавать\
	file, header, err := req.FormFile("myFile") // в  index смотри что на выходе
	if err != nil {
		http.Error(res, "Не удалось обработать данные формы", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError) //ПЕРЕДЕЛАЙ НА ЯНДЕКСЕ ЕСТЬ КАК СЧИТАТЬ ФАЙЛ ПО СТРОКАМ ЧТОБЫ НЕ ЗАГРУЖАТЬ ПАМЯТЬ
		return
	}

	dataOutput, err := service.AutoSelect(string(data))
	if err != nil {
		http.Error(res, "не удалось обработать создать файл результат", http.StatusInternalServerError)
	}
	ext := filepath.Ext(header.Filename)
	name := time.Now().UTC().String() + ext
	resFile, err := os.Create(name)
	if err != nil {
		http.Error(res, "не удалось обработать создать файл результат", http.StatusInternalServerError)
	}
	defer resFile.Close()
	fmt.Fprintf(resFile, "%s \n", dataOutput)
}
