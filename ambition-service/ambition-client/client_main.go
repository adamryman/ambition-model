package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	grpcclient "github.com/adamryman/ambition-model/ambition-service/generated/client/grpc"
	httpclient "github.com/adamryman/ambition-model/ambition-service/generated/client/http"
	clientHandler "github.com/adamryman/ambition-model/ambition-service/handlers/client"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterAmbitionServiceServer
)

func main() {
	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5040", "gRPC (HTTP) address of addsvc")
		method   = flag.String("method", "readactions", "readactions,readaction,createaction,readoccurrences,createoccurrence")
	)

	var (
		flagUserIdReadActions        = flag.Int64("readactions.userid", 0, "")
		flagActionIdReadAction       = flag.Int64("readaction.actionid", 0, "")
		flagActionNameReadAction     = flag.String("readaction.actionname", "", "")
		flagUserIdCreateAction       = flag.Int64("createaction.userid", 0, "")
		flagActionNameCreateAction   = flag.String("createaction.actionname", "", "")
		flagUserIdReadOccurrences    = flag.Int64("readoccurrences.userid", 0, "")
		flagActionIdReadOccurrences  = flag.Int64("readoccurrences.actionid", 0, "")
		flagActionIdCreateOccurrence = flag.Int64("createoccurrence.actionid", 0, "")
		flagDatetimeCreateOccurrence = flag.String("createoccurrence.datetime", "", "")
		flagDataCreateOccurrence     = flag.String("createoccurrence.data", "", "")
	)
	flag.Parse()

	var (
		service handler.Service
		err     error
	)
	if *httpAddr != "" {
		service, err = httpclient.New(*httpAddr)
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			os.Exit(1)
		}
		defer conn.Close()
		service = grpcclient.New(conn)
	} else {
		fmt.Fprintf(os.Stderr, "error: no remote address specified\n")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch *method {

	case "readactions":
		var err error
		UserIdReadActions := *flagUserIdReadActions
		request, err := clientHandler.ReadActions(UserIdReadActions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.ReadActions: %v\n", err)
			os.Exit(1)
		}

		v, err := service.ReadActions(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadActions: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(UserIdReadActions)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readaction":
		var err error
		ActionIdReadAction := *flagActionIdReadAction
		ActionNameReadAction := *flagActionNameReadAction
		request, err := clientHandler.ReadAction(ActionIdReadAction, ActionNameReadAction)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.ReadAction: %v\n", err)
			os.Exit(1)
		}

		v, err := service.ReadAction(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadAction: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(ActionIdReadAction, ActionNameReadAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "createaction":
		var err error
		UserIdCreateAction := *flagUserIdCreateAction
		ActionNameCreateAction := *flagActionNameCreateAction
		request, err := clientHandler.CreateAction(UserIdCreateAction, ActionNameCreateAction)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.CreateAction: %v\n", err)
			os.Exit(1)
		}

		v, err := service.CreateAction(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CreateAction: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(UserIdCreateAction, ActionNameCreateAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readoccurrences":
		var err error
		UserIdReadOccurrences := *flagUserIdReadOccurrences
		ActionIdReadOccurrences := *flagActionIdReadOccurrences
		request, err := clientHandler.ReadOccurrences(UserIdReadOccurrences, ActionIdReadOccurrences)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.ReadOccurrences: %v\n", err)
			os.Exit(1)
		}

		v, err := service.ReadOccurrences(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadOccurrences: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(UserIdReadOccurrences, ActionIdReadOccurrences)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "createoccurrence":
		var err error
		ActionIdCreateOccurrence := *flagActionIdCreateOccurrence
		DatetimeCreateOccurrence := *flagDatetimeCreateOccurrence
		DataCreateOccurrence := *flagDataCreateOccurrence
		request, err := clientHandler.CreateOccurrence(ActionIdCreateOccurrence, DatetimeCreateOccurrence, DataCreateOccurrence)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling clientHandler.CreateOccurrence: %v\n", err)
			os.Exit(1)
		}

		v, err := service.CreateOccurrence(context.Background(), request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CreateOccurrence: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Requested with:")
		fmt.Println(ActionIdCreateOccurrence, DatetimeCreateOccurrence, DataCreateOccurrence)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
