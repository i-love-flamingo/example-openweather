CONTEXT?=dev
REPLACE?=-replace flamingo.me/flamingo/v3=../../../flamingo/flamingo -replace flamingo.me/flamingo-commerce/v3=../../../flamingo/flamingo-commerce -replace go.aoe.com/flamingo-om3/v3=../../../flamingo/flamingo-om3 -replace flamingo.me/form=../../../flamingo/form
DROPREPLACE?=-dropreplace flamingo.me/flamingo/v3 -dropreplace flamingo.me/flamingo-commerce/v3 -dropreplace go.aoe.com/flamingo-om3/v3 -dropreplace flamingo.me/form

.PHONY: up localup update test serve frontend frontendbuild

serve:
	DEBUG=1 CONTEXT=$(CONTEXT) go run main.go serve

update:
	go get -u flamingo.me/flamingo/v3

local:
	git config filter.gomod-kso-flamingo.smudge 'go mod edit -fmt -print $(REPLACE) /dev/stdin'
	git config filter.gomod-kso-flamingo.clean 'go mod edit -fmt -print $(DROPREPLACE) /dev/stdin'
	git config filter.gomod-kso-flamingo.required true
	go mod edit -fmt $(REPLACE)

unlocal:
	git config filter.gomod-kso-flamingo.smudge ''
	git config filter.gomod-kso-flamingo.clean ''
	git config filter.gomod-kso-flamingo.required false
	go mod edit -fmt $(DROPREPLACE)

test:
	go test -test.vet "all" ./...

frontend-dev:
	bash -c 'cd frontend && npx flamingo-carotene dev'

frontend:
	bash -c 'cd frontend && npm ci && npx flamingo-carotene build'
