package test

import (
	"jwtTest/jwt"
	"testing"
)

func TestVerify(t *testing.T) {
	type args struct {
		token  string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test01", args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.skaja.byPBZwtv6AyHv2vIqr5yhEVta9DWvfdqZkBQBWtoPJI"}, true},
		{"test02", args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiY2hpaHVvIn0.byPBZwtv6AyHv2vIqr5yhEVta9DWvfdqZkBQBWtoPJI"}, false},
		{"test03", args{token: "sakf.skaja"}, true},
		{"test04", args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.IjEyMyI.FRmXiIk6k__trahUSlyFTYgqA6SvPvEi-4JlXpXRGrc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := jwt.Verify(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSign(t *testing.T) {
	type args struct {
		secret    string
		jwtHeader jwt.Header
		payload   jwt.Payload
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test01", args: args{jwtHeader: jwt.Header{Alg: "HS256", Typ: "JWT"}, payload: jwt.Payload{Issuer: "youngalone", IssuedAt: "1683461163", Expiration: "1683469160", Audience: "用户"}}, wantRes: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3N1ZXIiOiJ5b3VuZ2Fsb25lIiwiSXNzdWVkQXQiOiIxNjgzNDYxMTYzIiwiRXhwaXJhdGlvbiI6IjE2ODM0NjkxNjAiLCJBdWRpZW5jZSI6IueUqOaItyJ9.Didmsh6kUJy9fPG7xxYUcUT0HzoRjT40c6OtN2EZNZw", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := jwt.Sign(tt.args.jwtHeader, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Sign() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
