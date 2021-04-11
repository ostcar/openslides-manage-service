package manage

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

const startHelp = `Builds images and starts OpenSlides with Docker Compose

This command executes the following steps to start OpenSlides:
- Create a local docker-compose.yml.
- Create local secrets for the auth service.
- Run the docker compose file with the "up" command in daemonized mode.
- TODO ...
`

const appName = "openslides"

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

		dataPath := path.Join(xdg.DataHome, appName)

		// TODO: Only create file, if it does not exists.
		if err := createDockerComposeYML(ctx, dataPath); err != nil {
			return fmt.Errorf("creating Docker Compose YML: %w", err)
		}

		if err := createSecrets(dataPath); err != nil {
			return fmt.Errorf("creating secrets: %w", err)
		}

		// TODO: Start docker-compose. Maybe with another command like manage
		// compose XY where XY are arguments that the user can give.

		return nil
	}

	return cmd
}

// create Secrets creates random values uses as secrets in Docker Compose file.
func createSecrets(dataPath string) error {
	dataPath = path.Join(dataPath, "secrets")
	if err := os.MkdirAll(dataPath, fs.ModePerm); err != nil {
		return fmt.Errorf("creating directory `%s`: %w", dataPath, err)
	}

	secrets := []string{
		"auth_token_key",
		"auth_cookie_key",
	}
	for _, s := range secrets {
		f, err := os.Create(path.Join(dataPath, s))
		if err != nil {
			return fmt.Errorf("creating file `%s`: %w", path.Join(dataPath, s), err)
		}

		// @Norman (pls remove this comment): The defer is ok here. But there is
		// something to know about. The defer is only called when the function
		// exists. So if secrets would be a very long list with a million
		// entries, then it would open a million files and close all of them at
		// the end of the function call. To fix this, the code inside the
		// for-loop has to be moved inside a separat function. For example an
		// anonymous function:
		//
		// for {
		//   func() {
		//     f, err := open()
		//     defer f.Close
		//   }()
		// }
		//
		// If you have a function that blocks forever, then you have to do it
		// (or don't use defer).
		//
		// Also, there are optimizations about defer in go that do not work if
		// the defer is inside a for-loop. But does not matter at all. Its from
		// around 35ns to 3ns.
		defer f.Close()

		// This creates cryptographically secure random bytes. 32 Bytes means
		// 256bit. The output can contain zero bytes.
		if _, err := io.Copy(f, io.LimitReader(rand.Reader, 32)); err != nil {
			return fmt.Errorf("creating and writing secred: %w", err)
		}
	}

	return nil
}
