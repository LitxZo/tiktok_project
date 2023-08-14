package model

import (
	"gorm.io/gorm"
)

//	type Account struct {
//		gorm.Model
//		Id       int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`    // 帐户id
//		Name     string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"` //用户名称
//		Password string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
//	}
type User struct {
	gorm.Model
	Id              int    `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"`                   // 用户id
	UserName        string `protobuf:"bytes,2,req,name=user_name" json:"user_name,omitempty" gorm:"size:64;not null;unique" validate:"email"` //用户登录名
	Name            string `protobuf:"bytes,3,req,name=name" json:"name,omitempty" gorm:"size:64; not null"`                                  // 用户名称
	FollowCount     int    `protobuf:"varint,4,opt,name=follow_count,json=followCount" json:"follow_count,omitempty"`                         // 关注总数
	FollowerCount   int    `protobuf:"varint,5,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty"`                   // 粉丝总数
	Avatar          string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`                                                       //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty"`              //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`                                                 //个人简介
	TotalFavorited  int    `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited" json:"total_favorited,omitempty"`                //获赞数量
	WorkCount       int    `protobuf:"varint,10,opt,name=work_count,json=workCount" json:"work_count,omitempty"`                              //作品数量
	FavoriteCount   int    `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"`                  //点赞数量
	Password        string `protobuf:"bytes,12,req,name=password" json:"password,omitempty" gorm:"size:128; not null"`                        //密码
	// FollowId        IdGroup `protobuf:"bytes,13,opt,name=follow_id" jsopn:"follow_id,omitempty"`                                           //关注者id
	// FavoriteId      IdGroup `protobuf:"bytes,13,opt,name=favorite_id" jsopn:"favorite_id,omitempty"`
}

type Video struct {
	gorm.Model
	Id            int    `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"` // 视频唯一标识
	PlayUrl       string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int    `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int    `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,req,name=title" json:"title,omitempty"`                                       // 视频标题

	User     User `gorm:"foreignKey:Id; references:AuthorId"`
	AuthorId int  `json:"author_id"`
}

type Comment struct {
	Id         int    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                        // 视频评论id
	Content    string `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`               // 评论内容
	CreateDate string `protobuf:"bytes,4,req,name=create_date,json=createDate" json:"create_date"` // 评论发布日期，格式 mm-dd
	User       User   `gorm:"foreignKey:Id;referances:User"`
	UserId     int    `json:"user_id"` // 评论人id
	Video      Video  `gorm:"foreignKey:Id;referances:VideoId"`
	VideoId    int    `json:"video_id"` // 视频id
}

func (u User) GetTableName() string {
	return "users"
}

func (v Video) GetTableName() string {
	return "videos"
}

// type IdGroup []int64

// func (f *IdGroup) Scan(value interface{}) error {
// 	v, _ := value.([]byte)
// 	return json.Unmarshal(v, f)
// }

// func (f IdGroup) Value() (driver.Value, error) {
// 	return json.Marshal(f)
// }
