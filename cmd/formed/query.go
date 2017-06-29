package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/SimonRichardson/formed/pkg/fs"
	"github.com/SimonRichardson/formed/pkg/query"
	"github.com/SimonRichardson/formed/pkg/store"
	"github.com/SimonRichardson/formed/pkg/templates"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const (
	defaultFileStore = "./data/store.csv"
)

// runQuery creates all the dependencies required to create and run the query
// end point for the form component.
func runQuery(args []string) error {
	var (
		flagset = flag.NewFlagSet("query", flag.ExitOnError)

		debug     = flagset.Bool("debug", false, "debug logging")
		apiAddr   = flagset.String("api", defaultAPIAddr, "listen address for query API")
		fileStore = flagset.String("filestore", defaultFileStore, "location of where the file store")
		uiLocal   = flagset.Bool("ui.local", false, "ignores embedded files and goes straight to the filesystem")
	)

	flagset.Usage = usageFor(flagset, "query [flags]")
	if err := flagset.Parse(args); err != nil {
		return nil
	}

	// Setup the logger.
	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		if *debug {
			logLevel = level.AllowAll()
		}
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = level.NewFilter(logger, logLevel)
	}

	// Parse the apiNetwork and apiAddress from the flag set
	apiNetwork, apiAddress, err := parseAddr(*apiAddr, defaultAPIPort)
	if err != nil {
		return err
	}

	// Get all the templates for the query
	templates, err := gatherTemplates(*uiLocal)
	if err != nil {
		return err
	}

	// Create the api listener for the service
	apiListener, err := net.Listen(apiNetwork, apiAddress)
	if err != nil {
		return err
	}
	level.Debug(logger).Log("API", fmt.Sprintf("%s://%s", apiNetwork, apiAddress))

	// Execution group.
	defer apiListener.Close()

	// API that is going to handle the incoming requests.
	var (
		fsys     = fs.New()
		store    = store.New(fsys, *fileStore)
		injector = query.NewInjector(store, templates)
		api      = query.NewAPI(injector, log.With(logger, "component", "api"))
	)

	mux := http.NewServeMux()
	mux.Handle("/query/", http.StripPrefix("/query", api))

	return http.Serve(apiListener, mux)
}

func gatherTemplates(uiLocal bool) (*templates.Templates, error) {
	fallback, err := templates.NewErrorTemplate(uiLocal)
	if err != nil {
		return nil, errors.Wrap(err, "no fallback template")
	}

	formTemplate, err := templates.NewFormTemplate(uiLocal)
	if err != nil {
		return nil, errors.Wrap(err, "no form template")
	}

	templates := templates.NewTemplates(fallback)
	templates.Set(http.StatusOK, formTemplate)
	return templates, nil
}
