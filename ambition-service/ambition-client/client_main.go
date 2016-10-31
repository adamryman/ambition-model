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
		flagUserIdReadActions            = flag.Int64("readactions.userid", 0, "")
		flagActionIdReadAction           = flag.Int64("readaction.actionid", 0, "")
		flagUserIdReadAction             = flag.Int64("readaction.userid", 0, "")
		flagActionNameReadAction         = flag.String("readaction.actionname", "", "")
		flagTrelloIdReadAction           = flag.String("readaction.trelloid", "", "")
		flagActionIdCreateAction         = flag.Int64("createaction.actionid", 0, "")
		flagUserIdCreateAction           = flag.Int64("createaction.userid", 0, "")
		flagActionNameCreateAction       = flag.String("createaction.actionname", "", "")
		flagTrelloIdCreateAction         = flag.String("createaction.trelloid", "", "")
		flagOccurrenceIdReadOccurrences  = flag.Int64("readoccurrences.occurrenceid", 0, "")
		flagActionIdReadOccurrences      = flag.Int64("readoccurrences.actionid", 0, "")
		flagDatetimeReadOccurrences      = flag.String("readoccurrences.datetime", "", "")
		flagDataReadOccurrences          = flag.String("readoccurrences.data", "", "")
		flagOccurrenceIdCreateOccurrence = flag.Int64("createoccurrence.occurrenceid", 0, "")
		flagActionIdCreateOccurrence     = flag.Int64("createoccurrence.actionid", 0, "")
		flagDatetimeCreateOccurrence     = flag.String("createoccurrence.datetime", "", "")
		flagDataCreateOccurrence         = flag.String("createoccurrence.data", "", "")
	)
	flag.Parse()

	var (
		service pb.AmbitionServiceServer
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
		service, err = grpcclient.New(conn)
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
		UserIdReadAction := *flagUserIdReadAction
		ActionNameReadAction := *flagActionNameReadAction
		TrelloIdReadAction := *flagTrelloIdReadAction
		request, err := clientHandler.ReadAction(ActionIdReadAction, UserIdReadAction, ActionNameReadAction, TrelloIdReadAction)
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
		fmt.Println(ActionIdReadAction, UserIdReadAction, ActionNameReadAction, TrelloIdReadAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "createaction":
		var err error
		ActionIdCreateAction := *flagActionIdCreateAction
		UserIdCreateAction := *flagUserIdCreateAction
		ActionNameCreateAction := *flagActionNameCreateAction
		TrelloIdCreateAction := *flagTrelloIdCreateAction
		request, err := clientHandler.CreateAction(ActionIdCreateAction, UserIdCreateAction, ActionNameCreateAction, TrelloIdCreateAction)
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
		fmt.Println(ActionIdCreateAction, UserIdCreateAction, ActionNameCreateAction, TrelloIdCreateAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readoccurrences":
		var err error
		OccurrenceIdReadOccurrences := *flagOccurrenceIdReadOccurrences
		ActionIdReadOccurrences := *flagActionIdReadOccurrences
		DatetimeReadOccurrences := *flagDatetimeReadOccurrences
		DataReadOccurrences := *flagDataReadOccurrences
		request, err := clientHandler.ReadOccurrences(OccurrenceIdReadOccurrences, ActionIdReadOccurrences, DatetimeReadOccurrences, DataReadOccurrences)
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
		fmt.Println(OccurrenceIdReadOccurrences, ActionIdReadOccurrences, DatetimeReadOccurrences, DataReadOccurrences)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "createoccurrence":
		var err error
		OccurrenceIdCreateOccurrence := *flagOccurrenceIdCreateOccurrence
		ActionIdCreateOccurrence := *flagActionIdCreateOccurrence
		DatetimeCreateOccurrence := *flagDatetimeCreateOccurrence
		DataCreateOccurrence := *flagDataCreateOccurrence
		request, err := clientHandler.CreateOccurrence(OccurrenceIdCreateOccurrence, ActionIdCreateOccurrence, DatetimeCreateOccurrence, DataCreateOccurrence)
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
		fmt.Println(OccurrenceIdCreateOccurrence, ActionIdCreateOccurrence, DatetimeCreateOccurrence, DataCreateOccurrence)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
