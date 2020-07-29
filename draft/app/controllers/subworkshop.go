package controllers

import (
	"context"
	"fmt"
	"github.com/go-masonry/bricks/log"
	workshop "github.com/go-masonry/scaffolds/draft/api"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type SubWorkshopController interface {
	workshop.SubWorkshopServer
}

type subWorkshopControllerDeps struct {
	fx.In

	Logger log.Logger
}

type subWorkshopController struct {
	deps subWorkshopControllerDeps
}

func CreateSubWorkshopController(deps subWorkshopControllerDeps) SubWorkshopController {
	return &subWorkshopController{
		deps: deps,
	}
}

func (s *subWorkshopController) PaintCar(ctx context.Context, request *workshop.SubPaintCarRequest) (*empty.Empty, error) {
	// Paint car
	if err := s.doActualPaint(ctx, request.GetCar()); err != nil {
		return nil, err
	}
	// Dial back to caller
	conn, err := grpc.DialContext(ctx, request.GetCallbackServiceAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("car painted but we can't callback to %s, %w", request.GetCallbackServiceAddress(), err)
	}
	// Make client and call method
	workshopClient := workshop.NewWorkshopClient(conn)
	return workshopClient.CarPainted(ctx, &workshop.PaintFinishedRequest{CarId: request.GetCar().GetId(), DesiredColor: request.GetDesiredColor()})
}

func (s *subWorkshopController) doActualPaint(ctx context.Context, car *workshop.Car) error {
	// here be paint logic...
	// ...
	// ...
	return nil
}