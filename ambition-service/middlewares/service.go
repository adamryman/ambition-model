package middlewares

import (
	pb "github.com/adamryman/ambition-model/ambition-service"
)

func WrapService(in pb.AmbitionServer) pb.AmbitionServer {
	return in
}
