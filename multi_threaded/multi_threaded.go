package multi_threaded

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/pure-project/go-pdfium"
	"github.com/pure-project/go-pdfium/internal/commons"
)

type worker struct {
	plugin       commons.Pdfium
	pluginClient *plugin.Client
	rpcClient    plugin.ClientProtocol
}

type Config struct {
	LogCallback func(string)
	Command     Command
}

type Command struct {
	BinPath string
	Args    []string

	// StartTimeout is the timeout to wait for the plugin to say it
	// has started successfully.
	StartTimeout time.Duration
}

var (
	pluginClientMap sync.Map
	newInstance     func() (pdfium.Pdfium, error)
)

// Init will prepare pdfium subprocess argument.
func Init(config Config) {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"pdfium": &commons.PdfiumPlugin{},
	}

	// If we don't have a log callback, make the callback no-op.
	if config.LogCallback == nil {
		config.LogCallback = func(s string) {}
	}

	newInstance = func() (pdfium.Pdfium, error) {
		newWorker := &worker{}

		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(config.Command.BinPath, config.Command.Args...),
			Logger:          logger,
			StartTimeout:    config.Command.StartTimeout,
		})

		rpcClient, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpcClient.Dispense("pdfium")
		if err != nil {
			return nil, err
		}

		pdfium := raw.(commons.Pdfium)

		pong, err := pdfium.Ping()
		if err != nil {
			return nil, err
		}

		if pong != "Pong" {
			return nil, errors.New("Wrong ping/pong result")
		}

		newWorker.pluginClient = client
		newWorker.rpcClient = rpcClient
		newWorker.plugin = pdfium

		inst := &pdfiumInstance{
			worker: newWorker,
			lock:   &sync.Mutex{},
		}

		instanceRef := uuid.New()
		inst.instanceRef = instanceRef.String()

		pluginClientMap.Store(inst.instanceRef, newWorker.pluginClient)

		return inst, nil
	}
}

// Fini kill all pdfium worker subprocess
func Fini() {
	pluginClientMap.Range(func(k, v any) bool {
		client := v.(*plugin.Client)
		client.Kill()
		return true
	})
}

// NewInstance create new pdfium worker instance
func NewInstance() (pdfium.Pdfium, error) {
	return newInstance()
}

type pdfiumInstance struct {
	worker      *worker
	instanceRef string
	closed      bool
	lock        *sync.Mutex
}

// Close will close the instance and will clean up the underlying PDFium resources
// by calling i.worker.plugin.Close().
func (i *pdfiumInstance) Close() (err error) {
	i.lock.Lock()

	if i.closed {
		i.lock.Unlock()
		return errors.New("instance is already closed")
	}

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "Close", panicError)
		}
	}()

	pluginClientMap.Delete(i.instanceRef)

	i.worker.plugin.Close()
	i.worker.pluginClient.Kill()

	return nil
}

func (i *pdfiumInstance) GetImplementation() interface{} {
	return i.worker.plugin
}
