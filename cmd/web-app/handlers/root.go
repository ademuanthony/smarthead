package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subject"
	"remoteschool/smarthead/internal/subscription"
	"remoteschool/smarthead/internal/timetable"
	"remoteschool/smarthead/internal/webroute"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

// Root represents the Root API method handler set.
type Root struct {
	StudentRepo      *student.Repository
	SubscriptionRepo *subscription.Repository
	ClassRepo        *class.Repository
	SubjectRepo      *subject.Repository
	PeriodRepo       *period.Repository
	TimetableRepo    *timetable.Repository
	Renderer         web.Renderer
	Sitemap          *stm.Sitemap
	WebRoute         webroute.WebRoute
}

// Index determines if the user has authentication and loads the associated page.
func (h *Root) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	if claims, err := auth.ClaimsFromContext(ctx); err == nil && claims.HasAuth() {
		if claims.HasRole(auth.RoleAdmin) {
			return h.indexDashboard(ctx, w, r, params)
		} else if claims.HasRole(auth.RoleUser) {
			return h.studentsDashboard(ctx, w, r, params)
		}
	}

	return h.indexDefault(ctx, w, r, params)
}

// indexDashboard loads the dashboard for a user when they are authenticated.
func (h *Root) indexDashboard(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "root-dashboard.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, nil)
}

