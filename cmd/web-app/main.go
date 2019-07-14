package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"expvar"
	"fmt"
	"geeks-accelerator/oss/saas-starter-kit/internal/mid"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"html/template"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"time"

	"geeks-accelerator/oss/saas-starter-kit/cmd/web-app/handlers"
	"geeks-accelerator/oss/saas-starter-kit/internal/platform/devops"
	"geeks-accelerator/oss/saas-starter-kit/internal/platform/flag"
	img_resize "geeks-accelerator/oss/saas-starter-kit/internal/platform/img-resize"
	"geeks-accelerator/oss/saas-starter-kit/internal/platform/web"
	template_renderer "geeks-accelerator/oss/saas-starter-kit/internal/platform/web/template-renderer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"
	"github.com/lib/pq"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
	sqlxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

// service is the name of the program used for logging, tracing and the
// the prefix used for loading env variables
// ie: export WEB_APP_ENV=dev
var service = "WEB_APP"

func main() {

	// =========================================================================
	// Logging

	log := log.New(os.Stdout, service+" : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// =========================================================================
	// Configuration
	var cfg struct {
		Env  string `default:"dev" envconfig:"ENV"`
		HTTP struct {
			Host         string        `default:"0.0.0.0:3000" envconfig:"HOST"`
			ReadTimeout  time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
			WriteTimeout time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
		}
		HTTPS struct {
			Host         string        `default:"" envconfig:"HOST"`
			ReadTimeout  time.Duration `default:"5s" envconfig:"READ_TIMEOUT"`
			WriteTimeout time.Duration `default:"5s" envconfig:"WRITE_TIMEOUT"`
			DisableHTTP2 bool          `default:"false" envconfig:"DISABLE_HTTP2"`
		}
		Service struct {
			Name        string   `default:"web-app" envconfig:"NAME"`
			Project     string   `default:"" envconfig:"PROJECT"`
			BaseUrl     string   `default:"" envconfig:"BASE_URL"  example:"http://eproc.tech"`
			HostNames   []string `envconfig:"HOST_NAMES" example:"www.eproc.tech"`
			EnableHTTPS bool     `default:"false" envconfig:"ENABLE_HTTPS"`
			TemplateDir string   `default:"./templates" envconfig:"TEMPLATE_DIR"`
			StaticDir   string   `default:"./static" envconfig:"STATIC_DIR"`
			StaticS3    struct {
				S3Enabled         bool   `envconfig:"ENABLED"`
				S3KeyPrefix       string `default:"public/web_app/static" envconfig:"KEY_PREFIX"`
				CloudFrontEnabled bool   `envconfig:"CLOUDFRONT_ENABLED"`
				ImgResizeEnabled  bool   `envconfig:"IMG_RESIZE_ENABLED"`
			}
			DebugHost       string        `default:"0.0.0.0:4000" envconfig:"DEBUG_HOST"`
			ShutdownTimeout time.Duration `default:"5s" envconfig:"SHUTDOWN_TIMEOUT"`
		}
		Redis struct {
			Host            string        `default:":6379" envconfig:"HOST"`
			DB              int           `default:"1" envconfig:"DB"`
			DialTimeout     time.Duration `default:"5s" envconfig:"DIAL_TIMEOUT"`
			MaxmemoryPolicy string        `envconfig:"MAXMEMORY_POLICY"`
		}
		DB struct {
			Host       string `default:"127.0.0.1:5433" envconfig:"HOST"`
			User       string `default:"postgres" envconfig:"USER"`
			Pass       string `default:"postgres" envconfig:"PASS" json:"-"` // don't print
			Database   string `default:"shared" envconfig:"DATABASE"`
			Driver     string `default:"postgres" envconfig:"DRIVER"`
			Timezone   string `default:"utc" envconfig:"TIMEZONE"`
			DisableTLS bool   `default:"true" envconfig:"DISABLE_TLS"`
		}
		Trace struct {
			Host          string  `default:"127.0.0.1" envconfig:"DD_TRACE_AGENT_HOSTNAME"`
			Port          int     `default:"8126" envconfig:"DD_TRACE_AGENT_PORT"`
			AnalyticsRate float64 `default:"0.10" envconfig:"ANALYTICS_RATE"`
		}
		Aws struct {
			AccessKeyID                string `envconfig:"AWS_ACCESS_KEY_ID"`              // WEB_API_AWS_AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY_ID
			SecretAccessKey            string `envconfig:"AWS_SECRET_ACCESS_KEY" json:"-"` // don't print
			Region                     string `default:"us-east-1" envconfig:"AWS_REGION"`
			S3BucketPrivate            string `envconfig:"S3_BUCKET_PRIVATE"`
			S3BucketPublic             string `envconfig:"S3_BUCKET_PUBLIC"`
			SecretsManagerConfigPrefix string `default:"" envconfig:"SECRETS_MANAGER_CONFIG_PREFIX"`

			// Get an AWS session from an implicit source if no explicit
			// configuration is provided. This is useful for taking advantage of
			// EC2/ECS instance roles.
			UseRole bool `envconfig:"AWS_USE_ROLE"`
		}
		Auth struct {
			UseAwsSecretManager bool          `default:"false" envconfig:"USE_AWS_SECRET_MANAGER"`
			KeyExpiration       time.Duration `default:"3600s" envconfig:"KEY_EXPIRATION"`
		}
		BuildInfo struct {
			CiCommitRefName     string `envconfig:"CI_COMMIT_REF_NAME"`
			CiCommitRefSlug     string `envconfig:"CI_COMMIT_REF_SLUG"`
			CiCommitSha         string `envconfig:"CI_COMMIT_SHA"`
			CiCommitTag         string `envconfig:"CI_COMMIT_TAG"`
			CiCommitTitle       string `envconfig:"CI_COMMIT_TITLE"`
			CiCommitDescription string `envconfig:"CI_COMMIT_DESCRIPTION"`
			CiJobId             string `envconfig:"CI_COMMIT_JOB_ID"`
			CiJobUrl            string `envconfig:"CI_COMMIT_JOB_URL"`
			CiPipelineId        string `envconfig:"CI_COMMIT_PIPELINE_ID"`
			CiPipelineUrl       string `envconfig:"CI_COMMIT_PIPELINE_URL"`
		}
	}

	// For additional details refer to https://github.com/kelseyhightower/envconfig
	if err := envconfig.Process(service, &cfg); err != nil {
		log.Fatalf("main : Parsing Config : %+v", err)
	}

	if err := flag.Process(&cfg); err != nil {
		if err != flag.ErrHelp {
			log.Fatalf("main : Parsing Command Line : %+v", err)
		}
		return // We displayed help.
	}

	// =========================================================================
	// Config Validation & Defaults

	// AWS access keys are required, if roles are enabled, remove any placeholders.
	if cfg.Aws.UseRole {
		cfg.Aws.AccessKeyID = ""
		cfg.Aws.SecretAccessKey = ""
	}

	// Set the default AWS Secrets Manager prefix used for name to store config files that will be persisted across
	// deployments and distributed to each instance of the service running.
	if cfg.Aws.SecretsManagerConfigPrefix == "" {
		var pts []string
		if cfg.Service.Project != "" {
			pts = append(pts, cfg.Service.Project)
		}
		pts = append(pts, cfg.Env, cfg.Service.Name)

		cfg.Aws.SecretsManagerConfigPrefix = filepath.Join(pts...)
	}

	// If base URL is empty, set the default value from the HTTP Host
	if cfg.Service.BaseUrl == "" {
		baseUrl := cfg.HTTP.Host
		if !strings.HasPrefix(baseUrl, "http") {
			if strings.HasPrefix(baseUrl, "0.0.0.0:") {
				pts := strings.Split(baseUrl, ":")
				pts[0] = "127.0.0.1"
				baseUrl = strings.Join(pts, ":")
			} else if strings.HasPrefix(baseUrl, ":") {
				baseUrl = "127.0.0.1" + baseUrl
			}
			baseUrl = "http://" + baseUrl
		}
		cfg.Service.BaseUrl = baseUrl
	}

	// When HTTPS is not specifically enabled, but an HTTP host is set, enable HTTPS.
	if !cfg.Service.EnableHTTPS && cfg.HTTPS.Host != "" {
		cfg.Service.EnableHTTPS = true
	}

	// Determine the primary host by parsing host from the base app URL.
	baseSiteUrl, err := url.Parse(cfg.Service.BaseUrl)
	if err != nil {
		log.Fatalf("main : Parse service base URL : %s : %+v", cfg.Service.BaseUrl, err)
	}

	// Drop any ports from the base app URL.
	var primaryServiceHost string
	if strings.Contains(baseSiteUrl.Host, ":") {
		primaryServiceHost, _, err = net.SplitHostPort(baseSiteUrl.Host)
		if err != nil {
			log.Fatalf("main : SplitHostPort : %s : %+v", baseSiteUrl.Host, err)
		}
	} else {
		primaryServiceHost = baseSiteUrl.Host
	}

	// =========================================================================
	// Log Service Info

	// Print the build version for our logs. Also expose it under /debug/vars.
	expvar.NewString("build").Set(build)
	log.Printf("main : Started : Service Initializing version %q", build)
	defer log.Println("main : Completed")

	// Print the config for our logs. It's important to any credentials in the config
	// that could expose a security risk are excluded from being json encoded by
	// applying the tag `json:"-"` to the struct var.
	{
		cfgJSON, err := json.MarshalIndent(cfg, "", "    ")
		if err != nil {
			log.Fatalf("main : Marshalling Config to JSON : %+v", err)
		}
		log.Printf("main : Config : %v\n", string(cfgJSON))
	}

	// =========================================================================
	// Init AWS Session
	var awsSession *session.Session
	if cfg.Aws.UseRole {
		// Get an AWS session from an implicit source if no explicit
		// configuration is provided. This is useful for taking advantage of
		// EC2/ECS instance roles.
		awsSession = session.Must(session.NewSession())
		if cfg.Aws.Region != "" {
			awsSession.Config.WithRegion(cfg.Aws.Region)
		}

		log.Printf("main : AWS : Using role.\n")

	} else if cfg.Aws.AccessKeyID != "" {
		creds := credentials.NewStaticCredentials(cfg.Aws.AccessKeyID, cfg.Aws.SecretAccessKey, "")
		awsSession = session.New(&aws.Config{Region: aws.String(cfg.Aws.Region), Credentials: creds})

		log.Printf("main : AWS : Using static credentials\n")
	}

	// Wrap the AWS session to enable tracing.
	if awsSession != nil {
		awsSession = awstrace.WrapSession(awsSession)
	}

	// =========================================================================
	// Start Redis
	// Ensure the eviction policy on the redis cluster is set correctly.
	// 		AWS Elastic cache redis clusters by default have the volatile-lru.
	// 			volatile-lru: evict keys by trying to remove the less recently used (LRU) keys first, but only among keys that have an expire set, in order to make space for the new data added.
	// 			allkeys-lru: evict keys by trying to remove the less recently used (LRU) keys first, in order to make space for the new data added.
	//		Recommended to have eviction policy set to allkeys-lru
	log.Println("main : Started : Initialize Redis")
	redisClient := redistrace.NewClient(&redis.Options{
		Addr:        cfg.Redis.Host,
		DB:          cfg.Redis.DB,
		DialTimeout: cfg.Redis.DialTimeout,
	})
	defer redisClient.Close()

	evictPolicyConfigKey := "maxmemory-policy"

	// if the maxmemory policy is set for redis, make sure its set on the cluster
	// default not set and will based on the redis config values defined on the server
	if cfg.Redis.MaxmemoryPolicy != "" {
		err := redisClient.ConfigSet(evictPolicyConfigKey, cfg.Redis.MaxmemoryPolicy).Err()
		if err != nil && !strings.Contains(err.Error(), "unknown command") {
			log.Fatalf("main : redis : ConfigSet maxmemory-policy : %+v", err)
		}
	} else {
		evictPolicy, err := redisClient.ConfigGet(evictPolicyConfigKey).Result()
		if err != nil && !strings.Contains(err.Error(), "unknown command") {
			log.Fatalf("main : redis : ConfigGet maxmemory-policy : %+v", err)
		} else if evictPolicy != nil && len(evictPolicy) > 0 && evictPolicy[1] != "allkeys-lru" {
			log.Printf("main : redis : ConfigGet maxmemory-policy : recommended to be set to allkeys-lru to avoid OOM")
		}
	}

	// =========================================================================
	// Start Database
	var dbUrl url.URL
	{
		// Query parameters.
		var q url.Values = make(map[string][]string)

		// Handle SSL Mode
		if cfg.DB.DisableTLS {
			q.Set("sslmode", "disable")
		} else {
			q.Set("sslmode", "require")
		}

		q.Set("timezone", cfg.DB.Timezone)

		// Construct url.
		dbUrl = url.URL{
			Scheme:   cfg.DB.Driver,
			User:     url.UserPassword(cfg.DB.User, cfg.DB.Pass),
			Host:     cfg.DB.Host,
			Path:     cfg.DB.Database,
			RawQuery: q.Encode(),
		}
	}
	log.Println("main : Started : Initialize Database")

	// Register informs the sqlxtrace package of the driver that we will be using in our program.
	// It uses a default service name, in the below case "postgres.db". To use a custom service
	// name use RegisterWithServiceName.
	sqltrace.Register(cfg.DB.Driver, &pq.Driver{}, sqltrace.WithServiceName(service))
	masterDb, err := sqlxtrace.Open(cfg.DB.Driver, dbUrl.String())
	if err != nil {
		log.Fatalf("main : Register DB : %s : %+v", cfg.DB.Driver, err)
	}
	defer masterDb.Close()

	// =========================================================================
	// Load middlewares that need to be configured specific for the service.

	var serviceMiddlewares []web.Middleware

	// Init redirect middleware to ensure all requests go to the primary domain contained in the base URL.
	if primaryServiceHost != "127.0.0.0" && primaryServiceHost != "localhost" {
		redirect := mid.DomainNameRedirect(mid.DomainNameRedirectConfig{
			RedirectConfig: mid.RedirectConfig{
				Code: http.StatusMovedPermanently,
				Skipper: func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) bool {
					if r.URL.Path == "/ping" {
						return true
					}
					return false
				},
			},
			DomainName:   primaryServiceHost,
			HTTPSEnabled: cfg.Service.EnableHTTPS,
		})
		serviceMiddlewares = append(serviceMiddlewares, redirect)
	}

	// =========================================================================
	// URL Formatter
	// s3UrlFormatter is a help function used by to convert an s3 key to
	// a publicly available image URL.
	var staticS3UrlFormatter func(string) string
	if cfg.Service.StaticS3.S3Enabled || cfg.Service.StaticS3.CloudFrontEnabled || cfg.Service.StaticS3.ImgResizeEnabled {
		s3UrlFormatter, err := devops.S3UrlFormatter(awsSession, cfg.Aws.S3BucketPublic, cfg.Service.StaticS3.S3KeyPrefix, cfg.Service.StaticS3.CloudFrontEnabled)
		if err != nil {
			log.Fatalf("main : S3UrlFormatter failed : %+v", err)
		}

		staticS3UrlFormatter = func(p string) string {
			// When the path starts with a forward slash its referencing a local file,
			// make sure the static file prefix is included
			if strings.HasPrefix(p, "/") {
				p = filepath.Join(cfg.Service.StaticS3.S3KeyPrefix, p)
			}
			return s3UrlFormatter(p)
		}
	} else {
		baseUrl, err := url.Parse(cfg.Service.BaseUrl)
		if err != nil {
			log.Fatalf("main : url Parse(%s) : %+v", cfg.Service.BaseUrl, err)
		}

		staticS3UrlFormatter = func(p string) string {
			baseUrl.Path = p
			return baseUrl.String()
		}
	}

	// staticUrlFormatter is a help function used by template functions defined below.
	// If the app has an S3 bucket defined for the static directory, all references in the app
	// templates should be updated to use a fully qualified URL for either the public file on S3
	// on from the cloudfront distribution.
	var staticUrlFormatter func(string) string
	if cfg.Service.StaticS3.S3Enabled || cfg.Service.StaticS3.CloudFrontEnabled {
		staticUrlFormatter = staticS3UrlFormatter
	} else {
		baseUrl, err := url.Parse(cfg.Service.BaseUrl)
		if err != nil {
			log.Fatalf("main : url Parse(%s) : %+v", cfg.Service.BaseUrl, err)
		}

		staticUrlFormatter = func(p string) string {
			baseUrl.Path = p
			return baseUrl.String()
		}
	}

	// =========================================================================
	// Template Renderer
	// Implements interface web.Renderer to support alternative renderer

	// Append query string value to break browser cache used for services
	// that render responses for a browser with the following:
	// 	1. when env=dev, the current timestamp will be used to ensure every
	// 		request will skip browser cache.
	// 	2. all other envs, ie stage and prod. The commit hash will be used to
	// 		ensure that all cache will be reset with each new deployment.
	browserCacheBusterQueryString := func() string {
		var v string
		if cfg.Env == "dev" {
			// On dev always break cache.
			v = fmt.Sprintf("%d", time.Now().UTC().Unix())
		} else {
			// All other envs, use the current commit hash for the build
			v = cfg.BuildInfo.CiCommitSha
		}
		return v
	}

	// Helper method for appending the browser cache buster as a query string to
	// support breaking browser cache when necessary
	browserCacheBusterFunc := browserCacheBuster(browserCacheBusterQueryString)

	// Need defined functions below since they require config values, able to add additional functions
	// here to extend functionality.
	tmplFuncs := template.FuncMap{
		"BuildInfo": func(k string) string {
			r := reflect.ValueOf(cfg.BuildInfo)
			f := reflect.Indirect(r).FieldByName(k)
			return f.String()
		},
		"SiteBaseUrl": func(p string) string {
			u, err := url.Parse(cfg.HTTP.Host)
			if err != nil {
				return "?"
			}
			u.Path = p
			return u.String()
		},
		"AssetUrl": func(p string) string {
			var u string
			if staticUrlFormatter != nil {
				u = staticUrlFormatter(p)
			} else {
				if !strings.HasPrefix(p, "/") {
					p = "/" + p
				}
				u = p
			}

			u = browserCacheBusterFunc(u)

			return u
		},
		"SiteAssetUrl": func(p string) string {
			var u string
			if staticUrlFormatter != nil {
				u = staticUrlFormatter(filepath.Join(cfg.Service.Name, p))
			} else {
				if !strings.HasPrefix(p, "/") {
					p = "/" + p
				}
				u = p
			}

			u = browserCacheBusterFunc(u)

			return u
		},
		"SiteS3Url": func(p string) string {
			var u string
			if staticUrlFormatter != nil {
				u = staticUrlFormatter(filepath.Join(cfg.Service.Name, p))
			} else {
				u = p
			}
			return u
		},
		"S3Url": func(p string) string {
			var u string
			if staticUrlFormatter != nil {
				u = staticUrlFormatter(p)
			} else {
				u = p
			}
			return u
		},
	}

	// Image Formatter - additional functions exposed to templates for resizing images
	// to support response web applications.
	imgResizeS3KeyPrefix := filepath.Join(cfg.Service.StaticS3.S3KeyPrefix, "images/responsive")

	imgSrcAttr := func(ctx context.Context, p string, sizes []int, includeOrig bool) template.HTMLAttr {
		u := staticUrlFormatter(p)
		var srcAttr string
		if cfg.Service.StaticS3.ImgResizeEnabled {
			srcAttr, _ = img_resize.S3ImgSrc(ctx, redisClient, staticS3UrlFormatter, awsSession, cfg.Aws.S3BucketPublic, imgResizeS3KeyPrefix, u, sizes, includeOrig)
		} else {
			srcAttr = fmt.Sprintf("src=\"%s\"", u)
		}
		return template.HTMLAttr(srcAttr)
	}

	tmplFuncs["S3ImgSrcLarge"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320, 480, 800}, true)
	}
	tmplFuncs["S3ImgThumbSrcLarge"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320, 480, 800}, false)
	}
	tmplFuncs["S3ImgSrcMedium"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320, 640}, true)
	}
	tmplFuncs["S3ImgThumbSrcMedium"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320, 640}, false)
	}
	tmplFuncs["S3ImgSrcSmall"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320}, true)
	}
	tmplFuncs["S3ImgThumbSrcSmall"] = func(ctx context.Context, p string) template.HTMLAttr {
		return imgSrcAttr(ctx, p, []int{320}, false)
	}
	tmplFuncs["S3ImgSrc"] = func(ctx context.Context, p string, sizes []int) template.HTMLAttr {
		return imgSrcAttr(ctx, p, sizes, true)
	}
	tmplFuncs["S3ImgUrl"] = func(ctx context.Context, p string, size int) string {
		imgUrl := staticUrlFormatter(p)
		if cfg.Service.StaticS3.ImgResizeEnabled {
			imgUrl, _ = img_resize.S3ImgUrl(ctx, redisClient, staticS3UrlFormatter, awsSession, cfg.Aws.S3BucketPublic, imgResizeS3KeyPrefix, imgUrl, size)
		}
		return imgUrl
	}

	//
	t := template_renderer.NewTemplate(tmplFuncs)

	// global variables exposed for rendering of responses with templates
	gvd := map[string]interface{}{
		"_App": map[string]interface{}{
			"ENV":          cfg.Env,
			"BuildInfo":    cfg.BuildInfo,
			"BuildVersion": build,
		},
	}

	// Custom error handler to support rendering user friendly error page for improved web experience.
	eh := func(ctx context.Context, w http.ResponseWriter, r *http.Request, renderer web.Renderer, statusCode int, er error) error {
		data := map[string]interface{}{}

		return renderer.Render(ctx, w, r,
			"base.tmpl",  // base layout file to be used for rendering of errors
			"error.tmpl", // generic format for errors, could select based on status code
			web.MIMETextHTMLCharsetUTF8,
			http.StatusOK,
			data,
		)
	}

	// Enable template renderer to reload and parse template files when generating a response of dev
	// for a more developer friendly process. Any changes to the template files will be included
	// without requiring re-build/re-start of service.
	// This only supports files that already exist, if a new template file is added, then the
	// serivce needs to be restarted, but not rebuilt.
	enableHotReload := cfg.Env == "dev"

	// Template Renderer used to generate HTML response for web experience.
	renderer, err := template_renderer.NewTemplateRenderer(cfg.Service.TemplateDir, enableHotReload, gvd, t, eh)
	if err != nil {
		log.Fatalf("main : Marshalling Config to JSON : %+v", err)
	}

	// =========================================================================
	// Start Tracing Support
	th := fmt.Sprintf("%s:%d", cfg.Trace.Host, cfg.Trace.Port)
	log.Printf("main : Tracing Started : %s", th)
	sr := tracer.NewRateSampler(cfg.Trace.AnalyticsRate)
	tracer.Start(tracer.WithAgentAddr(th), tracer.WithSampler(sr))
	defer tracer.Stop()

	// =========================================================================
	// Start Debug Service. Not concerned with shutting this down when the
	// application is being shutdown.
	//
	// /debug/vars - Added to the default mux by the expvars package.
	// /debug/pprof - Added to the default mux by the net/http/pprof package.
	if cfg.Service.DebugHost != "" {
		go func() {
			log.Printf("main : Debug Listening %s", cfg.Service.DebugHost)
			log.Printf("main : Debug Listener closed : %v", http.ListenAndServe(cfg.Service.DebugHost, http.DefaultServeMux))
		}()
	}

	// =========================================================================
	// ECS Task registration for services that don't use an AWS Elastic Load Balancer.
	err = devops.EcsServiceTaskInit(log, awsSession)
	if err != nil {
		log.Fatalf("main : Ecs Service Task init : %+v", err)
	}

	// =========================================================================
	// Start APP Service

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Make an list of HTTP servers for both HTTP and HTTPS requests.
	var httpServers []http.Server

	// Start the HTTP service listening for requests.
	if cfg.HTTP.Host != "" {
		api := http.Server{
			Addr:           cfg.HTTP.Host,
			Handler:        handlers.APP(shutdown, log, cfg.Service.StaticDir, cfg.Service.TemplateDir, masterDb, redisClient, renderer, serviceMiddlewares...),
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}
		httpServers = append(httpServers, api)

		go func() {
			log.Printf("main : APP Listening %s", cfg.HTTP.Host)
			serverErrors <- api.ListenAndServe()
		}()
	}

	// Start the HTTPS service listening for requests with an SSL Cert auto generated with Let's Encrypt.
	if cfg.HTTPS.Host != "" {
		api := http.Server{
			Addr:           cfg.HTTPS.Host,
			Handler:        handlers.APP(shutdown, log, cfg.Service.StaticDir, cfg.Service.TemplateDir, masterDb, redisClient, renderer, serviceMiddlewares...),
			ReadTimeout:    cfg.HTTPS.ReadTimeout,
			WriteTimeout:   cfg.HTTPS.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		// Generate a unique list of hostnames.
		var hosts []string
		if primaryServiceHost != "" {
			hosts = append(hosts, primaryServiceHost)
		}
		for _, h := range cfg.Service.HostNames {
			h = strings.TrimSpace(h)
			if h != "" && h != primaryServiceHost {
				hosts = append(hosts, h)
			}
		}

		// Enable autocert to store certs via Secret Manager.
		secretPrefix := filepath.Join(cfg.Aws.SecretsManagerConfigPrefix, "autocert")

		// Local file cache to reduce requests hitting Secret Manager.
		localCache := autocert.DirCache(os.TempDir())

		cache, err := devops.NewSecretManagerAutocertCache(log, awsSession, secretPrefix, localCache)
		if err != nil {
			log.Fatalf("main : HTTPS : %+v", err)
		}

		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(hosts...),
			Cache:      cache,
		}
		api.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		api.TLSConfig.NextProtos = append(api.TLSConfig.NextProtos, acme.ALPNProto)
		if !cfg.HTTPS.DisableHTTP2 {
			api.TLSConfig.NextProtos = append(api.TLSConfig.NextProtos, "h2")
		}

		httpServers = append(httpServers, api)

		go func() {
			log.Printf("main : APP Listening %s with SSL cert for hosts %s", cfg.HTTPS.Host, strings.Join(hosts, ", "))
			serverErrors <- api.ListenAndServeTLS("", "")
		}()
	}

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("main : Error starting server: %+v", err)

	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown..", sig)

		// Create context for Shutdown call.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Service.ShutdownTimeout)
		defer cancel()

		// Handle closing connections for both possible HTTP servers.
		for _, api := range httpServers {

			// Asking listener to shutdown and load shed.
			err := api.Shutdown(ctx)
			if err != nil {
				log.Printf("main : Graceful shutdown did not complete in %v : %v", cfg.Service.ShutdownTimeout, err)
				err = api.Close()
			}

		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			log.Fatal("main : Integrity issue caused shutdown")
		case err != nil:
			log.Fatalf("main : Could not stop server gracefully : %+v", err)
		}
	}
}

// browserCacheBuster appends a the query string param v to a given url with
// a value based on the value returned from cacheBusterValueFunc
func browserCacheBuster(cacheBusterValueFunc func() string) func(uri string) string {
	f := func(uri string) string {
		v := cacheBusterValueFunc()
		if v == "" {
			return uri
		}

		u, err := url.Parse(uri)
		if err != nil {
			return ""
		}
		q := u.Query()
		q.Set("v", v)
		u.RawQuery = q.Encode()

		return u.String()
	}

	return f
}