package cmd

import "github.com/lsegal/gogcli/internal/googleapi"

var newCalendarService = googleapi.NewCalendar

const (
	scopeAll    = literalAll
	scopeSingle = "single"
	scopeFuture = "future"
)
