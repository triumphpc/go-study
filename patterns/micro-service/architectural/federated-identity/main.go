// Паттерн "Federated Identity" позволяет пользователям аутентифицироваться в одном месте и использовать эту аутентификацию
//для доступа к нескольким системам или сервисам. Это достигается за счет доверительных отношений между различными системами,
//которые доверяют одному или нескольким поставщикам идентификационных данных (Identity Providers, IdP).

// Когда использовать Federated Identity
//Мультисервисные системы: Когда пользователи должны иметь доступ к нескольким независимым системам или сервисам.
//Упрощение аутентификации: Когда необходимо упростить процесс аутентификации для пользователей, позволяя им использовать одну учетную запись для доступа к нескольким системам.
//Интеграция с внешними системами: Когда требуется интеграция с внешними системами, которые уже используют определенного поставщика идентификационных данных.

//Плюсы Federated Identity
//Удобство для пользователей: Пользователи могут использовать одну учетную запись для доступа к нескольким системам, что упрощает процесс аутентификации.
//Централизованное управление: Упрощает управление учетными записями и доступом, так как все учетные данные хранятся у одного или нескольких поставщиков идентификационных данных.
//Безопасность: Уменьшает количество мест, где хранятся учетные данные, что снижает риск утечки данных.
//Гибкость: Позволяет легко интегрировать новые системы и сервисы, которые доверяют тем же поставщикам идентификационных данных.

//Минусы Federated Identity
//Зависимость от IdP: Система становится зависимой от доступности и безопасности поставщика идентификационных данных.
//Сложность настройки: Требует настройки доверительных отношений между системами и поставщиками идентификационных данных.
//Потенциальные задержки: Может возникнуть задержка при аутентификации, если IdP находится в другой сети или регионе.
//Управление отказами: Необходимо предусмотреть механизмы для обработки отказов в случае недоступности IdP.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauthStateString {
		http.Error(w, "State is invalid", http.StatusBadRequest)
		return
	}

	token, err := oauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Code exchange failed", http.StatusInternalServerError)
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "Login successful")
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/callback", handleCallback)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
