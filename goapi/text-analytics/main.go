package main

import (
	"context"
	"fmt"
	"log"
	"os"

	textanalytics "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/textanalytics"
	autorest "github.com/Azure/go-autorest/autorest"
)

// StringPointer returns a pointer to the string value passed in.
func StringPointer(v string) *string {
	return &v
}

// BoolPointer returns a pointer to the bool value passed in.
func BoolPointer(v bool) *bool {
	return &v
}

// ExtractEntities ...
func ExtractEntities(textAnalyticsclient textanalytics.BaseClient) {

	ctx := context.Background()
	inputDocuments := []textanalytics.MultiLanguageInput{
		textanalytics.MultiLanguageInput{
			Language: StringPointer("en"),
			ID:       StringPointer("0"),
			Text:     StringPointer("Microsoft was founded by Bill Gates and Paul Allen on April 4, 1975, to develop and sell BASIC interpreters for the Altair 8800."),
		},
	}

	batchInput := textanalytics.MultiLanguageBatchInput{Documents: &inputDocuments}

	result, err := textAnalyticsclient.Entities(ctx, BoolPointer(false), &batchInput)
	if err != nil {
		log.Println(err)
	}

	// Printing extracted entities results
	for _, document := range *result.Documents {
		fmt.Printf("Document ID: %s\n", *document.ID)
		fmt.Printf("\tExtracted Entities:\n")
		for _, entity := range *document.Entities {
			fmt.Printf("\t\tName: %s\tType: %s", *entity.Name, *entity.Type)
			if entity.SubType != nil {
				fmt.Printf("\tSub-Type: %s\n", *entity.SubType)
			}
			fmt.Println()
			for _, match := range *entity.Matches {
				fmt.Printf("\t\t\tOffset: %v\tLength: %v\tScore: %f\n", *match.Offset, *match.Length, *match.EntityTypeScore)
			}
		}
		fmt.Println()
	}

	// Printing document errors
	fmt.Println("Document Errors")
	for _, error := range *result.Errors {
		fmt.Printf("Document ID: %s Message : %s\n", *error.ID, *error.Message)
	}
}

// ExtrctEnttsStrt ...
func ExtrctEnttsStrt(text, subscriptionKey, endpoint string) {

	textAnalyticsClient := textanalytics.New(endpoint)
	textAnalyticsClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(subscriptionKey)

	// log.Println(textAnalyticsClient)
	// log.Println(textAnalyticsClient.Authorizer)

	ExtractEntities(textAnalyticsClient)
}

func main() {
	// This sample assumes you have created an environment variable for your key
	subscriptionKey := os.Getenv("TEXT_ANALYTICS_SUBSCRIPTION_KEY")
	// replace this endpoint with the correct one for your Azure resource.
	endpoint := "https://test-text-anbima.cognitiveservices.azure.com"

	ExtrctEnttsStrt("start", subscriptionKey, endpoint)
}
