package server

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/vul-db/pkg/db"
	"github.com/khulnasoft/tunnel/pkg/commands/operation"
	"github.com/khulnasoft/tunnel/pkg/flag"
	"github.com/khulnasoft/tunnel/pkg/log"
	"github.com/khulnasoft/tunnel/pkg/module"
	rpcServer "github.com/khulnasoft/tunnel/pkg/rpc/server"
	"github.com/khulnasoft/tunnel/pkg/utils/fsutils"
)

// Run runs the scan
func Run(ctx context.Context, opts flag.Options) (err error) {
	if err = log.InitLogger(opts.Debug, opts.Quiet); err != nil {
		return xerrors.Errorf("failed to initialize a logger: %w", err)
	}

	// configure cache dir
	fsutils.SetCacheDir(opts.CacheDir)
	cache, err := operation.NewCache(opts.CacheOptions)
	if err != nil {
		return xerrors.Errorf("server cache error: %w", err)
	}
	defer cache.Close()
	log.Logger.Debugf("cache dir:  %s", fsutils.CacheDir())

	if opts.Reset {
		return cache.ClearDB()
	}

	// download the database file
	if err = operation.DownloadDB(ctx, opts.AppVersion, opts.CacheDir, opts.DBRepository,
		true, opts.SkipDBUpdate, opts.RegistryOpts()); err != nil {
		return err
	}

	if opts.DownloadDBOnly {
		return nil
	}

	if err = db.Init(opts.CacheDir); err != nil {
		return xerrors.Errorf("error in vulnerability DB initialize: %w", err)
	}

	// Initialize WASM modules
	m, err := module.NewManager(ctx, module.Options{
		Dir:            opts.ModuleDir,
		EnabledModules: opts.EnabledModules,
	})
	if err != nil {
		return xerrors.Errorf("WASM module error: %w", err)
	}
	m.Register()

	server := rpcServer.NewServer(opts.AppVersion, opts.Listen, opts.CacheDir, opts.Token, opts.TokenHeader,
		opts.DBRepository, opts.RegistryOpts())
	return server.ListenAndServe(cache, opts.SkipDBUpdate)
}