// studentsDashboard loads deshboard for register student
func (h *Root) studentsDashboard(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	ctxValue, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	data := map[string]interface{}{}

	currentStudent, err := h.StudentRepo.CurrentStudent(ctx, claims)
	if err != nil {
		return err
	}
	data["student"] = currentStudent.Response(ctx)
	data["sssThreeStudent"] = strings.Contains(currentStudent.Class.Name, "SSS 3")
	r.ParseForm()
	dDay := currentStudent.CreatedAt.Sub(time.Now()).Hours()
	data["isNew"] = math.Abs(dDay) < 24*5

	classes, err := h.ClassRepo.Find(ctx, class.FindRequest{
		Order: []string{models.ClassColumns.SchoolOrder, models.ClassColumns.Name},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["classes"] = classes

	myTimetable, err := h.TimetableRepo.StudentsTimetables(ctx, currentStudent.ID)
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["timetables"] = myTimetable.Response(ctx)

	// wget -qO- https://ubuntu.bigbluebutton.org/bbb-install.sh | bash -s platform.smarthead.com.ng
	// ./bbb-install.sh -v xenial-220 -s platform.smarthead.com.ng -e ademuanthony@gmail.com
	// nortwestsouthnigeriaafricaebulokoimankonig
	// /etc/letsencrypt/live/turn.smarthead.com.ng/fullchain.pem
	// /etc/letsencrypt/live/turn.smarthead.com.ng/privkey.pem
	// "ECDH+AESGCM:ECDH+CHACHA20:DH+AESGCM:ECDH+AES256:DH+AES256:ECDH+AES128:DH+AES:RSA+AESGCM:RSA+AES:!aNULL:!MD5:!DSS"

	// /etc/letsencrypt/live/app.smarthead.com.ng/
	/*
		listen 443 ssl;
		listen [::]:443 ssl;

		ssl_certificate /etc/letsencrypt/live/app.smarthead.com.ng/fullchain.pem;
		ssl_certificate_key /etc/letsencrypt/live/app.smarthead.com.ng/privkey.pem;
		ssl_session_cache shared:SSL:10m;
		ssl_session_timeout 10m;
		ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
		ssl_ciphers "ECDH+AESGCM:DH+AESGCM:ECDH+AES256:DH+AES256:ECDH+AES128:DH+AES:ECDH+3DES:DH+3DES:RSA+AESGCM:RSA+AES:RSA+3DES:!aNULL:!MD5:!DSS:!AES256";
		ssl_prefer_server_ciphers on;
		ssl_dhparam /etc/nginx/ssl/dhp-4096.pem;
	*/

	periods, err := h.PeriodRepo.Find(ctx, claims, period.FindRequest{
		Order: []string{"start_hour"},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["periods"] = periods

	subjects, err := h.SubjectRepo.Find(ctx, claims, subject.FindRequest{
		Order: []string{"name"},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["subjects"] = subjects

	var limit uint = 5
	activeSubscriptions, err := h.SubscriptionRepo.Find(ctx, claims, subscription.FindRequest{
		Where: "student_id = ? AND end_date > ?",
		Args:  []interface{}{currentStudent.ID, ctxValue.Now.UTC().Unix()},
		Order: []string{"end_date desc"},
		Limit: &limit,
	})

	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}

	data["subscriptions"] = activeSubscriptions.Response(ctx)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "root-dashboard-students.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// indexDefault loads the root index page when a user has no authentication.
func (h *Root) indexDefault(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	classes, err := h.ClassRepo.Find(ctx, class.FindRequest{
		Order: []string{"school_order", "name"},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data := map[string]interface{}{}
	data["classes"] = classes
	return h.Renderer.Render(ctx, w, r, tmplLayoutSite, "dtox-index.html", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// SitePage loads the page with the layout for site instead of the app base.
func (h *Root) SitePage(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	data := make(map[string]interface{})

	var tmpName string
	switch r.RequestURI {
	case "/":
		tmpName = "dtox-index.html"
	case "/api":
		tmpName = "site-api.gohtml"

		// http://127.0.0.1:3001/docs/doc.json
		swaggerJsonUrl := h.WebRoute.ApiDocsJson(true)

		// Load the json file from the API service.
		res, err := pester.Get(swaggerJsonUrl)
		if err != nil {
			return errors.WithMessagef(err, "Failed to load url '%s' for api documentation.", swaggerJsonUrl)
		}

		// Read the entire response body.
		dat, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return errors.WithStack(err)
		}

		// Define the basic JSON struct for the JSON file.
		type swaggerInfo struct {
			Description string `json:"description"`
			Title       string `json:"title"`
			Version     string `json:"version"`
		}
		type swaggerDoc struct {
			Schemes  []string    `json:"schemes"`
			Swagger  string      `json:"swagger"`
			Info     swaggerInfo `json:"info"`
			Host     string      `json:"host"`
			BasePath string      `json:"basePath"`
		}

		// JSON decode the response body.
		var doc swaggerDoc
		err = json.Unmarshal(dat, &doc)
		if err != nil {
			return errors.WithStack(err)
		}

		data["urlApiBaseUri"] = h.WebRoute.WebApiUrl(doc.BasePath)
		data["urlApiDocs"] = h.WebRoute.ApiDocs()

	case "/about":
		tmpName = "dtox-about.html"
	case "/thank-you":
		tmpName = "dtox-thank-you.html"
	case "/contact":
		tmpName = "dtox-contact.html"
	case "/legal/privacy":
		tmpName = "dtox-privacy.gohtml"
	case "/legal/terms":
		tmpName = "dtox-terms.gohtml"
	default:
		return web.Redirect(ctx, w, r, "/", http.StatusFound)
	}

	return h.Renderer.Render(ctx, w, r, tmplLayoutSite, tmpName, web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// IndexHtml redirects /index.html to the website root page.
func (h *Root) IndexHtml(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	return web.Redirect(ctx, w, r, "/", http.StatusMovedPermanently)
}

// RobotHandler returns a robots.txt response.
func (h *Root) RobotTxt(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	if webcontext.ContextEnv(ctx) != webcontext.Env_Prod {
		txt := "User-agent: *\nDisallow: /"
		return web.RespondText(ctx, w, txt, http.StatusOK)
	}

	sitemapUrl := h.WebRoute.WebAppUrl("/sitemap.xml")

	txt := fmt.Sprintf("User-agent: *\nDisallow: /ping\nDisallow: /status\nDisallow: /debug/\nSitemap: %s", sitemapUrl)
	return web.RespondText(ctx, w, txt, http.StatusOK)
}

func (h *Root) SearchConsoleVerificationPage(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	fmt.Fprint(w, "google-site-verification: google6058e3992c01a0e3.html")
	return nil
}

type SiteMap struct {
	Pages []SiteMapPage `json:"pages"`
}

type SiteMapPage struct {
	Loc        string  `json:"loc" xml:"loc"`
	File       string  `json:"file" xml:"file"`
	Changefreq string  `json:"changefreq" xml:"changefreq"`
	Mobile     bool    `json:"mobile" xml:"mobile"`
	Priority   float64 `json:"priority" xml:"priority"`
	Lastmod    string  `json:"lastmod" xml:"lastmod"`
}

// SitemapXml returns a robots.txt response.
func (h *Root) SitemapXml(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	w.Write(h.Sitemap.XMLContent())
	return nil
}
