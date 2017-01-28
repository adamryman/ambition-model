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
		method   = flag.String("method", "createaction", "createaction,createoccurrence,readactions,readaction")
	)

	var (
		flagUserIDCreateOccurrence     = flag.Int64("createoccurrence.userid", 0, "")
		flagOccurrenceCreateOccurrence = flag.String("createoccurrence.occurrence", "", "")
		flagUserIDReadActions          = flag.Int64("readactions.userid", 0, "")
		flagIDReadAction               = flag.Int64("readaction.id", 0, "")
		flagUserIDReadAction           = flag.Int64("readaction.userid", 0, "")
		flagActionNameReadAction       = flag.String("readaction.actionname", "", "")
		flagIDCreateAction             = flag.Int64("createaction.id", 0, "")
		flagUserIDCreateAction         = flag.Int64("createaction.userid", 0, "")
		flagActionNameCreateAction     = flag.String("createaction.actionname", "", "")
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
		ActionNameCreateAction := *flagActionNameCreateAction
		request, err := clientHandler.CreateAction(IDCreateAction, UserIDCreateAction, ActionNameCreateAction)
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
		fmt.Println(IDCreateAction, UserIDCreateAction, ActionNameCreateAction)
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
	case "readaction":
		var err error
		IDReadAction := *flagIDReadAction
		UserIDReadAction := *flagUserIDReadAction
		ActionNameReadAction := *flagActionNameReadAction
		request, err := clientHandler.ReadAction(IDReadAction, UserIDReadAction, ActionNameReadAction)
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
		fmt.Println(IDReadAction, UserIDReadAction, ActionNameReadAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
