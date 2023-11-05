package router

import (
	"AdHub/internal/pkg/entities"
	"encoding/json"
	"net/http"
)

func (tr *TargetRouter) CreateTargetHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		//Token  string `json:"token"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		MinAge int    `json:"min_age"`
		MaxAge int    `json:"max_age"`
		// Никита, тебе приходят теги в формате "tag1, tag2, tag3" и т.д. Ты должке их разделить по "," и получить массив строк
		/* func splitTags(input string) []string {
			tags := strings.Split(input, ",")
			for i := range tags {
				tags[i] = strings.TrimSpace(tags[i])
			}
			return tags
		} */
		// Также с регионами, интересами и ключевыми словами
		// Держи функцию
	}

	// Получение данных из запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.logger.Error("Invalid request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Получение айди пользователя из сессии
	//ownerId, err := tr.Session.GetUserId(request.Token)
	//if err != nil {
	//	tr.logger.Error("Error getting user ID from session: " + err.Error())
	//	http.Error(w, "Error getting user ID", http.StatusBadRequest)
	//	return
	//}

	uidAny := r.Context().Value("userid")
	ownerId, ok := uidAny.(int)
	if !ok {
		tr.logger.Error("user id is not an integer")
		http.Error(w, "auth error", http.StatusInternalServerError)
		return
	}

	newTarget := entities.Target{
		Name:     request.Name,
		Owner_id: ownerId,
		Gender:   request.Gender,
		Min_age:  request.MinAge,
		Max_age:  request.MaxAge,
	}

	// Сохранение в бд
	targetCreated, err := tr.Target.TargetCreate(&newTarget)
	if err != nil {
		tr.logger.Error("Error creating new target: " + err.Error())
		http.Error(w, "Error creating new target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targetCreated); err != nil {
		tr.logger.Error("Error encoding response: " + err.Error())
	}
}
