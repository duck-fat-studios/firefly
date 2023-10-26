# The URL where godoc serves the documentation
DOCS_URL=http://localhost:6060/pkg/github.com/duckfatstudios/firefly/

# The output directory where we'll save the documentation
OUTPUT_DIR=./docs

all: doc

# Start godoc in the background, fetch the documentation, then kill godoc
doc:
	@# No garuntees
	# pkill godoc

	@# Start godoc in the background. The '&' puts it in the background.
	godoc -http=:6060  &

	@# Give godoc a moment to start up
	sleep 5 ;\

	@# Fetch the docs with wget and keep the structure flat
	wget --recursive --no-clobber --page-requisites --convert-links --no-parent --no-directories -P $(OUTPUT_DIR) $(DOCS_URL) ;\

	@# Kill the godoc server
	pkill godoc;

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: doc clean

test:
	go test -v -cover