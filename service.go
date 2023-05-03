func main() {
  app := &App{}
  app.startLogging()
  app.monitorOperatingSystemSignals()
  
  go app.processLabResults()
  
  select {} // block, so the program stays resident
}
