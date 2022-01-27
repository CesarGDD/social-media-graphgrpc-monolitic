package main

import (
	"cesargdd/social-media-grpc/pb"
	"cesargdd/social-media-grpc/pg"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UsersServiceServer
	pb.PostsServiceServer
	pb.CommentsServiceServer
	pb.PostLikesServiceServer
	pb.CommentLikesServiceServer
	pb.HashtagsServiceServer
	pb.FollowersServiceServer
	pb.HashtagPostsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Users

func (*server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("Create user request")
	userReq := req.GetUser()
	user, err := db.CreateUser(ctx, pg.CreateUserParams{
		Username:  userReq.GetUsername(),
		Bio:       userReq.GetBio(),
		Avatar:    userReq.GetAvatar(),
		Email:     userReq.GetEmail(),
		Password:  userReq.GetPassword(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error Creating User", err)
	}
	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Bio:       user.Bio,
			Email:     user.Email,
			Password:  user.Password,
			Avatar:    user.Avatar,
		},
	}, nil
}

func (*server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Println("Get user request")
	getUser, err := db.GetUserById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting the user", err)
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        getUser.Id,
			CreatedAt: getUser.CreatedAt,
			UpdatedAt: getUser.UpdatedAt,
			Username:  getUser.Username,
			Bio:       getUser.Bio,
			Email:     getUser.Email,
			Password:  getUser.Password,
			Avatar:    getUser.Avatar,
		},
	}, nil
}
func (*server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	updateUser, err := db.UpdateUser(ctx, pg.UpdateUserParams{
		Id:        req.GetId(),
		Bio:       req.GetBio(),
		Avatar:    req.GetAvatar(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("error updating user", err)
	}
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:        updateUser.Id,
			CreatedAt: updateUser.CreatedAt,
			UpdatedAt: updateUser.UpdatedAt,
			Username:  updateUser.Username,
			Bio:       updateUser.Bio,
			Email:     updateUser.Email,
			Password:  updateUser.Password,
			Avatar:    updateUser.Avatar,
		},
	}, nil
}
func (*server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deleteUser, err := db.DeleteUser(ctx, req.GetId())
	if err != nil {
		fmt.Println("error deleting user", err)
	}
	return &pb.DeleteUserResponse{
		User: &pb.User{
			Id:       deleteUser.Id,
			Username: deleteUser.Username,
			Email:    deleteUser.Email,
		},
	}, nil
}

func (*server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	fmt.Println("List user request")
	users, err := db.ListUsers(ctx)
	if err != nil {
		fmt.Println("error listing users", err)
	}
	data := &pb.ListUsersResponse{}
	copier.Copy(&data.User, &users)
	return &pb.ListUsersResponse{
		User: data.User,
	}, nil
}

func (*server) ListUsersById(ctx context.Context, req *pb.ListUsersByIdRequest) (*pb.ListUsersByIdResponse, error) {
	users, err := db.ListUsersById(ctx, req.GetId())
	if err != nil {
		fmt.Println("error listing users by id", err)
	}
	data := &pb.ListUsersByIdResponse{}
	copier.Copy(&data.User, &users)
	return &pb.ListUsersByIdResponse{
		User: data.User,
	}, nil
}

// Posts

