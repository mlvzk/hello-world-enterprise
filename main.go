package main

import (
	"log"
	"os"
	"strings"
)

func printText(text string) {
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.Println("printText was called with argument: ", text)
	}
	println(text)
}

type HelloWorldService struct {
	translationService ITranslationService
}

type IHelloWorldService interface {
	getHelloText() string
}

var _ IHelloWorldService = (*HelloWorldService)(nil)

func (h HelloWorldService) getHelloText() string {
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.Println("getHelloText was called.")
	}

	if h.translationService.getHelloText() == "" {
		panic(nil)
	}
	return h.translationService.getHelloText()
}

func (h *HelloWorldService) createHelloWorldService(translationService ITranslationService) HelloWorldService {
	return HelloWorldService{translationService}
}

type Command interface {
	Do()
}

type HelloWorldCommand struct {
	helloService IHelloWorldService
}

func (h HelloWorldCommand) Do() {
	printText(h.helloService.getHelloText())
}

func (h *HelloWorldCommand) createHelloWorldCommand(helloService IHelloWorldService) HelloWorldCommand {
	return HelloWorldCommand{helloService}
}

type TranslationService struct {
	helloText string
}

type ITranslationService interface {
	getHelloText() string
}

func (ts *TranslationService) createTranslationService(language string) TranslationService {
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.Println("createTranslationService was called with argument: ", language)
	}

	if _, ok := languageToHelloWorldText[language]; !ok {
		main()
	}
	return TranslationService{
		helloText: languageToHelloWorldText[language],
	}
}

func (ts TranslationService) getHelloText() string { return ts.helloText }

var commands map[string]Command

func init() {
	args := os.Args
	var language string
	for i, arg := range args[1:] {
		if arg == "--language" {
			language = args[1:][i+1]
		}
	}
	if locale := os.Getenv("LANG"); language == "" {
		language = strings.Split(locale, "_")[0]
		if _, ok := languageToHelloWorldText[language]; !ok {
			main()
		}
	}
	translationService := (*TranslationService)(nil).createTranslationService(language)
	helloService := (*HelloWorldService)(nil).createHelloWorldService(translationService)
	commands = map[string]Command{
		"helloWorld": (*HelloWorldCommand)(nil).createHelloWorldCommand(helloService),
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("You probably passed a language that is not supported. Here is a list of supported languages:\n")
			for language := range languageToHelloWorldText {
				println(language)
			}
			println("\nIf you want your language supported, fork the project and add it yourself.")
			os.Exit(1)
		}
	}()
	command := commands["helloWorld"]
	command.Do()
}

var languageToHelloWorldText = map[string]string{
	"af": "Hallo wêreld",
	"ar": "مرحبا بالعالم",
	"bg": "Здравей, свят",
	"bn": "হ্যালো ওয়ার্ল্ড",
	"ca": "Hola món",
	"cs": "Ahoj světe",
	"da": "Hej verden",
	"hr": "Zdravo svijete",
	"nl": "Hallo mensen",
	"de": "Hallo Welt",
	"el": "Γεια σου κόσμε",
	"en": "Hello world",
	"et": "Tere maailm",
	"fi": "Moi maailma",
	"fj": "Vuravura ni bula",
	"fr": "Salut tout le monde",
	"gu": "કેમ છો દુનિયા",
	"hi": "हैलो वर्ल्ड",
	"hu": "helló világ",
	"ga": "Haigh ansin",
	"is": "Halló heimur",
	"it": "Salve, mondo",
	"ja": "ハローワールド",
	"kk": "Сәлемет пе әлем",
	"kn": "ಹಲೋ ಪ್ರಪಂಚ",
	"ko": "전 세계 여러분 안녕하세요",
	"lt": "sveikas, pasauli",
	"lv": "sveika, pasaule",
	"mg": "Manahoana tontolo",
	"fa": "سلام به دنیا",
	"mi": "Kia ora e te ao whānui.",
	"ml": "ഹലോ വേള്‍ഡ്",
	"mr": "विश्वाला नमस्कार",
	"ms": "Helo dunia",
	"mt": "Bonġu dinja",
	"nb": "hallo verden",
	"pa": "ਸਤਿ ਸ਼੍ਰੀ ਅਕਾਲ ਦੁਨੀਆਂ",
	"pl": "Cześć ludzie",
	"pt": "Olá, mundo",
	"es": "Hola mundo",
	"ro": "bună oameni buni",
	"ru": "Всем привет",
	"sk": "čaute všetci",
	"sl": "pozdravljeni vsi skupaj",
	"sm": "Talofa le lalolagi",
	"sv": "Hej världen",
	"sw": "Vipi dunia",
	"ta": "ஹலோ உலகம்",
}
