package ffuf

// import (
// 	"time"
// )

// // MatcherManager provides functions for managing matchers and filters
// type MatcherManager interface {
// 	SetCalibrated(calibrated bool)
// 	SetCalibratedForHost(host string, calibrated bool)
// 	AddFilter(name string, option string, replace bool) error
// 	AddPerDomainFilter(domain string, name string, option string) error
// 	RemoveFilter(name string)
// 	AddMatcher(name string, option string) error
// 	GetFilters() map[string]FilterProvider
// 	GetMatchers() map[string]FilterProvider
// 	FiltersForDomain(domain string) map[string]FilterProvider
// 	CalibratedForDomain(domain string) bool
// 	Calibrated() bool
// }

// // FilterProvider is a generic interface for both Matchers and Filters
// type FilterProvider interface {
// 	Filter(response *http.Response) (bool, error)
// 	Repr() string
// 	ReprVerbose() string
// }

// RunnerProvider is an interface for request executors
// type RunnerProvider interface {
// 	Prepare(input map[string][]byte, basereq *http.Request) (http.Request, error)
// 	Execute(req *http.Request) (http.Response, error)
// }

// InputProvider interface handles the input data for RunnerProvider
// type InputProvider interface {
// 	ActivateKeywords([]string)
// 	AddProvider(config.InputProviderConfig) error
// 	Keywords() []string
// 	Next() bool
// 	Position() int
// 	Reset()
// 	Value() map[string][]byte
// 	Total() int
// }

// // InternalInputProvider interface handles providing input data to InputProvider
// type InternalInputProvider interface {
// 	Keyword() string
// 	Next() bool
// 	Position() int
// 	ResetPosition()
// 	IncrementPosition()
// 	Value() []byte
// 	Total() int
// 	Active() bool
// 	Enable()
// 	Disable()
// }

// OutputProvider is responsible of providing output from the RunnerProvider
// type OutputProvider interface {
// 	Banner()
// 	Finalize() error
// 	Progress(status Progress)
// 	Info(infostring string)
// 	Error(errstring string)
// 	Raw(output string)
// 	Warning(warnstring string)
// 	Result(resp http.Response)
// 	PrintResult(res Result)
// 	SaveFile(filename, format string) error
// 	GetCurrentResults() []Result
// 	SetCurrentResults(results []Result)
// 	Reset()
// 	Cycle()
// }

// type Result struct {
// 	Input            map[string][]byte `json:"input"`
// 	Position         int               `json:"position"`
// 	StatusCode       int64             `json:"status"`
// 	ContentLength    int64             `json:"length"`
// 	ContentWords     int64             `json:"words"`
// 	ContentLines     int64             `json:"lines"`
// 	ContentType      string            `json:"content-type"`
// 	RedirectLocation string            `json:"redirectlocation"`
// 	Url              string            `json:"url"`
// 	Duration         time.Duration     `json:"duration"`
// 	ResultFile       string            `json:"resultfile"`
// 	Host             string            `json:"host"`
// 	HTMLColor        string            `json:"-"`
// }
