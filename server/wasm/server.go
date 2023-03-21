package wasm

import (
	"fmt"
	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/encoding"

	"github.com/streamingfast/dgrpc/server"
	"github.com/streamingfast/dgrpc/server/standard"
	sserver "github.com/streamingfast/substreams-sink-kv/server"
	"go.uber.org/zap"
)

var _ sserver.Serveable = (*Server)(nil)

func NewServer(config *Config, wasmEngine *Engine, protoCodec Codec, logger *zap.Logger) (*Server, error) {
	encoding.RegisterCodec(protoCodec)
	srv := standard.NewServer(server.NewOptions())

	grpcServer := srv.GrpcServer()

	var v interface{}
	grpcService := &grpc.ServiceDesc{
		ServiceName: config.FQGRPCServiceName,
		HandlerType: wasmEngine,
		Methods:     []grpc.MethodDesc{},
	}

	for _, methodConfig := range config.Methods {
		handler, err := wasmEngine.GetHandler(methodConfig, protoCodec, logger)
		if err != nil {
			return nil, fmt.Errorf("failed to get handler: %w", err)
		}

		grpcService.Streams = append(grpcService.Streams, grpc.StreamDesc{
			StreamName:    methodConfig.Name,
			Handler:       handler.handle,
			ServerStreams: true,
		})
	}

	grpcServer.RegisterService(grpcService, v)

	s := &Server{
		logger: logger,
		srv:    srv,
	}

	return s, nil
}

type Server struct {
	logger *zap.Logger
	srv    *standard.StandardServer
}

func (s *Server) Shutdown() {
	s.logger.Info("wasn server received shutdown, shutting down server")
	s.srv.Shutdown(5 * time.Second)
}

func (s *Server) Serve(listenAddr string) error {
	s.srv.Launch(listenAddr)
	return nil
}
