package model

import "gorm.io/gorm"

//	type Account struct {
//		gorm.Model
//		Id       int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`    // 帐户id
//		Name     string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"` //用户名称
//		Password string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
//	}
type User struct {
	gorm.Model
	Id              int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"`      // 用户id
	Name            string `protobuf:"bytes,2,req,name=name" json:"name,omitempty" gorm:"size:64; not null"`                     // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty"`      // 粉丝总数
	Avatar          string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`                                          //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`                                    //个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited" json:"total_favorited,omitempty"`   //获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount" json:"work_count,omitempty"`                 //作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"`     //点赞数量
	Password        string `protobuf:"bytes,2,req,name=password" json:"password,omitempty" gorm:"size:128; not null"`            //密码
}

type Video struct {
	gorm.Model
	Id            int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty" gorm:"not null;primary_key;auto_increment"` // 视频唯一标识
	PlayUrl       string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,req,name=title" json:"title,omitempty"`                                       // 视频标题

	User     User   `gorm:"foreignKey:AutorId"`
	AuthorId string `protobuf:"varint,1,req,name=author_id" json:"author_id,omitempty"`
}
