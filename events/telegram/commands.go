package telegram

import (
	"errors"
	"log"
	"net/url"
	e "read-adviser-bot/clients/telegram/lib"
	"read-adviser-bot/events/telegram/storage"

	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string,keyword string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command '%s' from '%s'", text, username)
	//add page: http////
	//rnd page: /rnd
	//help: /help
	//start: /start  hi+help
	if isAddCmd(text) {
		//TODO: AddPage()
		return p.savePage(chatID, text, username,keyword)
	}
	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)

	}
}
func (p *Processor) savePage(chatID int, pageURL string, username string, keyword string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
		KeyWord: keyword,
	}
	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}
	if err := p.storage.Save(page); err != nil {
		return err
	}
	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}
func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't fo command : can't send random", err) }()
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}
	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}
	return p.storage.Remove(page)
}
func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}
func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
func isAddCmd(text string) bool {
	return isURL(text)
}
func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
