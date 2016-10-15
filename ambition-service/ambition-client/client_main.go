package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	//"github.com/lightstep/lightstep-tracer-go"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//zipkin "github.com/openzipkin/zipkin-go-opentracing"
	//appdashot "github.com/sourcegraph/appdash/opentracing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"sourcegraph.com/sourcegraph/appdash"

	"github.com/pkg/errors"

	//"github.com/go-kit/kit/log"

	// This Service
	pb "github.com/adamryman/ambition-truss/ambition-service"
	grpcclient "github.com/adamryman/ambition-truss/ambition-service/generated/client/grpc"
	httpclient "github.com/adamryman/ambition-truss/ambition-service/generated/client/http"
	clientHandler "github.com/adamryman/ambition-truss/ambition-service/handlers/client"
	handler "github.com/adamryman/ambition-truss/ambition-service/handlers/server"
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
	// provide the direct address of an addsvc. This presumption is reflected in
	// the addcli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings. For an
	// example service with a client built on top of a service discovery system,
	// see profilesvc.

	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", "", "gRPC (HTTP) address of addsvc")
		//zipkinAddr     = flag.String("zipkin.addr", "", "Enable Zipkin tracing via a Kafka Collector host:port")
		//appdashAddr    = flag.String("appdash.addr", "", "Enable Appdash tracing via an Appdash server host:port")
		//lightstepToken = flag.String("lightstep.token", "", "Enable LightStep tracing via a LightStep access token")
		method = flag.String("method", "readactions", "readactions,createaction,readoccurrences,createoccurrence")
	)

	var (
		flagUserIdReadActions        = flag.Int64("readactions.userid", 0, "")
		flagUserIdCreateAction       = flag.Int64("createaction.userid", 0, "")
		flagActionNameCreateAction   = flag.String("createaction.actionname", "", "")
		flagUserIdReadOccurrences    = flag.Int64("readoccurrences.userid", 0, "")
		flagActionIdReadOccurrences  = flag.Int64("readoccurrences.actionid", 0, "")
		flagActionIdCreateOccurrence = flag.Int64("createoccurrence.actionid", 0, "")
		flagEpocTimeCreateOccurrence = flag.Int64("createoccurrence.epoctime", 0, "")
	)
	flag.Parse()

	// This is a demonstration client, which supports multiple tracers.
	// Your clients will probably just use one tracer.
	//var tracer stdopentracing.Tracer
	//{
	//if *zipkinAddr != "" {
	//collector, err := zipkin.NewKafkaCollector(
	//strings.Split(*zipkinAddr, ","),
	//zipkin.KafkaLogger(log.NewNopLogger()),
	//)
	//if err != nil {
	//fmt.Fprintf(os.Stderr, "%v\n", err)
	//os.Exit(1)
	//}
	//tracer, err = zipkin.NewTracer(
	//zipkin.NewRecorder(collector, false, "localhost:8000", "addcli"),
	//)
	//if err != nil {
	//fmt.Fprintf(os.Stderr, "%v\n", err)
	//os.Exit(1)
	//}
	//} else if *appdashAddr != "" {
	//tracer = appdashot.NewTracer(appdash.NewRemoteCollector(*appdashAddr))
	//} else if *lightstepToken != "" {
	//tracer = lightstep.NewTracer(lightstep.Options{
	//AccessToken: *lightstepToken,
	//})
	//defer lightstep.FlushLightStepTracer(tracer)
	//} else {
	//tracer = stdopentracing.GlobalTracer() // no-op
	//}
	//}

	// This is a demonstration client, which supports multiple transports.
	// Your clients will probably just define and stick with 1 transport.

	var (
		service handler.Service
		err     error
	)
	if *httpAddr != "" {
		//service, err = httpclient.New(*httpAddr, tracer, log.NewNopLogger())
		service, err = httpclient.New(*httpAddr)
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			os.Exit(1)
		}
		defer conn.Close()
		service = grpcclient.New(conn /*, tracer, log.NewNopLogger()*/)
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
		EpocTimeCreateOccurrence := *flagEpocTimeCreateOccurrence
		request, err := clientHandler.CreateOccurrence(ActionIdCreateOccurrence, EpocTimeCreateOccurrence)
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
		fmt.Println(ActionIdCreateOccurrence, EpocTimeCreateOccurrence)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", method)
		os.Exit(1)
	}
}
