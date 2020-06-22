package handlers

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subscription"
	"remoteschool/smarthead/internal/timetable"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Periods represents the Periods API method handler set.
type Lessons struct {
	PeriodRepo 	 *period.Repository
	TimetableRepo *timetable.Repository
	StudentRepo *student.Repository
	SubscriptionRepo *subscription.Repository

	AGORA_APP_ID               string
	AGORA_CUSTOMER_ID          string
	AGORA_CUSTOMER_CERTIFICATE string

	Redis    *redis.Client
	Renderer web.Renderer
}

func (h *Lessons) agoraBasicCredentials() string {
	return b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", h.AGORA_CUSTOMER_ID, h.AGORA_CUSTOMER_CERTIFICATE)))
}

// View handles displaying a period.
func (h *Lessons) Join(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return web.Redirect(ctx, w, r, "/", 320)
	}

	ctxValue, err := webcontext.ContextValues(ctx)
	if err != nil {
		return web.Redirect(ctx, w, r, "/", 320)
	}
	
	tID := params["timetable_id"]
	timetable, err := h.TimetableRepo.ReadByID(ctx, claims, tID)
	if err != nil {
		return web.Redirect(ctx, w, r, "/", 320)
	}

	currentStudent, err := h.StudentRepo.CurrentStudent(ctx, claims)
	if err != nil {
		return web.Redirect(ctx, w, r, "/", 320)
	}

	has, err := h.SubscriptionRepo.StudentHasSubscription(ctx, currentStudent.ID, timetable.SubjectID, ctxValue.Now)
	if err != nil {
		return web.Redirect(ctx, w, r, "/", 320)
	}

	if !has {
		webcontext.SessionFlashError(ctx, "Access Denied", "You do not have access to this class")
		return web.Redirect(ctx, w, r, "/", 320)
	}

	// TODO: check that its time for the lesson
	http.Redirect(w, r, timetable.Subclass.Link, 301)
	// return web.Redirect(ctx, w, r, timetable.Subclass.Link, 301)
	return nil
}

func (h *Lessons) Certificate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error  {
	return web.RespondJson(ctx, w, h.agoraBasicCredentials(), 200)
}