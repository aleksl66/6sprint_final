package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler возвращает содержимое index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла и конвертацию
func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Парсим форму (максимум 10 МБ)
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			logger.Println("Ошибка парсинга формы:", err)
			http.Error(w, "Ошибка обработки формы", http.StatusInternalServerError)
			return
		}

		// Получаем файл из формы
		file, header, err := r.FormFile("myFile")
		if err != nil {
			logger.Println("Ошибка получения файла:", err)
			http.Error(w, "Файл не найден в форме", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Читаем содержимое файла
		data, err := io.ReadAll(file)
		if err != nil {
			logger.Println("Ошибка чтения файла:", err)
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		// Конвертируем
		result := service.AutoConvert(string(data))

		// Генерируем имя выходного файла
		ext := filepath.Ext(header.Filename)
		baseName := header.Filename[:len(header.Filename)-len(ext)]
		outputName := baseName + "_" + time.Now().UTC().Format("20060102_150405") + ext

		// Сохраняем результат в локальный файл
		if err := os.WriteFile(outputName, []byte(result), 0644); err != nil {
			logger.Println("Ошибка сохранения файла:", err)
			http.Error(w, "Ошибка сохранения результата", http.StatusInternalServerError)
			return
		}

		logger.Printf("Файл сохранён: %s\n", outputName)

		// Возвращаем результат
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
