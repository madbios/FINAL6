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
	file, err := os.Open("../index.html")
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
	file, header, err := req.FormFile("myFile") //
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
		http.Error(res, "не удалось обработать  файл ", http.StatusInternalServerError)
	}
	ext := filepath.Ext(header.Filename)
	name := time.Now().UTC().Format("2006-01-02_15-04-05") + ext
	resFile, err := os.Create(name)
	if err != nil {
		http.Error(res, "не удалось создать файл результат", http.StatusInternalServerError)
	}
	defer resFile.Close()
	fmt.Fprintf(resFile, "%s \n", dataOutput)
	res.Write([]byte(dataOutput))
}
