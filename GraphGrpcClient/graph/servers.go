package graph

import (
	"cesargdd/social-media-grpcGraphClient/pb"
	"log"

	"google.golang.org/grpc"
)

func Server() (pb.UsersServiceClient, pb.PostsServiceClient, pb.CommentsServiceClient, pb.FollowersServiceClient, pb.HashtagsServiceClient, pb.PostLikesServiceClient, pb.CommentLikesServiceClient, pb.HashtagPostsServiceClient, *grpc.ClientConn) {
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()

	u := pb.NewUsersServiceClient(cc)
	p := pb.NewPostsServiceClient(cc)
	c := pb.NewCommentsServiceClient(cc)
	f := pb.NewFollowersServiceClient(cc)
	h := pb.NewHashtagsServiceClient(cc)
	pl := pb.NewPostLikesServiceClient(cc)
	cl := pb.NewCommentLikesServiceClient(cc)
	hp := pb.NewHashtagPostsServiceClient(cc)
	connection := cc
	return u, p, c, f, h, pl, cl, hp, connection
}
