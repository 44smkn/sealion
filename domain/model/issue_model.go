package model

type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Summary   string `json:"summary"`
		Issuetype struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
		} `json:"issuetype"`
		Project struct {
			Self            string `json:"self"`
			ID              string `json:"id"`
			Key             string `json:"key"`
			Name            string `json:"name"`
			ProjectCategory struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"projectCategory"`
		} `json:"project"`
		Status struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			Name           string `json:"name"`
			ID             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		DueDate string `json:"duedate"`
	} `json:"fields"`
}
