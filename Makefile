VERSION ?= $(patsubst v%,%,$(shell git describe))

bin/pulumi-resource-fivetran: cmd/pulumi-resource-fivetran/schema.json
	go build -o bin/pulumi-resource-fivetran ./cmd/pulumi-resource-fivetran

bin/pulumi-tfgen-fivetran: cmd/pulumi-tfgen-fivetran/*.go go.sum provider/*.go
	go build -o bin/pulumi-tfgen-fivetran ./cmd/pulumi-tfgen-fivetran

cmd/pulumi-resource-fivetran/schema.json: bin/pulumi-tfgen-fivetran
	bin/pulumi-tfgen-fivetran $(VERSION) schema --out ./cmd/pulumi-resource-fivetran

schema: cmd/pulumi-resource-fivetran/schema.json

python-sdk: bin/pulumi-tfgen-fivetran
	rm -rf sdk
	bin/pulumi-tfgen-fivetran $(VERSION) python
	cp README.md sdk/python/
	cd sdk/python/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" setup.py && \
		rm setup.py.bak
