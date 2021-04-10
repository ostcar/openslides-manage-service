package manage

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const startHelp = `Builds images and starts OpenSlides with Docker Compose

This command executes the following steps to start OpenSlides:
- Create a local docker-compose.yml.
- Create local secrets for the auth service.
- Run the docker compose file with the "up" command in daemonized mode.
- TODO ...
`

// CmdStart creates docker-compose.yml, secrets, runs docker-compose up in daemonized mode and ... TODO
func CmdStart(cfg *ClientConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Builds images and starts OpenSlides with Docker Compose.",
		Long:  startHelp,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		if err := createDockerComposeYML(ctx); err != nil {
			return fmt.Errorf("creating Docker Compose YML: %w", err)
		}

		return nil
	}

	return cmd
}

// createDockerComposeYML creates a docker-compose.yml file in the current working directory
// using a template. In non local mode it uses the GitHub API to fetch the required commit IDs
// of all services.
func createDockerComposeYML(ctx context.Context) error {
	filename := "docker-compose.yml"
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file `%s`: %w", filename, err)
	}
	defer f.Close()

	if err := writeDockerComposeYML(ctx, f); err != nil {
		return fmt.Errorf("writing content to file `%s`: %w", filename, err)
	}

	return nil
}

//go:embed docker-compose.yml.tpl
var defaultDockerCompose string

// writeDockerComposeYML writes the populated template to the given writer.
func writeDockerComposeYML(ctx context.Context, w io.Writer) error {
	// TODO:
	// * Use services.env https://github.com/OpenSlides/OpenSlides/blob/openslides4-dev/docker/services.env
	// TODO: Local case

	composeTPL, err := template.New("compose").Parse(defaultDockerCompose)
	if err != nil {
		return fmt.Errorf("creating Docker Compose template: %w", err)
	}
	composeTPL.Option("missingkey=error")

	var tplData struct {
		ExternalHTTPPort   string
		ExternalManagePort string
		CommitID           map[string]string
		Ref                string
	}

	tplData.ExternalHTTPPort = "8000"
	tplData.ExternalManagePort = "9008"
	tplData.Ref = "openslides4-dev"

	c, err := getCommitIDs(ctx, tplData.Ref)
	if err != nil {
		return fmt.Errorf("getting commit IDs: %w", err)
	}
	tplData.CommitID = c

	if err := composeTPL.Execute(w, tplData); err != nil {
		return fmt.Errorf("writing Docker Compose file: %w", err)
	}
	return nil
}

// getCommitIDs fetches the commit IDs for all services from GitHub API.
func getCommitIDs(ctx context.Context, ref string) (map[string]string, error) {
	addr := "https://api.github.com/repos/OpenSlides/OpenSlides/contents?ref=" + ref
	req, err := http.NewRequestWithContext(ctx, "GET", addr, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request to GitHub API: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request to GitHub API at %s: %w", addr, err)
	}
	defer resp.Body.Close()

	var apiBody []struct {
		Name string `json:"name"`
		SHA  string `json:"sha"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiBody); err != nil {
		return nil, fmt.Errorf("reading and decoding body from GitHub API: %w", err)
	}

	services := map[string]string{
		"openslides-client":             "client",
		"openslides-backend":            "backend",
		"openslides-datastore-service":  "datastore",
		"openslides-autoupdate-service": "autoupdate",
		"openslides-auth-service":       "auth",
		"openslides-media-service":      "media",
		"openslides-manage-service":     "manage",
		"openslides-permission-service": "permission", // TODO: Remove this line after permission service is removed.
	}

	commitIDs := make(map[string]string, len(services))
	for _, apiElement := range apiBody {
		tplName, ok := services[apiElement.Name]
		if ok {
			commitIDs[tplName] = apiElement.SHA
		}
	}

	return commitIDs, nil
}
