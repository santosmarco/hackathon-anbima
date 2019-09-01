package main

import "testing"

func TestExtrctEnttsStrt(t *testing.T) {
	type args struct {
		text            string
		subscriptionKey string
		endpoint        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"start",
			args{
				text:            "oi",
				subscriptionKey: "xxxxxx",
				endpoint:        "https://test-text-anbima.cognitiveservices.azure.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExtrctEnttsStrt(tt.args.text, tt.args.subscriptionKey, tt.args.endpoint)
		})
	}
}
