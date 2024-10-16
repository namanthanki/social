package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
		SELECT c.id, c.content, c.user_id, c.post_id, c.created_at, c.updated_at, u.username, u.id
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = $1
		ORDER BY c.created_at ASC
	`

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.UserID,
			&comment.PostID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.User.Username,
			&comment.User.ID,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
