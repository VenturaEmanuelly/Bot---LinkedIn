package main

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {

	if err := playwright.Install(); err != nil {
		log.Fatalf("erro ao instalar browsers: %v", err)
	}

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("erro ao iniciar playWright : %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})

	if err != nil {
		log.Fatalf("erro ao abrir o navegador : %v", err)
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		StorageStatePath: playwright.String("auth.json"),
	})

	if err != nil {
		log.Fatalf("erro ao criar contexto : %v", err)
	}

	page, err := context.NewPage()

	if err != nil {
		log.Fatalf("erro ao criar página: %v", err)
	}

	_, err = page.Goto("https://www.linkedin.com/jobs")

	if err != nil {
		log.Fatalf("erro ao abrir site : %v", err)
	}

	log.Println("LinkedIn Jobs aberto. Faça login (se necessário).")

	select {}
}
