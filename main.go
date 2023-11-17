package main

import (
	"context"
	"log"
	"net"
	"syscall"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-contact-service/config"
	"github.com/go-contact-service/config/database"
	v1 "github.com/go-contact-service/delivery/http/v1"
	"github.com/go-contact-service/entity"
	"github.com/go-contact-service/lib/logger"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/go-contact-service/docs"

	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

// @title ERP Contact Service API Documentation.
// @version 1.0
// @description This is a sample api documentation.

// @host localhost:2326
// @BasePath /api/v1

func main() {
	conf := config.NewConfig("config.env")
	appLogger := logger.NewApiLogger(conf)

	appLogger.InitLogger()
	appLogger.Info("Starting the API Server")
	db := database.NewDB(conf)

	accessKey := conf.Aws.AwsAccessKeyId
	accessSecret := conf.Aws.AwsSecretAccessKey
	defaultRegion := conf.Aws.AwsDefaultRegion
	bucketName := conf.Aws.AwsStorageBucketName
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
		appLogger.Info("aws configuration missing")
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String(defaultRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	e := echo.New()
	//e.Logger.SetLevel(log.INFO)
	// Eanble HTTP compression
	e.Use(middleware.Gzip())

	// Recover from panics
	e.Use(middleware.Recover())

	// Allow requests from *
	e.Use(middleware.CORS())

	// Print http request and response log to stdout if debug is enabled
	if conf.Debug {
		e.Use(middleware.Logger())
	}

	// JWT Middleware
	jwtConfig := middleware.JWTConfig{
		Claims:       &entity.JwtClaim{},
		SigningKey:   []byte(conf.JwtSecret),
		ErrorHandler: v1.InvalidJwt,
	}

	v1.SetupRouters(e, conf, db, jwtConfig, appLogger, sess)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go httpServer(e, conf.HTTP)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	go grpcServer(s, conf.GRPC, db)

	//e.Logger.Fatal(e.Start(conf.HTTPAddress))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Println("Shutting down gRPC server...")

	s.GracefulStop()

	log.Println("gRPC server stopped!")

	log.Println("Shutting down HTTP server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("HTTP server stopped!")
}
func httpServer(e *echo.Echo, httpConfig config.HTTP) {
	if err := e.Start(httpConfig.HTTPAddress); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}

func grpcServer(g *grpc.Server, grpcConfig config.GRPC, db *bun.DB) {

	lis, err := net.Listen("tcp", grpcConfig.GrpcPort)
	if err != nil {
		log.Fatalf("GRPC: failed to listen: %v", err)
	}
	/*productRepository := repository.NewProductRepo(db)
	productService := product.NewService(productRepository, nil)

	branchRepository := repository.NewBranchRepo(db)
	branchService := branch.NewService(branchRepository, nil)
	handler := grpc_handler.Server{
		ProductService: productService,
		BranchService:  branchService,
	}
	proto.RegisterBranchServer(g, &handler)
	proto.RegisterProductServer(g, &handler)
	*/

	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := g.Serve(lis); err != nil {
		log.Fatalf("GRPC: failed to serve: %v", err)
	}
}
