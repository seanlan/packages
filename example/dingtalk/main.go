package main

import "github.com/seanlan/packages/dingtalk"

func main() {
	token := "f1385dbba7dd77d2e1b2daa9de047050383252b6eb44be5a28a113a42d663a8f"
	secret := "SEC67f0b6b05ecb5bc9bd170238806a69f65a19e9bacc0dd63462769a83abd5b508"
	client := dingtalk.InitDingTalk(token, secret)
	client.SendMessage("123123")
}
