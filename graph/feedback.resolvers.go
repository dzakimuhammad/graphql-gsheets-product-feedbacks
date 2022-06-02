package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/dzakimuhammad/gql-sheets-exploration/graph/model"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

func (r *mutationResolver) FeedbackPost(ctx context.Context, input model.FeedbackPostRequest) (*model.FeedbackResponse, error) {
	newFeedbackData := &model.FeedbackResponse{
		FeedbackID: uuid.NewString(),
		ProductID:  input.ProductID,
		Rating:     input.Rating,
		Feedbacks:  input.Feedbacks,
		CreatedAt:  time.Now().UTC(),
	}

	const (
		client_secret_path = "./credentials.json"
	)

	// Create a service object for Google sheets
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(client_secret_path), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		return nil, err
	}

	// Change the Spreadsheet Id with yours
	spreadsheetId := ""

	// Define the Sheet Name and fields to select
	writeRange := "A:E"
	values := []interface{}{
		newFeedbackData.FeedbackID,
		newFeedbackData.ProductID,
		newFeedbackData.Rating,
		newFeedbackData.Feedbacks,
		newFeedbackData.CreatedAt,
	}
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, values)
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()

	if err != nil {
		return nil, err
	}

	return newFeedbackData, nil
}
