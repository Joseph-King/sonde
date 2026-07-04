# Sonde
Sonde is a CLI based tool writtin in Go. It's main purpose is to send logs from your terminal and docker/podman/kube containers, throw them into an LLM, and give you a human readable insights.

CURRENTLY IN ACTIVE DEVELOPMENT

Roadmap:
- [ ] Phase 1: Initial MVP
  - [X] Pipe in logs from your terminal
  - [ ] Retrieve logs from previously ran commands
    - [X] terminal history to rerun commands
      - [X] fish
      - [ ] bash
      - [ ] zsh
    - [ ] tmux and others to get logs without rerunning commands
  - [ ] Initial connections to local LLM

- [ ] Initial Release!

- [ ] Phase 2: Container log retrieval
  - [ ] Retrieve logs from docker containers
  - [ ] Retrieve logs from podman containers
  - [ ] Retrieve logs from kubernetes pods

- [ ] Phase 3: reworking commands
  - [ ] sonde summarize (summarize the logs for you... in case you don't like reading)
    - [ ] summarize
    - [ ] s (shorthand)
  - [ ] sonde debug (Help you debug the logs, give actionable insights)
    - [ ] debug
    - [ ] d (shorthand)
  - [ ] sonde cleanup (Suggestions for you to cleanup the logs, remove sensitive info, duplicates, etc.)
    - [ ] cleanup
    - [ ] c (shorthand)
  - [ ] sonde metrics (Show you metrics about the logs, like how many errors, warnings, etc.)
    - [ ] metrics
    - [ ] m (shorthand)
  - [ ] sonde explain (Explain the logs for you, in case you don't understand what's going on)
    - [ ] explain
    - [ ] e (shorthand)
  - [ ] sonde audit (Looks at logs and look for active/escalating anomalies. Looking for potential security issues)
    - [ ] audit
    - [ ] a (shorthand)
  - [ ] sonde translate (Translate the logs into a file format (e.g. JSON, CSV)
    - [ ] translate
    - [ ] t (shorthand)
  - [ ] sonde inspect (Inspect the logs against standard best practices)
    - [ ] inspect
    - [ ] i (shorthand)

- [ ] Phase 4: Add flags for commands
  - [ ] Output to a file (Overwrites or creates a new file)
    - [ ] --output
    - [ ] -o (shorthand)
  - [ ] Append to a file
    - [ ] --append
    - [ ] -a (shorthand)
  - [ ] Model (change LLM model on the fly without updating the config)
    - [ ] --model
    - [ ] -m (shorthand)
  - [ ] Raw (Bypass sonde filtering, directly take all the logs as-is)
    - [ ] --raw
    - [ ] -r (shorthand)
  - [ ] Temperature (Set the temperature for the LLM model)
    - [ ] --temperature
    - [ ] -t (shorthand)
  - [ ] Filter (Apply a filter to the logs)
    - [ ] --filter
    - [ ] -f (shorthand)
  - [ ] Lines (Number of lines from the bottom of the logs to send)
    - [ ] --lines
    - [ ] -l (shorthand)
  - [ ] Silent (Suppress output to stdout)
    - [ ] --silent
    - [ ] -s (shorthand)
  - [ ] Provider (Specify the provider for the LLM model)
    - [ ] --provider
    - [ ] -p (shorthand)
  - [ ] Config (Override the config file for the LLM model)
    - [ ] --config
    - [ ] -c (shorthand)
  - [ ] Api Key (Override the API key for the LLM model)
    - [ ] --api-key
    - [ ] -k (shorthand)

- [ ] Phase 5: Rolling out more LLMs
  - [ ] LocalAI
  - [ ] Anthropic
  - [ ] OpenAI
  - [ ] DeepSeek
  - [ ] Google AI?
  - [ ] Amazon Bedrock?
  - [ ] GitHub?

- [ ] Phase 6: Caching and Context
  - [ ] Research caching and context strategies
  - [ ] Implement caching and context strategies
  - [ ] Research a way to link to frontend of chosen llm provider

Other things to consider
- Auto-completion
