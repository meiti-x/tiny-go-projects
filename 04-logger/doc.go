/*
Package pocketlog exposes an API to log your work.
First, instantiate a logger with pocketlog.New, and giving it a threshold le
Messages of lesser criticality won't be logged.
Sharing the logger is the responsibility of the caller.
The logger can be called to log messages on three levels:
- Debug: mostly used to debug code, follow step-by-step processes
- Info: valuable messages providing insights to the milestones of a proce
- Error: error messages to understand what went wrong
*/
package main