func (*server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	postReq := req.GetPost()
	createPost, err := db.CreatePost(ctx, pg.CreatePostParams{
		Caption:   postReq.GetCaption(),
		Url:       postReq.GetUrl(),
		UserId:    postReq.GetUserId(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("error creating Post", err)
	}
	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        createPost.Id,
			CreatedAt: createPost.CreatedAt,
			UpdatedAt: createPost.UpdatedAt,
			Url:       createPost.Url,
			Caption:   createPost.Caption,
			UserId:    createPost.UserId,
		},
	}, nil
}
func (*server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := db.GetPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Post", err)
	}
	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Url:       post.Url,
			Caption:   post.Caption,
			UserId:    post.UserId,
		},
	}, nil
}
func (*server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	updatePost, err := db.UpdatePost(ctx, pg.UpdatePostParams{
		Id:        req.GetPostId(),
		Url:       req.GetUrl(),
		Caption:   req.GetCaption(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error updating Post")
	}
	return &pb.UpdatePostResponse{
		Post: &pb.Post{
			Id:        updatePost.Id,
			CreatedAt: updatePost.CreatedAt,
			UpdatedAt: updatePost.UpdatedAt,
			Url:       updatePost.Url,
			Caption:   updatePost.Caption,
			UserId:    updatePost.UserId,
		},
	}, nil
}
func (*server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	delPost, err := db.DeletePost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Post")
	}
	return &pb.DeletePostResponse{
		Post: &pb.Post{
			Id:        delPost.Id,
			CreatedAt: delPost.CreatedAt,
			UpdatedAt: delPost.UpdatedAt,
			Url:       delPost.Url,
			Caption:   delPost.Caption,
			UserId:    delPost.UserId,
		},
	}, nil
}
func (*server) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	listPost, err := db.ListPosts(ctx)
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsById(ctx context.Context, req *pb.ListPostsByIdRequest) (*pb.ListPostsByIdResponse, error) {
	listPost, err := db.ListPostsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsByIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsByIdResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsByUserId(ctx context.Context, req *pb.ListPostsByUserIdRequest) (*pb.ListPostsByUserIdResponse, error) {
	listPost, err := db.ListPostsByUserId(ctx, req.GetUserId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsByUserIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsByUserIdResponse{
		Post: data.Post,
	}, nil
}

// Comments

func (*server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	createComment, err := db.CreateComment(ctx, pg.CreateCommentParams{
		Contents:  req.GetComment().GetContents(),
		UserId:    req.GetComment().GetUserId(),
		PostId:    req.GetComment().GetPostId(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Comment", err)
	}
	return &pb.CreateCommentResponse{
		Comment: &pb.Comment{
			Id:        createComment.Id,
			CreatedAt: createComment.CreatedAt,
			UpdatedAt: createComment.UpdatedAt,
			Contents:  createComment.Contents,
			UserId:    createComment.UserId,
			PostId:    createComment.UserId,
		},
	}, nil
}
func (*server) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	comment, err := db.GetComment(ctx, req.GetId())
	if err != nil {
		fmt.Println("Can not get comment", err)
	}
	return &pb.GetCommentResponse{
		Comment: &pb.Comment{
			Id:        comment.Id,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			Contents:  comment.Contents,
			UserId:    comment.UserId,
			PostId:    comment.UserId,
		},
	}, nil
}
func (*server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	updateComment, err := db.UpdateComment(ctx, pg.UpdateCommentParams{
		Id:        req.GetId(),
		Contents:  req.GetContents(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Can not update comment")
	}
	return &pb.UpdateCommentResponse{
		Comment: &pb.Comment{
			Id:        updateComment.Id,
			CreatedAt: updateComment.CreatedAt,
			UpdatedAt: updateComment.UpdatedAt,
			Contents:  updateComment.Contents,
			UserId:    updateComment.UserId,
			PostId:    updateComment.UserId,
		},
	}, nil
}
func (*server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	delComment, err := db.DeleteComment(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Comment", err)
	}
	return &pb.DeleteCommentResponse{
		Comment: &pb.Comment{
			Id:        delComment.Id,
			CreatedAt: delComment.CreatedAt,
			UpdatedAt: delComment.UpdatedAt,
			Contents:  delComment.Contents,
			UserId:    delComment.UserId,
			PostId:    delComment.UserId,
		},
	}, nil
}
func (*server) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	comments, err := db.ListComments(ctx)
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsResponse{
		Comment: data.Comment,
	}, nil
}

func (*server) ListCommentsById(ctx context.Context, req *pb.ListCommentsByIdRequest) (*pb.ListCommentsByIdResponse, error) {
	comments, err := db.ListCommentsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsByIdResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsByIdResponse{
		Comment: data.Comment,
	}, nil
}
func (*server) ListCommentsByPostId(ctx context.Context, req *pb.ListCommentsByPostIdRequest) (*pb.ListCommentsByPostIdResponse, error) {
	comments, err := db.ListCommentsByPostId(ctx, req.GetPostId())
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsByPostIdResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsByPostIdResponse{
		Comment: data.Comment,
	}, nil
}

// PostLikes

func (*server) CreatePostLike(ctx context.Context, req *pb.CreatePostLikeRequest) (*pb.CreatePostLikeResponse, error) {
	createPostLike, err := db.CreatePostLike(ctx, pg.CreatePostLikeParams{
		UserId:    req.GetPostLike().GetUserId(),
		PostId:    req.GetPostLike().GetPostId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating postlike", err)
	}
	return &pb.CreatePostLikeResponse{
		PostLike: &pb.PostLike{
			Id:        createPostLike.Id,
			CreatedAt: createPostLike.CreatedAt,
			UserId:    createPostLike.UserId,
			PostId:    createPostLike.PostId,
		},
	}, nil
}
func (*server) GetPostLike(ctx context.Context, req *pb.GetPostLikeRequest) (*pb.GetPostLikeResponse, error) {
	postLike, err := db.GetPostLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting postLike", err)
	}
	return &pb.GetPostLikeResponse{
		PostLike: &pb.PostLike{
			Id:        postLike.Id,
			CreatedAt: postLike.CreatedAt,
			UserId:    postLike.UserId,
			PostId:    postLike.PostId,
		},
	}, nil
}
func (*server) DeletePostLike(ctx context.Context, req *pb.DeletePostLikeRequest) (*pb.DeletePostLikeResponse, error) {
	delPostLike, err := db.DeletePostLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting postLike", err)
	}
	return &pb.DeletePostLikeResponse{
		PostLike: &pb.PostLike{
			Id:        delPostLike.Id,
			CreatedAt: delPostLike.CreatedAt,
			UserId:    delPostLike.UserId,
			PostId:    delPostLike.PostId,
		},
	}, nil
}
func (*server) ListPostLikes(ctx context.Context, req *pb.ListPostLikesRequest) (*pb.ListPostLikesResponse, error) {
	postLikes, err := db.ListPostLikes(ctx)
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesById(ctx context.Context, req *pb.ListPostLikesByIdRequest) (*pb.ListPostLikesByIdResponse, error) {
	postLikes, err := db.ListPostLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesByIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesByIdResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesByPostId(ctx context.Context, req *pb.ListPostLikesByPostIdRequest) (*pb.ListPostLikesByPostIdResponse, error) {
	postLikes, err := db.ListPostLikesByPostId(ctx, req.GetPostId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesByPostIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesByPostIdResponse{
		PostLike: data.PostLike,
	}, nil
}

// CommentLikes

func (*server) CreateCommentLike(ctx context.Context, req *pb.CreateCommentLikeRequest) (*pb.CreateCommentLikeResponse, error) {
	createCommentLike, err := db.CreateCommentLike(ctx, pg.CreateCommentLikeParams{
		UserId:    req.GetCommentLike().GetUserId(),
		CommentId: req.GetCommentLike().GetCommentId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Commentlike", err)
	}
	return &pb.CreateCommentLikeResponse{
		CommentLike: &pb.CommentLike{
			Id:        createCommentLike.Id,
			CreatedAt: createCommentLike.CreatedAt,
			UserId:    createCommentLike.UserId,
			CommentId: createCommentLike.CommentId,
		},
	}, nil
}
func (*server) GetCommentLike(ctx context.Context, req *pb.GetCommentLikeRequest) (*pb.GetCommentLikeResponse, error) {
	commentLike, err := db.GetCommentLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting CommentLike", err)
	}
	return &pb.GetCommentLikeResponse{
		CommentLike: &pb.CommentLike{
			Id:        commentLike.Id,
			CreatedAt: commentLike.CreatedAt,
			UserId:    commentLike.UserId,
			CommentId: commentLike.CommentId,
		},
	}, nil
}
func (*server) DeleteCommentLike(ctx context.Context, req *pb.DeleteCommentLikeRequest) (*pb.DeleteCommentLikeResponse, error) {
	delCommentLike, err := db.DeleteCommentLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting CommentLike", err)
	}
	return &pb.DeleteCommentLikeResponse{
		CommentLike: &pb.CommentLike{
			Id:        delCommentLike.Id,
			CreatedAt: delCommentLike.CreatedAt,
			UserId:    delCommentLike.UserId,
			CommentId: delCommentLike.CommentId,
		},
	}, nil
}
func (*server) ListCommentLikes(ctx context.Context, req *pb.ListCommentLikesRequest) (*pb.ListCommentLikesResponse, error) {
	commentLikes, err := db.ListCommentLikes(ctx)
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesResponse{
		CommentLike: data.CommentLike,
	}, nil
}

func (*server) ListCommentLikesById(ctx context.Context, req *pb.ListCommentLikesByIdRequest) (*pb.ListCommentLikesByIdResponse, error) {
	commentLikes, err := db.ListCommentLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesByIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesByIdResponse{
		CommentLike: data.CommentLike,
	}, nil
}
func (*server) ListCommentLikesByCommentId(ctx context.Context, req *pb.ListCommentLikesByCommentIdRequest) (*pb.ListCommentLikesByCommentIdResponse, error) {
	commentLikes, err := db.ListCommentLikesByCommentId(ctx, req.GetCommentId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesByCommentIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesByCommentIdResponse{
		CommentLike: data.CommentLike,
	}, nil
}

// Hashtag

func (*server) CreateHashtag(ctx context.Context, req *pb.CreateHashtagRequest) (*pb.CreateHashtagResponse, error) {
	createHashtag, err := db.CreateHashtag(ctx, pg.CreateHashtagParams{
		Title:     req.GetHashtag().GetTitle(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Hashtag", err)
	}
	return &pb.CreateHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        createHashtag.Id,
			CreatedAt: createHashtag.CreatedAt,
			Title:     createHashtag.Title,
		},
	}, nil
}
func (*server) GetHashtagByTitle(ctx context.Context, req *pb.GetHashtagByTitleRequest) (*pb.GetHashtagByTitleResponse, error) {
	hashtag, err := db.GetHashtagByTitle(ctx, req.GetTitle())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &pb.GetHashtagByTitleResponse{
		Hashtag: &pb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) GetHashtag(ctx context.Context, req *pb.GetHashtagRequest) (*pb.GetHashtagResponse, error) {
	hashtag, err := db.GetHashtagById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &pb.GetHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) UpdateHashtag(ctx context.Context, req *pb.UpdateHashtagRequest) (*pb.UpdateHashtagResponse, error) {
	updateHashtag, err := db.UpdateHashtag(ctx, pg.UpdateHashtagParams{
		Id:    req.GetId(),
		Title: req.GetTitle(),
	})
	if err != nil {
		fmt.Println("Can not update Hashtag")
	}
	return &pb.UpdateHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        updateHashtag.Id,
			CreatedAt: updateHashtag.CreatedAt,
			Title:     updateHashtag.Title,
		},
	}, nil
}
func (*server) DeleteHashtag(ctx context.Context, req *pb.DeleteHashtagRequest) (*pb.DeleteHashtagResponse, error) {
	delHashtag, err := db.DeleteHashtag(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Hashtag", err)
	}
	return &pb.DeleteHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        delHashtag.Id,
			CreatedAt: delHashtag.CreatedAt,
			Title:     delHashtag.Title,
		},
	}, nil
}
func (*server) ListHashtags(ctx context.Context, req *pb.ListHashtagsRequest) (*pb.ListHashtagsResponse, error) {
	hashtags, err := db.ListHashtags(ctx)
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &pb.ListHashtagsResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &pb.ListHashtagsResponse{
		Hashtag: data.Hashtag,
	}, nil
}
func (*server) ListHashtagsById(ctx context.Context, req *pb.ListHashtagsByIdRequest) (*pb.ListHashtagsByIdResponse, error) {
	hashtags, err := db.ListHashtagsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &pb.ListHashtagsByIdResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &pb.ListHashtagsByIdResponse{
		Hashtag: data.Hashtag,
	}, nil
}

// HashtagPost
func (*server) CreateHashtagPost(ctx context.Context, req *pb.CreateHashtagPostRequest) (*pb.CreateHashtagPostResponse, error) {
	createHasPost, err := db.CreateHashtagPost(ctx, pg.CreateHashtagPostParams{
		HashtagId: req.GetHashtagPost().GetHashtagId(),
		PostId:    req.GetHashtagPost().GetPostId(),
	})
	if err != nil {
		fmt.Println("Error creating hashtagPost", err)
	}
	return &pb.CreateHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        createHasPost.Id,
			HashtagId: createHasPost.HashtagId,
			PostId:    createHasPost.PostId,
		},
	}, nil
}
func (*server) GetHashtagPost(ctx context.Context, req *pb.GetHashtagPostRequest) (*pb.GetHashtagPostResponse, error) {
	hashPost, err := db.GetHashtagPostById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting hashtagPost", err)
	}
	return &pb.GetHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        hashPost.Id,
			HashtagId: hashPost.HashtagId,
			PostId:    hashPost.PostId,
		},
	}, nil
}
func (*server) DeleteHashtagPost(ctx context.Context, req *pb.DeleteHashtagPostRequest) (*pb.DeleteHashtagPostResponse, error) {
	delHasPost, err := db.DeleteHashtagPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting hashtagPost", err)
	}
	return &pb.DeleteHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        delHasPost.Id,
			HashtagId: delHasPost.HashtagId,
			PostId:    delHasPost.PostId,
		},
	}, nil
}
func (*server) ListHashtagPosts(ctx context.Context, req *pb.ListHashtagPostsRequest) (*pb.ListHashtagPostsResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &pb.ListHashtagPostsResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &pb.ListHashtagPostsResponse{
		HashtagPost: data.HashtagPost,
	}, nil
}
func (*server) ListHashtagPostsById(ctx context.Context, req *pb.ListHashtagPostsByIdRequest) (*pb.ListHashtagPostsByIdResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &pb.ListHashtagPostsByIdResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &pb.ListHashtagPostsByIdResponse{
		HashtagPost: data.HashtagPost,
	}, nil
}

