package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Обработчик HTTP-запросов для корневого пути ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок Content-Type для HTML
		w.Header().Set("Content-Type", "text/html")

		// HTML-код, который мы хотим отдать
		html := `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Пример HTML-страницы</title>
            </head>
            <body>
			<div id="banner-container"></div>

			<script>
			document.addEventListener('DOMContentLoaded', function() {
				fetchBanner();
			});
			
			function fetchBanner() {
				fetch("http://127.0.0.1:8080/api/v1/getad?id=3")
				.then(response => {
					console.log(response)
					return response.text()
				})
				.then(data => {
					console.log(data)
					const bannerContainer = document.getElementById('banner-container');
					bannerContainer.innerHTML = data;
				})
				.catch(error => {
					console.error('Fetch error:', error);
				});
			}
			</script>
            </body>
            </html>
        `

		// Отправляем HTML-код клиенту
		fmt.Fprintf(w, html)
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на порту 8089")
	http.ListenAndServe(":8089", nil)
}
