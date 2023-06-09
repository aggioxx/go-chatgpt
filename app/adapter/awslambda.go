package adapter

import "context"

type LambdaInterface interface {
	Handler(ctx context.Context, req LambdaRequest) (LambdaResponse, error)
}

type LambdaRequest struct {
	Input string `json:"input"`
}

type LambdaResponse struct {
	Text string `json:"text"`
}

func Handler(ctx context.Context, req LambdaRequest) (LambdaResponse, error) {
	response, err := SendChat(req.Input)
	if err != nil {
		return LambdaResponse{}, err
	}

	return LambdaResponse{
		Text: response,
	}, nil
}
