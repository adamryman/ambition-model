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
	clientHandler "github.com/adamryman/ambition-model/ambition-service/generated/cli/handlers"
	grpcclient "github.com/adamryman/ambition-model/ambition-service/generated/client/grpc"
	httpclient "github.com/adamryman/ambition-model/ambition-service/generated/client/http"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterAmbitionServer
)

func main() {
	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5040", "gRPC (HTTP) address of addsvc")
		method   = flag.String("method", "createaction", "createaction,createoccurrence,readaction,readactions,readoccurrences")
	)

	var (
		flagIDCreateAction             = flag.Int64("createaction.id", 0, "")
		flagUserIDCreateAction         = flag.Int64("createaction.userid", 0, "")
		flagNameCreateAction           = flag.String("createaction.name", "", "")
		flagUserIDCreateOccurrence     = flag.Int64("createoccurrence.userid", 0, "")
		flagOccurrenceCreateOccurrence = flag.String("createoccurrence.occurrence", "", "")
		flagIDReadAction               = flag.Int64("readaction.id", 0, "")
		flagUserIDReadAction           = flag.Int64("readaction.userid", 0, "")
		flagNameReadAction             = flag.String("readaction.name", "", "")
		flagUserIDReadActions          = flag.Int64("readactions.userid", 0, "")
		flagIDReadOccurrences          = flag.Int64("readoccurrences.id", 0, "")
		flagUserIDReadOccurrences      = flag.Int64("readoccurrences.userid", 0, "")
		flagNameReadOccurrences        = flag.String("readoccurrences.name", "", "")
	)
	flag.Parse()

	var (
		service pb.AmbitionServer
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

	case "createaction":
		var err error
		IDCreateAction := *flagIDCreateAction
		UserIDCreateAction := *flagUserIDCreateAction
		NameCreateAction := *flagNameCreateAction
		request, err := clientHandler.CreateAction(IDCreateAction, UserIDCreateAction, NameCreateAction)
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
		fmt.Println(IDCreateAction, UserIDCreateAction, NameCreateAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "createoccurrence":
		var err error
		UserIDCreateOccurrence := *flagUserIDCreateOccurrence

		var OccurrenceCreateOccurrence pb.Occurrence
		if flagOccurrenceCreateOccurrence != nil && len(*flagOccurrenceCreateOccurrence) > 0 {
			err = json.Unmarshal([]byte(*flagOccurrenceCreateOccurrence), &OccurrenceCreateOccurrence)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling OccurrenceCreateOccurrence from %v:", flagOccurrenceCreateOccurrence))
			}
		}

		request, err := clientHandler.CreateOccurrence(UserIDCreateOccurrence, OccurrenceCreateOccurrence)
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
		fmt.Println(UserIDCreateOccurrence, OccurrenceCreateOccurrence)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readaction":
		var err error
		IDReadAction := *flagIDReadAction
		UserIDReadAction := *flagUserIDReadAction
		NameReadAction := *flagNameReadAction
		request, err := clientHandler.ReadAction(IDReadAction, UserIDReadAction, NameReadAction)
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
		fmt.Println(IDReadAction, UserIDReadAction, NameReadAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readactions":
		var err error
		UserIDReadActions := *flagUserIDReadActions
		request, err := clientHandler.ReadActions(UserIDReadActions)
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
		fmt.Println(UserIDReadActions)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	case "readoccurrences":
		var err error
		IDReadOccurrences := *flagIDReadOccurrences
		UserIDReadOccurrences := *flagUserIDReadOccurrences
		NameReadOccurrences := *flagNameReadOccurrences
		request, err := clientHandler.ReadOccurrences(IDReadOccurrences, UserIDReadOccurrences, NameReadOccurrences)
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
		fmt.Println(IDReadOccurrences, UserIDReadOccurrences, NameReadOccurrences)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
