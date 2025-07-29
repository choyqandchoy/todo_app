package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      int    `json:"user_id"`
}

func CreateTask(task Task) error {
	query := `INSERT INTO tasks (title, description, completed, user_id) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, task.Title, task.Description, task.Completed, task.UserID)
	return err
}

func GetTasksByUserID(userID int) ([]Task, error) {
	query := `SELECT id, title, description, completed, user_id FROM tasks WHERE user_id = $1`
	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTaskByID(taskID, userID int) (Task, error) {
	var task Task
	query := `SELECT id, title, description, completed, user_id FROM tasks WHERE id = $1 AND user_id = $2`
	err := DB.QueryRow(query, taskID, userID).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.UserID)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
