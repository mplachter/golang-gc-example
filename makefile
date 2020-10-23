GC_TRACE_1:
	export GODEBUG=gctrace=1; go run .

GC_TRACE_PACER_1:
	export GODEBUG=gctrace=1,gcpacertrace=1; go run .

BUILD_EXAMPLE_1_GC_FLAG:
	go build -gcflags='-m' -o example_main_01 .

BUILD_EXAMPLE_1:
	go build -o example_main_01 .

EXAMPLE_1_TRACE_VIEW:
	go tool trace trace.out

EXAMPLE_1_RUN:
	./example_main_01