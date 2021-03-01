package controllers

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
)

func CreateInitControllersFn(wc *WebController) fasthttp.RequestHandler {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		path := ctx.URI().PathOriginal()

		// GET
		if ctx.IsGet() {
			switch utils.ByteToString(path) {
			case "/":
				wc.Index(ctx)
			case "/echo":
				wc.Echo(ctx)
			case "/Robots.txt":
				wc.Robots(ctx)
			default:
				pageRe := regexp.MustCompile(`^/page/[\w\-]+$`)
				if pageRe.Match(path) {
					wc.Page(ctx)
					return
				}

				appRe := regexp.MustCompile(`^/app/\w+$`)
				if appRe.Match(path) {
					wc.App(ctx)
					return
				}

				oauthRe := regexp.MustCompile(`^/oauth/\w+$`)
				if oauthRe.Match(path) {
					wc.OAuth(ctx)
					return
				}

				qrRe := regexp.MustCompile(`^/qr/(.*)$`)
				if qrRe.Match(path) {
					wc.Qr(ctx)
					return
				}
				// auth
				if checkUUID(ctx.Path(), wc.midClient) {
					memoRe := regexp.MustCompile(`^/memo/[\w\-]+$`)
					if memoRe.Match(path) {
						wc.Memo(ctx)
						return
					}
					appsRe := regexp.MustCompile(`^/apps/[\w\-]+$`)
					if appsRe.Match(path) {
						wc.Apps(ctx)
						return
					}
					credentialsRe := regexp.MustCompile(`^/credentials/[\w\-]+$`)
					if credentialsRe.Match(path) {
						wc.Credentials(ctx)
						return
					}
					credentialsCreateRe := regexp.MustCompile(`^/credentials/[\w\-]+/create$`)
					if credentialsCreateRe.Match(path) {
						wc.CredentialsCreate(ctx)
						return
					}
					settingRe := regexp.MustCompile(`^/setting/[\w\-]+$`)
					if settingRe.Match(path) {
						wc.Setting(ctx)
						return
					}
					settingCreateRe := regexp.MustCompile(`^/setting/[\w\-]+/create$`)
					if settingCreateRe.Match(path) {
						wc.SettingCreate(ctx)
						return
					}
					scriptsRe := regexp.MustCompile(`^/scripts/[\w\-]+$`)
					if scriptsRe.Match(path) {
						wc.Scripts(ctx)
						return
					}
					scriptCreateRe := regexp.MustCompile(`^/script/[\w\-]+/create$`)
					if scriptCreateRe.Match(path) {
						wc.ScriptCreate(ctx)
						return
					}
					scriptRunRe := regexp.MustCompile(`^/script/[\w\-]+/run$`)
					if scriptRunRe.Match(path) {
						wc.ScriptRun(ctx)
						return
					}
					actionRe := regexp.MustCompile(`^/action/[\w\-]+$`)
					if actionRe.Match(path) {
						wc.Action(ctx)
						return
					}
					actionCreateRe := regexp.MustCompile(`^/action/[\w\-]+/create$`)
					if actionCreateRe.Match(path) {
						wc.ActionCreate(ctx)
						return
					}
					actionRunRe := regexp.MustCompile(`^/action/[\w\-]+/run$`)
					if actionRunRe.Match(path) {
						wc.ActionRun(ctx)
						return
					}
				} else {
					ctx.Error("Forbidden", fasthttp.StatusForbidden)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}

		// POST
		if ctx.IsPost() {
			switch utils.ByteToString(path) {
			case "/":
				wc.Index(ctx)
			default:
				// auth
				if checkUUID(ctx.Path(), wc.midClient) {
					credentialsCreateRe := regexp.MustCompile(`^/credentials/[\w\-]+/store$`)
					if credentialsCreateRe.Match(path) {
						wc.CredentialsStore(ctx)
						return
					}
					settingCreateRe := regexp.MustCompile(`^/setting/[\w\-]+/store$`)
					if settingCreateRe.Match(path) {
						wc.SettingStore(ctx)
						return
					}
					scriptCreateRe := regexp.MustCompile(`^/script/[\w\-]+/store$`)
					if scriptCreateRe.Match(path) {
						wc.ScriptStore(ctx)
						return
					}
					actionCreateRe := regexp.MustCompile(`^/action/[\w\-]+/store$`)
					if actionCreateRe.Match(path) {
						wc.ActionStore(ctx)
						return
					}
				} else {
					ctx.Error("Forbidden", fasthttp.StatusForbidden)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}

func checkUUID(path []byte, midClient pb.MiddleClient) bool {
	uuid := utils.ExtractUUID(utils.ByteToString(path))
	if uuid == "" {
		return false
	}

	reply, err := midClient.Authorization(context.Background(), &pb.TextRequest{
		Text: uuid,
	})
	if err != nil {
		return false
	}

	return reply.GetState()
}
