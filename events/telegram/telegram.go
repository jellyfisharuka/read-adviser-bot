package telegram

import "read-adviser-bot/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	//storage

}
fun New(client *telegram.Client) {
	
}