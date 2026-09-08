package mllp

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {

	got := Encode("MSH|^~\\&|APP")
	want := "\x0bMSH|^~\\&|APP\x1c\x0d"

	if string(got) != want {
		t.Fatalf("Encode()=%q, want %q", got, want)
	}
}

func TestReadMessage(t *testing.T) {

	var tests = []struct {
		description string
		stream      string
		want        []string
		wantFramed  bool
		wantErr     string
	}{
		{
			description: "Single Framed Message",
			stream:      "\x0bMSH|one\x1c\x0d",
			want:        []string{"MSH|one"},
			wantFramed:  true,
		},
		{
			description: "Several Messages On One Connection",
			stream:      "\x0bMSH|one\x1c\x0d\x0bMSH|two\x1c\x0d\x0bMSH|three\x1c\x0d",
			want:        []string{"MSH|one", "MSH|two", "MSH|three"},
			wantFramed:  true,
		},
		{
			description: "Carriage Returns Inside The Message Survive",
			stream:      "\x0bMSH|one\rEVN|A01\rPID|1\x1c\x0d",
			want:        []string{"MSH|one\rEVN|A01\rPID|1"},
			wantFramed:  true,
		},
		{
			description: "Unframed Message Is Still Returned",
			stream:      "MSH|one\rEVN|A01",
			want:        []string{"MSH|one\rEVN|A01"},
			wantFramed:  false,
		},
		{
			description: "Missing End Block",
			stream:      "\x0bMSH|one",
			wantErr:     "unterminated MLLP block",
		},
		{
			description: "Wrong Trailer After End Block",
			stream:      "\x0bMSH|one\x1cX",
			wantErr:     "malformed MLLP block",
		},
		{
			description: "End Block With No Trailer",
			stream:      "\x0bMSH|one\x1c",
			wantErr:     "missing its carriage return",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			reader := NewReader(strings.NewReader(tc.stream))

			if tc.wantErr != "" {
				_, _, err := reader.ReadMessage()
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("ReadMessage() error=%v, want error containing %q", err, tc.wantErr)
				}
				return
			}

			for i, want := range tc.want {
				got, framed, err := reader.ReadMessage()
				if err != nil {
					t.Fatalf("ReadMessage() #%d error=%v, want nil", i, err)
				}

				if got != want {
					t.Errorf("ReadMessage() #%d=%q, want %q", i, got, want)
				}

				if framed != tc.wantFramed {
					t.Errorf("ReadMessage() #%d framed=%v, want %v", i, framed, tc.wantFramed)
				}
			}

			_, _, err := reader.ReadMessage()
			if !errors.Is(err, io.EOF) {
				t.Errorf("ReadMessage() after the last message=%v, want io.EOF", err)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {

	want := "MSH|^~\\&|PLACEBO\rEVN|A01\rPID|1"

	var wire bytes.Buffer
	wire.Write(Encode(want))

	got, framed, err := NewReader(&wire).ReadMessage()
	if err != nil {
		t.Fatalf("ReadMessage() error=%v", err)
	}

	if !framed {
		t.Error("ReadMessage() framed=false, want true")
	}

	if got != want {
		t.Fatalf("round trip=%q, want %q", got, want)
	}
}
