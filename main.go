package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/szerookii/goquobuz/qobuz"
	"github.com/szerookii/goquobuz/qobuz/types"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var p *tea.Program

type progressWriter struct {
	total      int
	downloaded int
	file       *os.File
	reader     io.Reader
	onProgress func(float64)
}

func (pw *progressWriter) Start() {
	_, err := io.Copy(pw.file, io.TeeReader(pw.reader, pw))
	if err != nil {
		p.Send(errors.New("failed to download track"))
	}
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)
	if pw.total > 0 && pw.onProgress != nil {
		pw.onProgress(float64(pw.downloaded) / float64(pw.total))
	}
	return len(p), nil
}

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url) // nolint:gosec
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("receiving status of %d for url: %s", resp.StatusCode, url)
	}
	return resp, nil
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

const (
	padding  = 2
	maxWidth = 80
)

type progressMsg float64

type progressErrMsg struct{ err error }

func finalPause() tea.Cmd {
	return tea.Tick(time.Millisecond*750, func(_ time.Time) tea.Msg {
		return nil
	})
}

type model struct {
	pw       *progressWriter
	progress progress.Model
	err      error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case progressErrMsg:
		m.err = msg.err
		return m, tea.Quit

	case progressMsg:
		var cmds []tea.Cmd

		if msg >= 1.0 {
			cmds = append(cmds, tea.Sequence(finalPause(), tea.Quit))
		}

		cmds = append(cmds, m.progress.SetPercent(float64(msg)))
		return m, tea.Batch(cmds...)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	if m.err != nil {
		return "Error downloading: " + m.err.Error() + "\n"
	}

	pad := strings.Repeat(" ", padding)
	renderedProgress := m.progress.View()
	renderedProgress = strings.ReplaceAll(renderedProgress, "░", "—")
	renderedProgress = strings.ReplaceAll(renderedProgress, "█", "—")

	return "\n" +
		pad + renderedProgress + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

type Config struct {
	Email          string
	Password       string
	DownloadFolder string
}

func readConfig() (*Config, error) {
	var config *Config

	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		fmt.Println("Config file does not exist, creating a new one...")

		config = new(Config)
		config.DownloadFolder = "downloads"

		if err := writeConfig("config.json", config); err != nil {
			return nil, fmt.Errorf("failed to create empty config file: %v", err)
		}

		return config, nil
	}

	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func writeConfig(configPath string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New("failed to marshal config to JSON")
	}

	if err := os.WriteFile("config.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("   ▄████████  ▄█     ▄███████▄ ▀█████████▄  ███    █▄   ▄███████▄  ")
	fmt.Println("  ███    ███ ███    ███    ███   ███    ███ ███    ███ ██▀     ▄██ ")
	fmt.Println("  ███    ███ ███▌   ███    ███   ███    ███ ███    ███       ▄███▀ ")
	fmt.Println(" ▄███▄▄▄▄██▀ ███▌   ███    ███  ▄███▄▄▄██▀  ███    ███  ▀█▀▄███▀▄▄ ")
	fmt.Println("▀▀███▀▀▀▀▀   ███▌ ▀█████████▀  ▀▀███▀▀▀██▄  ███    ███   ▄███▀   ▀ ")
	fmt.Println("▀███████████ ███    ███          ███    ██▄ ███    ███ ▄███▀       ")
	fmt.Println("  ███    ███ ███    ███          ███    ███ ███    ███ ███▄     ▄█ ")
	fmt.Println("  ███    ███ █▀    ▄████▀      ▄█████████▀  ████████▀   ▀████████▀ ")
	fmt.Println("  ███    ███                                                       ")
	fmt.Println("")
	config, err := readConfig()
	if err != nil {
		panic(err)
	}

	//login:
	if config.Email == "" || config.Password == "" {
		var email, password string
		if err := huh.NewInput().Title("Enter your Qobuz email").Description("This is required to login to Qobuz.").Value(&email).Run(); err != nil {
			return
		}

		if err := huh.NewInput().Title("Enter your Qobuz password").Description("This is required to login to Qobuz.").EchoMode(huh.EchoModePassword).Value(&password).Run(); err != nil {
			return
		}

		config.Email = email
		config.Password = password

		if err := writeConfig("config.json", config); err != nil {
			panic(err)
		}
	}

	os.Mkdir(config.DownloadFolder, 0755)

	client, err := qobuz.NewFromAuthToken("***REMOVED***")
	if err != nil {
		fmt.Println("Failed to create Qobuz client:", err)
		return
	}

	var mode int
	if err := huh.NewSelect[int]().Title("Select a mode").Description("Choose a mode to continue.").Options(
		huh.NewOption("Download track", 1),
		huh.NewOption("Download album", 2),
	).Value(&mode).Run(); err != nil {
		return
	}

	switch mode {
	case 1:
		downloadTrack(client, config)
		break
	case 2:
		downloadAlbum(client, config)
		break
	}
}

func downloadAlbum(client *qobuz.QobuzClient, config *Config) {
	var query string
	if err := huh.NewInput().Title("Enter an album name/url").Description("This is required to search for an album.").Value(&query).Run(); err != nil {
		return
	}

	regex := regexp.MustCompile(`https://open.qobuz.com/album/(\d+)`)
	matches := regex.FindStringSubmatch(query)

	var album *types.Album
	if len(matches) > 1 {
		albumId := matches[1]

		fmt.Println("Album ID in link:", albumId)
	} else {
		var albums []types.Album
		if err := spinner.New().Title("Searching for albums...").Action(func() {
			response, err := client.Search(query)
			if err != nil {
				fmt.Println("Failed to search for albums:", err)
				return
			}

			albums = response.Albums.Items
		}).Run(); err != nil {
			return
		}

		if len(albums) == 0 {
			fmt.Println("No albums found.")
			return
		}

		var options []huh.Option[int]
		for i, album := range albums {
			options = append(options, huh.NewOption[int](album.Title+" by "+album.Artist.Name, i))
		}

		var selectedAlbum int
		if err := huh.NewSelect[int]().Title("Select an album").Description("Choose an album to download.").Options(options...).Value(&selectedAlbum).Run(); err != nil {
			return
		}

		album = &albums[selectedAlbum]

		albumInfo, err := client.Album(album.Id)
		if err != nil {
			fmt.Println("Failed to get album info:", err)
			return
		}

		var confirm bool
		if err := huh.NewConfirm().Title("Download album").Description(fmt.Sprintf("Are you sure you want to download %s by %s ? It will download %d tracks.", albumInfo.Title, albumInfo.Artist.Name, albumInfo.TracksCount)).Value(&confirm).Run(); err != nil {
			return
		}

		if !confirm {
			return
		}

		os.Mkdir(filepath.Join(config.DownloadFolder, albumInfo.Title), 0755)

		for _, track := range albumInfo.Tracks.Items {
		searchurl:
			trackURL, err := client.DownloadFileLink(fmt.Sprintf("%d", track.Id), 27)
			if err != nil {
				fmt.Println("Failed to get download link:", err)
				return
			}

			if trackURL == nil || trackURL.Url == "" {
				time.Sleep(1 * time.Second)
				goto searchurl
			}

			filePath := filepath.Join(config.DownloadFolder, albumInfo.Title, fmt.Sprintf("%d - %s.%s", track.TrackNumber, track.Title, strings.Split(trackURL.MimeType, "/")[1]))

			fmt.Println("Downloading", track.Title+"...")

			if err := downloadProgress(trackURL.Url, filePath); err != nil {
				fmt.Println("Failed to download track:", err)
				return
			}

			fmt.Printf("Downloaded %s (%d Bit / %.2f kHz).\n", track.Title, trackURL.BitDepth, trackURL.SamplingRate)
		}

		fmt.Printf("\nDownloaded %s.\n", albumInfo.Title)
	}
}

func downloadTrack(client *qobuz.QobuzClient, config *Config) {
	var query string
	if err := huh.NewInput().Title("Enter a track name/url").Description("This is required to search for a track.").Value(&query).Run(); err != nil {
		return
	}

	regex := regexp.MustCompile(`https://open.qobuz.com/track/(\d+)`)
	matches := regex.FindStringSubmatch(query)

	var track *types.Track
	if len(matches) > 1 {
		trackId := matches[1]

		fmt.Println("Track ID in link:", trackId)
	} else {
		var tracks []types.Track
		if err := spinner.New().Title("Searching for tracks...").Action(func() {
			response, err := client.Search(query)
			if err != nil {
				fmt.Println("Failed to search for tracks:", err)
				return
			}

			tracks = response.Tracks.Items
		}).Run(); err != nil {
			return
		}

		if len(tracks) == 0 {
			fmt.Println("No tracks found.")
			return
		}

		var options []huh.Option[int]
		for i, track := range tracks {
			options = append(options, huh.NewOption[int](track.Title+" by "+track.Performer.Name, i))
		}

		var selectedTrack int
		if err := huh.NewSelect[int]().Title("Select a track").Description("Choose a track to download.").Options(options...).Value(&selectedTrack).Run(); err != nil {
			return
		}

		track = &tracks[selectedTrack]
	}

	if track == nil {
		fmt.Println("No track found.")
		return
	}

	var confirm bool
	if err := huh.NewConfirm().Title("Download track").Description(fmt.Sprintf("Are you sure you want to download %s by %s ?", track.Title, track.Performer.Name)).Value(&confirm).Run(); err != nil {
		return
	}

	if !confirm {
		return
	}

searchurl:
	trackURL, err := client.DownloadFileLink(fmt.Sprintf("%d", track.Id), 27)
	if err != nil {
		fmt.Println("Failed to get download link:", err)
		return
	}

	if trackURL == nil || trackURL.Url == "" {
		time.Sleep(1 * time.Second)
		goto searchurl
	}

	filename := filepath.Base(track.Title + " - " + track.Performer.Name + "." + strings.Split(trackURL.MimeType, "/")[1])

	if err := downloadProgress(trackURL.Url, filepath.Join(config.DownloadFolder, filename)); err != nil {
		fmt.Println("Failed to download track:", err)
		return
	}

	fmt.Printf("Downloaded %s to %s (%d Bit / %.2f kHz)\n", filename, config.DownloadFolder, trackURL.BitDepth, trackURL.SamplingRate)
}

func downloadProgress(url, outputPath string) error {
	resp, err := getResponse(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close() // nolint:errcheck

	if resp.ContentLength <= 0 {
		return errors.New("could not get content length")
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer file.Close()

	pw := &progressWriter{
		total:  int(resp.ContentLength),
		file:   file,
		reader: resp.Body,
		onProgress: func(ratio float64) {
			p.Send(progressMsg(ratio))
		},
	}

	m := model{
		pw:       pw,
		progress: progress.New(progress.WithDefaultGradient()),
	}

	p = tea.NewProgram(m)

	go pw.Start()

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