// Followers

func (*server) CreateFollower(ctx context.Context, req *pb.CreateFollowerRequest) (*pb.CreateFollowerResponse, error) {
	createFollower, err := db.CreateFollower(ctx, pg.CreateFollowerParams{
		LeaderId:   req.GetFollower().GetLeaderId(),
		FollowerId: req.GetFollower().GetFollowerId(),
		CreatedAt:  time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Follower", err)
	}
	return &pb.CreateFollowerResponse{
		Follower: &pb.Follower{
			Id:         createFollower.Id,
			CreatedAt:  createFollower.CreatedAt,
			LeaderId:   createFollower.LeaderId,
			FollowerId: createFollower.FollowerId,
		},
	}, nil
}
func (*server) GetFollower(ctx context.Context, req *pb.GetFollowerRequest) (*pb.GetFollowerResponse, error) {
	follower, err := db.GetFollower(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting Follower", err)
	}
	return &pb.GetFollowerResponse{
		Follower: &pb.Follower{
			Id:         follower.Id,
			CreatedAt:  follower.CreatedAt,
			LeaderId:   follower.LeaderId,
			FollowerId: follower.FollowerId,
		},
	}, nil
}
func (*server) DeleteFollower(ctx context.Context, req *pb.DeleteFollowerRequest) (*pb.DeleteFollowerResponse, error) {
	delFollower, err := db.DeleteFollower(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Follower", err)
	}
	return &pb.DeleteFollowerResponse{
		Follower: &pb.Follower{
			Id:         delFollower.Id,
			CreatedAt:  delFollower.CreatedAt,
			LeaderId:   delFollower.LeaderId,
			FollowerId: delFollower.FollowerId,
		},
	}, nil
}
func (*server) ListFollowers(ctx context.Context, req *pb.ListFollowersRequest) (*pb.ListFollowersResponse, error) {
	Followers, err := db.ListFollowers(ctx)
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersResponse{
		Follower: data.Follower,
	}, nil
}
func (*server) ListFollowersById(ctx context.Context, req *pb.ListFollowersByIdRequest) (*pb.ListFollowersByIdResponse, error) {
	Followers, err := db.ListFollowersById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersByIdResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersByIdResponse{
		Follower: data.Follower,
	}, nil
}

func (*server) ListFollowersByLeaderId(ctx context.Context, req *pb.ListFollowersByLeaderIdRequest) (*pb.ListFollowersByLeaderIdResponse, error) {
	Followers, err := db.ListFollowersByLeaderId(ctx, req.GetLeaderId())
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersByLeaderIdResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersByLeaderIdResponse{
		Follower: data.Follower,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Blog service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterUsersServiceServer(s, &server{})
	pb.RegisterPostsServiceServer(s, &server{})
	pb.RegisterCommentsServiceServer(s, &server{})
	pb.RegisterPostLikesServiceServer(s, &server{})
	pb.RegisterCommentLikesServiceServer(s, &server{})
	pb.RegisterFollowersServiceServer(s, &server{})
	pb.RegisterHashtagsServiceServer(s, &server{})
	pb.RegisterHashtagPostsServiceServer(s, &server{})

	//Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	//Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//BLock until a signal is received
	<-ch

	fmt.Println("Closing Conection Connection")
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of program")
}
