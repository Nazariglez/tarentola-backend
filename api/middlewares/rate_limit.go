// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/utils"
	"net/http"
	"sync"
	"time"
)

var rateLimitIps = make(map[string]time.Time)
var rateLimitMu = sync.Mutex{}
var rateLimitCleaner *time.Ticker
var rateLimitRPS time.Duration

//from server side sdk send the header X-User-Final-IP with the real user ip (not the server addr)
func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	initRateLimitCleaner()

	return func(w http.ResponseWriter, r *http.Request) {
		ip := utils.GetIPAddr(r)
		if ip != "" {

			last, ok := getLastTime(ip)
			if ok {
				now := time.Now()

				rateLimitMu.Lock()
				expired := isRateLimitExpired(now, last)
				rateLimitMu.Unlock()

				if !expired {
					t := last.Add(rateLimitRPS)
					setLastTime(ip, t)
					time.Sleep(t.Sub(now))
				}
			} else {
				setLastTime(ip, time.Now())
			}
		}

		next.ServeHTTP(w, r)
	}
}

func isRateLimitExpired(now, t time.Time) bool {
	diff := now.Sub(t)
	return diff >= rateLimitRPS
}

func getLastTime(ip string) (time.Time, bool) {
	rateLimitMu.Lock()
	v, ok := rateLimitIps[ip]
	rateLimitMu.Unlock()
	return v, ok
}

func setLastTime(ip string, t time.Time) {
	rateLimitMu.Lock()
	rateLimitIps[ip] = t
	rateLimitMu.Unlock()
}

func initRateLimitCleaner() {
	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	if rateLimitCleaner != nil {
		return
	}

	rateLimitRPS = time.Second / time.Duration(config.Data.Middlewares.RateLimitRPS)

	rateLimitCleaner = time.NewTicker(time.Second * 15)
	go func() {
		for _ = range rateLimitCleaner.C {
			cleanRateLimitIps()
		}
	}()
}

func cleanRateLimitIps() {
	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	toDelete := []string{}
	now := time.Now()
	for k, t := range rateLimitIps {
		if isRateLimitExpired(now, t) {
			toDelete = append(toDelete, k)
		}
	}

	for _, k := range toDelete {
		delete(rateLimitIps, k)
	}
}
