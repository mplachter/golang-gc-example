GC_TRACE_1:
	export GODEBUG=gctrace=1; cd example_01; go run .

GC_TRACE_PACER_1:
	export GODEBUG=gctrace=1,gcpacertrace=1; cd example_01; go run .

BUILD_EXAMPLE_1_GC_FLAG:
	cd example_01; go build -gcflags='-m' -o example_main_01 main.go 

BUILD_EXAMPLE_1:
	cd  example_01; go build .

EXAMPLE_1_TRACE_VIEW:
	go tool trace example_01/trace.out

EXAMPLE_1_RUN:
	cd example_01; ./example_main_01