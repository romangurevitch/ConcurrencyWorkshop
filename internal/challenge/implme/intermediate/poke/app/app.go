package app

import (
	"log"
	"log/slog"
	"net/http"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/intermediate/poke/client"
)

const (
	windowWidth      = 400
	windowHeight     = 600
	defaultURL       = "https://golangify.com/wp-content/uploads/2020/04/go-read.png"
	notFoundImageURL = "https://miro.medium.com/v2/resize:fit:460/1*1Yf_9BPNftL1gdCTMr9Exw.png"
)

type PokeAPP interface {
	Start()
}

type pokeAPP struct {
	pokeClient client.PokeClient

	header *widget.Label
	img    *canvas.Image
	input  *widget.Entry
}

func NewPokeApp(pokeClient client.PokeClient) PokeAPP {
	return &pokeAPP{
		pokeClient: pokeClient,
		header:     createHeader(),
		img:        imageFromURL(defaultURL),
		input:      widget.NewEntry(),
	}
}

func (p *pokeAPP) Start() {
	myApp := app.New()
	myWindow := myApp.NewWindow("PokeGUI")
	myWindow.Resize(fyne.NewSize(windowWidth, windowHeight))

	p.input.SetPlaceHolder("e.g., pikachu or 25")
	// TODO: impl and use OnChangedNonBlocking to improve responsiveness of the app
	p.input.OnChanged = p.OnChanged

	content := container.NewStack(container.NewVBox(p.header, p.input), p.img)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func (p *pokeAPP) OnChanged(ID string) {
	_, err := p.fetchAndUpdatePokemon(ID)
	if err != nil {
		slog.Error("fetchAndUpdatePokemon", "error", err)
	}
}

// OnChangedNonBlocking triggers a non-blocking operation to fetch and update
// the Pokémon information based on the provided ID. This method improves the
// responsiveness of the application by not blocking the UI while performing
// the fetch operation. When the ID is entered, this function should be called,
// and it will handle the update asynchronously.
//
// The function is currently not implemented and will panic if used.
// TODO: Implement OnChangedNonBlocking to fetch and update Pokémon details asynchronously.
func (p *pokeAPP) OnChangedNonBlocking(ID string) {
	panic("implement me!")
}

func (p *pokeAPP) fetchAndUpdatePokemon(ID string) (bool, error) {
	if ID == "" {
		p.setImage(defaultURL)
		return true, nil
	}

	poke, err := p.pokeClient.FetchPokemon(ID)
	if err != nil {
		p.setName("Not Found")
		p.setImage(notFoundImageURL)
		return false, nil
	}

	p.setName(poke.Name)
	p.setImage(poke.Sprites.FrontDefault)
	return true, nil
}

func (p *pokeAPP) setName(header string) {
	p.header.SetText("Pokémon: " + header)
}

func (p *pokeAPP) setImage(url string) {
	if !isValidURL(url) {
		url = defaultURL
	}

	p.img.Resource = imageFromURL(url).Resource
	p.img.Refresh()
}

func createHeader() *widget.Label {
	header := widget.NewLabel("Enter Pokémon ID or Name")
	header.TextStyle = fyne.TextStyle{Bold: true}
	return header
}

func imageFromURL(url string) *canvas.Image {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			slog.Error("Close", "error", err)
		}
	}()

	img := canvas.NewImageFromReader(response.Body, "Pokemon")
	img.FillMode = canvas.ImageFillOriginal

	return img
}

func isValidURL(urlStr string) bool {
	// Parse the URL and ensure there were no errors.
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Check if the URL scheme and host are non-empty to consider this a valid URL.
	// You might want to extend this with more checks depending on your use case.
	return u.Scheme != "" && u.Host != ""
}
