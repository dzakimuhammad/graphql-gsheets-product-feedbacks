# GraphQL - Google Sheets for Product Feedbacks

API for storing product feedbacks in Google Sheets using GraphQL

## Requirement
1. Go
    
    https://go.dev/doc/install
    
2. Google Cloud Project
  
    https://developers.google.com/workspace/guides/create-project

3. GCP Service Account
    
    https://developers.google.com/workspace/guides/create-credentials#service-account

4. Google Sheets

    Create a spreadsheet to store the feedbacks. Make sure to add the GCP Service Account email as an editor to the spreadsheet so it can write on the spreadsheet.

## How to Use
1. Change the `credentials.json` file to suits your GCP service account credentials
2. Change the `spreadsheetId` in line 36 of `/graph/feedback.resolvers.go` to your spreadsheet ID
3. Run `go run main.go` to start the GraphQL server
4. Hit the endpoint `http://localhost:8080/query` with the following GraphQL request body format :
    ```
    mutation {
      feedbackPost(input: {
        productId: "product-001",
        rating: 4,
        feedbacks: [
          "Good Product",
          "Nice color"
        ]
      }) {
        feedbackId
        productId
        rating
        feedbacks
        createdAt
      }
    }
    ```
5. Spreadsheet should be appended with the new feedbacks

## Framework & Library
1. gqlgen

    https://gqlgen.com/

2. sheets package

    https://pkg.go.dev/google.golang.org/api/sheets/v4

## Author
- Dzaki Muhammad - https://github.com/dzakimuhammad