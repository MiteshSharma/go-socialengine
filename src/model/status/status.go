package status

import(
	"time"
)

type UserStatus struct {
	UserStatusId	int	   `gorm:"primary_key;auto_increment" json:"userStatusId"`
	UserId      	int    `sql:"not null" json:"userId"`            
	Status    		string `sql:"DEFAULT:NULL" json:"status"`				   
	Type       		string `sql:"DEFAULT:NULL" json:"type"`				   
	LikesCount 		int		`sql:"DEFAULT:0" json:"likeCount"`				   
	CommentCount    int 	`sql:"DEFAULT:0" json:"commentCount"`				   
	IsDeleted   	int 	`sql:"DEFAULT:0"`
	CreatedAt   	time.Time	`json:"createdAt"`
	UpdatedAt   	time.Time	`json:"updatedAt"`
	StatusLikes     []UserStatusLike   	`json:"userStatusLikes"` // One-To-Many relationship (has many)
	StatusComments  []UserStatusComment `json:"userStatusComments"` // One-To-Many relationship (has many)
}

type UserStatusLike struct {
	ID				 int	`gorm:"primary_key;auto_increment" json:"userStatusLikeId"`
	UserStatusId     int    `sql:"not null" json:"userStatusId"`
	UserId    		 int    `sql:"not null" json:"userId"`
	IsDeleted   	 int 	`sql:"DEFAULT:0"`
	CreatedAt   	 time.Time									`json:"createdAt"`
}

type UserStatusComment struct {
	ID					 int	`gorm:"primary_key;auto_increment" json:"userStatusCommentId"`
	UserStatusId     	int    `sql:"not null" json:"userStatusId"`
	UserId    		 	int    `sql:"not null" json:"userId"`
	Comment    		 	string `sql:"not null" json:"comment"`
	IsDeleted   	 	int 	`sql:"DEFAULT:0"`
	CreatedAt   	 	time.Time									`json:"createdAt"`
	UpdatedAt  	 		time.Time									`json:"updatedAt"`
}